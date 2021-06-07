package property

import "github.com/reearth/reearth-backend/pkg/id"

type MigrationPlans []MigrationPlan

type MigrationPlan struct {
	FromItem id.PropertySchemaFieldID
	From     id.PropertySchemaFieldID
	ToItem   id.PropertySchemaFieldID
	To       id.PropertySchemaFieldID
}
