package mongodoc

import (
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/plugin"
	"go.mongodb.org/mongo-driver/bson"
)

type PluginDocument struct {
	ID            string
	Name          map[string]string
	Author        string
	Description   map[string]string
	RepositoryURL string
	Extensions    []PluginExtensionDocument
	Schema        *string
	Scene         *string `bson:",omitempty"`
}

type PluginExtensionDocument struct {
	ID           string
	Type         string
	Name         map[string]string
	Description  map[string]string
	Icon         string
	Schema       string
	Visualizer   string `bson:",omitempty"`
	WidgetLayout *WidgetLayoutDocument
}

type WidgetLayoutDocument struct {
	Extendable      *WidgetExtendableDocument
	Extended        bool
	Floating        bool
	DefaultLocation *WidgetLocationDocument
}

type WidgetExtendableDocument struct {
	Vertically   bool
	Horizontally bool
}

type WidgetLocationDocument struct {
	Zone    string
	Section string
	Area    string
}

type PluginConsumer struct {
	Rows []*plugin.Plugin
}

func (c *PluginConsumer) Consume(raw bson.Raw) error {
	if raw == nil {
		return nil
	}

	var doc PluginDocument
	if err := bson.Unmarshal(raw, &doc); err != nil {
		return err
	}
	plugin, err := doc.Model()
	if err != nil {
		return err
	}
	c.Rows = append(c.Rows, plugin)
	return nil
}

func NewPlugin(plugin *plugin.Plugin) (*PluginDocument, string) {
	if plugin == nil {
		return nil, ""
	}

	extensions := plugin.Extensions()
	extensionsDoc := make([]PluginExtensionDocument, 0, len(extensions))
	for _, e := range extensions {
		extensionsDoc = append(extensionsDoc, PluginExtensionDocument{
			ID:           string(e.ID()),
			Type:         string(e.Type()),
			Name:         e.Name(),
			Description:  e.Description(),
			Icon:         e.Icon(),
			Schema:       e.Schema().String(),
			Visualizer:   string(e.Visualizer()),
			WidgetLayout: NewWidgetLayout(e.WidgetLayout()),
		})
	}

	pid := plugin.ID().String()
	return &PluginDocument{
		ID:            pid,
		Name:          plugin.Name(),
		Description:   plugin.Description(),
		Author:        plugin.Author(),
		RepositoryURL: plugin.RepositoryURL(),
		Extensions:    extensionsDoc,
		Schema:        plugin.Schema().StringRef(),
		Scene:         plugin.ID().Scene().StringRef(),
	}, pid
}

func (d *PluginDocument) Model() (*plugin.Plugin, error) {
	if d == nil {
		return nil, nil
	}

	pid, err := id.PluginIDFrom(d.ID)
	if err != nil {
		return nil, err
	}

	extensions := make([]*plugin.Extension, 0, len(d.Extensions))
	for _, e := range d.Extensions {
		psid, err := id.PropertySchemaIDFrom(e.Schema)
		if err != nil {
			return nil, err
		}
		extension, err := plugin.NewExtension().
			ID(id.PluginExtensionID(e.ID)).
			Type(plugin.ExtensionType(e.Type)).
			Name(e.Name).
			Description(e.Description).
			Icon(e.Icon).
			WidgetLayout(e.WidgetLayout.Model()).
			Schema(psid).
			Build()
		if err != nil {
			return nil, err
		}
		extensions = append(extensions, extension)
	}

	return plugin.New().
		ID(pid).
		Name(d.Name).
		Description(d.Description).
		Author(d.Author).
		RepositoryURL(d.RepositoryURL).
		Extensions(extensions).
		Schema(id.PropertySchemaIDFromRef(d.Schema)).
		Build()
}

func NewWidgetLayout(l *plugin.WidgetLayout) *WidgetLayoutDocument {
	if l == nil {
		return nil
	}

	return &WidgetLayoutDocument{
		Extendable: &WidgetExtendableDocument{
			Vertically:   l.VerticallyExtendable(),
			Horizontally: l.HorizontallyExtendable(),
		},
		Extended: l.Extended(),
		Floating: l.Floating(),
		DefaultLocation: &WidgetLocationDocument{
			Zone:    string(l.DefaultLocation().Zone),
			Section: string(l.DefaultLocation().Section),
			Area:    string(l.DefaultLocation().Area),
		},
	}
}

func (d *WidgetLayoutDocument) Model() *plugin.WidgetLayout {
	if d == nil {
		return nil
	}

	var loc *plugin.WidgetLocation
	if d.DefaultLocation != nil {
		loc = &plugin.WidgetLocation{
			Zone:    plugin.WidgetZoneType(d.DefaultLocation.Zone),
			Section: plugin.WidgetSectionType(d.DefaultLocation.Section),
			Area:    plugin.WidgetAreaType(d.DefaultLocation.Area),
		}
	}

	return plugin.NewWidgetLayout(
		d.Extendable.Horizontally,
		d.Extendable.Vertically,
		d.Extended,
		d.Floating,
		loc,
	).Ref()
}
