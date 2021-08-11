package tag

import (
	"errors"
	"testing"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/stretchr/testify/assert"
)

func TestBuilder_Build(t *testing.T) {
	tid := id.NewTagID()
	sid := id.NewSceneID()
	testCases := []struct {
		Name, Label string
		Id          id.TagID
		Scene       id.SceneID
		Expected    struct {
			Tag   tag
			Error error
		}
	}{
		{
			Name:  "fail: nil tag ID",
			Label: "xxx",
			Scene: id.NewSceneID(),
			Expected: struct {
				Tag   tag
				Error error
			}{
				Error: id.ErrInvalidID,
			},
		},
		{
			Name:  "fail: empty label",
			Id:    id.NewTagID(),
			Scene: id.NewSceneID(),
			Expected: struct {
				Tag   tag
				Error error
			}{
				Error: id.ErrInvalidID,
			},
		},
		{
			Name:  "fail: nil scene ID",
			Label: "xxx",
			Id:    id.NewTagID(),
			Expected: struct {
				Tag   tag
				Error error
			}{
				Error: ErrInvalidSceneID,
			},
		},
		{
			Name:  "success",
			Id:    tid,
			Label: "xxx",
			Scene: sid,
			Expected: struct {
				Tag   tag
				Error error
			}{
				Tag: tag{
					id:      tid,
					label:   "xxx",
					sceneId: sid,
				},
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.Name, func(tt *testing.T) {
			tt.Parallel()
			res, err := New().
				ID(tc.Id).
				Scene(tc.Scene).
				Label(tc.Label).
				Build()
			if err == nil {
				assert.Equal(tt, tc.Expected.Tag.ID(), res.ID())
				assert.Equal(tt, tc.Expected.Tag.Scene(), res.Scene())
				assert.Equal(tt, tc.Expected.Tag.Label(), res.Label())
			} else {
				assert.True(tt, errors.As(tc.Expected.Error, &err))
			}
		})
	}
}

func TestBuilder_NewID(t *testing.T) {
	b := New().NewID()
	assert.False(t, id.ID(b.t.id).IsNil())
}
