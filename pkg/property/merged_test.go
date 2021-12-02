package property

import (
	"testing"

	"github.com/reearth/reearth-backend/pkg/dataset"
	"github.com/reearth/reearth-backend/pkg/id"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	ds := id.NewDatasetSchemaID()
	df := id.NewDatasetSchemaFieldID()
	d := id.NewDatasetID()
	d2 := id.NewDatasetID()
	opid := id.NewPropertyID()
	ppid := id.NewPropertyID()
	psid := id.MustPropertySchemaID("hoge~0.1.0/fff")
	psid2 := id.MustPropertySchemaID("hoge~0.1.0/aaa")
	psgid1 := id.PropertySchemaGroupID("group1")
	psgid2 := id.PropertySchemaGroupID("group2")
	psgid3 := id.PropertySchemaGroupID("group3")
	psgid4 := id.PropertySchemaGroupID("group4")
	i1id := id.NewPropertyItemID()
	i2id := id.NewPropertyItemID()
	i3id := id.NewPropertyItemID()
	i4id := id.NewPropertyItemID()
	i5id := id.NewPropertyItemID()
	i6id := id.NewPropertyItemID()
	i7id := id.NewPropertyItemID()
	i8id := id.NewPropertyItemID()

	fields1 := []*Field{
		NewFieldUnsafe().FieldUnsafe(id.PropertySchemaFieldID("a")).ValueUnsafe(OptionalValueFrom(ValueTypeString.ValueFrom("a"))).Build(),
		NewFieldUnsafe().FieldUnsafe(id.PropertySchemaFieldID("b")).ValueUnsafe(OptionalValueFrom(ValueTypeString.ValueFrom("b"))).Build(),
		NewFieldUnsafe().FieldUnsafe(id.PropertySchemaFieldID("e")).ValueUnsafe(NewOptionalValue(ValueTypeString, nil)).LinksUnsafe(dataset.NewGraphPointer([]*dataset.Pointer{dataset.PointAt(d2, ds, df)})).Build(),
		NewFieldUnsafe().FieldUnsafe(id.PropertySchemaFieldID("f")).ValueUnsafe(NewOptionalValue(ValueTypeNumber, nil)).Build(),
	}

	fields2 := []*Field{
		NewFieldUnsafe().FieldUnsafe(id.PropertySchemaFieldID("a")).ValueUnsafe(OptionalValueFrom(ValueTypeString.ValueFrom("1"))).Build(),
		NewFieldUnsafe().FieldUnsafe(id.PropertySchemaFieldID("c")).ValueUnsafe(OptionalValueFrom(ValueTypeString.ValueFrom("2"))).Build(),
		NewFieldUnsafe().FieldUnsafe(id.PropertySchemaFieldID("d")).ValueUnsafe(NewOptionalValue(ValueTypeString, nil)).LinksUnsafe(dataset.NewGraphPointer([]*dataset.Pointer{dataset.PointAtField(ds, df)})).Build(),
		NewFieldUnsafe().FieldUnsafe(id.PropertySchemaFieldID("f")).ValueUnsafe(NewOptionalValue(ValueTypeString, nil)).Build(),
	}

	groups1 := []*Group{
		NewGroup().ID(i7id).Schema(psid, psgid1).Fields(fields1).MustBuild(),
	}

	groups2 := []*Group{
		NewGroup().ID(i8id).Schema(psid, psgid1).Fields(fields2).MustBuild(),
	}

	items1 := []Item{
		NewGroupList().ID(i1id).Schema(psid, psgid1).Groups(groups1).MustBuild(),
		NewGroup().ID(i2id).Schema(psid, psgid2).Fields(fields1).MustBuild(),
		NewGroup().ID(i3id).Schema(psid, psgid3).Fields(fields1).MustBuild(),
	}

	items2 := []Item{
		NewGroupList().ID(i4id).Schema(psid, psgid1).Groups(groups2).MustBuild(),
		NewGroup().ID(i5id).Schema(psid, psgid2).Fields(fields2).MustBuild(),
		NewGroup().ID(i6id).Schema(psid, psgid4).Fields(fields2).MustBuild(),
	}

	sid := id.NewSceneID()
	op := New().ID(opid).Scene(sid).Schema(psid).Items(items1).MustBuild()
	ppempty := New().NewID().Scene(sid).Schema(psid2).MustBuild()
	pp := New().ID(ppid).Scene(sid).Schema(psid).Items(items2).MustBuild()

	// Merge(op, pp, &d)
	expected1 := &Merged{
		Original:      opid.Ref(),
		Parent:        ppid.Ref(),
		Schema:        psid,
		LinkedDataset: &d,
		Groups: []*MergedGroup{
			{
				Original:      &i1id,
				Parent:        &i4id,
				SchemaGroup:   psgid1,
				LinkedDataset: &d,
				Groups: []*MergedGroup{
					{
						Original:      &i7id,
						Parent:        nil,
						SchemaGroup:   psgid1,
						LinkedDataset: &d,
						Fields: []*MergedField{
							{
								ID:    id.PropertySchemaFieldID("a"),
								Value: ValueTypeString.ValueFrom("a"),
								Type:  ValueTypeString,
							},
							{
								ID:    id.PropertySchemaFieldID("b"),
								Value: ValueTypeString.ValueFrom("b"),
								Type:  ValueTypeString,
							},
							{
								ID:    id.PropertySchemaFieldID("e"),
								Links: dataset.NewGraphPointer([]*dataset.Pointer{dataset.PointAt(d2, ds, df)}),
								Type:  ValueTypeString,
							},
							{
								ID:   id.PropertySchemaFieldID("f"),
								Type: ValueTypeNumber,
							},
						},
					},
				},
			},
			{
				Original:      &i2id,
				Parent:        &i5id,
				SchemaGroup:   psgid2,
				LinkedDataset: &d,
				Fields: []*MergedField{
					{
						ID:         id.PropertySchemaFieldID("a"),
						Value:      ValueTypeString.ValueFrom("a"),
						Type:       ValueTypeString,
						Overridden: true,
					},
					{
						ID:    id.PropertySchemaFieldID("b"),
						Value: ValueTypeString.ValueFrom("b"),
						Type:  ValueTypeString,
					},
					{
						ID:    id.PropertySchemaFieldID("e"),
						Links: dataset.NewGraphPointer([]*dataset.Pointer{dataset.PointAt(d2, ds, df)}),
						Type:  ValueTypeString,
					},
					{
						ID:    id.PropertySchemaFieldID("c"),
						Value: ValueTypeString.ValueFrom("2"),
						Type:  ValueTypeString,
					},
					{
						ID:    id.PropertySchemaFieldID("d"),
						Links: dataset.NewGraphPointer([]*dataset.Pointer{dataset.PointAt(d, ds, df)}),
						Type:  ValueTypeString,
					},
				},
			},
			{
				Original:      &i3id,
				Parent:        nil,
				SchemaGroup:   psgid3,
				LinkedDataset: &d,
				Fields: []*MergedField{
					{
						ID:    id.PropertySchemaFieldID("a"),
						Value: ValueTypeString.ValueFrom("a"),
						Type:  ValueTypeString,
					},
					{
						ID:    id.PropertySchemaFieldID("b"),
						Value: ValueTypeString.ValueFrom("b"),
						Type:  ValueTypeString,
					},
					{
						ID:    id.PropertySchemaFieldID("e"),
						Links: dataset.NewGraphPointer([]*dataset.Pointer{dataset.PointAt(d2, ds, df)}),
						Type:  ValueTypeString,
					},
					{
						ID:   id.PropertySchemaFieldID("f"),
						Type: ValueTypeNumber,
					},
				},
			},
			{
				Original:      nil,
				Parent:        &i6id,
				SchemaGroup:   psgid4,
				LinkedDataset: &d,
				Fields: []*MergedField{
					{
						ID:    id.PropertySchemaFieldID("a"),
						Value: ValueTypeString.ValueFrom("1"),
						Type:  ValueTypeString,
					},
					{
						ID:    id.PropertySchemaFieldID("c"),
						Value: ValueTypeString.ValueFrom("2"),
						Type:  ValueTypeString,
					},
					{
						ID:    id.PropertySchemaFieldID("d"),
						Links: dataset.NewGraphPointer([]*dataset.Pointer{dataset.PointAt(d, ds, df)}),
						Type:  ValueTypeString,
					},
					{
						ID:   id.PropertySchemaFieldID("f"),
						Type: ValueTypeString,
					},
				},
			},
		},
	}

	// Merge(op, nil, &d)
	expected2 := &Merged{
		Original:      opid.Ref(),
		Parent:        nil,
		Schema:        psid,
		LinkedDataset: &d,
		Groups: []*MergedGroup{
			{
				Original:      &i1id,
				Parent:        nil,
				SchemaGroup:   psgid1,
				LinkedDataset: &d,
				Groups: []*MergedGroup{
					{
						Original:      &i7id,
						Parent:        nil,
						SchemaGroup:   psgid1,
						LinkedDataset: &d,
						Fields: []*MergedField{
							{
								ID:    id.PropertySchemaFieldID("a"),
								Value: ValueTypeString.ValueFrom("a"),
								Type:  ValueTypeString,
							},
							{
								ID:    id.PropertySchemaFieldID("b"),
								Value: ValueTypeString.ValueFrom("b"),
								Type:  ValueTypeString,
							},
							{
								ID:    id.PropertySchemaFieldID("e"),
								Links: dataset.NewGraphPointer([]*dataset.Pointer{dataset.PointAt(d2, ds, df)}),
								Type:  ValueTypeString,
							},
							{
								ID:   id.PropertySchemaFieldID("f"),
								Type: ValueTypeNumber,
							},
						},
					},
				},
			},
			{
				Original:      &i2id,
				Parent:        nil,
				SchemaGroup:   psgid2,
				LinkedDataset: &d,
				Fields: []*MergedField{
					{
						ID:    id.PropertySchemaFieldID("a"),
						Value: ValueTypeString.ValueFrom("a"),
						Type:  ValueTypeString,
					},
					{
						ID:    id.PropertySchemaFieldID("b"),
						Value: ValueTypeString.ValueFrom("b"),
						Type:  ValueTypeString,
					},
					{
						ID:    id.PropertySchemaFieldID("e"),
						Links: dataset.NewGraphPointer([]*dataset.Pointer{dataset.PointAt(d2, ds, df)}),
						Type:  ValueTypeString,
					},
					{
						ID:   id.PropertySchemaFieldID("f"),
						Type: ValueTypeNumber,
					},
				},
			},
			{
				Original:      &i3id,
				Parent:        nil,
				SchemaGroup:   psgid3,
				LinkedDataset: &d,
				Fields: []*MergedField{
					{
						ID:    id.PropertySchemaFieldID("a"),
						Value: ValueTypeString.ValueFrom("a"),
						Type:  ValueTypeString,
					},
					{
						ID:    id.PropertySchemaFieldID("b"),
						Value: ValueTypeString.ValueFrom("b"),
						Type:  ValueTypeString,
					},
					{
						ID:    id.PropertySchemaFieldID("e"),
						Links: dataset.NewGraphPointer([]*dataset.Pointer{dataset.PointAt(d2, ds, df)}),
						Type:  ValueTypeString,
					},
					{
						ID:   id.PropertySchemaFieldID("f"),
						Type: ValueTypeNumber,
					},
				},
			},
		},
	}

	// Merge(nil, pp, &d)
	expected3 := &Merged{
		Original:      nil,
		Parent:        ppid.Ref(),
		Schema:        psid,
		LinkedDataset: &d,
		Groups: []*MergedGroup{
			{
				Original:      nil,
				Parent:        &i4id,
				SchemaGroup:   psgid1,
				LinkedDataset: &d,
				Groups: []*MergedGroup{
					{
						Original:      nil,
						Parent:        &i8id,
						SchemaGroup:   psgid1,
						LinkedDataset: &d,
						Fields: []*MergedField{
							{
								ID:    id.PropertySchemaFieldID("a"),
								Value: ValueTypeString.ValueFrom("1"),
								Type:  ValueTypeString,
							},
							{
								ID:    id.PropertySchemaFieldID("c"),
								Value: ValueTypeString.ValueFrom("2"),
								Type:  ValueTypeString,
							},
							{
								ID:    id.PropertySchemaFieldID("d"),
								Links: dataset.NewGraphPointer([]*dataset.Pointer{dataset.PointAt(d, ds, df)}),
								Type:  ValueTypeString,
							},
							{
								ID:   id.PropertySchemaFieldID("f"),
								Type: ValueTypeString,
							},
						},
					},
				},
			},
			{
				Original:      nil,
				Parent:        &i5id,
				SchemaGroup:   psgid2,
				LinkedDataset: &d,
				Fields: []*MergedField{
					{
						ID:    id.PropertySchemaFieldID("a"),
						Value: ValueTypeString.ValueFrom("1"),
						Type:  ValueTypeString,
					},
					{
						ID:    id.PropertySchemaFieldID("c"),
						Value: ValueTypeString.ValueFrom("2"),
						Type:  ValueTypeString,
					},
					{
						ID:    id.PropertySchemaFieldID("d"),
						Links: dataset.NewGraphPointer([]*dataset.Pointer{dataset.PointAt(d, ds, df)}),
						Type:  ValueTypeString,
					},
					{
						ID:   id.PropertySchemaFieldID("f"),
						Type: ValueTypeString,
					},
				},
			},
			{
				Original:      nil,
				Parent:        &i6id,
				SchemaGroup:   psgid4,
				LinkedDataset: &d,
				Fields: []*MergedField{
					{
						ID:    id.PropertySchemaFieldID("a"),
						Value: ValueTypeString.ValueFrom("1"),
						Type:  ValueTypeString,
					},
					{
						ID:    id.PropertySchemaFieldID("c"),
						Value: ValueTypeString.ValueFrom("2"),
						Type:  ValueTypeString,
					},
					{
						ID:    id.PropertySchemaFieldID("d"),
						Links: dataset.NewGraphPointer([]*dataset.Pointer{dataset.PointAt(d, ds, df)}),
						Type:  ValueTypeString,
					},
					{
						ID:   id.PropertySchemaFieldID("f"),
						Type: ValueTypeString,
					},
				},
			},
		},
	}

	tests := []struct {
		name string
		o    *Property
		p    *Property
		l    *id.DatasetID
		want *Merged
	}{
		{
			name: "nil",
			o:    nil,
			p:    nil,
			l:    nil,
			want: nil,
		},
		{
			name: "empty parent",
			o:    op,
			p:    ppempty,
			l:    nil,
			want: nil,
		},
		{
			name: "ok",
			o:    op,
			p:    pp,
			l:    &d,
			want: expected1,
		},
		{
			name: "original only",
			o:    op,
			p:    nil,
			l:    &d,
			want: expected2,
		},
		{
			name: "parent only",
			o:    nil,
			p:    pp,
			l:    &d,
			want: expected3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Merge(tt.o, tt.p, tt.l)
			assert.Equal(t, tt.want, actual)
		})
	}
}
