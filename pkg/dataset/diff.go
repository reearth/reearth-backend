package dataset

import "github.com/reearth/reearth-backend/pkg/id"

type Diff struct {
	Added   List
	Removed List
	Others  Map
}

type MigrationMap struct {
	Datasets       map[ID]ID
	Schemas        map[SchemaID]SchemaID
	Fields         map[FieldID]FieldID
	SchemaRev      map[SchemaID]SchemaID
	OldDatasets    Map
	OldSchemas     SchemaMap
	Deleted        *IDSet
	DeletedSchemas *SchemaIDSet
	Diff           map[SchemaID]Diff
}

func NewMigrationMap(newdsl SchemaList, newdl List, getOldSchemas func(*Schema) (SchemaList, error), getDatasets func(*Schema) (List, error)) (MigrationMap, error) {
	deleted := NewIDSet()
	deletedSchemas := NewSchemaIDDset()

	datasetMapOldNew := map[id.DatasetID]id.DatasetID{}
	schemaMapOldNew := map[id.DatasetSchemaID]id.DatasetSchemaID{}
	schemaMapNewOld := map[id.DatasetSchemaID]id.DatasetSchemaID{}
	schemaFieldMap := map[id.DatasetSchemaFieldID]id.DatasetSchemaFieldID{}
	oldSchemaMap := map[id.DatasetSchemaID]*Schema{}
	diffMap := map[id.DatasetSchemaID]Diff{}

	for _, newds := range newdsl {
		olddsl, err := getOldSchemas(newds)
		if err != nil {
			return MigrationMap{}, err
		}
		oldds := olddsl.Find(func(s *Schema) bool {
			return s.ID() != newds.ID()
		})
		if oldds == nil {
			continue
		}

		oldSchemaMap[oldds.ID()] = oldds
		schemaMapNewOld[newds.ID()] = oldds.ID()
		schemaMapOldNew[oldds.ID()] = newds.ID()

		fieldDiff := oldds.FieldDiffBySource(newds)
		for of, f := range fieldDiff.Replaced {
			schemaFieldMap[of] = f.ID()
		}

		olddl, err := getDatasets(oldds)
		if err != nil {
			return MigrationMap{}, err
		}

		deletedSchemas.Add(oldds.ID())
		for _, oldd := range olddl {
			deleted.Add(oldd.ID())
		}

		currentNewdl := newdl.FilterByDatasetSchema(newds.ID())
		diff := List(olddl).DiffBySource(currentNewdl)
		diffMap[newds.ID()] = diff
		for od, d := range diff.Others {
			datasetMapOldNew[od] = d.ID()
		}
	}

	return MigrationMap{
		Datasets:       datasetMapOldNew,
		Schemas:        schemaMapOldNew,
		Fields:         schemaFieldMap,
		SchemaRev:      schemaMapNewOld,
		OldSchemas:     oldSchemaMap,
		Deleted:        deleted,
		DeletedSchemas: deletedSchemas,
		Diff:           diffMap,
	}, nil
}

func (mm MigrationMap) OldSchema(newsid SchemaID) *Schema {
	oldsid, ok := mm.SchemaRev[newsid]
	if !ok {
		return nil
	}
	oldds, ok := mm.OldSchemas[oldsid]
	if !ok {
		return nil
	}
	return oldds
}

func (mm MigrationMap) Migrator() *Migrator {
	return MigratorFrom(mm.Datasets, mm.Schemas, mm.Fields)
}
