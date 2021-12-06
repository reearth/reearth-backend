package property

import "github.com/reearth/reearth-backend/pkg/id"

type ID = id.PropertyID
type ItemID = id.PropertyItemID
type FieldID = id.PropertySchemaFieldID
type SchemaID = id.PropertySchemaID
type SchemaGroupID = id.PropertySchemaGroupID

type IDSet = id.PropertyIDSet
type ItemIDSet = id.PropertyItemIDSet

var NewID = id.NewPropertyID
var NewItemID = id.NewPropertyItemID

var IDFrom = id.PropertyIDFrom
var SchemaIDFrom = id.PropertySchemaIDFrom
var FieldIDFrom = id.PropertySchemaFieldIDFrom
var ItemIDFrom = id.PropertyItemIDFrom

var MustID = id.MustPropertyID
var MustSchemaID = id.MustPropertySchemaID
var MustItemID = id.MustPropertyItemID

var NewIDSet = id.NewPropertyIDSet
var NewItemIDSet = id.NewPropertyItemIDSet
