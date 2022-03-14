// Code generated by gen, DO NOT EDIT.

package id

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/oklog/ulid"
	"github.com/stretchr/testify/assert"
)

func TestNewAuthRequestID(t *testing.T) {
	id := NewAuthRequestID()
	assert.NotNil(t, id)
	ulID, err := ulid.Parse(id.String())

	assert.NotNil(t, ulID)
	assert.Nil(t, err)
}

func TestAuthRequestIDFrom(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    string
		expected struct {
			result AuthRequestID
			err    error
		}
	}{
		{
			name:  "Fail:Not valid string",
			input: "testMustFail",
			expected: struct {
				result AuthRequestID
				err    error
			}{
				AuthRequestID{},
				ErrInvalidID,
			},
		},
		{
			name:  "Fail:Not valid string",
			input: "",
			expected: struct {
				result AuthRequestID
				err    error
			}{
				AuthRequestID{},
				ErrInvalidID,
			},
		},
		{
			name:  "success:valid string",
			input: "01f2r7kg1fvvffp0gmexgy5hxy",
			expected: struct {
				result AuthRequestID
				err    error
			}{
				AuthRequestID{ulid.MustParse("01f2r7kg1fvvffp0gmexgy5hxy")},
				nil,
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			result, err := AuthRequestIDFrom(tc.input)
			assert.Equal(tt, tc.expected.result, result)
			if err != nil {
				assert.True(tt, errors.As(tc.expected.err, &err))
			}
		})
	}
}

func TestMustAuthRequestID(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name        string
		input       string
		shouldPanic bool
		expected    AuthRequestID
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
			expected:    AuthRequestID{ulid.MustParse("01f2r7kg1fvvffp0gmexgy5hxy")},
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
			result := MustAuthRequestID(tc.input)
			assert.Equal(tt, tc.expected, result)
		})
	}
}

func TestAuthRequestIDFromRef(t *testing.T) {
	testCases := []struct {
		name     string
		input    string
		expected *AuthRequestID
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
			expected: &AuthRequestID{ulid.MustParse("01f2r7kg1fvvffp0gmexgy5hxy")},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			result := AuthRequestIDFromRef(&tc.input)
			assert.Equal(tt, tc.expected, result)
			if tc.expected != nil {
				assert.Equal(tt, *tc.expected, *result)
			}
		})
	}
}

func TestAuthRequestIDFromRefID(t *testing.T) {
	id := New()

	subId := AuthRequestIDFromRefID(&id)

	assert.NotNil(t, subId)
	assert.Equal(t, subId.id, id.id)
}

func TestAuthRequestID_ID(t *testing.T) {
	id := New()
	subId := AuthRequestIDFromRefID(&id)

	idOrg := subId.ID()

	assert.Equal(t, id, idOrg)
}

func TestAuthRequestID_String(t *testing.T) {
	id := New()
	subId := AuthRequestIDFromRefID(&id)

	assert.Equal(t, subId.String(), id.String())
}

func TestAuthRequestID_GoString(t *testing.T) {
	id := New()
	subId := AuthRequestIDFromRefID(&id)

	assert.Equal(t, subId.GoString(), "id.AuthRequestID("+id.String()+")")
}

func TestAuthRequestID_RefString(t *testing.T) {
	id := New()
	subId := AuthRequestIDFromRefID(&id)

	refString := subId.StringRef()

	assert.NotNil(t, refString)
	assert.Equal(t, *refString, id.String())
}

func TestAuthRequestID_Ref(t *testing.T) {
	id := New()
	subId := AuthRequestIDFromRefID(&id)

	subIdRef := subId.Ref()

	assert.Equal(t, *subId, *subIdRef)
}

func TestAuthRequestID_Contains(t *testing.T) {
	id := NewAuthRequestID()
	id2 := NewAuthRequestID()
	assert.True(t, id.Contains([]AuthRequestID{id, id2}))
	assert.False(t, id.Contains([]AuthRequestID{id2}))
}

func TestAuthRequestID_CopyRef(t *testing.T) {
	id := New()
	subId := AuthRequestIDFromRefID(&id)

	subIdCopyRef := subId.CopyRef()

	assert.Equal(t, *subId, *subIdCopyRef)
	assert.NotSame(t, subId, subIdCopyRef)
}

func TestAuthRequestID_IDRef(t *testing.T) {
	id := New()
	subId := AuthRequestIDFromRefID(&id)

	assert.Equal(t, id, *subId.IDRef())
}

func TestAuthRequestID_StringRef(t *testing.T) {
	id := New()
	subId := AuthRequestIDFromRefID(&id)

	assert.Equal(t, *subId.StringRef(), id.String())
}

func TestAuthRequestID_MarhsalJSON(t *testing.T) {
	id := New()
	subId := AuthRequestIDFromRefID(&id)

	res, err := subId.MarhsalJSON()
	exp, _ := json.Marshal(subId.String())

	assert.Nil(t, err)
	assert.Equal(t, exp, res)
}

func TestAuthRequestID_UnmarhsalJSON(t *testing.T) {
	jsonString := "\"01f3zhkysvcxsnzepyyqtq21fb\""

	subId := &AuthRequestID{}

	err := subId.UnmarhsalJSON([]byte(jsonString))

	assert.Nil(t, err)
	assert.Equal(t, "01f3zhkysvcxsnzepyyqtq21fb", subId.String())
}

func TestAuthRequestID_MarshalText(t *testing.T) {
	id := New()
	subId := AuthRequestIDFromRefID(&id)

	res, err := subId.MarshalText()

	assert.Nil(t, err)
	assert.Equal(t, []byte(id.String()), res)
}

func TestAuthRequestID_UnmarshalText(t *testing.T) {
	text := []byte("01f3zhcaq35403zdjnd6dcm0t2")

	subId := &AuthRequestID{}

	err := subId.UnmarshalText(text)

	assert.Nil(t, err)
	assert.Equal(t, "01f3zhcaq35403zdjnd6dcm0t2", subId.String())

}

func TestAuthRequestID_IsNil(t *testing.T) {
	subId := AuthRequestID{}

	assert.True(t, subId.IsNil())

	id := New()
	subId = *AuthRequestIDFromRefID(&id)

	assert.False(t, subId.IsNil())
}

func TestAuthRequestIDToKeys(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    []AuthRequestID
		expected []string
	}{
		{
			name:     "Empty slice",
			input:    make([]AuthRequestID, 0),
			expected: make([]string, 0),
		},
		{
			name:     "1 element",
			input:    []AuthRequestID{MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t2")},
			expected: []string{"01f3zhcaq35403zdjnd6dcm0t2"},
		},
		{
			name: "multiple elements",
			input: []AuthRequestID{
				MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t3"),
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
			assert.Equal(tt, tc.expected, AuthRequestIDToKeys(tc.input))
		})
	}

}

func TestAuthRequestIDsFrom(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    []string
		expected struct {
			res []AuthRequestID
			err error
		}
	}{
		{
			name:  "Empty slice",
			input: make([]string, 0),
			expected: struct {
				res []AuthRequestID
				err error
			}{
				res: make([]AuthRequestID, 0),
				err: nil,
			},
		},
		{
			name:  "1 element",
			input: []string{"01f3zhcaq35403zdjnd6dcm0t2"},
			expected: struct {
				res []AuthRequestID
				err error
			}{
				res: []AuthRequestID{MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t2")},
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
				res []AuthRequestID
				err error
			}{
				res: []AuthRequestID{
					MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t3"),
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
				res []AuthRequestID
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
				_, err := AuthRequestIDsFrom(tc.input)
				assert.True(tt, errors.As(ErrInvalidID, &err))
			} else {
				res, err := AuthRequestIDsFrom(tc.input)
				assert.Equal(tt, tc.expected.res, res)
				assert.Nil(tt, err)
			}

		})
	}
}

func TestAuthRequestIDsFromID(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		name     string
		input    []ID
		expected []AuthRequestID
	}{
		{
			name:     "Empty slice",
			input:    make([]ID, 0),
			expected: make([]AuthRequestID, 0),
		},
		{
			name:     "1 element",
			input:    []ID{MustBeID("01f3zhcaq35403zdjnd6dcm0t2")},
			expected: []AuthRequestID{MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t2")},
		},
		{
			name: "multiple elements",
			input: []ID{
				MustBeID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustBeID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustBeID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: []AuthRequestID{
				MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			res := AuthRequestIDsFromID(tc.input)
			assert.Equal(tt, tc.expected, res)
		})
	}
}

func TestAuthRequestIDsFromIDRef(t *testing.T) {
	t.Parallel()

	id1 := MustBeID("01f3zhcaq35403zdjnd6dcm0t1")
	id2 := MustBeID("01f3zhcaq35403zdjnd6dcm0t2")
	id3 := MustBeID("01f3zhcaq35403zdjnd6dcm0t3")

	testCases := []struct {
		name     string
		input    []*ID
		expected []AuthRequestID
	}{
		{
			name:     "Empty slice",
			input:    make([]*ID, 0),
			expected: make([]AuthRequestID, 0),
		},
		{
			name:     "1 element",
			input:    []*ID{&id1},
			expected: []AuthRequestID{MustAuthRequestID(id1.String())},
		},
		{
			name:  "multiple elements",
			input: []*ID{&id1, &id2, &id3},
			expected: []AuthRequestID{
				MustAuthRequestID(id1.String()),
				MustAuthRequestID(id2.String()),
				MustAuthRequestID(id3.String()),
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			res := AuthRequestIDsFromIDRef(tc.input)
			assert.Equal(tt, tc.expected, res)
		})
	}
}

func TestAuthRequestIDsToID(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []AuthRequestID
		expected []ID
	}{
		{
			name:     "Empty slice",
			input:    make([]AuthRequestID, 0),
			expected: make([]ID, 0),
		},
		{
			name:     "1 element",
			input:    []AuthRequestID{MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t2")},
			expected: []ID{MustBeID("01f3zhcaq35403zdjnd6dcm0t2")},
		},
		{
			name: "multiple elements",
			input: []AuthRequestID{
				MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t3"),
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

			res := AuthRequestIDsToID(tc.input)
			assert.Equal(tt, tc.expected, res)
		})
	}
}

func TestAuthRequestIDsToIDRef(t *testing.T) {
	t.Parallel()

	id1 := MustBeID("01f3zhcaq35403zdjnd6dcm0t1")
	subId1 := MustAuthRequestID(id1.String())
	id2 := MustBeID("01f3zhcaq35403zdjnd6dcm0t2")
	subId2 := MustAuthRequestID(id2.String())
	id3 := MustBeID("01f3zhcaq35403zdjnd6dcm0t3")
	subId3 := MustAuthRequestID(id3.String())

	testCases := []struct {
		name     string
		input    []*AuthRequestID
		expected []*ID
	}{
		{
			name:     "Empty slice",
			input:    make([]*AuthRequestID, 0),
			expected: make([]*ID, 0),
		},
		{
			name:     "1 element",
			input:    []*AuthRequestID{&subId1},
			expected: []*ID{&id1},
		},
		{
			name:     "multiple elements",
			input:    []*AuthRequestID{&subId1, &subId2, &subId3},
			expected: []*ID{&id1, &id2, &id3},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			res := AuthRequestIDsToIDRef(tc.input)
			assert.Equal(tt, tc.expected, res)
		})
	}
}

func TestNewAuthRequestIDSet(t *testing.T) {
	AuthRequestIdSet := NewAuthRequestIDSet()

	assert.NotNil(t, AuthRequestIdSet)
	assert.Empty(t, AuthRequestIdSet.m)
	assert.Empty(t, AuthRequestIdSet.s)
}

func TestAuthRequestIDSet_Add(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    []AuthRequestID
		expected *AuthRequestIDSet
	}{
		{
			name:  "Empty slice",
			input: make([]AuthRequestID, 0),
			expected: &AuthRequestIDSet{
				m: map[AuthRequestID]struct{}{},
				s: nil,
			},
		},
		{
			name:  "1 element",
			input: []AuthRequestID{MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1")},
			expected: &AuthRequestIDSet{
				m: map[AuthRequestID]struct{}{MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []AuthRequestID{MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
		{
			name: "multiple elements",
			input: []AuthRequestID{
				MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: &AuthRequestIDSet{
				m: map[AuthRequestID]struct{}{
					MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t2"): {},
					MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t3"): {},
				},
				s: []AuthRequestID{
					MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
		},
		{
			name: "multiple elements with duplication",
			input: []AuthRequestID{
				MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t3"),
			},
			expected: &AuthRequestIDSet{
				m: map[AuthRequestID]struct{}{
					MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t3"): {},
				},
				s: []AuthRequestID{
					MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			set := NewAuthRequestIDSet()
			set.Add(tc.input...)
			assert.Equal(tt, tc.expected, set)
		})
	}
}

func TestAuthRequestIDSet_AddRef(t *testing.T) {
	t.Parallel()

	AuthRequestId := MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1")

	testCases := []struct {
		name     string
		input    *AuthRequestID
		expected *AuthRequestIDSet
	}{
		{
			name:  "Empty slice",
			input: nil,
			expected: &AuthRequestIDSet{
				m: nil,
				s: nil,
			},
		},
		{
			name:  "1 element",
			input: &AuthRequestId,
			expected: &AuthRequestIDSet{
				m: map[AuthRequestID]struct{}{MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []AuthRequestID{MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()

			set := NewAuthRequestIDSet()
			set.AddRef(tc.input)
			assert.Equal(tt, tc.expected, set)
		})
	}
}

func TestAuthRequestIDSet_Has(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input struct {
			AuthRequestIDSet
			AuthRequestID
		}
		expected bool
	}{
		{
			name: "Empty Set",
			input: struct {
				AuthRequestIDSet
				AuthRequestID
			}{AuthRequestIDSet: AuthRequestIDSet{}, AuthRequestID: MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1")},
			expected: false,
		},
		{
			name: "Set Contains the element",
			input: struct {
				AuthRequestIDSet
				AuthRequestID
			}{AuthRequestIDSet: AuthRequestIDSet{
				m: map[AuthRequestID]struct{}{MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []AuthRequestID{MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1")},
			}, AuthRequestID: MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1")},
			expected: true,
		},
		{
			name: "Set does not Contains the element",
			input: struct {
				AuthRequestIDSet
				AuthRequestID
			}{AuthRequestIDSet: AuthRequestIDSet{
				m: map[AuthRequestID]struct{}{MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []AuthRequestID{MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1")},
			}, AuthRequestID: MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t2")},
			expected: false,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(tt *testing.T) {
			tt.Parallel()
			assert.Equal(tt, tc.expected, tc.input.AuthRequestIDSet.Has(tc.input.AuthRequestID))
		})
	}
}

func TestAuthRequestIDSet_Clear(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    AuthRequestIDSet
		expected AuthRequestIDSet
	}{
		{
			name:  "Empty Set",
			input: AuthRequestIDSet{},
			expected: AuthRequestIDSet{
				m: nil,
				s: nil,
			},
		},
		{
			name: "Set Contains the element",
			input: AuthRequestIDSet{
				m: map[AuthRequestID]struct{}{MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []AuthRequestID{MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			expected: AuthRequestIDSet{
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

func TestAuthRequestIDSet_All(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *AuthRequestIDSet
		expected []AuthRequestID
	}{
		{
			name: "Empty slice",
			input: &AuthRequestIDSet{
				m: map[AuthRequestID]struct{}{},
				s: nil,
			},
			expected: make([]AuthRequestID, 0),
		},
		{
			name: "1 element",
			input: &AuthRequestIDSet{
				m: map[AuthRequestID]struct{}{MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []AuthRequestID{MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			expected: []AuthRequestID{MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1")},
		},
		{
			name: "multiple elements",
			input: &AuthRequestIDSet{
				m: map[AuthRequestID]struct{}{
					MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t2"): {},
					MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t3"): {},
				},
				s: []AuthRequestID{
					MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
			expected: []AuthRequestID{
				MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1"),
				MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t2"),
				MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t3"),
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

func TestAuthRequestIDSet_Clone(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		input    *AuthRequestIDSet
		expected *AuthRequestIDSet
	}{
		{
			name:     "nil set",
			input:    nil,
			expected: NewAuthRequestIDSet(),
		},
		{
			name:     "Empty set",
			input:    NewAuthRequestIDSet(),
			expected: NewAuthRequestIDSet(),
		},
		{
			name: "1 element",
			input: &AuthRequestIDSet{
				m: map[AuthRequestID]struct{}{MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []AuthRequestID{MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
			expected: &AuthRequestIDSet{
				m: map[AuthRequestID]struct{}{MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []AuthRequestID{MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
		{
			name: "multiple elements",
			input: &AuthRequestIDSet{
				m: map[AuthRequestID]struct{}{
					MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t2"): {},
					MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t3"): {},
				},
				s: []AuthRequestID{
					MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t3"),
				},
			},
			expected: &AuthRequestIDSet{
				m: map[AuthRequestID]struct{}{
					MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t2"): {},
					MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t3"): {},
				},
				s: []AuthRequestID{
					MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t2"),
					MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t3"),
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

func TestAuthRequestIDSet_Merge(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		input struct {
			a *AuthRequestIDSet
			b *AuthRequestIDSet
		}
		expected *AuthRequestIDSet
	}{
		{
			name: "Empty Set",
			input: struct {
				a *AuthRequestIDSet
				b *AuthRequestIDSet
			}{
				a: &AuthRequestIDSet{},
				b: &AuthRequestIDSet{},
			},
			expected: &AuthRequestIDSet{},
		},
		{
			name: "1 Empty Set",
			input: struct {
				a *AuthRequestIDSet
				b *AuthRequestIDSet
			}{
				a: &AuthRequestIDSet{
					m: map[AuthRequestID]struct{}{MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
					s: []AuthRequestID{MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1")},
				},
				b: &AuthRequestIDSet{},
			},
			expected: &AuthRequestIDSet{
				m: map[AuthRequestID]struct{}{MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
				s: []AuthRequestID{MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1")},
			},
		},
		{
			name: "2 non Empty Set",
			input: struct {
				a *AuthRequestIDSet
				b *AuthRequestIDSet
			}{
				a: &AuthRequestIDSet{
					m: map[AuthRequestID]struct{}{MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1"): {}},
					s: []AuthRequestID{MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1")},
				},
				b: &AuthRequestIDSet{
					m: map[AuthRequestID]struct{}{MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t2"): {}},
					s: []AuthRequestID{MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t2")},
				},
			},
			expected: &AuthRequestIDSet{
				m: map[AuthRequestID]struct{}{
					MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1"): {},
					MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t2"): {},
				},
				s: []AuthRequestID{
					MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t1"),
					MustAuthRequestID("01f3zhcaq35403zdjnd6dcm0t2"),
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