// Code generated by gen, DO NOT EDIT.

package id

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/oklog/ulid"
	"github.com/stretchr/testify/assert"
)

func TestNewTagID(t *testing.T) {
	id := NewTagID()
	assert.NotNil(t, id)
	ulID, err := ulid.Parse(id.String())

	assert.NotNil(t, ulID)
	assert.Nil(t, err)
}

func TestTagIDFrom(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    string
		expected struct {
			result TagID
			err    error
		}
	}{
		{
			name:  "Fail:Not valid string",
			input: "testMustFail",
			expected: struct {
				result TagID
				err    error
			}{
				TagID{},
				ErrInvalidID,
			},
		},
		{
			name:  "Fail:Not valid string",
			input: "",
			expected: struct {
				result TagID
				err    error
			}{
				TagID{},
				ErrInvalidID,
			},
		},
		{
			name:  "success:valid string",
			input: "01f2r7kg1fvvffp0gmexgy5hxy",
			expected: struct {
				result TagID
				err    error
			}{
				TagID{ulid.MustParse("01f2r7kg1fvvffp0gmexgy5hxy")},
				nil,
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			result, err := TagIDFrom(tc.input)
			assert.Equal(tt, tc.expected.result, result)
			if err != nil {
				assert.True(tt, errors.As(tc.expected.err, &err))
			}
		})
	}
}

func TestMustTagID(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name        string
		input       string
		shouldPanic bool
		expected    TagID
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
			expected:    TagID{ulid.MustParse("01f2r7kg1fvvffp0gmexgy5hxy")},
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
			result := MustTagID(tc.input)
			assert.Equal(tt, tc.expected, result)
		})
	}
}

func TestTagIDFromRef(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected *TagID
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
			expected: &TagID{ulid.MustParse("01f2r7kg1fvvffp0gmexgy5hxy")},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			result := TagIDFromRef(&tc.input)
			assert.Equal(tt, tc.expected, result)
			if tc.expected != nil {
				assert.Equal(tt, *tc.expected, *result)
			}
		})
	}
}

func TestTagIDFromRefID(t *testing.T) {
	id := New()

	subId := TagIDFromRefID(&id)

	assert.NotNil(t, subId)
	assert.Equal(t, subId.id, id.id)
}

func TestTagID_ID(t *testing.T) {
	id := New()
	subId := TagIDFromRefID(&id)

	idOrg := subId.ID()

	assert.Equal(t, id, idOrg)
}

func TestTagID_String(t *testing.T) {
	id := New()
	subId := TagIDFromRefID(&id)

	assert.Equal(t, subId.String(), id.String())
}

func TestTagID_GoString(t *testing.T) {
	id := New()
	subId := TagIDFromRefID(&id)

	assert.Equal(t, subId.GoString(), "id.TagID("+id.String()+")")
}

func TestTagID_RefString(t *testing.T) {
	id := New()
	subId := TagIDFromRefID(&id)

	refString := subId.StringRef()

	assert.NotNil(t, refString)
	assert.Equal(t, *refString, id.String())
}

func TestTagID_Ref(t *testing.T) {
	id := New()
	subId := TagIDFromRefID(&id)

	subIdRef := subId.Ref()

	assert.Equal(t, *subId, *subIdRef)
}

func TestTagID_Contains(t *testing.T) {
	id := NewTagID()
	id2 := NewTagID()
	assert.True(t, id.Contains([]TagID{id, id2}))
	assert.False(t, id.Contains([]TagID{id2}))
}

func TestTagID_CopyRef(t *testing.T) {
	id := New()
	subId := TagIDFromRefID(&id)

	subIdCopyRef := subId.CopyRef()

	assert.Equal(t, *subId, *subIdCopyRef)
	assert.NotSame(t, subId, subIdCopyRef)
}

func TestTagID_IDRef(t *testing.T) {
	id := New()
	subId := TagIDFromRefID(&id)

	assert.Equal(t, id, *subId.IDRef())
}

func TestTagID_StringRef(t *testing.T) {
	id := New()
	subId := TagIDFromRefID(&id)

	assert.Equal(t, *subId.StringRef(), id.String())
}

func TestTagID_MarhsalJSON(t *testing.T) {
	id := New()
	subId := TagIDFromRefID(&id)

	res, err := subId.MarhsalJSON()
	exp, _ := json.Marshal(subId.String())

	assert.Nil(t, err)
	assert.Equal(t, exp, res)
}

func TestTagID_UnmarhsalJSON(t *testing.T) {
	jsonString := "\"01f3zhkysvcxsnzepyyqtq21fb\""

	subId := &TagID{}

	err := subId.UnmarhsalJSON([]byte(jsonString))

	assert.Nil(t, err)
	assert.Equal(t, "01f3zhkysvcxsnzepyyqtq21fb", subId.String())
}

func TestTagID_MarshalText(t *testing.T) {
	id := New()
	subId := TagIDFromRefID(&id)

	res, err := subId.MarshalText()

	assert.Nil(t, err)
	assert.Equal(t, []byte(id.String()), res)
}

func TestTagID_UnmarshalText(t *testing.T) {
	text := []byte("01f3zhcaq35403zdjnd6dcm0t2")

	subId := &TagID{}

	err := subId.UnmarshalText(text)

	assert.Nil(t, err)
	assert.Equal(t, "01f3zhcaq35403zdjnd6dcm0t2", subId.String())

}

func TestTagID_IsNil(t *testing.T) {
	subId := TagID{}

	assert.True(t, subId.IsNil())

	id := New()
	subId = *TagIDFromRefID(&id)

	assert.False(t, subId.IsNil())
}

func TestTagIDToKeys(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    []TagID
		expected []string
	}{
		{
			name:     "Empty slice",
			input:    make([]TagID, 0),
			expected: make([]string, 0),
		},
		{
			name:     "1 element",
			input:    []TagID{MustTagID("01f3zhcaq35403zdjnd6dcm0t2")},
			expected: []string{"01f3zhcaq35403zdjnd6dcm0t2"},
		},
		{
			name: "multiple elements",
			input: []TagID{
				MustTagID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustTagID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustTagID("01f3zhcaq35403zdjnd6dcm0t3"),
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
			assert.Equal(tt, tc.expected, TagIDToKeys(tc.input))
		})
	}

}

func TestTagIDsFrom(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    []string
		expected struct {
			res []TagID
			err error
		}
	}{
		{
			name:  "Empty slice",
			input: make([]string, 0),
			expected: struct {
				res []TagID
				err error
			}{
				res: make([]TagID, 0),
				err: nil,
			},
		},
		{
			name:  "1 element",
			input: []string{"01f3zhcaq35403zdjnd6dcm0t2"},
			expected: struct {
				res []TagID
				err error
			}{
				res: []TagID{MustTagID("01f3zhcaq35403zdjnd6dcm0t2")},
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
				res []TagID
				err error
			}{
				res: []TagID{
					MustTagID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustTagID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustTagID("01f3zhcaq35403zdjnd6dcm0t3"),
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
				res []TagID
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
				_, err := TagIDsFrom(tc.input)
				assert.True(tt, errors.As(ErrInvalidID, &err))
			} else {
				res, err := TagIDsFrom(tc.input)
				assert.Equal(tt, tc.expected.res, res)
				assert.Nil(tt, err)
			}

		})
	}
}

func TestTagIDsFromID(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    []ID
		expected []TagID
	}{
		{
			name:     "Empty slice",
			input:    make([]ID, 0),
			expected: make([]TagID, 0),
		},
		{
			name:     "1 element",
			input:    []ID{MustBeID("01f3zhcaq35403zdjnd6dcm0t2")},
			expected: []TagID{MustTagID("01f3zhcaq35403zdjnd6dcm0t2")},
		},
		{
			name: "multiple elements",
			input: []ID{
				MustBeID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustBeID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustBeID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: []TagID{
				MustTagID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustTagID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustTagID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			res := TagIDsFromID(tc.input)
			assert.Equal(tt, tc.expected, res)
		})
	}
}

func TestTagIDsFromIDRef(t *testing.T) {
	t.Parallel()

	id1 := MustBeID("01f3zhcaq35403zdjnd6dcm0t1")
	id2 := MustBeID("01f3zhcaq35403zdjnd6dcm0t2")
	id3 := MustBeID("01f3zhcaq35403zdjnd6dcm0t3")

	testCases := []struct {
		name     string
		input    []*ID
		expected []TagID
	}{
		{
			name:     "Empty slice",
			input:    make([]*ID, 0),
			expected: make([]TagID, 0),
		},
		{
			name:     "1 element",
			input:    []*ID{&id1},
			expected: []TagID{MustTagID(id1.String())},
		},
		{
			name:  "multiple elements",
			input: []*ID{&id1, &id2, &id3},
			expected: []TagID{
				MustTagID(id1.String()),
				MustTagID(id2.String()),
				MustTagID(id3.String()),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			res := TagIDsFromIDRef(tc.input)
			assert.Equal(tt, tc.expected, res)
		})
	}
}

func TestTagIDsToID(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []TagID
		expected []ID
	}{
		{
			name:     "Empty slice",
			input:    make([]TagID, 0),
			expected: make([]ID, 0),
		},
		{
			name:     "1 element",
			input:    []TagID{MustTagID("01f3zhcaq35403zdjnd6dcm0t2")},
			expected: []ID{MustBeID("01f3zhcaq35403zdjnd6dcm0t2")},
		},
		{
			name: "multiple elements",
			input: []TagID{
				MustTagID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustTagID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustTagID("01f3zhcaq35403zdjnd6dcm0t3"),
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

			res := TagIDsToID(tc.input)
			assert.Equal(tt, tc.expected, res)
		})
	}
}

func TestTagIDsToIDRef(t *testing.T) {
	t.Parallel()

	id1 := MustBeID("01f3zhcaq35403zdjnd6dcm0t1")
	subId1 := MustTagID(id1.String())
	id2 := MustBeID("01f3zhcaq35403zdjnd6dcm0t2")
	subId2 := MustTagID(id2.String())
	id3 := MustBeID("01f3zhcaq35403zdjnd6dcm0t3")
	subId3 := MustTagID(id3.String())

	testCases := []struct {
		name     string
		input    []*TagID
		expected []*ID
	}{
		{
			name:     "Empty slice",
			input:    make([]*TagID, 0),
			expected: make([]*ID, 0),
		},
		{
			name:     "1 element",
			input:    []*TagID{&subId1},
			expected: []*ID{&id1},
		},
		{
			name:     "multiple elements",
			input:    []*TagID{&subId1, &subId2, &subId3},
			expected: []*ID{&id1, &id2, &id3},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			res := TagIDsToIDRef(tc.input)
			assert.Equal(tt, tc.expected, res)
		})
	}
}

func TestNewTagIDSet(t *testing.T) {
	TagIdSet := NewTagIDSet()

	assert.NotNil(t, TagIdSet)
	assert.Empty(t, TagIdSet.m)
	assert.Empty(t, TagIdSet.s)
}

func TestTagIDSet_Add(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []TagID
		expected *TagIDSet
	}{
		{
			name:  "Empty slice",
			input: make([]TagID, 0),
			expected: &TagIDSet{
				m: map[TagID]struct{}{},
				s: nil,
			},
		},
		{
			name:  "1 element",
			input: []TagID{MustTagID("01f3zhcaq35403zdjnd6dcm0t1")},
			expected: &TagIDSet{
				m: map[TagID]struct{}{MustTagID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []TagID{MustTagID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
		{
			name: "multiple elements",
			input: []TagID{
				MustTagID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustTagID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustTagID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: &TagIDSet{
				m: map[TagID]struct{}{
					MustTagID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustTagID("01f3zhcaq35403zdjnd6dcm0t2"): {},
					MustTagID("01f3zhcaq35403zdjnd6dcm0t3"): {},
				},
				s: []TagID{
					MustTagID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustTagID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustTagID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
		},
		{
			name: "multiple elements with duplication",
			input: []TagID{
				MustTagID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustTagID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustTagID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: &TagIDSet{
				m: map[TagID]struct{}{
					MustTagID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustTagID("01f3zhcaq35403zdjnd6dcm0t3"): {},
				},
				s: []TagID{
					MustTagID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustTagID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			set := NewTagIDSet()
			set.Add(tc.input...)
			assert.Equal(tt, tc.expected, set)
		})
	}
}

func TestTagIDSet_AddRef(t *testing.T) {
	t.Parallel()

	TagId := MustTagID("01f3zhcaq35403zdjnd6dcm0t1")

	testCases := []struct {
		name     string
		input    *TagID
		expected *TagIDSet
	}{
		{
			name:  "Empty slice",
			input: nil,
			expected: &TagIDSet{
				m: nil,
				s: nil,
			},
		},
		{
			name:  "1 element",
			input: &TagId,
			expected: &TagIDSet{
				m: map[TagID]struct{}{MustTagID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []TagID{MustTagID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			set := NewTagIDSet()
			set.AddRef(tc.input)
			assert.Equal(tt, tc.expected, set)
		})
	}
}

func TestTagIDSet_Has(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input struct {
			TagIDSet
			TagID
		}
		expected bool
	}{
		{
			name: "Empty Set",
			input: struct {
				TagIDSet
				TagID
			}{TagIDSet: TagIDSet{}, TagID: MustTagID("01f3zhcaq35403zdjnd6dcm0t1")},
			expected: false,
		},
		{
			name: "Set Contains the element",
			input: struct {
				TagIDSet
				TagID
			}{TagIDSet: TagIDSet{
				m: map[TagID]struct{}{MustTagID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []TagID{MustTagID("01f3zhcaq35403zdjnd6dcm0t1")},
			}, TagID: MustTagID("01f3zhcaq35403zdjnd6dcm0t1")},
			expected: true,
		},
		{
			name: "Set does not Contains the element",
			input: struct {
				TagIDSet
				TagID
			}{TagIDSet: TagIDSet{
				m: map[TagID]struct{}{MustTagID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []TagID{MustTagID("01f3zhcaq35403zdjnd6dcm0t1")},
			}, TagID: MustTagID("01f3zhcaq35403zdjnd6dcm0t2")},
			expected: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			assert.Equal(tt, tc.expected, tc.input.TagIDSet.Has(tc.input.TagID))
		})
	}
}

func TestTagIDSet_Clear(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    TagIDSet
		expected TagIDSet
	}{
		{
			name:  "Empty Set",
			input: TagIDSet{},
			expected: TagIDSet{
				m: nil,
				s: nil,
			},
		},
		{
			name: "Set Contains the element",
			input: TagIDSet{
				m: map[TagID]struct{}{MustTagID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []TagID{MustTagID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			expected: TagIDSet{
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

func TestTagIDSet_All(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *TagIDSet
		expected []TagID
	}{
		{
			name: "Empty slice",
			input: &TagIDSet{
				m: map[TagID]struct{}{},
				s: nil,
			},
			expected: make([]TagID, 0),
		},
		{
			name: "1 element",
			input: &TagIDSet{
				m: map[TagID]struct{}{MustTagID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []TagID{MustTagID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			expected: []TagID{MustTagID("01f3zhcaq35403zdjnd6dcm0t1")},
		},
		{
			name: "multiple elements",
			input: &TagIDSet{
				m: map[TagID]struct{}{
					MustTagID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustTagID("01f3zhcaq35403zdjnd6dcm0t2"): {},
					MustTagID("01f3zhcaq35403zdjnd6dcm0t3"): {},
				},
				s: []TagID{
					MustTagID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustTagID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustTagID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
			expected: []TagID{
				MustTagID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustTagID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustTagID("01f3zhcaq35403zdjnd6dcm0t3"),
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

func TestTagIDSet_Clone(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *TagIDSet
		expected *TagIDSet
	}{
		{
			name:     "nil set",
			input:    nil,
			expected: NewTagIDSet(),
		},
		{
			name:     "Empty set",
			input:    NewTagIDSet(),
			expected: NewTagIDSet(),
		},
		{
			name: "1 element",
			input: &TagIDSet{
				m: map[TagID]struct{}{MustTagID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []TagID{MustTagID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			expected: &TagIDSet{
				m: map[TagID]struct{}{MustTagID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []TagID{MustTagID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
		{
			name: "multiple elements",
			input: &TagIDSet{
				m: map[TagID]struct{}{
					MustTagID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustTagID("01f3zhcaq35403zdjnd6dcm0t2"): {},
					MustTagID("01f3zhcaq35403zdjnd6dcm0t3"): {},
				},
				s: []TagID{
					MustTagID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustTagID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustTagID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
			expected: &TagIDSet{
				m: map[TagID]struct{}{
					MustTagID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustTagID("01f3zhcaq35403zdjnd6dcm0t2"): {},
					MustTagID("01f3zhcaq35403zdjnd6dcm0t3"): {},
				},
				s: []TagID{
					MustTagID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustTagID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustTagID("01f3zhcaq35403zdjnd6dcm0t3"),
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

func TestTagIDSet_Merge(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input struct {
			a *TagIDSet
			b *TagIDSet
		}
		expected *TagIDSet
	}{
		{
			name: "Empty Set",
			input: struct {
				a *TagIDSet
				b *TagIDSet
			}{
				a: &TagIDSet{},
				b: &TagIDSet{},
			},
			expected: &TagIDSet{},
		},
		{
			name: "1 Empty Set",
			input: struct {
				a *TagIDSet
				b *TagIDSet
			}{
				a: &TagIDSet{
					m: map[TagID]struct{}{MustTagID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
					s: []TagID{MustTagID("01f3zhcaq35403zdjnd6dcm0t1")},
				},
				b: &TagIDSet{},
			},
			expected: &TagIDSet{
				m: map[TagID]struct{}{MustTagID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []TagID{MustTagID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
		{
			name: "2 non Empty Set",
			input: struct {
				a *TagIDSet
				b *TagIDSet
			}{
				a: &TagIDSet{
					m: map[TagID]struct{}{MustTagID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
					s: []TagID{MustTagID("01f3zhcaq35403zdjnd6dcm0t1")},
				},
				b: &TagIDSet{
					m: map[TagID]struct{}{MustTagID("01f3zhcaq35403zdjnd6dcm0t2"): {}},
					s: []TagID{MustTagID("01f3zhcaq35403zdjnd6dcm0t2")},
				},
			},
			expected: &TagIDSet{
				m: map[TagID]struct{}{
					MustTagID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustTagID("01f3zhcaq35403zdjnd6dcm0t2"): {},
				},
				s: []TagID{
					MustTagID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustTagID("01f3zhcaq35403zdjnd6dcm0t2"),
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