// Code generated by gen, DO NOT EDIT.

package id

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/oklog/ulid"
	"github.com/stretchr/testify/assert"
)

func TestNewDatasetSchemaID(t *testing.T) {
	id := NewDatasetSchemaID()
	assert.NotNil(t, id)
	ulID, err := ulid.Parse(id.String())

	assert.NotNil(t, ulID)
	assert.Nil(t, err)
}

func TestDatasetSchemaIDFrom(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    string
		expected struct {
			result DatasetSchemaID
			err    error
		}
	}{
		{
			name:  "Fail:Not valid string",
			input: "testMustFail",
			expected: struct {
				result DatasetSchemaID
				err    error
			}{
				DatasetSchemaID{},
				ErrInvalidID,
			},
		},
		{
			name:  "Fail:Not valid string",
			input: "",
			expected: struct {
				result DatasetSchemaID
				err    error
			}{
				DatasetSchemaID{},
				ErrInvalidID,
			},
		},
		{
			name:  "success:valid string",
			input: "01f2r7kg1fvvffp0gmexgy5hxy",
			expected: struct {
				result DatasetSchemaID
				err    error
			}{
				DatasetSchemaID{ulid.MustParse("01f2r7kg1fvvffp0gmexgy5hxy")},
				nil,
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			result, err := DatasetSchemaIDFrom(tc.input)
			assert.Equal(tt, tc.expected.result, result)
			if err != nil {
				assert.True(tt, errors.As(tc.expected.err, &err))
			}
		})
	}
}

func TestMustDatasetSchemaID(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name        string
		input       string
		shouldPanic bool
		expected    DatasetSchemaID
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
			expected:    DatasetSchemaID{ulid.MustParse("01f2r7kg1fvvffp0gmexgy5hxy")},
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
			result := MustDatasetSchemaID(tc.input)
			assert.Equal(tt, tc.expected, result)
		})
	}
}

func TestDatasetSchemaIDFromRef(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected *DatasetSchemaID
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
			expected: &DatasetSchemaID{ulid.MustParse("01f2r7kg1fvvffp0gmexgy5hxy")},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			result := DatasetSchemaIDFromRef(&tc.input)
			assert.Equal(tt, tc.expected, result)
			if tc.expected != nil {
				assert.Equal(tt, *tc.expected, *result)
			}
		})
	}
}

func TestDatasetSchemaIDFromRefID(t *testing.T) {
	id := New()

	subId := DatasetSchemaIDFromRefID(&id)

	assert.NotNil(t, subId)
	assert.Equal(t, subId.id, id.id)
}

func TestDatasetSchemaID_ID(t *testing.T) {
	id := New()
	subId := DatasetSchemaIDFromRefID(&id)

	idOrg := subId.ID()

	assert.Equal(t, id, idOrg)
}

func TestDatasetSchemaID_String(t *testing.T) {
	id := New()
	subId := DatasetSchemaIDFromRefID(&id)

	assert.Equal(t, subId.String(), id.String())
}

func TestDatasetSchemaID_GoString(t *testing.T) {
	id := New()
	subId := DatasetSchemaIDFromRefID(&id)

	assert.Equal(t, subId.GoString(), "id.DatasetSchemaID("+id.String()+")")
}

func TestDatasetSchemaID_RefString(t *testing.T) {
	id := New()
	subId := DatasetSchemaIDFromRefID(&id)

	refString := subId.StringRef()

	assert.NotNil(t, refString)
	assert.Equal(t, *refString, id.String())
}

func TestDatasetSchemaID_Ref(t *testing.T) {
	id := New()
	subId := DatasetSchemaIDFromRefID(&id)

	subIdRef := subId.Ref()

	assert.Equal(t, *subId, *subIdRef)
}

func TestDatasetSchemaID_Contains(t *testing.T) {
	id := NewDatasetSchemaID()
	id2 := NewDatasetSchemaID()
	assert.True(t, id.Contains([]DatasetSchemaID{id, id2}))
	assert.False(t, id.Contains([]DatasetSchemaID{id2}))
}

func TestDatasetSchemaID_CopyRef(t *testing.T) {
	id := New()
	subId := DatasetSchemaIDFromRefID(&id)

	subIdCopyRef := subId.CopyRef()

	assert.Equal(t, *subId, *subIdCopyRef)
	assert.NotSame(t, subId, subIdCopyRef)
}

func TestDatasetSchemaID_IDRef(t *testing.T) {
	id := New()
	subId := DatasetSchemaIDFromRefID(&id)

	assert.Equal(t, id, *subId.IDRef())
}

func TestDatasetSchemaID_StringRef(t *testing.T) {
	id := New()
	subId := DatasetSchemaIDFromRefID(&id)

	assert.Equal(t, *subId.StringRef(), id.String())
}

func TestDatasetSchemaID_MarhsalJSON(t *testing.T) {
	id := New()
	subId := DatasetSchemaIDFromRefID(&id)

	res, err := subId.MarhsalJSON()
	exp, _ := json.Marshal(subId.String())

	assert.Nil(t, err)
	assert.Equal(t, exp, res)
}

func TestDatasetSchemaID_UnmarhsalJSON(t *testing.T) {
	jsonString := "\"01f3zhkysvcxsnzepyyqtq21fb\""

	subId := &DatasetSchemaID{}

	err := subId.UnmarhsalJSON([]byte(jsonString))

	assert.Nil(t, err)
	assert.Equal(t, "01f3zhkysvcxsnzepyyqtq21fb", subId.String())
}

func TestDatasetSchemaID_MarshalText(t *testing.T) {
	id := New()
	subId := DatasetSchemaIDFromRefID(&id)

	res, err := subId.MarshalText()

	assert.Nil(t, err)
	assert.Equal(t, []byte(id.String()), res)
}

func TestDatasetSchemaID_UnmarshalText(t *testing.T) {
	text := []byte("01f3zhcaq35403zdjnd6dcm0t2")

	subId := &DatasetSchemaID{}

	err := subId.UnmarshalText(text)

	assert.Nil(t, err)
	assert.Equal(t, "01f3zhcaq35403zdjnd6dcm0t2", subId.String())

}

func TestDatasetSchemaID_IsNil(t *testing.T) {
	subId := DatasetSchemaID{}

	assert.True(t, subId.IsNil())

	id := New()
	subId = *DatasetSchemaIDFromRefID(&id)

	assert.False(t, subId.IsNil())
}

func TestDatasetSchemaIDToKeys(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    []DatasetSchemaID
		expected []string
	}{
		{
			name:     "Empty slice",
			input:    make([]DatasetSchemaID, 0),
			expected: make([]string, 0),
		},
		{
			name:     "1 element",
			input:    []DatasetSchemaID{MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t2")},
			expected: []string{"01f3zhcaq35403zdjnd6dcm0t2"},
		},
		{
			name: "multiple elements",
			input: []DatasetSchemaID{
				MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t3"),
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
			assert.Equal(tt, tc.expected, DatasetSchemaIDToKeys(tc.input))
		})
	}

}

func TestDatasetSchemaIDsFrom(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    []string
		expected struct {
			res []DatasetSchemaID
			err error
		}
	}{
		{
			name:  "Empty slice",
			input: make([]string, 0),
			expected: struct {
				res []DatasetSchemaID
				err error
			}{
				res: make([]DatasetSchemaID, 0),
				err: nil,
			},
		},
		{
			name:  "1 element",
			input: []string{"01f3zhcaq35403zdjnd6dcm0t2"},
			expected: struct {
				res []DatasetSchemaID
				err error
			}{
				res: []DatasetSchemaID{MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t2")},
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
				res []DatasetSchemaID
				err error
			}{
				res: []DatasetSchemaID{
					MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t3"),
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
				res []DatasetSchemaID
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
				_, err := DatasetSchemaIDsFrom(tc.input)
				assert.True(tt, errors.As(ErrInvalidID, &err))
			} else {
				res, err := DatasetSchemaIDsFrom(tc.input)
				assert.Equal(tt, tc.expected.res, res)
				assert.Nil(tt, err)
			}

		})
	}
}

func TestDatasetSchemaIDsFromID(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    []ID
		expected []DatasetSchemaID
	}{
		{
			name:     "Empty slice",
			input:    make([]ID, 0),
			expected: make([]DatasetSchemaID, 0),
		},
		{
			name:     "1 element",
			input:    []ID{MustBeID("01f3zhcaq35403zdjnd6dcm0t2")},
			expected: []DatasetSchemaID{MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t2")},
		},
		{
			name: "multiple elements",
			input: []ID{
				MustBeID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustBeID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustBeID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: []DatasetSchemaID{
				MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			res := DatasetSchemaIDsFromID(tc.input)
			assert.Equal(tt, tc.expected, res)
		})
	}
}

func TestDatasetSchemaIDsFromIDRef(t *testing.T) {
	t.Parallel()

	id1 := MustBeID("01f3zhcaq35403zdjnd6dcm0t1")
	id2 := MustBeID("01f3zhcaq35403zdjnd6dcm0t2")
	id3 := MustBeID("01f3zhcaq35403zdjnd6dcm0t3")

	testCases := []struct {
		name     string
		input    []*ID
		expected []DatasetSchemaID
	}{
		{
			name:     "Empty slice",
			input:    make([]*ID, 0),
			expected: make([]DatasetSchemaID, 0),
		},
		{
			name:     "1 element",
			input:    []*ID{&id1},
			expected: []DatasetSchemaID{MustDatasetSchemaID(id1.String())},
		},
		{
			name:  "multiple elements",
			input: []*ID{&id1, &id2, &id3},
			expected: []DatasetSchemaID{
				MustDatasetSchemaID(id1.String()),
				MustDatasetSchemaID(id2.String()),
				MustDatasetSchemaID(id3.String()),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			res := DatasetSchemaIDsFromIDRef(tc.input)
			assert.Equal(tt, tc.expected, res)
		})
	}
}

func TestDatasetSchemaIDsToID(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []DatasetSchemaID
		expected []ID
	}{
		{
			name:     "Empty slice",
			input:    make([]DatasetSchemaID, 0),
			expected: make([]ID, 0),
		},
		{
			name:     "1 element",
			input:    []DatasetSchemaID{MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t2")},
			expected: []ID{MustBeID("01f3zhcaq35403zdjnd6dcm0t2")},
		},
		{
			name: "multiple elements",
			input: []DatasetSchemaID{
				MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t3"),
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

			res := DatasetSchemaIDsToID(tc.input)
			assert.Equal(tt, tc.expected, res)
		})
	}
}

func TestDatasetSchemaIDsToIDRef(t *testing.T) {
	t.Parallel()

	id1 := MustBeID("01f3zhcaq35403zdjnd6dcm0t1")
	subId1 := MustDatasetSchemaID(id1.String())
	id2 := MustBeID("01f3zhcaq35403zdjnd6dcm0t2")
	subId2 := MustDatasetSchemaID(id2.String())
	id3 := MustBeID("01f3zhcaq35403zdjnd6dcm0t3")
	subId3 := MustDatasetSchemaID(id3.String())

	testCases := []struct {
		name     string
		input    []*DatasetSchemaID
		expected []*ID
	}{
		{
			name:     "Empty slice",
			input:    make([]*DatasetSchemaID, 0),
			expected: make([]*ID, 0),
		},
		{
			name:     "1 element",
			input:    []*DatasetSchemaID{&subId1},
			expected: []*ID{&id1},
		},
		{
			name:     "multiple elements",
			input:    []*DatasetSchemaID{&subId1, &subId2, &subId3},
			expected: []*ID{&id1, &id2, &id3},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			res := DatasetSchemaIDsToIDRef(tc.input)
			assert.Equal(tt, tc.expected, res)
		})
	}
}

func TestNewDatasetSchemaIDSet(t *testing.T) {
	DatasetSchemaIdSet := NewDatasetSchemaIDSet()

	assert.NotNil(t, DatasetSchemaIdSet)
	assert.Empty(t, DatasetSchemaIdSet.m)
	assert.Empty(t, DatasetSchemaIdSet.s)
}

func TestDatasetSchemaIDSet_Add(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []DatasetSchemaID
		expected *DatasetSchemaIDSet
	}{
		{
			name:  "Empty slice",
			input: make([]DatasetSchemaID, 0),
			expected: &DatasetSchemaIDSet{
				m: map[DatasetSchemaID]struct{}{},
				s: nil,
			},
		},
		{
			name:  "1 element",
			input: []DatasetSchemaID{MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1")},
			expected: &DatasetSchemaIDSet{
				m: map[DatasetSchemaID]struct{}{MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []DatasetSchemaID{MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
		{
			name: "multiple elements",
			input: []DatasetSchemaID{
				MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: &DatasetSchemaIDSet{
				m: map[DatasetSchemaID]struct{}{
					MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t2"): {},
					MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t3"): {},
				},
				s: []DatasetSchemaID{
					MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
		},
		{
			name: "multiple elements with duplication",
			input: []DatasetSchemaID{
				MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: &DatasetSchemaIDSet{
				m: map[DatasetSchemaID]struct{}{
					MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t3"): {},
				},
				s: []DatasetSchemaID{
					MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			set := NewDatasetSchemaIDSet()
			set.Add(tc.input...)
			assert.Equal(tt, tc.expected, set)
		})
	}
}

func TestDatasetSchemaIDSet_AddRef(t *testing.T) {
	t.Parallel()

	DatasetSchemaId := MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1")

	testCases := []struct {
		name     string
		input    *DatasetSchemaID
		expected *DatasetSchemaIDSet
	}{
		{
			name:  "Empty slice",
			input: nil,
			expected: &DatasetSchemaIDSet{
				m: nil,
				s: nil,
			},
		},
		{
			name:  "1 element",
			input: &DatasetSchemaId,
			expected: &DatasetSchemaIDSet{
				m: map[DatasetSchemaID]struct{}{MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []DatasetSchemaID{MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			set := NewDatasetSchemaIDSet()
			set.AddRef(tc.input)
			assert.Equal(tt, tc.expected, set)
		})
	}
}

func TestDatasetSchemaIDSet_Has(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input struct {
			DatasetSchemaIDSet
			DatasetSchemaID
		}
		expected bool
	}{
		{
			name: "Empty Set",
			input: struct {
				DatasetSchemaIDSet
				DatasetSchemaID
			}{DatasetSchemaIDSet: DatasetSchemaIDSet{}, DatasetSchemaID: MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1")},
			expected: false,
		},
		{
			name: "Set Contains the element",
			input: struct {
				DatasetSchemaIDSet
				DatasetSchemaID
			}{DatasetSchemaIDSet: DatasetSchemaIDSet{
				m: map[DatasetSchemaID]struct{}{MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []DatasetSchemaID{MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1")},
			}, DatasetSchemaID: MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1")},
			expected: true,
		},
		{
			name: "Set does not Contains the element",
			input: struct {
				DatasetSchemaIDSet
				DatasetSchemaID
			}{DatasetSchemaIDSet: DatasetSchemaIDSet{
				m: map[DatasetSchemaID]struct{}{MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []DatasetSchemaID{MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1")},
			}, DatasetSchemaID: MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t2")},
			expected: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			assert.Equal(tt, tc.expected, tc.input.DatasetSchemaIDSet.Has(tc.input.DatasetSchemaID))
		})
	}
}

func TestDatasetSchemaIDSet_Clear(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    DatasetSchemaIDSet
		expected DatasetSchemaIDSet
	}{
		{
			name:  "Empty Set",
			input: DatasetSchemaIDSet{},
			expected: DatasetSchemaIDSet{
				m: nil,
				s: nil,
			},
		},
		{
			name: "Set Contains the element",
			input: DatasetSchemaIDSet{
				m: map[DatasetSchemaID]struct{}{MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []DatasetSchemaID{MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			expected: DatasetSchemaIDSet{
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

func TestDatasetSchemaIDSet_All(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *DatasetSchemaIDSet
		expected []DatasetSchemaID
	}{
		{
			name: "Empty slice",
			input: &DatasetSchemaIDSet{
				m: map[DatasetSchemaID]struct{}{},
				s: nil,
			},
			expected: make([]DatasetSchemaID, 0),
		},
		{
			name: "1 element",
			input: &DatasetSchemaIDSet{
				m: map[DatasetSchemaID]struct{}{MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []DatasetSchemaID{MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			expected: []DatasetSchemaID{MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1")},
		},
		{
			name: "multiple elements",
			input: &DatasetSchemaIDSet{
				m: map[DatasetSchemaID]struct{}{
					MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t2"): {},
					MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t3"): {},
				},
				s: []DatasetSchemaID{
					MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
			expected: []DatasetSchemaID{
				MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t3"),
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

func TestDatasetSchemaIDSet_Clone(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *DatasetSchemaIDSet
		expected *DatasetSchemaIDSet
	}{
		{
			name:     "nil set",
			input:    nil,
			expected: NewDatasetSchemaIDSet(),
		},
		{
			name:     "Empty set",
			input:    NewDatasetSchemaIDSet(),
			expected: NewDatasetSchemaIDSet(),
		},
		{
			name: "1 element",
			input: &DatasetSchemaIDSet{
				m: map[DatasetSchemaID]struct{}{MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []DatasetSchemaID{MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			expected: &DatasetSchemaIDSet{
				m: map[DatasetSchemaID]struct{}{MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []DatasetSchemaID{MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
		{
			name: "multiple elements",
			input: &DatasetSchemaIDSet{
				m: map[DatasetSchemaID]struct{}{
					MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t2"): {},
					MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t3"): {},
				},
				s: []DatasetSchemaID{
					MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
			expected: &DatasetSchemaIDSet{
				m: map[DatasetSchemaID]struct{}{
					MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t2"): {},
					MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t3"): {},
				},
				s: []DatasetSchemaID{
					MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t3"),
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

func TestDatasetSchemaIDSet_Merge(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input struct {
			a *DatasetSchemaIDSet
			b *DatasetSchemaIDSet
		}
		expected *DatasetSchemaIDSet
	}{
		{
			name: "Empty Set",
			input: struct {
				a *DatasetSchemaIDSet
				b *DatasetSchemaIDSet
			}{
				a: &DatasetSchemaIDSet{},
				b: &DatasetSchemaIDSet{},
			},
			expected: &DatasetSchemaIDSet{},
		},
		{
			name: "1 Empty Set",
			input: struct {
				a *DatasetSchemaIDSet
				b *DatasetSchemaIDSet
			}{
				a: &DatasetSchemaIDSet{
					m: map[DatasetSchemaID]struct{}{MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
					s: []DatasetSchemaID{MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1")},
				},
				b: &DatasetSchemaIDSet{},
			},
			expected: &DatasetSchemaIDSet{
				m: map[DatasetSchemaID]struct{}{MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []DatasetSchemaID{MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
		{
			name: "2 non Empty Set",
			input: struct {
				a *DatasetSchemaIDSet
				b *DatasetSchemaIDSet
			}{
				a: &DatasetSchemaIDSet{
					m: map[DatasetSchemaID]struct{}{MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
					s: []DatasetSchemaID{MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1")},
				},
				b: &DatasetSchemaIDSet{
					m: map[DatasetSchemaID]struct{}{MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t2"): {}},
					s: []DatasetSchemaID{MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t2")},
				},
			},
			expected: &DatasetSchemaIDSet{
				m: map[DatasetSchemaID]struct{}{
					MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t2"): {},
				},
				s: []DatasetSchemaID{
					MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustDatasetSchemaID("01f3zhcaq35403zdjnd6dcm0t2"),
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
