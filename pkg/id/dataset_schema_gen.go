// Code generated by gen, DO NOT EDIT.

package id

import "encoding/json"

// DatasetSchemaID is an ID for DatasetSchema.
type DatasetSchemaID ID

// NewDatasetSchemaID generates a new DatasetSchemaId.
func NewDatasetSchemaID() DatasetSchemaID {
	return DatasetSchemaID(New())
}

// DatasetSchemaIDFrom generates a new DatasetSchemaID from a string.
func DatasetSchemaIDFrom(i string) (nid DatasetSchemaID, err error) {
	var did ID
	did, err = FromID(i)
	if err != nil {
		return
	}
	nid = DatasetSchemaID(did)
	return
}

// MustDatasetSchemaID generates a new DatasetSchemaID from a string, but panics if the string cannot be parsed.
func MustDatasetSchemaID(i string) DatasetSchemaID {
	did, err := FromID(i)
	if err != nil {
		panic(err)
	}
	return DatasetSchemaID(did)
}

// DatasetSchemaIDFromRef generates a new DatasetSchemaID from a string ref.
func DatasetSchemaIDFromRef(i *string) *DatasetSchemaID {
	did := FromIDRef(i)
	if did == nil {
		return nil
	}
	nid := DatasetSchemaID(*did)
	return &nid
}

// DatasetSchemaIDFromRefID generates a new DatasetSchemaID from a ref of a generic ID.
func DatasetSchemaIDFromRefID(i *ID) *DatasetSchemaID {
	if i == nil || i.IsNil() {
		return nil
	}
	nid := DatasetSchemaID(*i)
	return &nid
}

// ID returns a domain ID.
func (d DatasetSchemaID) ID() ID {
	return ID(d)
}

// String returns a string representation.
func (d DatasetSchemaID) String() string {
	return ID(d).String()
}

// GoString implements fmt.GoStringer interface.
func (d DatasetSchemaID) GoString() string {
	return "id.DatasetSchemaID(" + d.String() + ")"
}

// RefString returns a reference of string representation.
func (d DatasetSchemaID) RefString() *string {
	if d.IsNil() {
		return nil
	}
	id := ID(d).String()
	return &id
}

// Ref returns a reference.
func (d DatasetSchemaID) Ref() *DatasetSchemaID {
	if d.IsNil() {
		return nil
	}
	d2 := d
	return &d2
}

// Contains returns whether the id is contained in the slice.
func (d DatasetSchemaID) Contains(ids []DatasetSchemaID) bool {
	for _, i := range ids {
		if d.ID().Equal(i.ID()) {
			return true
		}
	}
	return false
}

// CopyRef returns a copy of a reference.
func (d *DatasetSchemaID) CopyRef() *DatasetSchemaID {
	if d == nil || d.IsNil() {
		return nil
	}
	d2 := *d
	return &d2
}

// IDRef returns a reference of a domain id.
func (d *DatasetSchemaID) IDRef() *ID {
	if d == nil || d.IsNil() {
		return nil
	}
	id := ID(*d)
	return &id
}

// StringRef returns a reference of a string representation.
func (d *DatasetSchemaID) StringRef() *string {
	if d == nil || d.IsNil() {
		return nil
	}
	id := ID(*d).String()
	return &id
}

// MarhsalJSON implements json.Marhsaler interface
func (d *DatasetSchemaID) MarhsalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

// UnmarhsalJSON implements json.Unmarshaler interface
func (d *DatasetSchemaID) UnmarhsalJSON(bs []byte) (err error) {
	var idstr string
	if err = json.Unmarshal(bs, &idstr); err != nil {
		return
	}
	*d, err = DatasetSchemaIDFrom(idstr)
	return
}

// MarshalText implements encoding.TextMarshaler interface
func (d *DatasetSchemaID) MarshalText() ([]byte, error) {
	if d == nil {
		return nil, nil
	}
	return []byte(d.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler interface
func (d *DatasetSchemaID) UnmarshalText(text []byte) (err error) {
	*d, err = DatasetSchemaIDFrom(string(text))
	return
}

// Ref returns true if a ID is nil or zero-value
func (d DatasetSchemaID) IsNil() bool {
	return ID(d).IsNil()
}

// DatasetSchemaIDToKeys converts IDs into a string slice.
func DatasetSchemaIDToKeys(ids []DatasetSchemaID) []string {
	keys := make([]string, 0, len(ids))
	for _, i := range ids {
		keys = append(keys, i.String())
	}
	return keys
}

// DatasetSchemaIDsFrom converts a string slice into a ID slice.
func DatasetSchemaIDsFrom(ids []string) ([]DatasetSchemaID, error) {
	dids := make([]DatasetSchemaID, 0, len(ids))
	for _, i := range ids {
		did, err := DatasetSchemaIDFrom(i)
		if err != nil {
			return nil, err
		}
		dids = append(dids, did)
	}
	return dids, nil
}

// DatasetSchemaIDsFromID converts a generic ID slice into a ID slice.
func DatasetSchemaIDsFromID(ids []ID) []DatasetSchemaID {
	dids := make([]DatasetSchemaID, 0, len(ids))
	for _, i := range ids {
		dids = append(dids, DatasetSchemaID(i))
	}
	return dids
}

// DatasetSchemaIDsFromIDRef converts a ref of a generic ID slice into a ID slice.
func DatasetSchemaIDsFromIDRef(ids []*ID) []DatasetSchemaID {
	dids := make([]DatasetSchemaID, 0, len(ids))
	for _, i := range ids {
		if i != nil {
			dids = append(dids, DatasetSchemaID(*i))
		}
	}
	return dids
}

// DatasetSchemaIDsToID converts a ID slice into a generic ID slice.
func DatasetSchemaIDsToID(ids []DatasetSchemaID) []ID {
	dids := make([]ID, 0, len(ids))
	for _, i := range ids {
		dids = append(dids, i.ID())
	}
	return dids
}

// DatasetSchemaIDsToIDRef converts a ID ref slice into a generic ID ref slice.
func DatasetSchemaIDsToIDRef(ids []*DatasetSchemaID) []*ID {
	dids := make([]*ID, 0, len(ids))
	for _, i := range ids {
		dids = append(dids, i.IDRef())
	}
	return dids
}

// DatasetSchemaIDSet represents a set of DatasetSchemaIDs
type DatasetSchemaIDSet struct {
	m map[DatasetSchemaID]struct{}
	s []DatasetSchemaID
}

// NewDatasetSchemaIDSet creates a new DatasetSchemaIDSet
func NewDatasetSchemaIDSet() *DatasetSchemaIDSet {
	return &DatasetSchemaIDSet{}
}

// Add adds a new ID if it does not exists in the set
func (s *DatasetSchemaIDSet) Add(p ...DatasetSchemaID) {
	if s == nil || p == nil {
		return
	}
	if s.m == nil {
		s.m = map[DatasetSchemaID]struct{}{}
	}
	for _, i := range p {
		if _, ok := s.m[i]; !ok {
			if s.s == nil {
				s.s = []DatasetSchemaID{}
			}
			s.m[i] = struct{}{}
			s.s = append(s.s, i)
		}
	}
}

// AddRef adds a new ID ref if it does not exists in the set
func (s *DatasetSchemaIDSet) AddRef(p *DatasetSchemaID) {
	if s == nil || p == nil {
		return
	}
	s.Add(*p)
}

// Has checks if the ID exists in the set
func (s *DatasetSchemaIDSet) Has(p DatasetSchemaID) bool {
	if s == nil || s.m == nil {
		return false
	}
	_, ok := s.m[p]
	return ok
}

// Clear clears all stored IDs
func (s *DatasetSchemaIDSet) Clear() {
	if s == nil {
		return
	}
	s.m = nil
	s.s = nil
}

// All returns stored all IDs as a slice
func (s *DatasetSchemaIDSet) All() []DatasetSchemaID {
	if s == nil {
		return nil
	}
	return append([]DatasetSchemaID{}, s.s...)
}

// Clone returns a cloned set
func (s *DatasetSchemaIDSet) Clone() *DatasetSchemaIDSet {
	if s == nil {
		return NewDatasetSchemaIDSet()
	}
	s2 := NewDatasetSchemaIDSet()
	s2.Add(s.s...)
	return s2
}

// Merge returns a merged set
func (s *DatasetSchemaIDSet) Merge(s2 *DatasetSchemaIDSet) *DatasetSchemaIDSet {
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
