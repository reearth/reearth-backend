// Code generated by gen, DO NOT EDIT.

package id

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/oklog/ulid"
	"github.com/stretchr/testify/assert"
)

func TestNewLayerID(t *testing.T) {
	id := NewLayerID()
	assert.NotNil(t, id)
	ulID, err := ulid.Parse(id.String())

	assert.NotNil(t, ulID)
	assert.Nil(t, err)
}

func TestLayerIDFrom(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    string
		expected struct {
			result LayerID
			err    error
		}
	}{
		{
			name:  "Fail:Not valid string",
			input: "testMustFail",
			expected: struct {
				result LayerID
				err    error
			}{
				LayerID{},
				ErrInvalidID,
			},
		},
		{
			name:  "Fail:Not valid string",
			input: "",
			expected: struct {
				result LayerID
				err    error
			}{
				LayerID{},
				ErrInvalidID,
			},
		},
		{
			name:  "success:valid string",
			input: "01f2r7kg1fvvffp0gmexgy5hxy",
			expected: struct {
				result LayerID
				err    error
			}{
				LayerID{ulid.MustParse("01f2r7kg1fvvffp0gmexgy5hxy")},
				nil,
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			result, err := LayerIDFrom(tc.input)
			assert.Equal(tt, tc.expected.result, result)
			if err != nil {
				assert.True(tt, errors.As(tc.expected.err, &err))
			}
		})
	}
}

func TestMustLayerID(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name        string
		input       string
		shouldPanic bool
		expected    LayerID
	}{
		{
			name:        "Fail:Not valid string",
			input:       "testMustFail",
			shouldPanic: true,
		},
		{
			name:        "Fail:Not valid string",
			input:       "",
			shouldPanic: true,
		},
		{
			name:        "success:valid string",
			input:       "01f2r7kg1fvvffp0gmexgy5hxy",
			shouldPanic: false,
			expected:    LayerID{ulid.MustParse("01f2r7kg1fvvffp0gmexgy5hxy")},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			if tc.shouldPanic {
				assert.Panics(tt, func() { MustBeID(tc.input) })
				return
			}
			result := MustLayerID(tc.input)
			assert.Equal(tt, tc.expected, result)
		})
	}
}

func TestLayerIDFromRef(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected *LayerID
	}{
		{
			name:     "Fail:Not valid string",
			input:    "testMustFail",
			expected: nil,
		},
		{
			name:     "Fail:Not valid string",
			input:    "",
			expected: nil,
		},
		{
			name:     "success:valid string",
			input:    "01f2r7kg1fvvffp0gmexgy5hxy",
			expected: &LayerID{ulid.MustParse("01f2r7kg1fvvffp0gmexgy5hxy")},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			result := LayerIDFromRef(&tc.input)
			assert.Equal(tt, tc.expected, result)
			if tc.expected != nil {
				assert.Equal(tt, *tc.expected, *result)
			}
		})
	}
}

func TestLayerIDFromRefID(t *testing.T) {
	id := New()

	subId := LayerIDFromRefID(&id)

	assert.NotNil(t, subId)
	assert.Equal(t, subId.id, id.id)
}

func TestLayerID_ID(t *testing.T) {
	id := New()
	subId := LayerIDFromRefID(&id)

	idOrg := subId.ID()

	assert.Equal(t, id, idOrg)
}

func TestLayerID_String(t *testing.T) {
	id := New()
	subId := LayerIDFromRefID(&id)

	assert.Equal(t, subId.String(), id.String())
}

func TestLayerID_GoString(t *testing.T) {
	id := New()
	subId := LayerIDFromRefID(&id)

	assert.Equal(t, subId.GoString(), "id.LayerID("+id.String()+")")
}

func TestLayerID_RefString(t *testing.T) {
	id := New()
	subId := LayerIDFromRefID(&id)

	refString := subId.StringRef()

	assert.NotNil(t, refString)
	assert.Equal(t, *refString, id.String())
}

func TestLayerID_Ref(t *testing.T) {
	id := New()
	subId := LayerIDFromRefID(&id)

	subIdRef := subId.Ref()

	assert.Equal(t, *subId, *subIdRef)
}

func TestLayerID_Contains(t *testing.T) {
	id := NewLayerID()
	id2 := NewLayerID()
	assert.True(t, id.Contains([]LayerID{id, id2}))
	assert.False(t, id.Contains([]LayerID{id2}))
}

func TestLayerID_CopyRef(t *testing.T) {
	id := New()
	subId := LayerIDFromRefID(&id)

	subIdCopyRef := subId.CopyRef()

	assert.Equal(t, *subId, *subIdCopyRef)
	assert.NotSame(t, subId, subIdCopyRef)
}

func TestLayerID_IDRef(t *testing.T) {
	id := New()
	subId := LayerIDFromRefID(&id)

	assert.Equal(t, id, *subId.IDRef())
}

func TestLayerID_StringRef(t *testing.T) {
	id := New()
	subId := LayerIDFromRefID(&id)

	assert.Equal(t, *subId.StringRef(), id.String())
}

func TestLayerID_MarhsalJSON(t *testing.T) {
	id := New()
	subId := LayerIDFromRefID(&id)

	res, err := subId.MarhsalJSON()
	exp, _ := json.Marshal(subId.String())

	assert.Nil(t, err)
	assert.Equal(t, exp, res)
}

func TestLayerID_UnmarhsalJSON(t *testing.T) {
	jsonString := "\"01f3zhkysvcxsnzepyyqtq21fb\""

	subId := &LayerID{}

	err := subId.UnmarhsalJSON([]byte(jsonString))

	assert.Nil(t, err)
	assert.Equal(t, "01f3zhkysvcxsnzepyyqtq21fb", subId.String())
}

func TestLayerID_MarshalText(t *testing.T) {
	id := New()
	subId := LayerIDFromRefID(&id)

	res, err := subId.MarshalText()

	assert.Nil(t, err)
	assert.Equal(t, []byte(id.String()), res)
}

func TestLayerID_UnmarshalText(t *testing.T) {
	text := []byte("01f3zhcaq35403zdjnd6dcm0t2")

	subId := &LayerID{}

	err := subId.UnmarshalText(text)

	assert.Nil(t, err)
	assert.Equal(t, "01f3zhcaq35403zdjnd6dcm0t2", subId.String())

}

func TestLayerID_IsNil(t *testing.T) {
	subId := LayerID{}

	assert.True(t, subId.IsNil())

	id := New()
	subId = *LayerIDFromRefID(&id)

	assert.False(t, subId.IsNil())
}

func TestLayerIDToKeys(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    []LayerID
		expected []string
	}{
		{
			name:     "Empty slice",
			input:    make([]LayerID, 0),
			expected: make([]string, 0),
		},
		{
			name:     "1 element",
			input:    []LayerID{MustLayerID("01f3zhcaq35403zdjnd6dcm0t2")},
			expected: []string{"01f3zhcaq35403zdjnd6dcm0t2"},
		},
		{
			name: "multiple elements",
			input: []LayerID{
				MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustLayerID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustLayerID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: []string{
				"01f3zhcaq35403zdjnd6dcm0t1",
				"01f3zhcaq35403zdjnd6dcm0t2",
				"01f3zhcaq35403zdjnd6dcm0t3",
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			assert.Equal(tt, tc.expected, LayerIDToKeys(tc.input))
		})
	}

}

func TestLayerIDsFrom(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    []string
		expected struct {
			res []LayerID
			err error
		}
	}{
		{
			name:  "Empty slice",
			input: make([]string, 0),
			expected: struct {
				res []LayerID
				err error
			}{
				res: make([]LayerID, 0),
				err: nil,
			},
		},
		{
			name:  "1 element",
			input: []string{"01f3zhcaq35403zdjnd6dcm0t2"},
			expected: struct {
				res []LayerID
				err error
			}{
				res: []LayerID{MustLayerID("01f3zhcaq35403zdjnd6dcm0t2")},
				err: nil,
			},
		},
		{
			name: "multiple elements",
			input: []string{
				"01f3zhcaq35403zdjnd6dcm0t1",
				"01f3zhcaq35403zdjnd6dcm0t2",
				"01f3zhcaq35403zdjnd6dcm0t3",
			},
			expected: struct {
				res []LayerID
				err error
			}{
				res: []LayerID{
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
				err: nil,
			},
		},
		{
			name: "multiple elements",
			input: []string{
				"01f3zhcaq35403zdjnd6dcm0t1",
				"01f3zhcaq35403zdjnd6dcm0t2",
				"01f3zhcaq35403zdjnd6dcm0t3",
			},
			expected: struct {
				res []LayerID
				err error
			}{
				res: nil,
				err: ErrInvalidID,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			if tc.expected.err != nil {
				_, err := LayerIDsFrom(tc.input)
				assert.True(tt, errors.As(ErrInvalidID, &err))
			} else {
				res, err := LayerIDsFrom(tc.input)
				assert.Equal(tt, tc.expected.res, res)
				assert.Nil(tt, err)
			}

		})
	}
}

func TestLayerIDsFromID(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    []ID
		expected []LayerID
	}{
		{
			name:     "Empty slice",
			input:    make([]ID, 0),
			expected: make([]LayerID, 0),
		},
		{
			name:     "1 element",
			input:    []ID{MustBeID("01f3zhcaq35403zdjnd6dcm0t2")},
			expected: []LayerID{MustLayerID("01f3zhcaq35403zdjnd6dcm0t2")},
		},
		{
			name: "multiple elements",
			input: []ID{
				MustBeID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustBeID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustBeID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: []LayerID{
				MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustLayerID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustLayerID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			res := LayerIDsFromID(tc.input)
			assert.Equal(tt, tc.expected, res)
		})
	}
}

func TestLayerIDsFromIDRef(t *testing.T) {
	t.Parallel()

	id1 := MustBeID("01f3zhcaq35403zdjnd6dcm0t1")
	id2 := MustBeID("01f3zhcaq35403zdjnd6dcm0t2")
	id3 := MustBeID("01f3zhcaq35403zdjnd6dcm0t3")

	testCases := []struct {
		name     string
		input    []*ID
		expected []LayerID
	}{
		{
			name:     "Empty slice",
			input:    make([]*ID, 0),
			expected: make([]LayerID, 0),
		},
		{
			name:     "1 element",
			input:    []*ID{&id1},
			expected: []LayerID{MustLayerID(id1.String())},
		},
		{
			name:  "multiple elements",
			input: []*ID{&id1, &id2, &id3},
			expected: []LayerID{
				MustLayerID(id1.String()),
				MustLayerID(id2.String()),
				MustLayerID(id3.String()),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			res := LayerIDsFromIDRef(tc.input)
			assert.Equal(tt, tc.expected, res)
		})
	}
}

func TestLayerIDsToID(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []LayerID
		expected []ID
	}{
		{
			name:     "Empty slice",
			input:    make([]LayerID, 0),
			expected: make([]ID, 0),
		},
		{
			name:     "1 element",
			input:    []LayerID{MustLayerID("01f3zhcaq35403zdjnd6dcm0t2")},
			expected: []ID{MustBeID("01f3zhcaq35403zdjnd6dcm0t2")},
		},
		{
			name: "multiple elements",
			input: []LayerID{
				MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustLayerID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustLayerID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: []ID{
				MustBeID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustBeID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustBeID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			res := LayerIDsToID(tc.input)
			assert.Equal(tt, tc.expected, res)
		})
	}
}

func TestLayerIDsToIDRef(t *testing.T) {
	t.Parallel()

	id1 := MustBeID("01f3zhcaq35403zdjnd6dcm0t1")
	subId1 := MustLayerID(id1.String())
	id2 := MustBeID("01f3zhcaq35403zdjnd6dcm0t2")
	subId2 := MustLayerID(id2.String())
	id3 := MustBeID("01f3zhcaq35403zdjnd6dcm0t3")
	subId3 := MustLayerID(id3.String())

	testCases := []struct {
		name     string
		input    []*LayerID
		expected []*ID
	}{
		{
			name:     "Empty slice",
			input:    make([]*LayerID, 0),
			expected: make([]*ID, 0),
		},
		{
			name:     "1 element",
			input:    []*LayerID{&subId1},
			expected: []*ID{&id1},
		},
		{
			name:     "multiple elements",
			input:    []*LayerID{&subId1, &subId2, &subId3},
			expected: []*ID{&id1, &id2, &id3},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			res := LayerIDsToIDRef(tc.input)
			assert.Equal(tt, tc.expected, res)
		})
	}
}

func TestNewLayerIDSet(t *testing.T) {
	LayerIdSet := NewLayerIDSet()

	assert.NotNil(t, LayerIdSet)
	assert.Empty(t, LayerIdSet.m)
	assert.Empty(t, LayerIdSet.s)
}

func TestLayerIDSet_Add(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []LayerID
		expected *LayerIDSet
	}{
		{
			name:  "Empty slice",
			input: make([]LayerID, 0),
			expected: &LayerIDSet{
				m: map[LayerID]struct{}{},
				s: nil,
			},
		},
		{
			name:  "1 element",
			input: []LayerID{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1")},
			expected: &LayerIDSet{
				m: map[LayerID]struct{}{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []LayerID{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
		{
			name: "multiple elements",
			input: []LayerID{
				MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustLayerID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustLayerID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: &LayerIDSet{
				m: map[LayerID]struct{}{
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t2"): {},
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t3"): {},
				},
				s: []LayerID{
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
		},
		{
			name: "multiple elements with duplication",
			input: []LayerID{
				MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustLayerID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: &LayerIDSet{
				m: map[LayerID]struct{}{
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t3"): {},
				},
				s: []LayerID{
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			set := NewLayerIDSet()
			set.Add(tc.input...)
			assert.Equal(tt, tc.expected, set)
		})
	}
}

func TestLayerIDSet_AddRef(t *testing.T) {
	t.Parallel()

	LayerId := MustLayerID("01f3zhcaq35403zdjnd6dcm0t1")

	testCases := []struct {
		name     string
		input    *LayerID
		expected *LayerIDSet
	}{
		{
			name:  "Empty slice",
			input: nil,
			expected: &LayerIDSet{
				m: nil,
				s: nil,
			},
		},
		{
			name:  "1 element",
			input: &LayerId,
			expected: &LayerIDSet{
				m: map[LayerID]struct{}{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []LayerID{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			set := NewLayerIDSet()
			set.AddRef(tc.input)
			assert.Equal(tt, tc.expected, set)
		})
	}
}

func TestLayerIDSet_Has(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input struct {
			LayerIDSet
			LayerID
		}
		expected bool
	}{
		{
			name: "Empty Set",
			input: struct {
				LayerIDSet
				LayerID
			}{LayerIDSet: LayerIDSet{}, LayerID: MustLayerID("01f3zhcaq35403zdjnd6dcm0t1")},
			expected: false,
		},
		{
			name: "Set Contains the element",
			input: struct {
				LayerIDSet
				LayerID
			}{LayerIDSet: LayerIDSet{
				m: map[LayerID]struct{}{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []LayerID{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1")},
			}, LayerID: MustLayerID("01f3zhcaq35403zdjnd6dcm0t1")},
			expected: true,
		},
		{
			name: "Set does not Contains the element",
			input: struct {
				LayerIDSet
				LayerID
			}{LayerIDSet: LayerIDSet{
				m: map[LayerID]struct{}{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []LayerID{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1")},
			}, LayerID: MustLayerID("01f3zhcaq35403zdjnd6dcm0t2")},
			expected: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			assert.Equal(tt, tc.expected, tc.input.LayerIDSet.Has(tc.input.LayerID))
		})
	}
}

func TestLayerIDSet_Clear(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    LayerIDSet
		expected LayerIDSet
	}{
		{
			name:  "Empty Set",
			input: LayerIDSet{},
			expected: LayerIDSet{
				m: nil,
				s: nil,
			},
		},
		{
			name: "Set Contains the element",
			input: LayerIDSet{
				m: map[LayerID]struct{}{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []LayerID{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			expected: LayerIDSet{
				m: nil,
				s: nil,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			set := tc.input
			p := &set
			p.Clear()
			assert.Equal(tt, tc.expected, *p)
		})
	}
}

func TestLayerIDSet_All(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *LayerIDSet
		expected []LayerID
	}{
		{
			name: "Empty slice",
			input: &LayerIDSet{
				m: map[LayerID]struct{}{},
				s: nil,
			},
			expected: make([]LayerID, 0),
		},
		{
			name: "1 element",
			input: &LayerIDSet{
				m: map[LayerID]struct{}{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []LayerID{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			expected: []LayerID{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1")},
		},
		{
			name: "multiple elements",
			input: &LayerIDSet{
				m: map[LayerID]struct{}{
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t2"): {},
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t3"): {},
				},
				s: []LayerID{
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
			expected: []LayerID{
				MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustLayerID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustLayerID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			assert.Equal(tt, tc.expected, tc.input.All())
		})
	}
}

func TestLayerIDSet_Clone(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *LayerIDSet
		expected *LayerIDSet
	}{
		{
			name:     "nil set",
			input:    nil,
			expected: NewLayerIDSet(),
		},
		{
			name:     "Empty set",
			input:    NewLayerIDSet(),
			expected: NewLayerIDSet(),
		},
		{
			name: "1 element",
			input: &LayerIDSet{
				m: map[LayerID]struct{}{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []LayerID{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			expected: &LayerIDSet{
				m: map[LayerID]struct{}{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []LayerID{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
		{
			name: "multiple elements",
			input: &LayerIDSet{
				m: map[LayerID]struct{}{
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t2"): {},
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t3"): {},
				},
				s: []LayerID{
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
			expected: &LayerIDSet{
				m: map[LayerID]struct{}{
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t2"): {},
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t3"): {},
				},
				s: []LayerID{
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			clone := tc.input.Clone()
			assert.Equal(tt, tc.expected, clone)
			assert.False(tt, tc.input == clone)
		})
	}
}

func TestLayerIDSet_Merge(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input struct {
			a *LayerIDSet
			b *LayerIDSet
		}
		expected *LayerIDSet
	}{
		{
			name: "Empty Set",
			input: struct {
				a *LayerIDSet
				b *LayerIDSet
			}{
				a: &LayerIDSet{},
				b: &LayerIDSet{},
			},
			expected: &LayerIDSet{},
		},
		{
			name: "1 Empty Set",
			input: struct {
				a *LayerIDSet
				b *LayerIDSet
			}{
				a: &LayerIDSet{
					m: map[LayerID]struct{}{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
					s: []LayerID{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1")},
				},
				b: &LayerIDSet{},
			},
			expected: &LayerIDSet{
				m: map[LayerID]struct{}{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []LayerID{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
		{
			name: "2 non Empty Set",
			input: struct {
				a *LayerIDSet
				b *LayerIDSet
			}{
				a: &LayerIDSet{
					m: map[LayerID]struct{}{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
					s: []LayerID{MustLayerID("01f3zhcaq35403zdjnd6dcm0t1")},
				},
				b: &LayerIDSet{
					m: map[LayerID]struct{}{MustLayerID("01f3zhcaq35403zdjnd6dcm0t2"): {}},
					s: []LayerID{MustLayerID("01f3zhcaq35403zdjnd6dcm0t2")},
				},
			},
			expected: &LayerIDSet{
				m: map[LayerID]struct{}{
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t2"): {},
				},
				s: []LayerID{
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustLayerID("01f3zhcaq35403zdjnd6dcm0t2"),
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			assert.Equal(tt, tc.expected, tc.input.a.Merge(tc.input.b))
		})
	}
}
