// Code generated by gen, DO NOT EDIT.

package id

import "encoding/json"

// LayerID is an ID for Layer.
type LayerID ID

// NewLayerID generates a new LayerId.
func NewLayerID() LayerID {
	return LayerID(New())
}

// LayerIDFrom generates a new LayerID from a string.
func LayerIDFrom(i string) (nid LayerID, err error) {
	var did ID
	did, err = FromID(i)
	if err != nil {
		return
	}
	nid = LayerID(did)
	return
}

// MustLayerID generates a new LayerID from a string, but panics if the string cannot be parsed.
func MustLayerID(i string) LayerID {
	did, err := FromID(i)
	if err != nil {
		panic(err)
	}
	return LayerID(did)
}

// LayerIDFromRef generates a new LayerID from a string ref.
func LayerIDFromRef(i *string) *LayerID {
	did := FromIDRef(i)
	if did == nil {
		return nil
	}
	nid := LayerID(*did)
	return &nid
}

// LayerIDFromRefID generates a new LayerID from a ref of a generic ID.
func LayerIDFromRefID(i *ID) *LayerID {
	if i == nil {
		return nil
	}
	nid := LayerID(*i)
	return &nid
}

// ID returns a domain ID.
func (d LayerID) ID() ID {
	return ID(d)
}

// String returns a string representation.
func (d LayerID) String() string {
	return ID(d).String()
}

// GoString implements fmt.GoStringer interface.
func (d LayerID) GoString() string {
	return "id.LayerID(" + d.String() + ")"
}

// RefString returns a reference of string representation.
func (d LayerID) RefString() *string {
	id := ID(d).String()
	return &id
}

// Ref returns a reference.
func (d LayerID) Ref() *LayerID {
	d2 := d
	return &d2
}

// CopyRef returns a copy of a reference.
func (d *LayerID) CopyRef() *LayerID {
	if d == nil {
		return nil
	}
	d2 := *d
	return &d2
}

// IDRef returns a reference of a domain id.
func (d *LayerID) IDRef() *ID {
	if d == nil {
		return nil
	}
	id := ID(*d)
	return &id
}

// StringRef returns a reference of a string representation.
func (d *LayerID) StringRef() *string {
	if d == nil {
		return nil
	}
	id := ID(*d).String()
	return &id
}

// MarhsalJSON implements json.Marhsaler interface
func (d *LayerID) MarhsalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

// UnmarhsalJSON implements json.Unmarshaler interface
func (d *LayerID) UnmarhsalJSON(bs []byte) (err error) {
	var idstr string
	if err = json.Unmarshal(bs, &idstr); err != nil {
		return
	}
	*d, err = LayerIDFrom(idstr)
	return
}

// MarshalText implements encoding.TextMarshaler interface
func (d *LayerID) MarshalText() ([]byte, error) {
	if d == nil {
		return nil, nil
	}
	return []byte(d.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler interface
func (d *LayerID) UnmarshalText(text []byte) (err error) {
	*d, err = LayerIDFrom(string(text))
	return
}

// Ref returns true if a ID is nil or zero-value
func (d LayerID) IsNil() bool {
	return ID(d).IsNil()
}

// LayerIDToKeys converts IDs into a string slice.
func LayerIDToKeys(ids []LayerID) []string {
	keys := make([]string, 0, len(ids))
	for _, i := range ids {
		keys = append(keys, i.String())
	}
	return keys
}

// LayerIDsFrom converts a string slice into a ID slice.
func LayerIDsFrom(ids []string) ([]LayerID, error) {
	dids := make([]LayerID, 0, len(ids))
	for _, i := range ids {
		did, err := LayerIDFrom(i)
		if err != nil {
			return nil, err
		}
		dids = append(dids, did)
	}
	return dids, nil
}

// LayerIDsFromID converts a generic ID slice into a ID slice.
func LayerIDsFromID(ids []ID) []LayerID {
	dids := make([]LayerID, 0, len(ids))
	for _, i := range ids {
		dids = append(dids, LayerID(i))
	}
	return dids
}

// LayerIDsFromIDRef converts a ref of a generic ID slice into a ID slice.
func LayerIDsFromIDRef(ids []*ID) []LayerID {
	dids := make([]LayerID, 0, len(ids))
	for _, i := range ids {
		if i != nil {
			dids = append(dids, LayerID(*i))
		}
	}
	return dids
}

// LayerIDsToID converts a ID slice into a generic ID slice.
func LayerIDsToID(ids []LayerID) []ID {
	dids := make([]ID, 0, len(ids))
	for _, i := range ids {
		dids = append(dids, i.ID())
	}
	return dids
}

// LayerIDsToIDRef converts a ID ref slice into a generic ID ref slice.
func LayerIDsToIDRef(ids []*LayerID) []*ID {
	dids := make([]*ID, 0, len(ids))
	for _, i := range ids {
		dids = append(dids, i.IDRef())
	}
	return dids
}

// LayerIDSet represents a set of LayerIDs
type LayerIDSet struct {
	m map[LayerID]struct{}
	s []LayerID
}

// NewLayerIDSet creates a new LayerIDSet
func NewLayerIDSet() *LayerIDSet {
	return &LayerIDSet{}
}

// Add adds a new ID if it does not exists in the set
func (s *LayerIDSet) Add(p ...LayerID) {
	if s == nil || p == nil {
		return
	}
	if s.m == nil {
		s.m = map[LayerID]struct{}{}
	}
	for _, i := range p {
		if _, ok := s.m[i]; !ok {
			if s.s == nil {
				s.s = []LayerID{}
			}
			s.m[i] = struct{}{}
			s.s = append(s.s, i)
		}
	}
}

// AddRef adds a new ID ref if it does not exists in the set
func (s *LayerIDSet) AddRef(p *LayerID) {
	if s == nil || p == nil {
		return
	}
	s.Add(*p)
}

// Has checks if the ID exists in the set
func (s *LayerIDSet) Has(p LayerID) bool {
	if s == nil || s.m == nil {
		return false
	}
	_, ok := s.m[p]
	return ok
}

// Clear clears all stored IDs
func (s *LayerIDSet) Clear() {
	if s == nil {
		return
	}
	s.m = nil
	s.s = nil
}

// All returns stored all IDs as a slice
func (s *LayerIDSet) All() []LayerID {
	if s == nil {
		return nil
	}
	return append([]LayerID{}, s.s...)
}

// Clone returns a cloned set
func (s *LayerIDSet) Clone() *LayerIDSet {
	if s == nil {
		return NewLayerIDSet()
	}
	s2 := NewLayerIDSet()
	s2.Add(s.s...)
	return s2
}

// Merge returns a merged set
func (s *LayerIDSet) Merge(s2 *LayerIDSet) *LayerIDSet {
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