// Code generated by gen, DO NOT EDIT.

package id

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/oklog/ulid"
	"github.com/stretchr/testify/assert"
)

func TestNewDatasetID(t *testing.T) {
	id := NewDatasetID()
	assert.NotNil(t, id)
	ulID, err := ulid.Parse(id.String())

	assert.NotNil(t, ulID)
	assert.Nil(t, err)
}

func TestDatasetIDFrom(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    string
		expected struct {
			result DatasetID
			err    error
		}
	}{
		{
			name:  "Fail:Not valid string",
			input: "testMustFail",
			expected: struct {
				result DatasetID
				err    error
			}{
				DatasetID{},
				ErrInvalidID,
			},
		},
		{
			name:  "Fail:Not valid string",
			input: "",
			expected: struct {
				result DatasetID
				err    error
			}{
				DatasetID{},
				ErrInvalidID,
			},
		},
		{
			name:  "success:valid string",
			input: "01f2r7kg1fvvffp0gmexgy5hxy",
			expected: struct {
				result DatasetID
				err    error
			}{
				DatasetID{ulid.MustParse("01f2r7kg1fvvffp0gmexgy5hxy")},
				nil,
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			result, err := DatasetIDFrom(tc.input)
			assert.Equal(tt, tc.expected.result, result)
			if err != nil {
				assert.True(tt, errors.As(tc.expected.err, &err))
			}
		})
	}
}

func TestMustDatasetID(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name        string
		input       string
		shouldPanic bool
		expected    DatasetID
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
			expected:    DatasetID{ulid.MustParse("01f2r7kg1fvvffp0gmexgy5hxy")},
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
			result := MustDatasetID(tc.input)
			assert.Equal(tt, tc.expected, result)
		})
	}
}

func TestDatasetIDFromRef(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected *DatasetID
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
			expected: &DatasetID{ulid.MustParse("01f2r7kg1fvvffp0gmexgy5hxy")},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			result := DatasetIDFromRef(&tc.input)
			assert.Equal(tt, tc.expected, result)
			if tc.expected != nil {
				assert.Equal(tt, *tc.expected, *result)
			}
		})
	}
}

func TestDatasetIDFromRefID(t *testing.T) {
	id := New()

	subId := DatasetIDFromRefID(&id)

	assert.NotNil(t, subId)
	assert.Equal(t, subId.id, id.id)
}

func TestDatasetID_ID(t *testing.T) {
	id := New()
	subId := DatasetIDFromRefID(&id)

	idOrg := subId.ID()

	assert.Equal(t, id, idOrg)
}

func TestDatasetID_String(t *testing.T) {
	id := New()
	subId := DatasetIDFromRefID(&id)

	assert.Equal(t, subId.String(), id.String())
}

func TestDatasetID_GoString(t *testing.T) {
	id := New()
	subId := DatasetIDFromRefID(&id)

	assert.Equal(t, subId.GoString(), "id.DatasetID("+id.String()+")")
}

func TestDatasetID_RefString(t *testing.T) {
	id := New()
	subId := DatasetIDFromRefID(&id)

	refString := subId.StringRef()

	assert.NotNil(t, refString)
	assert.Equal(t, *refString, id.String())
}

func TestDatasetID_Ref(t *testing.T) {
	id := New()
	subId := DatasetIDFromRefID(&id)

	subIdRef := subId.Ref()

	assert.Equal(t, *subId, *subIdRef)
}

func TestDatasetID_CopyRef(t *testing.T) {
	id := New()
	subId := DatasetIDFromRefID(&id)

	subIdCopyRef := subId.CopyRef()

	assert.Equal(t, *subId, *subIdCopyRef)
	assert.False(t, subId == subIdCopyRef)
}

func TestDatasetID_IDRef(t *testing.T) {
	id := New()
	subId := DatasetIDFromRefID(&id)

	assert.Equal(t, id, *subId.IDRef())
}

func TestDatasetID_StringRef(t *testing.T) {
	id := New()
	subId := DatasetIDFromRefID(&id)

	assert.Equal(t, *subId.StringRef(), id.String())
}

func TestDatasetID_MarhsalJSON(t *testing.T) {
	id := New()
	subId := DatasetIDFromRefID(&id)

	res, err := subId.MarhsalJSON()
	exp, _ := json.Marshal(subId.String())

	assert.Nil(t, err)
	assert.Equal(t, exp, res)
}

func TestDatasetID_UnmarhsalJSON(t *testing.T) {
	jsonString := "\"01f3zhkysvcxsnzepyyqtq21fb\""

	subId := &DatasetID{}

	err := subId.UnmarhsalJSON([]byte(jsonString))

	assert.Nil(t, err)
	assert.Equal(t, "01f3zhkysvcxsnzepyyqtq21fb", subId.String())
}

func TestDatasetID_MarshalText(t *testing.T) {
	id := New()
	subId := DatasetIDFromRefID(&id)

	res, err := subId.MarshalText()

	assert.Nil(t, err)
	assert.Equal(t, []byte(id.String()), res)
}

func TestDatasetID_UnmarshalText(t *testing.T) {
	text := []byte("01f3zhcaq35403zdjnd6dcm0t2")

	subId := &DatasetID{}

	err := subId.UnmarshalText(text)

	assert.Nil(t, err)
	assert.Equal(t, "01f3zhcaq35403zdjnd6dcm0t2", subId.String())

}

func TestDatasetID_IsNil(t *testing.T) {
	subId := DatasetID{}

	assert.True(t, subId.IsNil())

	id := New()
	subId = *DatasetIDFromRefID(&id)

	assert.False(t, subId.IsNil())
}

func TestDatasetIDToKeys(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    []DatasetID
		expected []string
	}{
		{
			name:     "Empty slice",
			input:    make([]DatasetID, 0),
			expected: make([]string, 0),
		},
		{
			name:     "1 element",
			input:    []DatasetID{MustDatasetID("01f3zhcaq35403zdjnd6dcm0t2")},
			expected: []string{"01f3zhcaq35403zdjnd6dcm0t2"},
		},
		{
			name: "multiple elements",
			input: []DatasetID{
				MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustDatasetID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustDatasetID("01f3zhcaq35403zdjnd6dcm0t3"),
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
			assert.Equal(tt, tc.expected, DatasetIDToKeys(tc.input))
		})
	}

}

func TestDatasetIDsFrom(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    []string
		expected struct {
			res []DatasetID
			err error
		}
	}{
		{
			name:  "Empty slice",
			input: make([]string, 0),
			expected: struct {
				res []DatasetID
				err error
			}{
				res: make([]DatasetID, 0),
				err: nil,
			},
		},
		{
			name:  "1 element",
			input: []string{"01f3zhcaq35403zdjnd6dcm0t2"},
			expected: struct {
				res []DatasetID
				err error
			}{
				res: []DatasetID{MustDatasetID("01f3zhcaq35403zdjnd6dcm0t2")},
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
				res []DatasetID
				err error
			}{
				res: []DatasetID{
					MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustDatasetID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustDatasetID("01f3zhcaq35403zdjnd6dcm0t3"),
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
				res []DatasetID
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
				_, err := DatasetIDsFrom(tc.input)
				assert.True(tt, errors.As(ErrInvalidID, &err))
			} else {
				res, err := DatasetIDsFrom(tc.input)
				assert.Equal(tt, tc.expected.res, res)
				assert.Nil(tt, err)
			}

		})
	}
}

func TestDatasetIDsFromID(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    []ID
		expected []DatasetID
	}{
		{
			name:     "Empty slice",
			input:    make([]ID, 0),
			expected: make([]DatasetID, 0),
		},
		{
			name:     "1 element",
			input:    []ID{MustBeID("01f3zhcaq35403zdjnd6dcm0t2")},
			expected: []DatasetID{MustDatasetID("01f3zhcaq35403zdjnd6dcm0t2")},
		},
		{
			name: "multiple elements",
			input: []ID{
				MustBeID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustBeID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustBeID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: []DatasetID{
				MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustDatasetID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustDatasetID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			res := DatasetIDsFromID(tc.input)
			assert.Equal(tt, tc.expected, res)
		})
	}
}

func TestDatasetIDsFromIDRef(t *testing.T) {
	t.Parallel()

	id1 := MustBeID("01f3zhcaq35403zdjnd6dcm0t1")
	id2 := MustBeID("01f3zhcaq35403zdjnd6dcm0t2")
	id3 := MustBeID("01f3zhcaq35403zdjnd6dcm0t3")

	testCases := []struct {
		name     string
		input    []*ID
		expected []DatasetID
	}{
		{
			name:     "Empty slice",
			input:    make([]*ID, 0),
			expected: make([]DatasetID, 0),
		},
		{
			name:     "1 element",
			input:    []*ID{&id1},
			expected: []DatasetID{MustDatasetID(id1.String())},
		},
		{
			name:  "multiple elements",
			input: []*ID{&id1, &id2, &id3},
			expected: []DatasetID{
				MustDatasetID(id1.String()),
				MustDatasetID(id2.String()),
				MustDatasetID(id3.String()),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			res := DatasetIDsFromIDRef(tc.input)
			assert.Equal(tt, tc.expected, res)
		})
	}
}

func TestDatasetIDsToID(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []DatasetID
		expected []ID
	}{
		{
			name:     "Empty slice",
			input:    make([]DatasetID, 0),
			expected: make([]ID, 0),
		},
		{
			name:     "1 element",
			input:    []DatasetID{MustDatasetID("01f3zhcaq35403zdjnd6dcm0t2")},
			expected: []ID{MustBeID("01f3zhcaq35403zdjnd6dcm0t2")},
		},
		{
			name: "multiple elements",
			input: []DatasetID{
				MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustDatasetID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustDatasetID("01f3zhcaq35403zdjnd6dcm0t3"),
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

			res := DatasetIDsToID(tc.input)
			assert.Equal(tt, tc.expected, res)
		})
	}
}

func TestDatasetIDsToIDRef(t *testing.T) {
	t.Parallel()

	id1 := MustBeID("01f3zhcaq35403zdjnd6dcm0t1")
	subId1 := MustDatasetID(id1.String())
	id2 := MustBeID("01f3zhcaq35403zdjnd6dcm0t2")
	subId2 := MustDatasetID(id2.String())
	id3 := MustBeID("01f3zhcaq35403zdjnd6dcm0t3")
	subId3 := MustDatasetID(id3.String())

	testCases := []struct {
		name     string
		input    []*DatasetID
		expected []*ID
	}{
		{
			name:     "Empty slice",
			input:    make([]*DatasetID, 0),
			expected: make([]*ID, 0),
		},
		{
			name:     "1 element",
			input:    []*DatasetID{&subId1},
			expected: []*ID{&id1},
		},
		{
			name:     "multiple elements",
			input:    []*DatasetID{&subId1, &subId2, &subId3},
			expected: []*ID{&id1, &id2, &id3},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			res := DatasetIDsToIDRef(tc.input)
			assert.Equal(tt, tc.expected, res)
		})
	}
}

func TestNewDatasetIDSet(t *testing.T) {
	DatasetIdSet := NewDatasetIDSet()

	assert.NotNil(t, DatasetIdSet)
	assert.Empty(t, DatasetIdSet.m)
	assert.Empty(t, DatasetIdSet.s)
}

func TestDatasetIDSet_Add(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []DatasetID
		expected *DatasetIDSet
	}{
		{
			name:  "Empty slice",
			input: make([]DatasetID, 0),
			expected: &DatasetIDSet{
				m: map[DatasetID]struct{}{},
				s: nil,
			},
		},
		{
			name:  "1 element",
			input: []DatasetID{MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1")},
			expected: &DatasetIDSet{
				m: map[DatasetID]struct{}{MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{}},
				s: []DatasetID{MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
		{
			name: "multiple elements",
			input: []DatasetID{
				MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustDatasetID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustDatasetID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: &DatasetIDSet{
				m: map[DatasetID]struct{}{
					MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{},
					MustDatasetID("01f3zhcaq35403zdjnd6dcm0t2"): struct{}{},
					MustDatasetID("01f3zhcaq35403zdjnd6dcm0t3"): struct{}{},
				},
				s: []DatasetID{
					MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustDatasetID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustDatasetID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
		},
		{
			name: "multiple elements with duplication",
			input: []DatasetID{
				MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustDatasetID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: &DatasetIDSet{
				m: map[DatasetID]struct{}{
					MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{},
					MustDatasetID("01f3zhcaq35403zdjnd6dcm0t3"): struct{}{},
				},
				s: []DatasetID{
					MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustDatasetID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			set := NewDatasetIDSet()
			set.Add(tc.input...)
			assert.Equal(tt, tc.expected, set)
		})
	}
}

func TestDatasetIDSet_AddRef(t *testing.T) {
	t.Parallel()

	DatasetId := MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1")

	testCases := []struct {
		name     string
		input    *DatasetID
		expected *DatasetIDSet
	}{
		{
			name:  "Empty slice",
			input: nil,
			expected: &DatasetIDSet{
				m: nil,
				s: nil,
			},
		},
		{
			name:  "1 element",
			input: &DatasetId,
			expected: &DatasetIDSet{
				m: map[DatasetID]struct{}{MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{}},
				s: []DatasetID{MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			set := NewDatasetIDSet()
			set.AddRef(tc.input)
			assert.Equal(tt, tc.expected, set)
		})
	}
}

func TestDatasetIDSet_Has(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input struct {
			DatasetIDSet
			DatasetID
		}
		expected bool
	}{
		{
			name: "Empty Set",
			input: struct {
				DatasetIDSet
				DatasetID
			}{DatasetIDSet: DatasetIDSet{}, DatasetID: MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1")},
			expected: false,
		},
		{
			name: "Set Contains the element",
			input: struct {
				DatasetIDSet
				DatasetID
			}{DatasetIDSet: DatasetIDSet{
				m: map[DatasetID]struct{}{MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{}},
				s: []DatasetID{MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1")},
			}, DatasetID: MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1")},
			expected: true,
		},
		{
			name: "Set does not Contains the element",
			input: struct {
				DatasetIDSet
				DatasetID
			}{DatasetIDSet: DatasetIDSet{
				m: map[DatasetID]struct{}{MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{}},
				s: []DatasetID{MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1")},
			}, DatasetID: MustDatasetID("01f3zhcaq35403zdjnd6dcm0t2")},
			expected: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			assert.Equal(tt, tc.expected, tc.input.DatasetIDSet.Has(tc.input.DatasetID))
		})
	}
}

func TestDatasetIDSet_Clear(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    DatasetIDSet
		expected DatasetIDSet
	}{
		{
			name:  "Empty Set",
			input: DatasetIDSet{},
			expected: DatasetIDSet{
				m: nil,
				s: nil,
			},
		},
		{
			name: "Set Contains the element",
			input: DatasetIDSet{
				m: map[DatasetID]struct{}{MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{}},
				s: []DatasetID{MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			expected: DatasetIDSet{
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

func TestDatasetIDSet_All(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *DatasetIDSet
		expected []DatasetID
	}{
		{
			name: "Empty slice",
			input: &DatasetIDSet{
				m: map[DatasetID]struct{}{},
				s: nil,
			},
			expected: make([]DatasetID, 0),
		},
		{
			name: "1 element",
			input: &DatasetIDSet{
				m: map[DatasetID]struct{}{MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{}},
				s: []DatasetID{MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			expected: []DatasetID{MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1")},
		},
		{
			name: "multiple elements",
			input: &DatasetIDSet{
				m: map[DatasetID]struct{}{
					MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{},
					MustDatasetID("01f3zhcaq35403zdjnd6dcm0t2"): struct{}{},
					MustDatasetID("01f3zhcaq35403zdjnd6dcm0t3"): struct{}{},
				},
				s: []DatasetID{
					MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustDatasetID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustDatasetID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
			expected: []DatasetID{
				MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustDatasetID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustDatasetID("01f3zhcaq35403zdjnd6dcm0t3"),
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

func TestDatasetIDSet_Clone(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *DatasetIDSet
		expected *DatasetIDSet
	}{
		{
			name:     "nil set",
			input:    nil,
			expected: NewDatasetIDSet(),
		},
		{
			name:     "Empty set",
			input:    NewDatasetIDSet(),
			expected: NewDatasetIDSet(),
		},
		{
			name: "1 element",
			input: &DatasetIDSet{
				m: map[DatasetID]struct{}{MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{}},
				s: []DatasetID{MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			expected: &DatasetIDSet{
				m: map[DatasetID]struct{}{MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{}},
				s: []DatasetID{MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
		{
			name: "multiple elements",
			input: &DatasetIDSet{
				m: map[DatasetID]struct{}{
					MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{},
					MustDatasetID("01f3zhcaq35403zdjnd6dcm0t2"): struct{}{},
					MustDatasetID("01f3zhcaq35403zdjnd6dcm0t3"): struct{}{},
				},
				s: []DatasetID{
					MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustDatasetID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustDatasetID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
			expected: &DatasetIDSet{
				m: map[DatasetID]struct{}{
					MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{},
					MustDatasetID("01f3zhcaq35403zdjnd6dcm0t2"): struct{}{},
					MustDatasetID("01f3zhcaq35403zdjnd6dcm0t3"): struct{}{},
				},
				s: []DatasetID{
					MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustDatasetID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustDatasetID("01f3zhcaq35403zdjnd6dcm0t3"),
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

func TestDatasetIDSet_Merge(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input struct {
			a *DatasetIDSet
			b *DatasetIDSet
		}
		expected *DatasetIDSet
	}{
		{
			name: "Empty Set",
			input: struct {
				a *DatasetIDSet
				b *DatasetIDSet
			}{
				a: &DatasetIDSet{},
				b: &DatasetIDSet{},
			},
			expected: &DatasetIDSet{},
		},
		{
			name: "1 Empty Set",
			input: struct {
				a *DatasetIDSet
				b *DatasetIDSet
			}{
				a: &DatasetIDSet{
					m: map[DatasetID]struct{}{MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{}},
					s: []DatasetID{MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1")},
				},
				b: &DatasetIDSet{},
			},
			expected: &DatasetIDSet{
				m: map[DatasetID]struct{}{MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{}},
				s: []DatasetID{MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
		{
			name: "2 non Empty Set",
			input: struct {
				a *DatasetIDSet
				b *DatasetIDSet
			}{
				a: &DatasetIDSet{
					m: map[DatasetID]struct{}{MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{}},
					s: []DatasetID{MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1")},
				},
				b: &DatasetIDSet{
					m: map[DatasetID]struct{}{MustDatasetID("01f3zhcaq35403zdjnd6dcm0t2"): struct{}{}},
					s: []DatasetID{MustDatasetID("01f3zhcaq35403zdjnd6dcm0t2")},
				},
			},
			expected: &DatasetIDSet{
				m: map[DatasetID]struct{}{
					MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1"): struct{}{},
					MustDatasetID("01f3zhcaq35403zdjnd6dcm0t2"): struct{}{},
				},
				s: []DatasetID{
					MustDatasetID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustDatasetID("01f3zhcaq35403zdjnd6dcm0t2"),
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
