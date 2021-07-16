package property

import "github.com/reearth/reearth-backend/pkg/id"

// GroupListBuilder _
type GroupListBuilder struct {
	p *GroupList
}

// NewGroupList _
func NewGroupList() *GroupListBuilder {
	return &GroupListBuilder{
		p: &GroupList{},
	}
}

// InitGroupListFrom _
func InitGroupListFrom(g *SchemaGroup) *GroupList {
	if g == nil || !g.IsList() {
		return nil
	}
	g2, _ := NewGroupList().NewID().Schema(g.Schema(), g.ID()).Build()
	return g2
}

// Build _
func (b *GroupListBuilder) Build() (*GroupList, error) {
	if id.ID(b.p.itemBase.ID).IsNil() {
		return nil, id.ErrInvalidID
	}
	return b.p, nil
}

// MustBuild _
func (b *GroupListBuilder) MustBuild() *GroupList {
	p, err := b.Build()
	if err != nil {
		panic(err)
	}
	return p
}

func (b *GroupListBuilder) base(base itemBase) *GroupListBuilder {
	b.p.itemBase = base
	return b
}

// ID _
func (b *GroupListBuilder) ID(id id.PropertyItemID) *GroupListBuilder {
	b.p.itemBase.ID = id
	return b
}

// NewID _
func (b *GroupListBuilder) NewID() *GroupListBuilder {
	b.p.itemBase.ID = id.NewPropertyItemID()
	return b
}

// Schema _
func (b *GroupListBuilder) Schema(s id.PropertySchemaID, g id.PropertySchemaFieldID) *GroupListBuilder {
	b.p.itemBase.Schema = s
	b.p.itemBase.SchemaGroup = g
	return b
}

// Groups _
func (b *GroupListBuilder) Groups(fields []*Group) *GroupListBuilder {
	newGroups := []*Group{}
	ids := map[id.PropertyItemID]struct{}{}
	for _, f := range fields {
		if f == nil {
			continue
		}
		if _, ok := ids[f.ID()]; ok {
			continue
		}
		ids[f.ID()] = struct{}{}
		newGroups = append(newGroups, f)
	}
	b.p.groups = newGroups
	return b
}