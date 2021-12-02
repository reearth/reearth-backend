package property

import (
	"context"

	"github.com/reearth/reearth-backend/pkg/dataset"
)

type Migrator struct {
	NewSchema *Schema
	Plans     []MigrationPlan
}

type MigrationPlan struct {
	From *Pointer
	To   *Pointer
}

// func (m Migrator) Migrate(from *Property) *Property {

// }

type DatasetMigrator struct {
	m *dataset.Migrator
	l dataset.GraphLoader
}

func NewDatasetMigrator(m *dataset.Migrator, l dataset.GraphLoader) *DatasetMigrator {
	return &DatasetMigrator{
		m: m,
		l: l,
	}
}

func (m *DatasetMigrator) MigrateProperty(ctx context.Context, p *Property) error {
	if m == nil {
		return nil
	}
	for _, i := range p.Items() {
		if g := ToGroup(i); g != nil {
			if err := m.MigrateGroup(ctx, g); err != nil {
				return err
			}
		} else if l := ToGroupList(i); l != nil {
			if err := m.MigrateGroupList(ctx, l); err != nil {
				return err
			}
		}
	}
	return nil
}

func (m *DatasetMigrator) MigrateGroupList(ctx context.Context, g *GroupList) error {
	if m == nil {
		return nil
	}
	for _, g := range g.Groups() {
		if err := m.MigrateGroup(ctx, g); err != nil {
			return err
		}
	}
	return nil
}

func (m *DatasetMigrator) MigrateGroup(ctx context.Context, g *Group) error {
	if m == nil {
		return nil
	}
	for _, f := range g.Fields() {
		if err := m.MigrateField(ctx, f); err != nil {
			return err
		}
	}
	return nil
}

func (m *DatasetMigrator) MigrateField(ctx context.Context, f *Field) error {
	if m == nil {
		return nil
	}
	nl, err := m.m.MigrateAndValidateGraphPointer(ctx, f.Links(), m.l)
	if err != nil {
		return err
	}
	f.Link(nl)
	return nil
}
