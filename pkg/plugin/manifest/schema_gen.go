package manifest

// generated by "/var/folders/lz/nhqy382n28g31wb4f_40gbmc0000gp/T/go-build3270329547/b001/exe/schematyper -o schema_gen.go --package manifest ../../../plugin_manifest_schema.json" -- DO NOT EDIT

type Choice struct {
	Icon  string `json:"icon,omitempty"`
	Key   string `json:"key"`
	Label string `json:"label,omitempty"`
}

type Extension struct {
	Description *string         `json:"description,omitempty"`
	ID          ID              `json:"id"`
	Icon        *string         `json:"icon,omitempty"`
	Schema      *PropertySchema `json:"schema,omitempty"`
	Title       string          `json:"title"`
	Type        string          `json:"type"`
	Visualizer  string          `json:"visualizer"`
}

type ID string

type Id string

type PropertyCondition struct {
	Field string      `json:"field"`
	Type  Valuetype   `json:"type"`
	Value interface{} `json:"value"`
}

type PropertyLinkableFields struct {
	Latlng *PropertyPointer `json:"latlng,omitempty"`
	URL    *PropertyPointer `json:"url,omitempty"`
}

type PropertyPointer struct {
	FieldID       string `json:"fieldId"`
	SchemaGroupID string `json:"schemaGroupId"`
}

type PropertySchema struct {
	Groups   []PropertySchemaGroup   `json:"groups,omitempty"`
	Linkable *PropertyLinkableFields `json:"linkable,omitempty"`
	Version  float64                 `json:"version,omitempty"`
}

type PropertySchemaField struct {
	AvailableIf  *PropertyCondition `json:"availableIf,omitempty"`
	Choices      []Choice           `json:"choices,omitempty"`
	DefaultValue interface{}        `json:"defaultValue,omitempty"`
	Description  *string            `json:"description,omitempty"`
	ID           ID                 `json:"id"`
	Max          *float64           `json:"max,omitempty"`
	Min          *float64           `json:"min,omitempty"`
	Prefix       *string            `json:"prefix,omitempty"`
	Suffix       *string            `json:"suffix,omitempty"`
	Title        *string            `json:"title"`
	Type         Valuetype          `json:"type"`
	UI           *string            `json:"ui,omitempty"`
}

type PropertySchemaGroup struct {
	AvailableIf         *PropertyCondition    `json:"availableIf,omitempty"`
	Description         *string               `json:"description,omitempty"`
	Fields              []PropertySchemaField `json:"fields,omitempty"`
	ID                  ID                    `json:"id"`
	List                bool                  `json:"list,omitempty"`
	RepresentativeField *Id                   `json:"representativeField,omitempty"`
	Title               string                `json:"title"`
}

type Root struct {
	Author      *string         `json:"author,omitempty"`
	Description *string         `json:"description,omitempty"`
	Extensions  []Extension     `json:"extensions,omitempty"`
	ID          ID              `json:"id"`
	Main        *string         `json:"main,omitempty"`
	Repository  *string         `json:"repository,omitempty"`
	Schema      *PropertySchema `json:"schema,omitempty"`
	System      bool            `json:"system,omitempty"`
	Title       string          `json:"title"`
	Version     string          `json:"version,omitempty"`
}

type Valuetype string
