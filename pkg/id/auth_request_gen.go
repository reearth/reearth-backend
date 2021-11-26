// Code generated by gen, DO NOT EDIT.

package id

import "encoding/json"

// AuthRequestID is an ID for AuthRequest.
type AuthRequestID ID

// NewAuthRequestID generates a new AuthRequestId.
func NewAuthRequestID() AuthRequestID {
	return AuthRequestID(New())
}

// AuthRequestIDFrom generates a new AuthRequestID from a string.
func AuthRequestIDFrom(i string) (nid AuthRequestID, err error) {
	var did ID
	did, err = FromID(i)
	if err != nil {
		return
	}
	nid = AuthRequestID(did)
	return
}

// MustAuthRequestID generates a new AuthRequestID from a string, but panics if the string cannot be parsed.
func MustAuthRequestID(i string) AuthRequestID {
	did, err := FromID(i)
	if err != nil {
		panic(err)
	}
	return AuthRequestID(did)
}

// AuthRequestIDFromRef generates a new AuthRequestID from a string ref.
func AuthRequestIDFromRef(i *string) *AuthRequestID {
	did := FromIDRef(i)
	if did == nil {
		return nil
	}
	nid := AuthRequestID(*did)
	return &nid
}

// AuthRequestIDFromRefID generates a new AuthRequestID from a ref of a generic ID.
func AuthRequestIDFromRefID(i *ID) *AuthRequestID {
	if i == nil {
		return nil
	}
	nid := AuthRequestID(*i)
	return &nid
}

// ID returns a domain ID.
func (d AuthRequestID) ID() ID {
	return ID(d)
}

// String returns a string representation.
func (d AuthRequestID) String() string {
	return ID(d).String()
}

// GoString implements fmt.GoStringer interface.
func (d AuthRequestID) GoString() string {
	return "id.AuthRequestID(" + d.String() + ")"
}

// RefString returns a reference of string representation.
func (d AuthRequestID) RefString() *string {
	id := ID(d).String()
	return &id
}

// Ref returns a reference.
func (d AuthRequestID) Ref() *AuthRequestID {
	d2 := d
	return &d2
}

// Contains returns whether the id is contained in the slice.
func (d AuthRequestID) Contains(ids []AuthRequestID) bool {
	for _, i := range ids {
		if d.ID().Equal(i.ID()) {
			return true
		}
	}
	return false
}

// CopyRef returns a copy of a reference.
func (d *AuthRequestID) CopyRef() *AuthRequestID {
	if d == nil {
		return nil
	}
	d2 := *d
	return &d2
}

// IDRef returns a reference of a domain id.
func (d *AuthRequestID) IDRef() *ID {
	if d == nil {
		return nil
	}
	id := ID(*d)
	return &id
}

// StringRef returns a reference of a string representation.
func (d *AuthRequestID) StringRef() *string {
	if d == nil {
		return nil
	}
	id := ID(*d).String()
	return &id
}

// MarhsalJSON implements json.Marhsaler interface
func (d *AuthRequestID) MarhsalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

// UnmarhsalJSON implements json.Unmarshaler interface
func (d *AuthRequestID) UnmarhsalJSON(bs []byte) (err error) {
	var idstr string
	if err = json.Unmarshal(bs, &idstr); err != nil {
		return
	}
	*d, err = AuthRequestIDFrom(idstr)
	return
}

// MarshalText implements encoding.TextMarshaler interface
func (d *AuthRequestID) MarshalText() ([]byte, error) {
	if d == nil {
		return nil, nil
	}
	return []byte(d.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler interface
func (d *AuthRequestID) UnmarshalText(text []byte) (err error) {
	*d, err = AuthRequestIDFrom(string(text))
	return
}

// Ref returns true if a ID is nil or zero-value
func (d AuthRequestID) IsNil() bool {
	return ID(d).IsNil()
}

// AuthRequestIDToKeys converts IDs into a string slice.
func AuthRequestIDToKeys(ids []AuthRequestID) []string {
	keys := make([]string, 0, len(ids))
	for _, i := range ids {
		keys = append(keys, i.String())
	}
	return keys
}

// AuthRequestIDsFrom converts a string slice into a ID slice.
func AuthRequestIDsFrom(ids []string) ([]AuthRequestID, error) {
	dids := make([]AuthRequestID, 0, len(ids))
	for _, i := range ids {
		did, err := AuthRequestIDFrom(i)
		if err != nil {
			return nil, err
		}
		dids = append(dids, did)
	}
	return dids, nil
}

// AuthRequestIDsFromID converts a generic ID slice into a ID slice.
func AuthRequestIDsFromID(ids []ID) []AuthRequestID {
	dids := make([]AuthRequestID, 0, len(ids))
	for _, i := range ids {
		dids = append(dids, AuthRequestID(i))
	}
	return dids
}

// AuthRequestIDsFromIDRef converts a ref of a generic ID slice into a ID slice.
func AuthRequestIDsFromIDRef(ids []*ID) []AuthRequestID {
	dids := make([]AuthRequestID, 0, len(ids))
	for _, i := range ids {
		if i != nil {
			dids = append(dids, AuthRequestID(*i))
		}
	}
	return dids
}

// AuthRequestIDsToID converts a ID slice into a generic ID slice.
func AuthRequestIDsToID(ids []AuthRequestID) []ID {
	dids := make([]ID, 0, len(ids))
	for _, i := range ids {
		dids = append(dids, i.ID())
	}
	return dids
}

// AuthRequestIDsToIDRef converts a ID ref slice into a generic ID ref slice.
func AuthRequestIDsToIDRef(ids []*AuthRequestID) []*ID {
	dids := make([]*ID, 0, len(ids))
	for _, i := range ids {
		dids = append(dids, i.IDRef())
	}
	return dids
}

// AuthRequestIDSet represents a set of AuthRequestIDs
type AuthRequestIDSet struct {
	m map[AuthRequestID]struct{}
	s []AuthRequestID
}

// NewAuthRequestIDSet creates a new AuthRequestIDSet
func NewAuthRequestIDSet() *AuthRequestIDSet {
	return &AuthRequestIDSet{}
}

// Add adds a new ID if it does not exists in the set
func (s *AuthRequestIDSet) Add(p ...AuthRequestID) {
	if s == nil || p == nil {
		return
	}
	if s.m == nil {
		s.m = map[AuthRequestID]struct{}{}
	}
	for _, i := range p {
		if _, ok := s.m[i]; !ok {
			if s.s == nil {
				s.s = []AuthRequestID{}
			}
			s.m[i] = struct{}{}
			s.s = append(s.s, i)
		}
	}
}

// AddRef adds a new ID ref if it does not exists in the set
func (s *AuthRequestIDSet) AddRef(p *AuthRequestID) {
	if s == nil || p == nil {
		return
	}
	s.Add(*p)
}

// Has checks if the ID exists in the set
func (s *AuthRequestIDSet) Has(p AuthRequestID) bool {
	if s == nil || s.m == nil {
		return false
	}
	_, ok := s.m[p]
	return ok
}

// Clear clears all stored IDs
func (s *AuthRequestIDSet) Clear() {
	if s == nil {
		return
	}
	s.m = nil
	s.s = nil
}

// All returns stored all IDs as a slice
func (s *AuthRequestIDSet) All() []AuthRequestID {
	if s == nil {
		return nil
	}
	return append([]AuthRequestID{}, s.s...)
}

// Clone returns a cloned set
func (s *AuthRequestIDSet) Clone() *AuthRequestIDSet {
	if s == nil {
		return NewAuthRequestIDSet()
	}
	s2 := NewAuthRequestIDSet()
	s2.Add(s.s...)
	return s2
}

// Merge returns a merged set
func (s *AuthRequestIDSet) Merge(s2 *AuthRequestIDSet) *AuthRequestIDSet {
	if s == nil {
		return nil
	}
	s3 := s.Clone()
	if s2 == nil {
		return s3
	}
	s3.Add(s2.s...)
	return s3
}
