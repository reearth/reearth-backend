//go:generate go run github.com/reearth/reearth-backend/tools/cmd/embed -all -n pluginManifestJSON -i manifest.yml -yaml2json

package builtin

import (
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/plugin"
	"github.com/reearth/reearth-backend/pkg/plugin/manifest"
	"github.com/reearth/reearth-backend/pkg/property"
	"github.com/reearth/reearth-backend/pkg/visualizer"
)

var pluginTranslationList = map[string]*manifest.TranslationRoot{"ja": manifest.MustParseTransSystemFromStaticJSON(pluginManifestJSON_ja)}
var pluginManifest = manifest.MergeManifestTranslation(manifest.MustParseSystemFromStaticJSON(pluginManifestJSON), pluginTranslationList)

// MUST NOT CHANGE
var PropertySchemaIDVisualizerCesium = id.MustPropertySchemaID("reearth/cesium")

// MUST NOT CHANGE
var PropertySchemaIDInfobox = id.MustPropertySchemaID("reearth/infobox")

func GetPropertySchemaByVisualizer(v visualizer.Visualizer) *property.Schema {
	for _, p := range pluginManifest.ExtensionSchema {
		if p.ID().String() == "reearth/"+string(v) {
			return p
		}
	}
	return nil
}

func MustPropertySchemaByVisualizer(v visualizer.Visualizer) *property.Schema {
	ps := GetPropertySchemaByVisualizer(v)
	if ps == nil {
		panic("property schema not found: " + v)
	}
	return ps
}

func GetPropertySchema(id id.PropertySchemaID) *property.Schema {
	for _, p := range pluginManifest.ExtensionSchema {
		if id == p.ID() {
			return p
		}
	}
	return nil
}

func Plugin() *plugin.Plugin {
	return pluginManifest.Plugin
}

func GetPlugin(id id.PluginID) *plugin.Plugin {
	if id.Equal(pluginManifest.Plugin.ID()) {
		return pluginManifest.Plugin
	}
	return nil
}