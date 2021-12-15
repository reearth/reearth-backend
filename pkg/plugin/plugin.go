package plugin

import (
	"github.com/blang/semver"
	"github.com/reearth/reearth-backend/pkg/i18n"
	"github.com/reearth/reearth-backend/pkg/id"
)

type Plugin struct {
	id             ID
	name           i18n.String
	author         string
	description    i18n.String
	repositoryURL  string
	extensions     map[ExtensionID]*Extension
	extensionOrder []ExtensionID
	schema         *id.PropertySchemaID
}

func (p *Plugin) ID() ID {
	if p == nil {
		return ID{}
	}
	return p.id
}

func (p *Plugin) Version() semver.Version {
	if p == nil {
		return semver.Version{}
	}
	return p.id.Version()
}

func (p *Plugin) Name() i18n.String {
	if p == nil {
		return nil
	}
	return p.name.Copy()
}

func (p *Plugin) Author() string {
	if p == nil {
		return ""
	}
	return p.author
}

func (p *Plugin) Description() i18n.String {
	if p == nil {
		return nil
	}
	return p.description.Copy()
}

func (p *Plugin) RepositoryURL() string {
	if p == nil {
		return ""
	}
	return p.repositoryURL
}

func (p *Plugin) Extensions() []*Extension {
	if p == nil {
		return nil
	}

	list := make([]*Extension, 0, len(p.extensions))
	for _, id := range p.extensionOrder {
		list = append(list, p.extensions[id])
	}
	return list
}

func (p *Plugin) Extension(id ExtensionID) *Extension {
	if p == nil {
		return nil
	}

	e, ok := p.extensions[id]
	if ok {
		return e
	}
	return nil
}

func (p *Plugin) Schema() *id.PropertySchemaID {
	if p == nil {
		return nil
	}

	return p.schema.CopyRef()
}

func (p *Plugin) PropertySchemas() []id.PropertySchemaID {
	if p == nil {
		return nil
	}

	ps := make([]id.PropertySchemaID, 0, len(p.extensions)+1)
	if p.schema != nil {
		ps = append(ps, *p.schema)
	}
	for _, e := range p.extensionOrder {
		ps = append(ps, p.extensions[e].Schema())
	}
	return ps
}

func (p *Plugin) Rename(name i18n.String) {
	p.name = name.Copy()
}

func (p *Plugin) SetDescription(des i18n.String) {
	p.description = des.Copy()
}
