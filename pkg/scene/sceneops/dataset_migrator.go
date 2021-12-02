package sceneops

import (
	"context"

	"github.com/reearth/reearth-backend/internal/usecase/repo"
	"github.com/reearth/reearth-backend/pkg/dataset"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/layer"
	"github.com/reearth/reearth-backend/pkg/layer/layerops"
	"github.com/reearth/reearth-backend/pkg/plugin"
	"github.com/reearth/reearth-backend/pkg/property"
)

// TODO: define new loader types and use them instead of repos
type DatasetMigrator struct {
	PropertyRepo      repo.Property
	LayerRepo         repo.Layer
	DatasetSchemaRepo repo.DatasetSchema
	DatasetRepo       repo.Dataset
	Plugin            plugin.Loader
}

type MigrateDatasetResult struct {
	Layers                layer.Map
	Properties            property.Map
	RemovedLayers         *id.LayerIDSet
	RemovedDatasetSchemas []id.DatasetSchemaID
	RemovedDatasets       []id.DatasetID
}

func (r MigrateDatasetResult) Merge(r2 MigrateDatasetResult) MigrateDatasetResult {
	return MigrateDatasetResult{
		Layers:        r.Layers.Merge(r2.Layers),
		Properties:    r.Properties.Merge(r2.Properties),
		RemovedLayers: r.RemovedLayers.Merge(r2.RemovedLayers),
	}
}

func (srv DatasetMigrator) Migrate(ctx context.Context, sid id.SceneID, newdsl []*dataset.Schema, newdl dataset.List) (MigrateDatasetResult, error) {
	scenes := []id.SceneID{sid}
	result := MigrateDatasetResult{}

	mm, err := dataset.NewMigrationMap(
		newdsl,
		newdl,
		func(s *dataset.Schema) (dataset.SchemaList, error) {
			return srv.DatasetSchemaRepo.FindBySceneAndSource(ctx, sid, s.Source())
		},
		func(s *dataset.Schema) (dataset.List, error) {
			olddl, _, err := srv.DatasetRepo.FindBySchema(ctx, s.ID(), scenes, nil)
			return olddl, err
		},
	)
	if err != nil {
		return MigrateDatasetResult{}, err
	}

	dm := property.NewDatasetMigrator(mm.Migrator(), dataset.GraphLoaderFromMap(newdl.Map()))

	propeties, err := srv.PropertyRepo.FindLinkedAll(ctx, sid)
	if err != nil {
		return MigrateDatasetResult{}, err
	}

	for _, p := range propeties {
		if err := dm.MigrateProperty(ctx, p); err != nil {
			return MigrateDatasetResult{}, err
		}
	}

	result.Properties = propeties.Map()

	for _, newds := range newdsl {
		oldds := mm.OldSchema(newds.ID())
		if oldds == nil {
			continue
		}

		diff, ok := mm.Diff[newds.ID()]
		if !ok {
			continue
		}

		result2, err := srv.migrateLayer(ctx, sid, oldds, newds, diff)
		if err != nil {
			return MigrateDatasetResult{}, err
		}

		result = result.Merge(result2)
	}

	result.RemovedDatasetSchemas = append(result.RemovedDatasetSchemas, mm.DeletedSchemas.All()...)
	result.RemovedDatasets = append(result.RemovedDatasets, mm.Deleted.All()...)
	return result, nil
}

func (srv DatasetMigrator) migrateLayer(ctx context.Context, sid id.SceneID, oldds, newds *dataset.Schema, diff dataset.Diff) (MigrateDatasetResult, error) {
	scenes := []id.SceneID{sid}

	// 前のデータセットスキーマに紐づいたレイヤーグループを取得
	layerGroups, err := srv.LayerRepo.FindGroupBySceneAndLinkedDatasetSchema(ctx, sid, oldds.ID())
	if err != nil {
		return MigrateDatasetResult{}, err
	}

	addedAndUpdatedLayers := layer.List{}
	addedProperties := property.List{}
	removedLayers := []id.LayerID{}

	for _, lg := range layerGroups {
		layers, err := srv.LayerRepo.FindByIDs(ctx, lg.Layers().Layers(), scenes)
		if err != nil {
			return MigrateDatasetResult{}, err
		}

		// スキーマが消滅した場合
		if newds == nil {
			// レイヤーグループ自体をアンリンク
			lg.Unlink()
			// 子レイヤーを全て削除
			for _, l := range layers {
				if l == nil {
					continue
				}
				lid := (*l).ID()
				removedLayers = append(removedLayers, lid)
			}
			lg.Layers().Empty()
			continue
		}

		// レイヤーグループのリンク張り替えと名前変更
		lg.Link(newds.ID())
		if lg.Name() == oldds.Name() {
			lg.Rename(newds.Name())
		}

		// 消えたデータセット→レイヤーを削除
		for _, d := range diff.Removed {
			if l := layers.FindByDataset(d.ID()); l != nil {
				lg.Layers().RemoveLayer(l.ID())
				removedLayers = append(removedLayers, l.ID())
			}
		}

		// 追加されたデータセット→レイヤーを作成して追加
		if len(diff.Added) > 0 {
			// プラグインを取得
			var plug *plugin.Plugin
			if pid := lg.Plugin(); pid != nil {
				plug2, err := srv.Plugin(ctx, []id.PluginID{*pid}, []id.SceneID{sid})
				if err != nil || len(plug2) < 1 {
					return MigrateDatasetResult{}, err
				}
				plug = plug2[0]
			}

			representativeFieldID := newds.RepresentativeFieldID()
			for _, added := range diff.Added {
				did := added.ID()

				name := ""
				if rf := added.FieldRef(representativeFieldID); rf != nil && rf.Type() == dataset.ValueTypeString {
					name = rf.Value().Value().(string)
				}

				layerItem, property, err := layerops.LayerItem{
					SceneID:         sid,
					ParentLayerID:   lg.ID(),
					LinkedDatasetID: &did,
					Plugin:          plug,
					ExtensionID:     lg.Extension(),
					Name:            name,
				}.Initialize()
				if err != nil {
					return MigrateDatasetResult{}, err
				}

				var l layer.Layer = layerItem
				lg.Layers().AddLayer(layerItem.ID(), -1)
				addedAndUpdatedLayers = append(addedAndUpdatedLayers, &l)
				addedProperties = append(addedProperties, property)
			}
		}

		// 残りのデータセット→レイヤーのリンクを張り替え
		for olddsid, newds := range diff.Others {
			if il := layers.FindByDataset(olddsid); il != nil {
				var il2 layer.Layer = il
				il.Link(newds.ID())
				addedAndUpdatedLayers = append(addedAndUpdatedLayers, &il2)
			}
		}
	}

	layers := append(
		addedAndUpdatedLayers,
		layerGroups.ToLayerList()...,
	)

	set := id.NewLayerIDSet()
	set.Add(removedLayers...)

	return MigrateDatasetResult{
		Layers:        layers.Map(),
		Properties:    addedProperties.Map(),
		RemovedLayers: set,
	}, nil
}
