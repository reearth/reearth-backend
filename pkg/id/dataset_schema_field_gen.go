// Code generated by gen, DO NOT EDIT.

package id

import "encoding/json"

// DatasetSchemaFieldID is an ID for DatasetSchemaField.
type DatasetSchemaFieldID ID

// NewDatasetSchemaFieldID generates a new DatasetSchemaFieldId.
func NewDatasetSchemaFieldID() DatasetSchemaFieldID {
	return DatasetSchemaFieldID(New())
}

// DatasetSchemaFieldIDFrom generates a new DatasetSchemaFieldID from a string.
func DatasetSchemaFieldIDFrom(i string) (nid DatasetSchemaFieldID, err error) {
	var did ID
	did, err = FromID(i)
	if err != nil {
		return
	}
	nid = DatasetSchemaFieldID(did)
	return
}

// MustDatasetSchemaFieldID generates a new DatasetSchemaFieldID from a string, but panics if the string cannot be parsed.
func MustDatasetSchemaFieldID(i string) DatasetSchemaFieldID {
	did, err := FromID(i)
	if err != nil {
		panic(err)
	}
	return DatasetSchemaFieldID(did)
}

// DatasetSchemaFieldIDFromRef generates a new DatasetSchemaFieldID from a string ref.
func DatasetSchemaFieldIDFromRef(i *string) *DatasetSchemaFieldID {
	did := FromIDRef(i)
	if did == nil {
		return nil
	}
	nid := DatasetSchemaFieldID(*did)
	return &nid
}

// DatasetSchemaFieldIDFromRefID generates a new DatasetSchemaFieldID from a ref of a generic ID.
func DatasetSchemaFieldIDFromRefID(i *ID) *DatasetSchemaFieldID {
	if i == nil {
		return nil
	}
	nid := DatasetSchemaFieldID(*i)
	return &nid
}

// ID returns a domain ID.
func (d DatasetSchemaFieldID) ID() ID {
	return ID(d)
}

// String returns a string representation.
func (d DatasetSchemaFieldID) String() string {
	return ID(d).String()
}

// GoString implements fmt.GoStringer interface.
func (d DatasetSchemaFieldID) GoString() string {
	return "id.DatasetSchemaFieldID(" + d.String() + ")"
}

// RefString returns a reference of string representation.
func (d DatasetSchemaFieldID) RefString() *string {
	id := ID(d).String()
	return &id
}

// Ref returns a reference.
func (d DatasetSchemaFieldID) Ref() *DatasetSchemaFieldID {
	d2 := d
	return &d2
}

// Contains returns whether the id is contained in the slice.
func (d DatasetSchemaFieldID) Contains(ids []DatasetSchemaFieldID) bool {
	for _, i := range ids {
		if d.ID().Equal(i.ID()) {
			return true
		}
	}
	return false
}

// CopyRef returns a copy of a reference.
func (d *DatasetSchemaFieldID) CopyRef() *DatasetSchemaFieldID {
	if d == nil {
		return nil
	}
	d2 := *d
	return &d2
}

// IDRef returns a reference of a domain id.
func (d *DatasetSchemaFieldID) IDRef() *ID {
	if d == nil {
		return nil
	}
	id := ID(*d)
	return &id
}

// StringRef returns a reference of a string representation.
func (d *DatasetSchemaFieldID) StringRef() *string {
	if d == nil {
		return nil
	}
	id := ID(*d).String()
	return &id
}

// MarhsalJSON implements json.Marhsaler interface
func (d *DatasetSchemaFieldID) MarhsalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

// UnmarhsalJSON implements json.Unmarshaler interface
func (d *DatasetSchemaFieldID) UnmarhsalJSON(bs []byte) (err error) {
	var idstr string
	if err = json.Unmarshal(bs, &idstr); err != nil {
		return
	}
	*d, err = DatasetSchemaFieldIDFrom(idstr)
	return
}

// MarshalText implements encoding.TextMarshaler interface
func (d *DatasetSchemaFieldID) MarshalText() ([]byte, error) {
	if d == nil {
		return nil, nil
	}
	return []byte(d.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler interface
func (d *DatasetSchemaFieldID) UnmarshalText(text []byte) (err error) {
	*d, err = DatasetSchemaFieldIDFrom(string(text))
	return
}

// Ref returns true if a ID is nil or zero-value
func (d DatasetSchemaFieldID) IsNil() bool {
	return ID(d).IsNil()
}

// DatasetSchemaFieldIDToKeys converts IDs into a string slice.
func DatasetSchemaFieldIDToKeys(ids []DatasetSchemaFieldID) []string {
	keys := make([]string, 0, len(ids))
	for _, i := range ids {
		keys = append(keys, i.String())
	}
	return keys
}

// DatasetSchemaFieldIDsFrom converts a string slice into a ID slice.
func DatasetSchemaFieldIDsFrom(ids []string) ([]DatasetSchemaFieldID, error) {
	dids := make([]DatasetSchemaFieldID, 0, len(ids))
	for _, i := range ids {
		did, err := DatasetSchemaFieldIDFrom(i)
		if err != nil {
			return nil, err
		}
		dids = append(dids, did)
	}
	return dids, nil
}

// DatasetSchemaFieldIDsFromID converts a generic ID slice into a ID slice.
func DatasetSchemaFieldIDsFromID(ids []ID) []DatasetSchemaFieldID {
	dids := make([]DatasetSchemaFieldID, 0, len(ids))
	for _, i := range ids {
		dids = append(dids, DatasetSchemaFieldID(i))
	}
	return dids
}

// DatasetSchemaFieldIDsFromIDRef converts a ref of a generic ID slice into a ID slice.
func DatasetSchemaFieldIDsFromIDRef(ids []*ID) []DatasetSchemaFieldID {
	dids := make([]DatasetSchemaFieldID, 0, len(ids))
	for _, i := range ids {
		if i != nil {
			dids = append(dids, DatasetSchemaFieldID(*i))
		}
	}
	return dids
}

// DatasetSchemaFieldIDsToID converts a ID slice into a generic ID slice.
func DatasetSchemaFieldIDsToID(ids []DatasetSchemaFieldID) []ID {
	dids := make([]ID, 0, len(ids))
	for _, i := range ids {
		dids = append(dids, i.ID())
	}
	return dids
}

// DatasetSchemaFieldIDsToIDRef converts a ID ref slice into a generic ID ref slice.
func DatasetSchemaFieldIDsToIDRef(ids []*DatasetSchemaFieldID) []*ID {
	dids := make([]*ID, 0, len(ids))
	for _, i := range ids {
		dids = append(dids, i.IDRef())
	}
	return dids
}

// DatasetSchemaFieldIDSet represents a set of DatasetSchemaFieldIDs
type DatasetSchemaFieldIDSet struct {
	m map[DatasetSchemaFieldID]struct{}
	s []DatasetSchemaFieldID
}

// NewDatasetSchemaFieldIDSet creates a new DatasetSchemaFieldIDSet
func NewDatasetSchemaFieldIDSet() *DatasetSchemaFieldIDSet {
	return &DatasetSchemaFieldIDSet{}
}

// Add adds a new ID if it does not exists in the set
func (s *DatasetSchemaFieldIDSet) Add(p ...DatasetSchemaFieldID) {
	if s == nil || p == nil {
		return
	}
	if s.m == nil {
		s.m = map[DatasetSchemaFieldID]struct{}{}
	}
	for _, i := range p {
		if _, ok := s.m[i]; !ok {
			if s.s == nil {
				s.s = []DatasetSchemaFieldID{}
			}
			s.m[i] = struct{}{}
			s.s = append(s.s, i)
		}
	}
}

// AddRef adds a new ID ref if it does not exists in the set
func (s *DatasetSchemaFieldIDSet) AddRef(p *DatasetSchemaFieldID) {
	if s == nil || p == nil {
		return
	}
	s.Add(*p)
}

// Has checks if the ID exists in the set
func (s *DatasetSchemaFieldIDSet) Has(p DatasetSchemaFieldID) bool {
	if s == nil || s.m == nil {
		return false
	}
	_, ok := s.m[p]
	return ok
}

// Clear clears all stored IDs
func (s *DatasetSchemaFieldIDSet) Clear() {
	if s == nil {
		return
	}
	s.m = nil
	s.s = nil
}

// All returns stored all IDs as a slice
func (s *DatasetSchemaFieldIDSet) All() []DatasetSchemaFieldID {
	if s == nil {
		return nil
	}
	return append([]DatasetSchemaFieldID{}, s.s...)
}

// Clone returns a cloned set
func (s *DatasetSchemaFieldIDSet) Clone() *DatasetSchemaFieldIDSet {
	if s == nil {
		return NewDatasetSchemaFieldIDSet()
	}
	s2 := NewDatasetSchemaFieldIDSet()
	s2.Add(s.s...)
	return s2
}

// Merge returns a merged set
func (s *DatasetSchemaFieldIDSet) Merge(s2 *DatasetSchemaFieldIDSet) *DatasetSchemaFieldIDSet {
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
