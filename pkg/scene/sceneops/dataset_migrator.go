package sceneops

import (
	"context"

	"github.com/reearth/reearth-backend/internal/usecase/repo"
	"github.com/reearth/reearth-backend/pkg/dataset"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/layer"
	"github.com/reearth/reearth-backend/pkg/layer/initializer"
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

// NOTE: DatasetSchemaの削除には対応していない（自動的に削除されない）
func (srv DatasetMigrator) Migrate(ctx context.Context, sid id.SceneID, newdsl []*dataset.Schema, newdl dataset.List) (MigrateDatasetResult, error) {
	scenes := []id.SceneID{sid}
	result := MigrateDatasetResult{}

	// 削除対象
	noLogerUsedDS := []id.DatasetSchemaID{}
	noLogerUsedD := []id.DatasetID{}

	// 古いDatasetSchema
	oldDatasetSchemaMap := map[id.DatasetSchemaID]*dataset.Schema{}
	// 新しいDatasetSchema
	newDatasetSchemaMap := map[id.DatasetSchemaID]*dataset.Schema{}
	// 新しいDatasetSchemaから古いDatasetSchemaIDへの対応
	datasetSchemaMapNewOld := map[id.DatasetSchemaID]id.DatasetSchemaID{}
	// 古いDatasetSchemaから新しいDatasetSchemaIDへの対応
	datasetSchemaMapOldNew := map[id.DatasetSchemaID]id.DatasetSchemaID{}
	// 古いDatasetFieldIDから新しいDatasetSchemaFieldIDへの対応
	datasetSchemaFieldIDMap := map[id.DatasetSchemaFieldID]id.DatasetSchemaFieldID{}
	// 古いDatasetから新しいDatasetへの対応
	newDatasetMap := map[id.DatasetID]*dataset.Dataset{}
	datasetMapOldNew := map[id.DatasetID]*dataset.Dataset{}
	datasetIDMapOldNew := map[id.DatasetID]id.DatasetID{}
	// 新しいDatasetSchemaからDatasetDiffへの対応
	datasetDiffMap := map[id.DatasetSchemaID]dataset.Diff{}

	// マップの作成
	for _, newds := range newdsl {
		newDatasetSchemaMap[newds.ID()] = newds

		// ソース元が同じ古いDSを取得
		olddsl, err := srv.DatasetSchemaRepo.FindBySceneAndSource(ctx, sid, newds.Source())
		if err != nil {
			return MigrateDatasetResult{}, err
		}

		// 古いデータセットを探す（新しく追加されたものも入り込んでいるので）
		var oldds *dataset.Schema
		for _, o := range olddsl {
			if o.ID() != newds.ID() {
				oldds = o
			}
		}
		if oldds == nil {
			// ないならリンクされているレイヤーやプロパティも作成されていないはずなので無視
			continue
		}

		oldDatasetSchemaMap[oldds.ID()] = oldds
		datasetSchemaMapNewOld[newds.ID()] = oldds.ID()
		datasetSchemaMapOldNew[oldds.ID()] = newds.ID()

		// フィールドの差分を取る
		fieldDiff := oldds.FieldDiffBySource(newds)
		for of, f := range fieldDiff.Replaced {
			datasetSchemaFieldIDMap[of] = f.ID()
		}

		// 古いDSのDを探し出す
		olddl, _, err := srv.DatasetRepo.FindBySchema(ctx, oldds.ID(), scenes, nil)
		if err != nil {
			return MigrateDatasetResult{}, err
		}

		// 削除対象に追加
		noLogerUsedDS = append(noLogerUsedDS, oldds.ID())
		for _, oldd := range olddl {
			noLogerUsedD = append(noLogerUsedD, oldd.ID())
		}

		// 新しいDSのDのみ抽出
		currentNewdl := newdl.FilterByDatasetSchema(newds.ID())

		// データセットの差分をとる
		diff := dataset.List(olddl).DiffBySource(currentNewdl)
		datasetDiffMap[newds.ID()] = diff
		for od, d := range diff.Others {
			datasetMapOldNew[od] = d
			datasetIDMapOldNew[od] = d.ID()
			newDatasetMap[d.ID()] = d
		}
	}

	// プロパティのマイグレーション
	propeties, err := srv.PropertyRepo.FindLinkedAll(ctx, sid)
	if err != nil {
		return MigrateDatasetResult{}, err
	}
	for _, p := range propeties {
		p.MigrateDataset(property.DatasetMigrationParam{
			OldDatasetSchemaMap: datasetSchemaMapOldNew,
			OldDatasetMap:       datasetIDMapOldNew,
			DatasetFieldIDMap:   datasetSchemaFieldIDMap,
			NewDatasetSchemaMap: newDatasetSchemaMap,
			NewDatasetMap:       newDatasetMap,
		})
	}
	result.Properties = propeties.Map()

	// 新しいDSでループ
	for _, newds := range newdsl {
		oldds := oldDatasetSchemaMap[datasetSchemaMapNewOld[newds.ID()]]
		if oldds == nil {
			// リンクされているレイヤーやプロパティも作成されていないはずなので無視
			continue
		}
		diff, ok := datasetDiffMap[newds.ID()]
		if !ok {
			continue
		}

		// レイヤーのマイグレーション
		result2, err := srv.migrateLayer(ctx, sid, oldds, newds, diff)
		if err != nil {
			return MigrateDatasetResult{}, err
		}

		result = result.Merge(result2)
	}

	result.RemovedDatasetSchemas = append(result.RemovedDatasetSchemas, noLogerUsedDS...)
	result.RemovedDatasets = append(result.RemovedDatasets, noLogerUsedD...)
	return result, nil
}

func (srv DatasetMigrator) migrateLayer(ctx context.Context, sid id.SceneID, oldds *dataset.Schema, newds *dataset.Schema, diff dataset.Diff) (MigrateDatasetResult, error) {
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
				plug2, err := srv.Plugin(ctx, *pid)
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

				layerItem, property, err := initializer.LayerItem{
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