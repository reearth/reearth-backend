// Code generated by gen, DO NOT EDIT.

package id

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/oklog/ulid"
	"github.com/stretchr/testify/assert"
)

func TestNewClusterID(t *testing.T) {
	id := NewClusterID()
	assert.NotNil(t, id)
	ulID, err := ulid.Parse(id.String())

	assert.NotNil(t, ulID)
	assert.Nil(t, err)
}

func TestClusterIDFrom(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    string
		expected struct {
			result ClusterID
			err    error
		}
	}{
		{
			name:  "Fail:Not valid string",
			input: "testMustFail",
			expected: struct {
				result ClusterID
				err    error
			}{
				ClusterID{},
				ErrInvalidID,
			},
		},
		{
			name:  "Fail:Not valid string",
			input: "",
			expected: struct {
				result ClusterID
				err    error
			}{
				ClusterID{},
				ErrInvalidID,
			},
		},
		{
			name:  "success:valid string",
			input: "01f2r7kg1fvvffp0gmexgy5hxy",
			expected: struct {
				result ClusterID
				err    error
			}{
				ClusterID{ulid.MustParse("01f2r7kg1fvvffp0gmexgy5hxy")},
				nil,
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			result, err := ClusterIDFrom(tc.input)
			assert.Equal(tt, tc.expected.result, result)
			if err != nil {
				assert.True(tt, errors.As(tc.expected.err, &err))
			}
		})
	}
}

func TestMustClusterID(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name        string
		input       string
		shouldPanic bool
		expected    ClusterID
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
			expected:    ClusterID{ulid.MustParse("01f2r7kg1fvvffp0gmexgy5hxy")},
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
			result := MustClusterID(tc.input)
			assert.Equal(tt, tc.expected, result)
		})
	}
}

func TestClusterIDFromRef(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected *ClusterID
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
			expected: &ClusterID{ulid.MustParse("01f2r7kg1fvvffp0gmexgy5hxy")},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			result := ClusterIDFromRef(&tc.input)
			assert.Equal(tt, tc.expected, result)
			if tc.expected != nil {
				assert.Equal(tt, *tc.expected, *result)
			}
		})
	}
}

func TestClusterIDFromRefID(t *testing.T) {
	id := New()

	subId := ClusterIDFromRefID(&id)

	assert.NotNil(t, subId)
	assert.Equal(t, subId.id, id.id)
}

func TestClusterID_ID(t *testing.T) {
	id := New()
	subId := ClusterIDFromRefID(&id)

	idOrg := subId.ID()

	assert.Equal(t, id, idOrg)
}

func TestClusterID_String(t *testing.T) {
	id := New()
	subId := ClusterIDFromRefID(&id)

	assert.Equal(t, subId.String(), id.String())
}

func TestClusterID_GoString(t *testing.T) {
	id := New()
	subId := ClusterIDFromRefID(&id)

	assert.Equal(t, subId.GoString(), "id.ClusterID("+id.String()+")")
}

func TestClusterID_RefString(t *testing.T) {
	id := New()
	subId := ClusterIDFromRefID(&id)

	refString := subId.StringRef()

	assert.NotNil(t, refString)
	assert.Equal(t, *refString, id.String())
}

func TestClusterID_Ref(t *testing.T) {
	id := New()
	subId := ClusterIDFromRefID(&id)

	subIdRef := subId.Ref()

	assert.Equal(t, *subId, *subIdRef)
}

func TestClusterID_Contains(t *testing.T) {
	id := NewClusterID()
	id2 := NewClusterID()
	assert.True(t, id.Contains([]ClusterID{id, id2}))
	assert.False(t, id.Contains([]ClusterID{id2}))
}

func TestClusterID_CopyRef(t *testing.T) {
	id := New()
	subId := ClusterIDFromRefID(&id)

	subIdCopyRef := subId.CopyRef()

	assert.Equal(t, *subId, *subIdCopyRef)
	assert.NotSame(t, subId, subIdCopyRef)
}

func TestClusterID_IDRef(t *testing.T) {
	id := New()
	subId := ClusterIDFromRefID(&id)

	assert.Equal(t, id, *subId.IDRef())
}

func TestClusterID_StringRef(t *testing.T) {
	id := New()
	subId := ClusterIDFromRefID(&id)

	assert.Equal(t, *subId.StringRef(), id.String())
}

func TestClusterID_MarhsalJSON(t *testing.T) {
	id := New()
	subId := ClusterIDFromRefID(&id)

	res, err := subId.MarhsalJSON()
	exp, _ := json.Marshal(subId.String())

	assert.Nil(t, err)
	assert.Equal(t, exp, res)
}

func TestClusterID_UnmarhsalJSON(t *testing.T) {
	jsonString := "\"01f3zhkysvcxsnzepyyqtq21fb\""

	subId := &ClusterID{}

	err := subId.UnmarhsalJSON([]byte(jsonString))

	assert.Nil(t, err)
	assert.Equal(t, "01f3zhkysvcxsnzepyyqtq21fb", subId.String())
}

func TestClusterID_MarshalText(t *testing.T) {
	id := New()
	subId := ClusterIDFromRefID(&id)

	res, err := subId.MarshalText()

	assert.Nil(t, err)
	assert.Equal(t, []byte(id.String()), res)
}

func TestClusterID_UnmarshalText(t *testing.T) {
	text := []byte("01f3zhcaq35403zdjnd6dcm0t2")

	subId := &ClusterID{}

	err := subId.UnmarshalText(text)

	assert.Nil(t, err)
	assert.Equal(t, "01f3zhcaq35403zdjnd6dcm0t2", subId.String())

}

func TestClusterID_IsNil(t *testing.T) {
	subId := ClusterID{}

	assert.True(t, subId.IsNil())

	id := New()
	subId = *ClusterIDFromRefID(&id)

	assert.False(t, subId.IsNil())
}

func TestClusterIDToKeys(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    []ClusterID
		expected []string
	}{
		{
			name:     "Empty slice",
			input:    make([]ClusterID, 0),
			expected: make([]string, 0),
		},
		{
			name:     "1 element",
			input:    []ClusterID{MustClusterID("01f3zhcaq35403zdjnd6dcm0t2")},
			expected: []string{"01f3zhcaq35403zdjnd6dcm0t2"},
		},
		{
			name: "multiple elements",
			input: []ClusterID{
				MustClusterID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustClusterID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustClusterID("01f3zhcaq35403zdjnd6dcm0t3"),
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
			assert.Equal(tt, tc.expected, ClusterIDToKeys(tc.input))
		})
	}

}

func TestClusterIDsFrom(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    []string
		expected struct {
			res []ClusterID
			err error
		}
	}{
		{
			name:  "Empty slice",
			input: make([]string, 0),
			expected: struct {
				res []ClusterID
				err error
			}{
				res: make([]ClusterID, 0),
				err: nil,
			},
		},
		{
			name:  "1 element",
			input: []string{"01f3zhcaq35403zdjnd6dcm0t2"},
			expected: struct {
				res []ClusterID
				err error
			}{
				res: []ClusterID{MustClusterID("01f3zhcaq35403zdjnd6dcm0t2")},
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
				res []ClusterID
				err error
			}{
				res: []ClusterID{
					MustClusterID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustClusterID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustClusterID("01f3zhcaq35403zdjnd6dcm0t3"),
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
				res []ClusterID
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
				_, err := ClusterIDsFrom(tc.input)
				assert.True(tt, errors.As(ErrInvalidID, &err))
			} else {
				res, err := ClusterIDsFrom(tc.input)
				assert.Equal(tt, tc.expected.res, res)
				assert.Nil(tt, err)
			}

		})
	}
}

func TestClusterIDsFromID(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    []ID
		expected []ClusterID
	}{
		{
			name:     "Empty slice",
			input:    make([]ID, 0),
			expected: make([]ClusterID, 0),
		},
		{
			name:     "1 element",
			input:    []ID{MustBeID("01f3zhcaq35403zdjnd6dcm0t2")},
			expected: []ClusterID{MustClusterID("01f3zhcaq35403zdjnd6dcm0t2")},
		},
		{
			name: "multiple elements",
			input: []ID{
				MustBeID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustBeID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustBeID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: []ClusterID{
				MustClusterID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustClusterID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustClusterID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			res := ClusterIDsFromID(tc.input)
			assert.Equal(tt, tc.expected, res)
		})
	}
}

func TestClusterIDsFromIDRef(t *testing.T) {
	t.Parallel()

	id1 := MustBeID("01f3zhcaq35403zdjnd6dcm0t1")
	id2 := MustBeID("01f3zhcaq35403zdjnd6dcm0t2")
	id3 := MustBeID("01f3zhcaq35403zdjnd6dcm0t3")

	testCases := []struct {
		name     string
		input    []*ID
		expected []ClusterID
	}{
		{
			name:     "Empty slice",
			input:    make([]*ID, 0),
			expected: make([]ClusterID, 0),
		},
		{
			name:     "1 element",
			input:    []*ID{&id1},
			expected: []ClusterID{MustClusterID(id1.String())},
		},
		{
			name:  "multiple elements",
			input: []*ID{&id1, &id2, &id3},
			expected: []ClusterID{
				MustClusterID(id1.String()),
				MustClusterID(id2.String()),
				MustClusterID(id3.String()),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			res := ClusterIDsFromIDRef(tc.input)
			assert.Equal(tt, tc.expected, res)
		})
	}
}

func TestClusterIDsToID(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []ClusterID
		expected []ID
	}{
		{
			name:     "Empty slice",
			input:    make([]ClusterID, 0),
			expected: make([]ID, 0),
		},
		{
			name:     "1 element",
			input:    []ClusterID{MustClusterID("01f3zhcaq35403zdjnd6dcm0t2")},
			expected: []ID{MustBeID("01f3zhcaq35403zdjnd6dcm0t2")},
		},
		{
			name: "multiple elements",
			input: []ClusterID{
				MustClusterID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustClusterID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustClusterID("01f3zhcaq35403zdjnd6dcm0t3"),
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

			res := ClusterIDsToID(tc.input)
			assert.Equal(tt, tc.expected, res)
		})
	}
}

func TestClusterIDsToIDRef(t *testing.T) {
	t.Parallel()

	id1 := MustBeID("01f3zhcaq35403zdjnd6dcm0t1")
	subId1 := MustClusterID(id1.String())
	id2 := MustBeID("01f3zhcaq35403zdjnd6dcm0t2")
	subId2 := MustClusterID(id2.String())
	id3 := MustBeID("01f3zhcaq35403zdjnd6dcm0t3")
	subId3 := MustClusterID(id3.String())

	testCases := []struct {
		name     string
		input    []*ClusterID
		expected []*ID
	}{
		{
			name:     "Empty slice",
			input:    make([]*ClusterID, 0),
			expected: make([]*ID, 0),
		},
		{
			name:     "1 element",
			input:    []*ClusterID{&subId1},
			expected: []*ID{&id1},
		},
		{
			name:     "multiple elements",
			input:    []*ClusterID{&subId1, &subId2, &subId3},
			expected: []*ID{&id1, &id2, &id3},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			res := ClusterIDsToIDRef(tc.input)
			assert.Equal(tt, tc.expected, res)
		})
	}
}

func TestNewClusterIDSet(t *testing.T) {
	ClusterIdSet := NewClusterIDSet()

	assert.NotNil(t, ClusterIdSet)
	assert.Empty(t, ClusterIdSet.m)
	assert.Empty(t, ClusterIdSet.s)
}

func TestClusterIDSet_Add(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []ClusterID
		expected *ClusterIDSet
	}{
		{
			name:  "Empty slice",
			input: make([]ClusterID, 0),
			expected: &ClusterIDSet{
				m: map[ClusterID]struct{}{},
				s: nil,
			},
		},
		{
			name:  "1 element",
			input: []ClusterID{MustClusterID("01f3zhcaq35403zdjnd6dcm0t1")},
			expected: &ClusterIDSet{
				m: map[ClusterID]struct{}{MustClusterID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []ClusterID{MustClusterID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
		{
			name: "multiple elements",
			input: []ClusterID{
				MustClusterID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustClusterID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustClusterID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: &ClusterIDSet{
				m: map[ClusterID]struct{}{
					MustClusterID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustClusterID("01f3zhcaq35403zdjnd6dcm0t2"): {},
					MustClusterID("01f3zhcaq35403zdjnd6dcm0t3"): {},
				},
				s: []ClusterID{
					MustClusterID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustClusterID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustClusterID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
		},
		{
			name: "multiple elements with duplication",
			input: []ClusterID{
				MustClusterID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustClusterID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustClusterID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: &ClusterIDSet{
				m: map[ClusterID]struct{}{
					MustClusterID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustClusterID("01f3zhcaq35403zdjnd6dcm0t3"): {},
				},
				s: []ClusterID{
					MustClusterID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustClusterID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			set := NewClusterIDSet()
			set.Add(tc.input...)
			assert.Equal(tt, tc.expected, set)
		})
	}
}

func TestClusterIDSet_AddRef(t *testing.T) {
	t.Parallel()

	ClusterId := MustClusterID("01f3zhcaq35403zdjnd6dcm0t1")

	testCases := []struct {
		name     string
		input    *ClusterID
		expected *ClusterIDSet
	}{
		{
			name:  "Empty slice",
			input: nil,
			expected: &ClusterIDSet{
				m: nil,
				s: nil,
			},
		},
		{
			name:  "1 element",
			input: &ClusterId,
			expected: &ClusterIDSet{
				m: map[ClusterID]struct{}{MustClusterID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []ClusterID{MustClusterID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			set := NewClusterIDSet()
			set.AddRef(tc.input)
			assert.Equal(tt, tc.expected, set)
		})
	}
}

func TestClusterIDSet_Has(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input struct {
			ClusterIDSet
			ClusterID
		}
		expected bool
	}{
		{
			name: "Empty Set",
			input: struct {
				ClusterIDSet
				ClusterID
			}{ClusterIDSet: ClusterIDSet{}, ClusterID: MustClusterID("01f3zhcaq35403zdjnd6dcm0t1")},
			expected: false,
		},
		{
			name: "Set Contains the element",
			input: struct {
				ClusterIDSet
				ClusterID
			}{ClusterIDSet: ClusterIDSet{
				m: map[ClusterID]struct{}{MustClusterID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []ClusterID{MustClusterID("01f3zhcaq35403zdjnd6dcm0t1")},
			}, ClusterID: MustClusterID("01f3zhcaq35403zdjnd6dcm0t1")},
			expected: true,
		},
		{
			name: "Set does not Contains the element",
			input: struct {
				ClusterIDSet
				ClusterID
			}{ClusterIDSet: ClusterIDSet{
				m: map[ClusterID]struct{}{MustClusterID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []ClusterID{MustClusterID("01f3zhcaq35403zdjnd6dcm0t1")},
			}, ClusterID: MustClusterID("01f3zhcaq35403zdjnd6dcm0t2")},
			expected: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			assert.Equal(tt, tc.expected, tc.input.ClusterIDSet.Has(tc.input.ClusterID))
		})
	}
}

func TestClusterIDSet_Clear(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    ClusterIDSet
		expected ClusterIDSet
	}{
		{
			name:  "Empty Set",
			input: ClusterIDSet{},
			expected: ClusterIDSet{
				m: nil,
				s: nil,
			},
		},
		{
			name: "Set Contains the element",
			input: ClusterIDSet{
				m: map[ClusterID]struct{}{MustClusterID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []ClusterID{MustClusterID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			expected: ClusterIDSet{
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

func TestClusterIDSet_All(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *ClusterIDSet
		expected []ClusterID
	}{
		{
			name: "Empty slice",
			input: &ClusterIDSet{
				m: map[ClusterID]struct{}{},
				s: nil,
			},
			expected: make([]ClusterID, 0),
		},
		{
			name: "1 element",
			input: &ClusterIDSet{
				m: map[ClusterID]struct{}{MustClusterID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []ClusterID{MustClusterID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			expected: []ClusterID{MustClusterID("01f3zhcaq35403zdjnd6dcm0t1")},
		},
		{
			name: "multiple elements",
			input: &ClusterIDSet{
				m: map[ClusterID]struct{}{
					MustClusterID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustClusterID("01f3zhcaq35403zdjnd6dcm0t2"): {},
					MustClusterID("01f3zhcaq35403zdjnd6dcm0t3"): {},
				},
				s: []ClusterID{
					MustClusterID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustClusterID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustClusterID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
			expected: []ClusterID{
				MustClusterID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustClusterID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustClusterID("01f3zhcaq35403zdjnd6dcm0t3"),
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

func TestClusterIDSet_Clone(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *ClusterIDSet
		expected *ClusterIDSet
	}{
		{
			name:     "nil set",
			input:    nil,
			expected: NewClusterIDSet(),
		},
		{
			name:     "Empty set",
			input:    NewClusterIDSet(),
			expected: NewClusterIDSet(),
		},
		{
			name: "1 element",
			input: &ClusterIDSet{
				m: map[ClusterID]struct{}{MustClusterID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []ClusterID{MustClusterID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			expected: &ClusterIDSet{
				m: map[ClusterID]struct{}{MustClusterID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []ClusterID{MustClusterID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
		{
			name: "multiple elements",
			input: &ClusterIDSet{
				m: map[ClusterID]struct{}{
					MustClusterID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustClusterID("01f3zhcaq35403zdjnd6dcm0t2"): {},
					MustClusterID("01f3zhcaq35403zdjnd6dcm0t3"): {},
				},
				s: []ClusterID{
					MustClusterID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustClusterID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustClusterID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
			expected: &ClusterIDSet{
				m: map[ClusterID]struct{}{
					MustClusterID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustClusterID("01f3zhcaq35403zdjnd6dcm0t2"): {},
					MustClusterID("01f3zhcaq35403zdjnd6dcm0t3"): {},
				},
				s: []ClusterID{
					MustClusterID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustClusterID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustClusterID("01f3zhcaq35403zdjnd6dcm0t3"),
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

func TestClusterIDSet_Merge(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input struct {
			a *ClusterIDSet
			b *ClusterIDSet
		}
		expected *ClusterIDSet
	}{
		{
			name: "Empty Set",
			input: struct {
				a *ClusterIDSet
				b *ClusterIDSet
			}{
				a: &ClusterIDSet{},
				b: &ClusterIDSet{},
			},
			expected: &ClusterIDSet{},
		},
		{
			name: "1 Empty Set",
			input: struct {
				a *ClusterIDSet
				b *ClusterIDSet
			}{
				a: &ClusterIDSet{
					m: map[ClusterID]struct{}{MustClusterID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
					s: []ClusterID{MustClusterID("01f3zhcaq35403zdjnd6dcm0t1")},
				},
				b: &ClusterIDSet{},
			},
			expected: &ClusterIDSet{
				m: map[ClusterID]struct{}{MustClusterID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []ClusterID{MustClusterID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
		{
			name: "2 non Empty Set",
			input: struct {
				a *ClusterIDSet
				b *ClusterIDSet
			}{
				a: &ClusterIDSet{
					m: map[ClusterID]struct{}{MustClusterID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
					s: []ClusterID{MustClusterID("01f3zhcaq35403zdjnd6dcm0t1")},
				},
				b: &ClusterIDSet{
					m: map[ClusterID]struct{}{MustClusterID("01f3zhcaq35403zdjnd6dcm0t2"): {}},
					s: []ClusterID{MustClusterID("01f3zhcaq35403zdjnd6dcm0t2")},
				},
			},
			expected: &ClusterIDSet{
				m: map[ClusterID]struct{}{
					MustClusterID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustClusterID("01f3zhcaq35403zdjnd6dcm0t2"): {},
				},
				s: []ClusterID{
					MustClusterID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustClusterID("01f3zhcaq35403zdjnd6dcm0t2"),
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
