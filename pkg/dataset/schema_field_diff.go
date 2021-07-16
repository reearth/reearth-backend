package dataset

import "github.com/reearth/reearth-backend/pkg/id"

// SchemaFieldDiff _
type SchemaFieldDiff struct {
	Added    []*SchemaField
	Removed  []*SchemaField
	Replaced map[id.DatasetSchemaFieldID]*SchemaField
}

// FieldDiffBySource _
func (d *Schema) FieldDiffBySource(d2 *Schema) SchemaFieldDiff {
	added := []*SchemaField{}
	removed := []*SchemaField{}
	// others := map[DatasetSource]DatasetDiffTouple{}
	others2 := map[id.DatasetSchemaFieldID]*SchemaField{}

	s1 := map[Source]*SchemaField{}
	for _, d1 := range d.fields {
		s1[d1.Source()] = d1
	}

	for _, d2 := range d2.fields {
		if d1, ok := s1[d2.Source()]; ok {
			others2[d1.ID()] = d2
		} else {
			// added
			added = append(added, d2)
		}
	}

	for _, d1 := range d.fields {
		if _, ok := others2[d1.ID()]; !ok {
			// removed
			removed = append(removed, d1)
		}
	}

	return SchemaFieldDiff{
		Added:    added,
		Removed:  removed,
		Replaced: others2,
	}
}