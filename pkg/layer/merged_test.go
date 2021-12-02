package layer

import (
	"testing"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/property"
	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	scene := id.NewSceneID()
	dataset1 := id.NewDatasetID()
	p := id.MustPluginID("xxx~1.1.1")
	e := id.PluginExtensionID("foo")

	itemProperty := id.NewPropertyID()
	groupProperty := id.NewPropertyID()
	ib1pr := id.NewPropertyID()
	ib2pr := id.NewPropertyID()
	f1pr := id.NewPropertyID()
	f2pr := id.NewPropertyID()
	f3pr := id.NewPropertyID()

	f1 := NewInfoboxField().NewID().Plugin(p).Extension(e).Property(f1pr).MustBuild()
	f2 := NewInfoboxField().NewID().Plugin(p).Extension(e).Property(f2pr).MustBuild()
	f3 := NewInfoboxField().NewID().Plugin(p).Extension(e).Property(f3pr).MustBuild()

	// no-infobox and no-linked
	itemLayer1 := NewItem().
		NewID().
		Scene(scene).
		Plugin(&p).
		Extension(&e).
		Property(&itemProperty).
		MustBuild()
	// no-infobox
	itemLayer2 := NewItem().
		NewID().
		Scene(scene).
		Plugin(&p).
		Extension(&e).
		Property(&itemProperty).
		LinkedDataset(&dataset1).
		MustBuild()
	// infobox
	itemLayer3 := NewItem().
		NewID().
		Scene(scene).
		Plugin(&p).
		Extension(&e).
		Property(&itemProperty).
		LinkedDataset(&dataset1).
		Infobox(NewInfobox([]*InfoboxField{f1, f3}, ib1pr)).
		MustBuild()
	// infobox but field is empty
	itemLayer4 := NewItem().
		NewID().
		Scene(scene).
		Plugin(&p).
		Extension(&e).
		Property(&itemProperty).
		LinkedDataset(&dataset1).
		Infobox(NewInfobox(nil, ib1pr)).
		MustBuild()
	// no-infobox
	groupLayer1 := NewGroup().
		NewID().
		Scene(scene).
		Plugin(&p).
		Extension(&e).
		Property(&groupProperty).
		MustBuild()
	// infobox
	groupLayer2 := NewGroup().
		NewID().
		Scene(scene).
		Plugin(&p).
		Extension(&e).
		Property(&groupProperty).
		Infobox(NewInfobox([]*InfoboxField{f2, f3}, ib2pr)).
		MustBuild()

	tests := []struct {
		name string
		o    Layer
		p    *Group
		want *Merged
	}{
		{
			name: "nil",
			o:    nil,
			p:    nil,
			want: nil,
		},
		{
			name: "parent only",
			o:    nil,
			p:    groupLayer1,
			want: nil,
		},
		{
			name: "only original without infobox and link",
			o:    itemLayer1,
			p:    nil,
			want: &Merged{
				Original:    itemLayer1.ID(),
				Parent:      nil,
				Scene:       scene,
				PluginID:    &p,
				ExtensionID: &e,
				Property: &property.MergedMetadata{
					Original:      &itemProperty,
					Parent:        nil,
					LinkedDataset: nil,
				},
			},
		},
		{
			name: "only original with infobox",
			o:    itemLayer3,
			p:    nil,
			want: &Merged{
				Original:    itemLayer3.ID(),
				Parent:      nil,
				Scene:       scene,
				PluginID:    &p,
				ExtensionID: &e,
				Property: &property.MergedMetadata{
					Original:      &itemProperty,
					Parent:        nil,
					LinkedDataset: &dataset1,
				},
				Infobox: &MergedInfobox{
					Property: &property.MergedMetadata{
						Original:      &ib1pr,
						Parent:        nil,
						LinkedDataset: &dataset1,
					},
					Fields: []*MergedInfoboxField{
						{
							ID:        f1.ID(),
							Plugin:    p,
							Extension: e,
							Property: &property.MergedMetadata{
								Original:      &f1pr,
								Parent:        nil,
								LinkedDataset: &dataset1,
							},
						},
						{
							ID:        f3.ID(),
							Plugin:    p,
							Extension: e,
							Property: &property.MergedMetadata{
								Original:      &f3pr,
								Parent:        nil,
								LinkedDataset: &dataset1,
							},
						},
					},
				},
			},
		},
		{
			name: "original without infobox, parent without infobox",
			o:    itemLayer2,
			p:    groupLayer1,
			want: &Merged{
				Original:    itemLayer2.ID(),
				Parent:      groupLayer1.IDRef(),
				Scene:       scene,
				PluginID:    &p,
				ExtensionID: &e,
				Property: &property.MergedMetadata{
					Original:      &itemProperty,
					Parent:        &groupProperty,
					LinkedDataset: &dataset1,
				},
			},
		},
		{
			name: "original with infobox, parent without infobox",
			o:    itemLayer3,
			p:    groupLayer1,
			want: &Merged{
				Original:    itemLayer3.ID(),
				Parent:      groupLayer1.IDRef(),
				Scene:       scene,
				PluginID:    &p,
				ExtensionID: &e,
				Property: &property.MergedMetadata{
					Original:      &itemProperty,
					Parent:        &groupProperty,
					LinkedDataset: &dataset1,
				},
				Infobox: &MergedInfobox{
					Property: &property.MergedMetadata{
						Original:      &ib1pr,
						Parent:        nil,
						LinkedDataset: &dataset1,
					},
					Fields: []*MergedInfoboxField{
						{
							ID:        f1.ID(),
							Plugin:    p,
							Extension: e,
							Property: &property.MergedMetadata{
								Original:      &f1pr,
								Parent:        nil,
								LinkedDataset: &dataset1,
							},
						},
						{
							ID:        f3.ID(),
							Plugin:    p,
							Extension: e,
							Property: &property.MergedMetadata{
								Original:      &f3pr,
								Parent:        nil,
								LinkedDataset: &dataset1,
							},
						},
					},
				},
			},
		},
		{
			name: "original without infobox, parent with infobox",
			o:    itemLayer2,
			p:    groupLayer2,
			want: &Merged{
				Original:    itemLayer2.ID(),
				Parent:      groupLayer2.IDRef(),
				Scene:       scene,
				PluginID:    &p,
				ExtensionID: &e,
				Property: &property.MergedMetadata{
					Original:      &itemProperty,
					Parent:        &groupProperty,
					LinkedDataset: &dataset1,
				},
				Infobox: &MergedInfobox{
					Property: &property.MergedMetadata{
						Original:      nil,
						Parent:        &ib2pr,
						LinkedDataset: &dataset1,
					},
					Fields: []*MergedInfoboxField{
						{
							ID:        f2.ID(),
							Plugin:    p,
							Extension: e,
							Property: &property.MergedMetadata{
								Original:      &f2pr,
								Parent:        nil,
								LinkedDataset: &dataset1,
							},
						},
						{
							ID:        f3.ID(),
							Plugin:    p,
							Extension: e,
							Property: &property.MergedMetadata{
								Original:      &f3pr,
								Parent:        nil,
								LinkedDataset: &dataset1,
							},
						},
					},
				},
			},
		},
		{
			name: "original with infobox, parent with infobox",
			o:    itemLayer3,
			p:    groupLayer2,
			want: &Merged{
				Original:    itemLayer3.ID(),
				Parent:      groupLayer2.IDRef(),
				Scene:       scene,
				PluginID:    &p,
				ExtensionID: &e,
				Property: &property.MergedMetadata{
					Original:      &itemProperty,
					Parent:        &groupProperty,
					LinkedDataset: &dataset1,
				},
				Infobox: &MergedInfobox{
					Property: &property.MergedMetadata{
						Original:      &ib1pr,
						Parent:        &ib2pr,
						LinkedDataset: &dataset1,
					},
					Fields: []*MergedInfoboxField{
						{
							ID:        f1.ID(),
							Plugin:    p,
							Extension: e,
							Property: &property.MergedMetadata{
								Original:      &f1pr,
								Parent:        nil,
								LinkedDataset: &dataset1,
							},
						},
						{
							ID:        f3.ID(),
							Plugin:    p,
							Extension: e,
							Property: &property.MergedMetadata{
								Original:      &f3pr,
								Parent:        nil,
								LinkedDataset: &dataset1,
							},
						},
					},
				},
			},
		},
		{
			name: "original with infobox but field is empty, parent with infobox",
			o:    itemLayer4,
			p:    groupLayer2,
			want: &Merged{
				Original:    itemLayer4.ID(),
				Parent:      groupLayer2.IDRef(),
				Scene:       scene,
				PluginID:    &p,
				ExtensionID: &e,
				Property: &property.MergedMetadata{
					Original:      &itemProperty,
					Parent:        &groupProperty,
					LinkedDataset: &dataset1,
				},
				Infobox: &MergedInfobox{
					Property: &property.MergedMetadata{
						Original:      &ib1pr,
						Parent:        &ib2pr,
						LinkedDataset: &dataset1,
					},
					Fields: []*MergedInfoboxField{},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := Merge(tt.o, tt.p)
			assert.Equal(t, tt.want, actual)
		})
	}
}

func TestMergedProperties(t *testing.T) {
	itemProperty := id.NewPropertyID()
	groupProperty := id.NewPropertyID()
	ib1pr := id.NewPropertyID()
	ib2pr := id.NewPropertyID()
	f1pr := id.NewPropertyID()
	f2pr := id.NewPropertyID()
	f3pr := id.NewPropertyID()

	merged := &Merged{
		Property: &property.MergedMetadata{
			Original: &itemProperty,
			Parent:   &groupProperty,
		},
		Infobox: &MergedInfobox{
			Property: &property.MergedMetadata{
				Original: &ib1pr,
				Parent:   &ib2pr,
			},
			Fields: []*MergedInfoboxField{
				{
					Property: &property.MergedMetadata{
						Original: &f1pr,
						Parent:   &f2pr,
					},
				},
				{
					Property: &property.MergedMetadata{
						Original: &f3pr,
						Parent:   nil,
					},
				},
			},
		},
	}

	assert.Equal(t, []id.PropertyID{
		itemProperty, groupProperty, ib1pr, ib2pr, f1pr, f2pr, f3pr,
	}, merged.Properties())
}
