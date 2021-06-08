// Code generated by gen, DO NOT EDIT.

package id

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/oklog/ulid"
	"github.com/stretchr/testify/assert"
)

func TestNewDatasetSchemaFieldID(t *testing.T) {
	id := NewDatasetSchemaFieldID()
	assert.NotNil(t, id)
	ulID, err := ulid.Parse(id.String())

	assert.NotNil(t, ulID)
	assert.Nil(t, err)
}

func TestDatasetSchemaFieldIDFrom(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    string
		expected struct {
			result DatasetSchemaFieldID
			err    error
		}
	}{
		{
			name:  "Fail:Not valid string",
			input: "testMustFail",
			expected: struct {
				result DatasetSchemaFieldID
				err    error
			}{
				DatasetSchemaFieldID{},
				ErrInvalidID,
			},
		},
		{
			name:  "Fail:Not valid string",
			input: "",
			expected: struct {
				result DatasetSchemaFieldID
				err    error
			}{
				DatasetSchemaFieldID{},
				ErrInvalidID,
			},
		},
		{
			name:  "success:valid string",
			input: "01f2r7kg1fvvffp0gmexgy5hxy",
			expected: struct {
				result DatasetSchemaFieldID
				err    error
			}{
				DatasetSchemaFieldID{ulid.MustParse("01f2r7kg1fvvffp0gmexgy5hxy")},
				nil,
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			result, err := DatasetSchemaFieldIDFrom(tc.input)
			assert.Equal(tt, tc.expected.result, result)
			if err != nil {
				assert.True(tt, errors.As(tc.expected.err, &err))
			}
		})
	}
}

func TestMustDatasetSchemaFieldID(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name        string
		input       string
		shouldPanic bool
		expected    DatasetSchemaFieldID
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
			expected:    DatasetSchemaFieldID{ulid.MustParse("01f2r7kg1fvvffp0gmexgy5hxy")},
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
			result := MustDatasetSchemaFieldID(tc.input)
			assert.Equal(tt, tc.expected, result)
		})
	}
}

func TestDatasetSchemaFieldIDFromRef(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected *DatasetSchemaFieldID
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
			expected: &DatasetSchemaFieldID{ulid.MustParse("01f2r7kg1fvvffp0gmexgy5hxy")},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			result := DatasetSchemaFieldIDFromRef(&tc.input)
			assert.Equal(tt, tc.expected, result)
			if tc.expected != nil {
				assert.Equal(tt, *tc.expected, *result)
			}
		})
	}
}

func TestDatasetSchemaFieldIDFromRefID(t *testing.T) {
	id := New()

	subId := DatasetSchemaFieldIDFromRefID(&id)

	assert.NotNil(t, subId)
	assert.Equal(t, subId.id, id.id)
}

func TestDatasetSchemaFieldID_ID(t *testing.T) {
	id := New()
	subId := DatasetSchemaFieldIDFromRefID(&id)

	idOrg := subId.ID()

	assert.Equal(t, id, idOrg)
}

func TestDatasetSchemaFieldID_String(t *testing.T) {
	id := New()
	subId := DatasetSchemaFieldIDFromRefID(&id)

	assert.Equal(t, subId.String(), id.String())
}

func TestDatasetSchemaFieldID_GoString(t *testing.T) {
	id := New()
	subId := DatasetSchemaFieldIDFromRefID(&id)

	assert.Equal(t, subId.GoString(), "id.DatasetSchemaFieldID("+id.String()+")")
}

func TestDatasetSchemaFieldID_RefString(t *testing.T) {
	id := New()
	subId := DatasetSchemaFieldIDFromRefID(&id)

	refString := subId.StringRef()

	assert.NotNil(t, refString)
	assert.Equal(t, *refString, id.String())
}

func TestDatasetSchemaFieldID_Ref(t *testing.T) {
	id := New()
	subId := DatasetSchemaFieldIDFromRefID(&id)

	subIdRef := subId.Ref()

	assert.Equal(t, *subId, *subIdRef)
}

func TestDatasetSchemaFieldID_CopyRef(t *testing.T) {
	id := New()
	subId := DatasetSchemaFieldIDFromRefID(&id)

	subIdCopyRef := subId.CopyRef()

	assert.Equal(t, *subId, *subIdCopyRef)
	assert.NotSame(t, subId, subIdCopyRef)
}

func TestDatasetSchemaFieldID_IDRef(t *testing.T) {
	id := New()
	subId := DatasetSchemaFieldIDFromRefID(&id)

	assert.Equal(t, id, *subId.IDRef())
}

func TestDatasetSchemaFieldID_StringRef(t *testing.T) {
	id := New()
	subId := DatasetSchemaFieldIDFromRefID(&id)

	assert.Equal(t, *subId.StringRef(), id.String())
}

func TestDatasetSchemaFieldID_MarhsalJSON(t *testing.T) {
	id := New()
	subId := DatasetSchemaFieldIDFromRefID(&id)

	res, err := subId.MarhsalJSON()
	exp, _ := json.Marshal(subId.String())

	assert.Nil(t, err)
	assert.Equal(t, exp, res)
}

func TestDatasetSchemaFieldID_UnmarhsalJSON(t *testing.T) {
	jsonString := "\"01f3zhkysvcxsnzepyyqtq21fb\""

	subId := &DatasetSchemaFieldID{}

	err := subId.UnmarhsalJSON([]byte(jsonString))

	assert.Nil(t, err)
	assert.Equal(t, "01f3zhkysvcxsnzepyyqtq21fb", subId.String())
}

func TestDatasetSchemaFieldID_MarshalText(t *testing.T) {
	id := New()
	subId := DatasetSchemaFieldIDFromRefID(&id)

	res, err := subId.MarshalText()

	assert.Nil(t, err)
	assert.Equal(t, []byte(id.String()), res)
}

func TestDatasetSchemaFieldID_UnmarshalText(t *testing.T) {
	text := []byte("01f3zhcaq35403zdjnd6dcm0t2")

	subId := &DatasetSchemaFieldID{}

	err := subId.UnmarshalText(text)

	assert.Nil(t, err)
	assert.Equal(t, "01f3zhcaq35403zdjnd6dcm0t2", subId.String())

}

func TestDatasetSchemaFieldID_IsNil(t *testing.T) {
	subId := DatasetSchemaFieldID{}

	assert.True(t, subId.IsNil())

	id := New()
	subId = *DatasetSchemaFieldIDFromRefID(&id)

	assert.False(t, subId.IsNil())
}

func TestDatasetSchemaFieldIDToKeys(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    []DatasetSchemaFieldID
		expected []string
	}{
		{
			name:     "Empty slice",
			input:    make([]DatasetSchemaFieldID, 0),
			expected: make([]string, 0),
		},
		{
			name:     "1 element",
			input:    []DatasetSchemaFieldID{MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t2")},
			expected: []string{"01f3zhcaq35403zdjnd6dcm0t2"},
		},
		{
			name: "multiple elements",
			input: []DatasetSchemaFieldID{
				MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t3"),
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
			assert.Equal(tt, tc.expected, DatasetSchemaFieldIDToKeys(tc.input))
		})
	}

}

func TestDatasetSchemaFieldIDsFrom(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    []string
		expected struct {
			res []DatasetSchemaFieldID
			err error
		}
	}{
		{
			name:  "Empty slice",
			input: make([]string, 0),
			expected: struct {
				res []DatasetSchemaFieldID
				err error
			}{
				res: make([]DatasetSchemaFieldID, 0),
				err: nil,
			},
		},
		{
			name:  "1 element",
			input: []string{"01f3zhcaq35403zdjnd6dcm0t2"},
			expected: struct {
				res []DatasetSchemaFieldID
				err error
			}{
				res: []DatasetSchemaFieldID{MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t2")},
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
				res []DatasetSchemaFieldID
				err error
			}{
				res: []DatasetSchemaFieldID{
					MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t3"),
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
				res []DatasetSchemaFieldID
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
				_, err := DatasetSchemaFieldIDsFrom(tc.input)
				assert.True(tt, errors.As(ErrInvalidID, &err))
			} else {
				res, err := DatasetSchemaFieldIDsFrom(tc.input)
				assert.Equal(tt, tc.expected.res, res)
				assert.Nil(tt, err)
			}

		})
	}
}

func TestDatasetSchemaFieldIDsFromID(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    []ID
		expected []DatasetSchemaFieldID
	}{
		{
			name:     "Empty slice",
			input:    make([]ID, 0),
			expected: make([]DatasetSchemaFieldID, 0),
		},
		{
			name:     "1 element",
			input:    []ID{MustBeID("01f3zhcaq35403zdjnd6dcm0t2")},
			expected: []DatasetSchemaFieldID{MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t2")},
		},
		{
			name: "multiple elements",
			input: []ID{
				MustBeID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustBeID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustBeID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: []DatasetSchemaFieldID{
				MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			res := DatasetSchemaFieldIDsFromID(tc.input)
			assert.Equal(tt, tc.expected, res)
		})
	}
}

func TestDatasetSchemaFieldIDsFromIDRef(t *testing.T) {
	t.Parallel()

	id1 := MustBeID("01f3zhcaq35403zdjnd6dcm0t1")
	id2 := MustBeID("01f3zhcaq35403zdjnd6dcm0t2")
	id3 := MustBeID("01f3zhcaq35403zdjnd6dcm0t3")

	testCases := []struct {
		name     string
		input    []*ID
		expected []DatasetSchemaFieldID
	}{
		{
			name:     "Empty slice",
			input:    make([]*ID, 0),
			expected: make([]DatasetSchemaFieldID, 0),
		},
		{
			name:     "1 element",
			input:    []*ID{&id1},
			expected: []DatasetSchemaFieldID{MustDatasetSchemaFieldID(id1.String())},
		},
		{
			name:  "multiple elements",
			input: []*ID{&id1, &id2, &id3},
			expected: []DatasetSchemaFieldID{
				MustDatasetSchemaFieldID(id1.String()),
				MustDatasetSchemaFieldID(id2.String()),
				MustDatasetSchemaFieldID(id3.String()),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			res := DatasetSchemaFieldIDsFromIDRef(tc.input)
			assert.Equal(tt, tc.expected, res)
		})
	}
}

func TestDatasetSchemaFieldIDsToID(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []DatasetSchemaFieldID
		expected []ID
	}{
		{
			name:     "Empty slice",
			input:    make([]DatasetSchemaFieldID, 0),
			expected: make([]ID, 0),
		},
		{
			name:     "1 element",
			input:    []DatasetSchemaFieldID{MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t2")},
			expected: []ID{MustBeID("01f3zhcaq35403zdjnd6dcm0t2")},
		},
		{
			name: "multiple elements",
			input: []DatasetSchemaFieldID{
				MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t3"),
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

			res := DatasetSchemaFieldIDsToID(tc.input)
			assert.Equal(tt, tc.expected, res)
		})
	}
}

func TestDatasetSchemaFieldIDsToIDRef(t *testing.T) {
	t.Parallel()

	id1 := MustBeID("01f3zhcaq35403zdjnd6dcm0t1")
	subId1 := MustDatasetSchemaFieldID(id1.String())
	id2 := MustBeID("01f3zhcaq35403zdjnd6dcm0t2")
	subId2 := MustDatasetSchemaFieldID(id2.String())
	id3 := MustBeID("01f3zhcaq35403zdjnd6dcm0t3")
	subId3 := MustDatasetSchemaFieldID(id3.String())

	testCases := []struct {
		name     string
		input    []*DatasetSchemaFieldID
		expected []*ID
	}{
		{
			name:     "Empty slice",
			input:    make([]*DatasetSchemaFieldID, 0),
			expected: make([]*ID, 0),
		},
		{
			name:     "1 element",
			input:    []*DatasetSchemaFieldID{&subId1},
			expected: []*ID{&id1},
		},
		{
			name:     "multiple elements",
			input:    []*DatasetSchemaFieldID{&subId1, &subId2, &subId3},
			expected: []*ID{&id1, &id2, &id3},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			res := DatasetSchemaFieldIDsToIDRef(tc.input)
			assert.Equal(tt, tc.expected, res)
		})
	}
}

func TestNewDatasetSchemaFieldIDSet(t *testing.T) {
	DatasetSchemaFieldIdSet := NewDatasetSchemaFieldIDSet()

	assert.NotNil(t, DatasetSchemaFieldIdSet)
	assert.Empty(t, DatasetSchemaFieldIdSet.m)
	assert.Empty(t, DatasetSchemaFieldIdSet.s)
}

func TestDatasetSchemaFieldIDSet_Add(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []DatasetSchemaFieldID
		expected *DatasetSchemaFieldIDSet
	}{
		{
			name:  "Empty slice",
			input: make([]DatasetSchemaFieldID, 0),
			expected: &DatasetSchemaFieldIDSet{
				m: map[DatasetSchemaFieldID]struct{}{},
				s: nil,
			},
		},
		{
			name:  "1 element",
			input: []DatasetSchemaFieldID{MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1")},
			expected: &DatasetSchemaFieldIDSet{
				m: map[DatasetSchemaFieldID]struct{}{MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{}},
				s: []DatasetSchemaFieldID{MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
		{
			name: "multiple elements",
			input: []DatasetSchemaFieldID{
				MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: &DatasetSchemaFieldIDSet{
				m: map[DatasetSchemaFieldID]struct{}{
					MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{},
					MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t2"): struct{}{},
					MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t3"): struct{}{},
				},
				s: []DatasetSchemaFieldID{
					MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
		},
		{
			name: "multiple elements with duplication",
			input: []DatasetSchemaFieldID{
				MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: &DatasetSchemaFieldIDSet{
				m: map[DatasetSchemaFieldID]struct{}{
					MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{},
					MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t3"): struct{}{},
				},
				s: []DatasetSchemaFieldID{
					MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			set := NewDatasetSchemaFieldIDSet()
			set.Add(tc.input...)
			assert.Equal(tt, tc.expected, set)
		})
	}
}

func TestDatasetSchemaFieldIDSet_AddRef(t *testing.T) {
	t.Parallel()

	DatasetSchemaFieldId := MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1")

	testCases := []struct {
		name     string
		input    *DatasetSchemaFieldID
		expected *DatasetSchemaFieldIDSet
	}{
		{
			name:  "Empty slice",
			input: nil,
			expected: &DatasetSchemaFieldIDSet{
				m: nil,
				s: nil,
			},
		},
		{
			name:  "1 element",
			input: &DatasetSchemaFieldId,
			expected: &DatasetSchemaFieldIDSet{
				m: map[DatasetSchemaFieldID]struct{}{MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{}},
				s: []DatasetSchemaFieldID{MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			set := NewDatasetSchemaFieldIDSet()
			set.AddRef(tc.input)
			assert.Equal(tt, tc.expected, set)
		})
	}
}

func TestDatasetSchemaFieldIDSet_Has(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input struct {
			DatasetSchemaFieldIDSet
			DatasetSchemaFieldID
		}
		expected bool
	}{
		{
			name: "Empty Set",
			input: struct {
				DatasetSchemaFieldIDSet
				DatasetSchemaFieldID
			}{DatasetSchemaFieldIDSet: DatasetSchemaFieldIDSet{}, DatasetSchemaFieldID: MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1")},
			expected: false,
		},
		{
			name: "Set Contains the element",
			input: struct {
				DatasetSchemaFieldIDSet
				DatasetSchemaFieldID
			}{DatasetSchemaFieldIDSet: DatasetSchemaFieldIDSet{
				m: map[DatasetSchemaFieldID]struct{}{MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{}},
				s: []DatasetSchemaFieldID{MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1")},
			}, DatasetSchemaFieldID: MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1")},
			expected: true,
		},
		{
			name: "Set does not Contains the element",
			input: struct {
				DatasetSchemaFieldIDSet
				DatasetSchemaFieldID
			}{DatasetSchemaFieldIDSet: DatasetSchemaFieldIDSet{
				m: map[DatasetSchemaFieldID]struct{}{MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{}},
				s: []DatasetSchemaFieldID{MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1")},
			}, DatasetSchemaFieldID: MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t2")},
			expected: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			assert.Equal(tt, tc.expected, tc.input.DatasetSchemaFieldIDSet.Has(tc.input.DatasetSchemaFieldID))
		})
	}
}

func TestDatasetSchemaFieldIDSet_Clear(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    DatasetSchemaFieldIDSet
		expected DatasetSchemaFieldIDSet
	}{
		{
			name:  "Empty Set",
			input: DatasetSchemaFieldIDSet{},
			expected: DatasetSchemaFieldIDSet{
				m: nil,
				s: nil,
			},
		},
		{
			name: "Set Contains the element",
			input: DatasetSchemaFieldIDSet{
				m: map[DatasetSchemaFieldID]struct{}{MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{}},
				s: []DatasetSchemaFieldID{MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			expected: DatasetSchemaFieldIDSet{
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

func TestDatasetSchemaFieldIDSet_All(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *DatasetSchemaFieldIDSet
		expected []DatasetSchemaFieldID
	}{
		{
			name: "Empty slice",
			input: &DatasetSchemaFieldIDSet{
				m: map[DatasetSchemaFieldID]struct{}{},
				s: nil,
			},
			expected: make([]DatasetSchemaFieldID, 0),
		},
		{
			name: "1 element",
			input: &DatasetSchemaFieldIDSet{
				m: map[DatasetSchemaFieldID]struct{}{MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{}},
				s: []DatasetSchemaFieldID{MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			expected: []DatasetSchemaFieldID{MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1")},
		},
		{
			name: "multiple elements",
			input: &DatasetSchemaFieldIDSet{
				m: map[DatasetSchemaFieldID]struct{}{
					MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{},
					MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t2"): struct{}{},
					MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t3"): struct{}{},
				},
				s: []DatasetSchemaFieldID{
					MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
			expected: []DatasetSchemaFieldID{
				MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t3"),
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

func TestDatasetSchemaFieldIDSet_Clone(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *DatasetSchemaFieldIDSet
		expected *DatasetSchemaFieldIDSet
	}{
		{
			name:     "nil set",
			input:    nil,
			expected: NewDatasetSchemaFieldIDSet(),
		},
		{
			name:     "Empty set",
			input:    NewDatasetSchemaFieldIDSet(),
			expected: NewDatasetSchemaFieldIDSet(),
		},
		{
			name: "1 element",
			input: &DatasetSchemaFieldIDSet{
				m: map[DatasetSchemaFieldID]struct{}{MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{}},
				s: []DatasetSchemaFieldID{MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			expected: &DatasetSchemaFieldIDSet{
				m: map[DatasetSchemaFieldID]struct{}{MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{}},
				s: []DatasetSchemaFieldID{MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
		{
			name: "multiple elements",
			input: &DatasetSchemaFieldIDSet{
				m: map[DatasetSchemaFieldID]struct{}{
					MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{},
					MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t2"): struct{}{},
					MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t3"): struct{}{},
				},
				s: []DatasetSchemaFieldID{
					MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
			expected: &DatasetSchemaFieldIDSet{
				m: map[DatasetSchemaFieldID]struct{}{
					MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{},
					MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t2"): struct{}{},
					MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t3"): struct{}{},
				},
				s: []DatasetSchemaFieldID{
					MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t3"),
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

func TestDatasetSchemaFieldIDSet_Merge(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input struct {
			a *DatasetSchemaFieldIDSet
			b *DatasetSchemaFieldIDSet
		}
		expected *DatasetSchemaFieldIDSet
	}{
		{
			name: "Empty Set",
			input: struct {
				a *DatasetSchemaFieldIDSet
				b *DatasetSchemaFieldIDSet
			}{
				a: &DatasetSchemaFieldIDSet{},
				b: &DatasetSchemaFieldIDSet{},
			},
			expected: &DatasetSchemaFieldIDSet{},
		},
		{
			name: "1 Empty Set",
			input: struct {
				a *DatasetSchemaFieldIDSet
				b *DatasetSchemaFieldIDSet
			}{
				a: &DatasetSchemaFieldIDSet{
					m: map[DatasetSchemaFieldID]struct{}{MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{}},
					s: []DatasetSchemaFieldID{MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1")},
				},
				b: &DatasetSchemaFieldIDSet{},
			},
			expected: &DatasetSchemaFieldIDSet{
				m: map[DatasetSchemaFieldID]struct{}{MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{}},
				s: []DatasetSchemaFieldID{MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
		{
			name: "2 non Empty Set",
			input: struct {
				a *DatasetSchemaFieldIDSet
				b *DatasetSchemaFieldIDSet
			}{
				a: &DatasetSchemaFieldIDSet{
					m: map[DatasetSchemaFieldID]struct{}{MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{}},
					s: []DatasetSchemaFieldID{MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1")},
				},
				b: &DatasetSchemaFieldIDSet{
					m: map[DatasetSchemaFieldID]struct{}{MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t2"): struct{}{}},
					s: []DatasetSchemaFieldID{MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t2")},
				},
			},
			expected: &DatasetSchemaFieldIDSet{
				m: map[DatasetSchemaFieldID]struct{}{
					MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{},
					MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t2"): struct{}{},
				},
				s: []DatasetSchemaFieldID{
					MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustDatasetSchemaFieldID("01f3zhcaq35403zdjnd6dcm0t2"),
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