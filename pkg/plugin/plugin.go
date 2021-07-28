package plugin

import (
	"github.com/blang/semver"
	"github.com/reearth/reearth-backend/pkg/i18n"
	"github.com/reearth/reearth-backend/pkg/id"
)

// Plugin _
type Plugin struct {
	id             id.PluginID
	name           i18n.String
	author         string
	description    i18n.String
	repositoryURL  string
	extensions     map[id.PluginExtensionID]*Extension
	extensionOrder []id.PluginExtensionID
	schema         *id.PropertySchemaID
	scene          *id.SceneID
}

func (p *Plugin) ID() id.PluginID {
	return p.id
}

func (p *Plugin) Version() semver.Version {
	return p.id.Version()
}

func (p *Plugin) Name() i18n.String {
	return p.name.Copy()
}

func (p *Plugin) Author() string {
	return p.author
}

func (p *Plugin) Description() i18n.String {
	return p.description.Copy()
}

func (p *Plugin) RepositoryURL() string {
	return p.repositoryURL
}

func (p *Plugin) Extensions() []*Extension {
	if p.extensionOrder == nil {
		return []*Extension{}
	}
	list := make([]*Extension, 0, len(p.extensions))
	for _, id := range p.extensionOrder {
		list = append(list, p.extensions[id])
	}
	return list
}

func (p *Plugin) Extension(id id.PluginExtensionID) *Extension {
	e, ok := p.extensions[id]
	if ok {
		return e
	}
	return nil
}

func (p *Plugin) Schema() *id.PropertySchemaID {
	return p.schema
}

func (p *Plugin) Rename(name i18n.String) {
	p.name = name.Copy()
}

func (p *Plugin) SetDescription(des i18n.String) {
	p.description = des.Copy()
}

// Scene returns scene ID of the plugin. If the scene ID is nil, it indicates that the plugin is public and can be used by anyone.
func (p *Plugin) Scene() *id.SceneID {
	return p.scene.CopyRef()
}
