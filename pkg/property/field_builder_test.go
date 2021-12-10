package property

import (
	"testing"

	"github.com/reearth/reearth-backend/pkg/dataset"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestFieldBuilder_Value(t *testing.T) {
	v := ValueTypeString.ValueFrom("vvv")
	b := NewField().Field("a").Value(OptionalValueFrom(v)).Build()
	assert.Equal(t, v, b.Value())
}

func TestFieldBuilder_Link(t *testing.T) {
	l := dataset.NewGraphPointer([]*dataset.Pointer{
		dataset.PointAt(id.NewDatasetID(), id.NewDatasetSchemaID(), id.NewDatasetSchemaFieldID()),
	})

	tests := []struct {
		Name  string
		Links *dataset.GraphPointer
	}{
		{
			Name:  "success",
			Links: l,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			res := NewField().Field(FieldID("a")).Value(NewOptionalValue(ValueTypeBool, nil)).Link(tt.Links).Build()
			assert.Equal(t, l, res.Links())
		})
	}
}
