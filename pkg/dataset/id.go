package dataset

import "github.com/reearth/reearth-backend/pkg/id"

type ID = id.DatasetID
type IDSet = id.DatasetIDSet

type SchemaID = id.DatasetSchemaID
type SchemaIDSet = id.DatasetSchemaIDSet

type FieldID = id.DatasetSchemaFieldID
type FieldIDSet = id.DatasetSchemaFieldIDSet

var NewID = id.NewDatasetID
var NewSchemaID = id.NewDatasetSchemaID
var NewFieldID = id.NewDatasetSchemaFieldID

var MustID = id.MustDatasetID
var MustSchemaID = id.MustDatasetSchemaID
var MustFieldID = id.MustDatasetSchemaFieldID

var IDFrom = id.DatasetIDFrom
var SchemaIDFrom = id.DatasetSchemaIDFrom
var FieldIDFrom = id.DatasetSchemaFieldIDFrom

var NewIDSet = id.NewDatasetIDSet
var NewSchemaIDDset = id.NewDatasetSchemaIDSet
var NewFieldIDDset = id.NewDatasetSchemaFieldIDSet
