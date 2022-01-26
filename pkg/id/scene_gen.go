// Code generated by gen, DO NOT EDIT.

package id

import "encoding/json"

// SceneID is an ID for Scene.
type SceneID ID

// NewSceneID generates a new SceneId.
func NewSceneID() SceneID {
	return SceneID(New())
}

// SceneIDFrom generates a new SceneID from a string.
func SceneIDFrom(i string) (nid SceneID, err error) {
	var did ID
	did, err = FromID(i)
	if err != nil {
		return
	}
	nid = SceneID(did)
	return
}

// MustSceneID generates a new SceneID from a string, but panics if the string cannot be parsed.
func MustSceneID(i string) SceneID {
	did, err := FromID(i)
	if err != nil {
		panic(err)
	}
	return SceneID(did)
}

// SceneIDFromRef generates a new SceneID from a string ref.
func SceneIDFromRef(i *string) *SceneID {
	did := FromIDRef(i)
	if did == nil {
		return nil
	}
	nid := SceneID(*did)
	return &nid
}

// SceneIDFromRefID generates a new SceneID from a ref of a generic ID.
func SceneIDFromRefID(i *ID) *SceneID {
	if i == nil || i.IsNil() {
		return nil
	}
	nid := SceneID(*i)
	return &nid
}

// ID returns a domain ID.
func (d SceneID) ID() ID {
	return ID(d)
}

// String returns a string representation.
func (d SceneID) String() string {
	if d.IsNil() {
		return ""
	}
	return ID(d).String()
}

// StringRef returns a reference of the string representation.
func (d SceneID) RefString() *string {
	if d.IsNil() {
		return nil
	}
	str := d.String()
	return &str
}

// GoString implements fmt.GoStringer interface.
func (d SceneID) GoString() string {
	return "SceneID(" + d.String() + ")"
}

// Ref returns a reference.
func (d SceneID) Ref() *SceneID {
	if d.IsNil() {
		return nil
	}
	d2 := d
	return &d2
}

// Contains returns whether the id is contained in the slice.
func (d SceneID) Contains(ids []SceneID) bool {
	if d.IsNil() {
		return false
	}
	for _, i := range ids {
		if d.ID().Equal(i.ID()) {
			return true
		}
	}
	return false
}

// CopyRef returns a copy of a reference.
func (d *SceneID) CopyRef() *SceneID {
	if d.IsNilRef() {
		return nil
	}
	d2 := *d
	return &d2
}

// IDRef returns a reference of a domain id.
func (d *SceneID) IDRef() *ID {
	if d.IsNilRef() {
		return nil
	}
	id := ID(*d)
	return &id
}

// StringRef returns a reference of a string representation.
func (d *SceneID) StringRef() *string {
	if d.IsNilRef() {
		return nil
	}
	id := ID(*d).String()
	return &id
}

// MarhsalJSON implements json.Marhsaler interface
func (d *SceneID) MarhsalJSON() ([]byte, error) {
	if d.IsNilRef() {
		return nil, nil
	}
	return json.Marshal(d.String())
}

// UnmarhsalJSON implements json.Unmarshaler interface
func (d *SceneID) UnmarhsalJSON(bs []byte) (err error) {
	var idstr string
	if err = json.Unmarshal(bs, &idstr); err != nil {
		return
	}
	*d, err = SceneIDFrom(idstr)
	return
}

// MarshalText implements encoding.TextMarshaler interface
func (d *SceneID) MarshalText() ([]byte, error) {
	if d.IsNilRef() {
		return nil, nil
	}
	return []byte(d.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler interface
func (d *SceneID) UnmarshalText(text []byte) (err error) {
	*d, err = SceneIDFrom(string(text))
	return
}

// IsNil returns true if a ID is zero-value
func (d SceneID) IsNil() bool {
	return ID(d).IsNil()
}

// IsNilRef returns true if a ID is nil or zero-value
func (d *SceneID) IsNilRef() bool {
	return d == nil || ID(*d).IsNil()
}

// SceneIDsToStrings converts IDs into a string slice.
func SceneIDsToStrings(ids []SceneID) []string {
	strs := make([]string, 0, len(ids))
	for _, i := range ids {
		strs = append(strs, i.String())
	}
	return strs
}

// SceneIDsFrom converts a string slice into a ID slice.
func SceneIDsFrom(ids []string) ([]SceneID, error) {
	dids := make([]SceneID, 0, len(ids))
	for _, i := range ids {
		did, err := SceneIDFrom(i)
		if err != nil {
			return nil, err
		}
		dids = append(dids, did)
	}
	return dids, nil
}

// SceneIDsFromID converts a generic ID slice into a ID slice.
func SceneIDsFromID(ids []ID) []SceneID {
	dids := make([]SceneID, 0, len(ids))
	for _, i := range ids {
		dids = append(dids, SceneID(i))
	}
	return dids
}

// SceneIDsFromIDRef converts a ref of a generic ID slice into a ID slice.
func SceneIDsFromIDRef(ids []*ID) []SceneID {
	dids := make([]SceneID, 0, len(ids))
	for _, i := range ids {
		if i != nil {
			dids = append(dids, SceneID(*i))
		}
	}
	return dids
}

// SceneIDsToID converts a ID slice into a generic ID slice.
func SceneIDsToID(ids []SceneID) []ID {
	dids := make([]ID, 0, len(ids))
	for _, i := range ids {
		dids = append(dids, i.ID())
	}
	return dids
}

// SceneIDsToIDRef converts a ID ref slice into a generic ID ref slice.
func SceneIDsToIDRef(ids []*SceneID) []*ID {
	dids := make([]*ID, 0, len(ids))
	for _, i := range ids {
		dids = append(dids, i.IDRef())
	}
	return dids
}

// SceneIDSet represents a set of SceneIDs
type SceneIDSet struct {
	m map[SceneID]struct{}
	s []SceneID
}

// NewSceneIDSet creates a new SceneIDSet
func NewSceneIDSet() *SceneIDSet {
	return &SceneIDSet{}
}

// Add adds a new ID if it does not exists in the set
func (s *SceneIDSet) Add(p ...SceneID) {
	if s == nil || p == nil {
		return
	}
	if s.m == nil {
		s.m = map[SceneID]struct{}{}
	}
	for _, i := range p {
		if _, ok := s.m[i]; !ok {
			if s.s == nil {
				s.s = []SceneID{}
			}
			s.m[i] = struct{}{}
			s.s = append(s.s, i)
		}
	}
}

// AddRef adds a new ID ref if it does not exists in the set
func (s *SceneIDSet) AddRef(p *SceneID) {
	if s == nil || p == nil {
		return
	}
	s.Add(*p)
}

// Has checks if the ID exists in the set
func (s *SceneIDSet) Has(p SceneID) bool {
	if s == nil || s.m == nil {
		return false
	}
	_, ok := s.m[p]
	return ok
}

// Clear clears all stored IDs
func (s *SceneIDSet) Clear() {
	if s == nil {
		return
	}
	s.m = nil
	s.s = nil
}

// All returns stored all IDs as a slice
func (s *SceneIDSet) All() []SceneID {
	if s == nil {
		return nil
	}
	return append([]SceneID{}, s.s...)
}

// Clone returns a cloned set
func (s *SceneIDSet) Clone() *SceneIDSet {
	if s == nil {
		return NewSceneIDSet()
	}
	s2 := NewSceneIDSet()
	s2.Add(s.s...)
	return s2
}

// Merge returns a merged set
func (s *SceneIDSet) Merge(s2 *SceneIDSet) *SceneIDSet {
	s3 := s.Clone()
	if s2 == nil {
		return s3
	}
	s3.Add(s2.s...)
	return s3
}
