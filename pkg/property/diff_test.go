package property

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSchemaDiffFrom(t *testing.T) {
	type args struct {
		old *Schema
		new *Schema
	}
	tests := []struct {
		name string
		args args
		want SchemaDiff
	}{
		{
			name: "diff",
			args: args{
				old: &Schema{
					groups: &SchemaGroupList{groups: []*SchemaGroup{
						{id: "a", fields: []*SchemaField{
							{id: "aa", propertyType: ValueTypeString}, // deleted
							{id: "ab", propertyType: ValueTypeString},
							{id: "ac", propertyType: ValueTypeString},
							{id: "ad", propertyType: ValueTypeString},
						}},
					}},
				},
				new: &Schema{
					groups: &SchemaGroupList{groups: []*SchemaGroup{
						{id: "a", fields: []*SchemaField{
							{id: "ab", propertyType: ValueTypeNumber}, // type changed
							{id: "ae", propertyType: ValueTypeString}, // added
						}},
						{id: "b", list: true, fields: []*SchemaField{
							{id: "ac", propertyType: ValueTypeString}, // moved
							{id: "ad", propertyType: ValueTypeNumber}, // moved and type changed
						}},
					}},
				},
			},
			want: SchemaDiff{
				Deleted: []SchemaDiffDeleted{
					{SchemaGroup: "a", Field: "aa"},
				},
				Moved: []SchemaDiffMoved{
					{From: SchemaFieldPointer{SchemaGroup: "a", Field: "ac"}, To: SchemaFieldPointer{SchemaGroup: "b", Field: "ac"}, ToList: true},
					{From: SchemaFieldPointer{SchemaGroup: "a", Field: "ad"}, To: SchemaFieldPointer{SchemaGroup: "b", Field: "ad"}, ToList: true},
				},
				TypeChanged: []SchemaDiffTypeChanged{
					{SchemaFieldPointer: SchemaFieldPointer{SchemaGroup: "a", Field: "ab"}, NewType: ValueTypeNumber},
					{SchemaFieldPointer: SchemaFieldPointer{SchemaGroup: "b", Field: "ad"}, NewType: ValueTypeNumber},
				},
			},
		},
		{
			name: "no diff",
			args: args{
				old: &Schema{
					groups: &SchemaGroupList{groups: []*SchemaGroup{
						{id: "a", fields: []*SchemaField{
							{id: "aa", propertyType: ValueTypeNumber},
						}},
					}},
				},
				new: &Schema{
					groups: &SchemaGroupList{groups: []*SchemaGroup{
						{id: "a", fields: []*SchemaField{
							{id: "aa", propertyType: ValueTypeNumber},
						}},
						{id: "b", list: true, fields: []*SchemaField{
							{id: "ba", propertyType: ValueTypeString}, // added
						}},
					}},
				},
			},
			want: SchemaDiff{},
		},
		{
			name: "same schemas",
			args: args{
				old: testSchema1,
				new: testSchema1,
			},
			want: SchemaDiff{},
		},
		{
			name: "nil",
			args: args{
				old: nil,
				new: nil,
			},
			want: SchemaDiff{},
		},
		{
			name: "old nil",
			args: args{
				old: nil,
				new: testSchema1,
			},
			want: SchemaDiff{},
		},
		{
			name: "new nil",
			args: args{
				old: testSchema1,
				new: nil,
			},
			want: SchemaDiff{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, SchemaDiffFrom(tt.args.old, tt.args.new))
		})
	}
}

func TestSchemaDiffFromProperty(t *testing.T) {
	type args struct {
		old *Property
		new *Schema
	}
	tests := []struct {
		name string
		args args
		want SchemaDiff
	}{
		{
			name: "diff",
			args: args{
				old: testProperty1,
				new: &Schema{
					groups: &SchemaGroupList{groups: []*SchemaGroup{
						{id: testSchemaGroup1.ID(), fields: []*SchemaField{
							{id: testSchemaField1.ID(), propertyType: ValueTypeNumber}, // type changed
							{id: testSchemaField3.ID(), propertyType: ValueTypeNumber}, // moved and type changed
							{id: "xxxx", propertyType: ValueTypeString},                // added
						}},
						{id: testSchemaGroup2.ID(), list: true, fields: []*SchemaField{}},
					}},
				},
			},
			want: SchemaDiff{
				Deleted: nil,
				Moved: []SchemaDiffMoved{
					{
						From: SchemaFieldPointer{SchemaGroup: testSchemaGroup2.ID(), Field: testSchemaField3.ID()},
						To:   SchemaFieldPointer{SchemaGroup: testSchemaGroup1.ID(), Field: testSchemaField3.ID()},
					},
				},
				TypeChanged: []SchemaDiffTypeChanged{
					{SchemaFieldPointer: SchemaFieldPointer{SchemaGroup: testSchemaGroup1.ID(), Field: testSchemaField1.ID()}, NewType: ValueTypeNumber},
					{SchemaFieldPointer: SchemaFieldPointer{SchemaGroup: testSchemaGroup1.ID(), Field: testSchemaField3.ID()}, NewType: ValueTypeNumber},
				},
			},
		},
		{
			name: "no diff",
			args: args{
				old: testProperty1,
				new: testSchema1,
			},
			want: SchemaDiff{},
		},
		{
			name: "nil",
			args: args{
				old: nil,
				new: nil,
			},
			want: SchemaDiff{},
		},
		{
			name: "old nil",
			args: args{
				old: nil,
				new: testSchema1,
			},
			want: SchemaDiff{},
		},
		{
			name: "new nil",
			args: args{
				old: testProperty1,
				new: nil,
			},
			want: SchemaDiff{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, SchemaDiffFromProperty(tt.args.old, tt.args.new))
		})
	}
}

func TestSchemaDiff_Migrate(t *testing.T) {
	itemID := NewItemID()
	NewItemID = func() ItemID { return itemID }

	tests := []struct {
		name         string
		target       SchemaDiff
		args         *Property
		want         bool
		wantProperty *Property
	}{
		{
			name: "deleted and type changed",
			target: SchemaDiff{
				Deleted: []SchemaDiffDeleted{
					{SchemaGroup: testGroup1.SchemaGroup(), Field: testField1.Field()},
				},
				TypeChanged: []SchemaDiffTypeChanged{
					{SchemaFieldPointer: SchemaFieldPointer{SchemaGroup: testGroupList1.SchemaGroup(), Field: testField2.Field()}, NewType: ValueTypeString},
				},
			},
			args: testProperty1.Clone(),
			want: true,
			wantProperty: &Property{
				id:     testProperty1.ID(),
				scene:  testProperty1.Scene(),
				schema: testProperty1.Schema(),
				items: []Item{
					&Group{
						itemBase: itemBase{
							ID:          testGroup1.ID(),
							SchemaGroup: testGroup1.SchemaGroup(),
						},
						fields: []*Field{
							// deleted
						},
					},
					&GroupList{
						itemBase: itemBase{
							ID:          testGroupList1.ID(),
							SchemaGroup: testGroupList1.SchemaGroup(),
						},
						groups: []*Group{
							{
								itemBase: itemBase{
									ID:          testGroup2.ID(),
									SchemaGroup: testGroup2.SchemaGroup(),
								},
								fields: []*Field{
									{field: testField2.Field(), v: NewOptionalValue(ValueTypeString, nil)}, // type changed
								},
							},
						},
					},
				},
			},
		},
		{
			name: "moved",
			target: SchemaDiff{
				Moved: []SchemaDiffMoved{
					{
						From: SchemaFieldPointer{SchemaGroup: testGroup1.SchemaGroup(), Field: testField1.Field()},
						To:   SchemaFieldPointer{SchemaGroup: "x", Field: testField1.Field()},
					},
				},
			},
			args: testProperty1.Clone(),
			want: true,
			wantProperty: &Property{
				id:     testProperty1.ID(),
				scene:  testProperty1.Scene(),
				schema: testProperty1.Schema(),
				items: []Item{
					&Group{
						itemBase: itemBase{
							ID:          testGroup1.ID(),
							SchemaGroup: testGroup1.SchemaGroup(),
						},
						fields: []*Field{
							// deleted
						},
					},
					testGroupList1,
					&Group{
						itemBase: itemBase{
							ID:          itemID,
							SchemaGroup: "x",
						},
						fields: []*Field{testField1},
					},
				},
			},
		},
		{
			name: "moved and type changed",
			target: SchemaDiff{
				Moved: []SchemaDiffMoved{
					{
						From: SchemaFieldPointer{SchemaGroup: testGroup1.SchemaGroup(), Field: testField1.Field()},
						To:   SchemaFieldPointer{SchemaGroup: "x", Field: testField1.Field()},
					},
				},
				TypeChanged: []SchemaDiffTypeChanged{
					{SchemaFieldPointer: SchemaFieldPointer{SchemaGroup: "x", Field: testField1.Field()}, NewType: ValueTypeNumber},
				},
			},
			args: testProperty1.Clone(),
			want: true,
			wantProperty: &Property{
				id:     testProperty1.ID(),
				scene:  testProperty1.Scene(),
				schema: testProperty1.Schema(),
				items: []Item{
					&Group{
						itemBase: itemBase{
							ID:          testGroup1.ID(),
							SchemaGroup: testGroup1.SchemaGroup(),
						},
						fields: []*Field{
							// deleted
						},
					},
					testGroupList1,
					&Group{
						itemBase: itemBase{
							ID:          itemID,
							SchemaGroup: "x",
						},
						fields: []*Field{
							{field: testField1.Field(), v: NewOptionalValue(ValueTypeNumber, nil)},
						},
					},
				},
			},
		},
		{
			name: "group -> list",
			target: SchemaDiff{
				Moved: []SchemaDiffMoved{
					{
						From: SchemaFieldPointer{SchemaGroup: testGroup1.SchemaGroup(), Field: testField1.Field()},
						To:   SchemaFieldPointer{SchemaGroup: testGroup2.SchemaGroup(), Field: testField1.Field()},
					},
				},
			},
			args: testProperty1.Clone(),
			want: true,
			wantProperty: &Property{
				id:     testProperty1.ID(),
				scene:  testProperty1.Scene(),
				schema: testProperty1.Schema(),
				items: []Item{
					&Group{
						itemBase: itemBase{
							ID:          testGroup1.ID(),
							SchemaGroup: testGroup1.SchemaGroup(),
						},
						fields: []*Field{
							// deleted
						},
					},
					testGroupList1,
				},
			},
		},
		{
			name: "list -> group",
			target: SchemaDiff{
				Moved: []SchemaDiffMoved{
					{
						From: SchemaFieldPointer{SchemaGroup: testGroup2.SchemaGroup(), Field: testField2.Field()},
						To:   SchemaFieldPointer{SchemaGroup: testGroup1.SchemaGroup(), Field: testField2.Field()},
					},
				},
			},
			args: testProperty1.Clone(),
			want: true,
			wantProperty: &Property{
				id:     testProperty1.ID(),
				scene:  testProperty1.Scene(),
				schema: testProperty1.Schema(),
				items: []Item{
					testGroup1,
					&GroupList{
						itemBase: itemBase{
							ID:          testGroupList1.ID(),
							SchemaGroup: testGroupList1.SchemaGroup(),
						},
						groups: []*Group{
							{
								itemBase: itemBase{
									ID:          testGroup2.ID(),
									SchemaGroup: testGroup2.SchemaGroup(),
								},
								fields: []*Field{
									// deleted
								},
							},
						},
					},
				},
			},
		},
		{
			name:         "empty",
			target:       SchemaDiff{},
			args:         testProperty1,
			want:         false,
			wantProperty: testProperty1,
		},
		{
			name: "nil property",
			target: SchemaDiff{
				Deleted: []SchemaDiffDeleted{{SchemaGroup: testGroup1.SchemaGroup(), Field: testField1.Field()}},
			},
			args:         nil,
			want:         false,
			wantProperty: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.target.Migrate(tt.args))
			assert.Equal(t, tt.wantProperty, tt.args)
		})
	}
}
