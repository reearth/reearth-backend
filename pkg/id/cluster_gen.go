// Code generated by gen, DO NOT EDIT.

package id

import "encoding/json"

// ClusterID is an ID for Cluster.
type ClusterID ID

// NewClusterID generates a new ClusterId.
func NewClusterID() ClusterID {
	return ClusterID(New())
}

// ClusterIDFrom generates a new ClusterID from a string.
func ClusterIDFrom(i string) (nid ClusterID, err error) {
	var did ID
	did, err = FromID(i)
	if err != nil {
		return
	}
	nid = ClusterID(did)
	return
}

// MustClusterID generates a new ClusterID from a string, but panics if the string cannot be parsed.
func MustClusterID(i string) ClusterID {
	did, err := FromID(i)
	if err != nil {
		panic(err)
	}
	return ClusterID(did)
}

// ClusterIDFromRef generates a new ClusterID from a string ref.
func ClusterIDFromRef(i *string) *ClusterID {
	did := FromIDRef(i)
	if did == nil {
		return nil
	}
	nid := ClusterID(*did)
	return &nid
}

// ClusterIDFromRefID generates a new ClusterID from a ref of a generic ID.
func ClusterIDFromRefID(i *ID) *ClusterID {
	if i == nil {
		return nil
	}
	nid := ClusterID(*i)
	return &nid
}

// ID returns a domain ID.
func (d ClusterID) ID() ID {
	return ID(d)
}

// String returns a string representation.
func (d ClusterID) String() string {
	return ID(d).String()
}

// GoString implements fmt.GoStringer interface.
func (d ClusterID) GoString() string {
	return "id.ClusterID(" + d.String() + ")"
}

// RefString returns a reference of string representation.
func (d ClusterID) RefString() *string {
	id := ID(d).String()
	return &id
}

// Ref returns a reference.
func (d ClusterID) Ref() *ClusterID {
	d2 := d
	return &d2
}

// Contains returns whether the id is contained in the slice.
func (d ClusterID) Contains(ids []ClusterID) bool {
	for _, i := range ids {
		if d.ID().Equal(i.ID()) {
			return true
		}
	}
	return false
}

// CopyRef returns a copy of a reference.
func (d *ClusterID) CopyRef() *ClusterID {
	if d == nil {
		return nil
	}
	d2 := *d
	return &d2
}

// IDRef returns a reference of a domain id.
func (d *ClusterID) IDRef() *ID {
	if d == nil {
		return nil
	}
	id := ID(*d)
	return &id
}

// StringRef returns a reference of a string representation.
func (d *ClusterID) StringRef() *string {
	if d == nil {
		return nil
	}
	id := ID(*d).String()
	return &id
}

// MarhsalJSON implements json.Marhsaler interface
func (d *ClusterID) MarhsalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

// UnmarhsalJSON implements json.Unmarshaler interface
func (d *ClusterID) UnmarhsalJSON(bs []byte) (err error) {
	var idstr string
	if err = json.Unmarshal(bs, &idstr); err != nil {
		return
	}
	*d, err = ClusterIDFrom(idstr)
	return
}

// MarshalText implements encoding.TextMarshaler interface
func (d *ClusterID) MarshalText() ([]byte, error) {
	if d == nil {
		return nil, nil
	}
	return []byte(d.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler interface
func (d *ClusterID) UnmarshalText(text []byte) (err error) {
	*d, err = ClusterIDFrom(string(text))
	return
}

// Ref returns true if a ID is nil or zero-value
func (d ClusterID) IsNil() bool {
	return ID(d).IsNil()
}

// ClusterIDToKeys converts IDs into a string slice.
func ClusterIDToKeys(ids []ClusterID) []string {
	keys := make([]string, 0, len(ids))
	for _, i := range ids {
		keys = append(keys, i.String())
	}
	return keys
}

// ClusterIDsFrom converts a string slice into a ID slice.
func ClusterIDsFrom(ids []string) ([]ClusterID, error) {
	dids := make([]ClusterID, 0, len(ids))
	for _, i := range ids {
		did, err := ClusterIDFrom(i)
		if err != nil {
			return nil, err
		}
		dids = append(dids, did)
	}
	return dids, nil
}

// ClusterIDsFromID converts a generic ID slice into a ID slice.
func ClusterIDsFromID(ids []ID) []ClusterID {
	dids := make([]ClusterID, 0, len(ids))
	for _, i := range ids {
		dids = append(dids, ClusterID(i))
	}
	return dids
}

// ClusterIDsFromIDRef converts a ref of a generic ID slice into a ID slice.
func ClusterIDsFromIDRef(ids []*ID) []ClusterID {
	dids := make([]ClusterID, 0, len(ids))
	for _, i := range ids {
		if i != nil {
			dids = append(dids, ClusterID(*i))
		}
	}
	return dids
}

// ClusterIDsToID converts a ID slice into a generic ID slice.
func ClusterIDsToID(ids []ClusterID) []ID {
	dids := make([]ID, 0, len(ids))
	for _, i := range ids {
		dids = append(dids, i.ID())
	}
	return dids
}

// ClusterIDsToIDRef converts a ID ref slice into a generic ID ref slice.
func ClusterIDsToIDRef(ids []*ClusterID) []*ID {
	dids := make([]*ID, 0, len(ids))
	for _, i := range ids {
		dids = append(dids, i.IDRef())
	}
	return dids
}

// ClusterIDSet represents a set of ClusterIDs
type ClusterIDSet struct {
	m map[ClusterID]struct{}
	s []ClusterID
}

// NewClusterIDSet creates a new ClusterIDSet
func NewClusterIDSet() *ClusterIDSet {
	return &ClusterIDSet{}
}

// Add adds a new ID if it does not exists in the set
func (s *ClusterIDSet) Add(p ...ClusterID) {
	if s == nil || p == nil {
		return
	}
	if s.m == nil {
		s.m = map[ClusterID]struct{}{}
	}
	for _, i := range p {
		if _, ok := s.m[i]; !ok {
			if s.s == nil {
				s.s = []ClusterID{}
			}
			s.m[i] = struct{}{}
			s.s = append(s.s, i)
		}
	}
}

// AddRef adds a new ID ref if it does not exists in the set
func (s *ClusterIDSet) AddRef(p *ClusterID) {
	if s == nil || p == nil {
		return
	}
	s.Add(*p)
}

// Has checks if the ID exists in the set
func (s *ClusterIDSet) Has(p ClusterID) bool {
	if s == nil || s.m == nil {
		return false
	}
	_, ok := s.m[p]
	return ok
}

// Clear clears all stored IDs
func (s *ClusterIDSet) Clear() {
	if s == nil {
		return
	}
	s.m = nil
	s.s = nil
}

// All returns stored all IDs as a slice
func (s *ClusterIDSet) All() []ClusterID {
	if s == nil {
		return nil
	}
	return append([]ClusterID{}, s.s...)
}

// Clone returns a cloned set
func (s *ClusterIDSet) Clone() *ClusterIDSet {
	if s == nil {
		return NewClusterIDSet()
	}
	s2 := NewClusterIDSet()
	s2.Add(s.s...)
	return s2
}

// Merge returns a merged set
func (s *ClusterIDSet) Merge(s2 *ClusterIDSet) *ClusterIDSet {
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