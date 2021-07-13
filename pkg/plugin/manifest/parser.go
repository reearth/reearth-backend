package manifest

//go:generate go run github.com/idubinskiy/schematyper -o schema_gen.go --package manifest ../../../plugin_manifest_schema.json

import (
	_ "embed"
	"errors"
	"fmt"
	"io"

	"github.com/xeipuuv/gojsonschema"
	"gopkg.in/yaml.v2"
)

var (
	ErrInvalidManifest       error = errors.New("invalid manifest")
	ErrFailedToParseManifest error = errors.New("failed to parse plugin manifest")
	ErrSystemManifest              = errors.New("cannot build system manifest")
	//go:embed plugin_manifest_schema.json
	SchemaJSON   []byte
	schemaLoader = gojsonschema.NewBytesLoader(SchemaJSON)
)

func Parse(source io.Reader) (*Manifest, error) {
	// TODO: When using gojsonschema.NewReaderLoader, gojsonschema.Validate returns io.EOF error.
	// doc, err := io.ReadAll(source)
	// if err != nil {
	// 	return nil, ErrFailedToParseManifest
	// }

	// TODO: gojsonschema does not support yaml
	// documentLoader := gojsonschema.NewBytesLoader(doc)
	// if err := validate(documentLoader); err != nil {
	// 	return nil, err
	// }

	root := Root{}
	if err := yaml.NewDecoder(source).Decode(&root); err != nil {
		return nil, ErrFailedToParseManifest
		// return nil, fmt.Errorf("failed to parse plugin manifest: %w", err)
	}

	manifest, err := root.manifest()
	if err != nil {
		return nil, err
	}
	if manifest.Plugin.ID().System() {
		return nil, ErrSystemManifest
	}

	return manifest, nil
}

func ParseSystemFromBytes(source []byte) (*Manifest, error) {
	// TODO: gojsonschema does not support yaml
	// documentLoader := gojsonschema.NewBytesLoader(src)
	// if err := validate(documentLoader); err != nil {
	// 	return nil, err
	// }

	root := Root{}
	if err := yaml.Unmarshal(source, &root); err != nil {
		return nil, ErrFailedToParseManifest
		// return nil, fmt.Errorf("failed to parse plugin manifest: %w", err)
	}

	manifest, err := root.manifest()
	if err != nil {
		return nil, err
	}

	return manifest, nil
}

func MustParseSystemFromBytes(source []byte) *Manifest {
	m, err := ParseSystemFromBytes(source)
	if err != nil {
		panic(err)
	}
	return m
}

func validate(ld gojsonschema.JSONLoader) error {
	// documentLoader, reader2 := gojsonschema.NewReaderLoader(source)
	result, err := gojsonschema.Validate(schemaLoader, ld)
	if err != nil {
		return ErrFailedToParseManifest
	}

	if !result.Valid() {
		var errstr string
		for i, e := range result.Errors() {
			if i > 0 {
				errstr += ", "
			}
			errstr += e.String()
		}
		return fmt.Errorf("invalid manifest: %w", errors.New(errstr))
	}

	return nil
}
