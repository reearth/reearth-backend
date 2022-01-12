package scene

import (
	"github.com/reearth/reearth-backend/pkg/id"
)

type Plugins struct {
	plugins []*Plugin
}

func NewPlugins(p []*Plugin) *Plugins {
	if len(p) == 0 {
		return &Plugins{}
	}

	p2 := make([]*Plugin, 0, len(p))
	for _, p1 := range p {
		if p1 == nil {
			continue
		}
		duplicated := false
		for _, p3 := range p2 {
			if p1.plugin.Equal(p3.plugin) {
				duplicated = true
				break
			}
		}
		if !duplicated {
			p3 := *p1
			p2 = append(p2, &p3)
		}
	}

	return &Plugins{plugins: p2}
}

func (p *Plugins) Plugins() []*Plugin {
	return append([]*Plugin{}, p.plugins...)
}

func (p *Plugins) Property(id id.PluginID) *id.PropertyID {
	for _, p := range p.plugins {
		if p.plugin.Equal(id) {
			return p.property.CopyRef()
		}
	}
	return nil
}

func (p *Plugins) Has(id id.PluginID) bool {
	for _, p2 := range p.plugins {
		if p2.plugin.Equal(id) {
			return true
		}
	}
	return false
}

func (p *Plugins) HasNamed(name string) bool {
	for _, p2 := range p.plugins {
		if p2.plugin.Name() == name {
			return true
		}
	}
	return false
}

func (p *Plugins) Add(sp *Plugin) {
	if sp == nil || p.HasNamed(sp.plugin.Name()) || sp.plugin.Equal(id.OfficialPluginID) {
		return
	}
	p.plugins = append(p.plugins, sp)
}

func (p *Plugins) Remove(pid id.PluginID) {
	if pid.Equal(id.OfficialPluginID) {
		return
	}
	for i, p2 := range p.plugins {
		if p2.plugin.Equal(pid) {
			p.plugins = append(p.plugins[:i], p.plugins[i+1:]...)
			return
		}
	}
}

func (p *Plugins) Upgrade(pid, newpid id.PluginID, pr *id.PropertyID, deleteProperty bool) {
	if p == nil || p.HasNamed(newpid.Name()) {
		return
	}

	for i, pp := range p.plugins {
		if pp.plugin.Equal(id.OfficialPluginID) {
			continue
		}
		if pp.plugin.Equal(pid) {
			if newpid.IsNil() {
				newpid = pp.plugin
			}
			newp := pp.property
			if pr != nil {
				newp = pr.CopyRef()
			}
			if deleteProperty {
				newp = nil
			}
			p.plugins[i] = NewPlugin(newpid, newp)
			return
		}
	}
}

func (p *Plugins) Properties() []id.PropertyID {
	if p == nil {
		return nil
	}
	res := make([]id.PropertyID, 0, len(p.plugins))
	for _, pp := range p.plugins {
		if pp.property != nil {
			res = append(res, *pp.property)
		}
	}
	return res
}

func (p *Plugins) Plugin(i id.PluginID) *Plugin {
	for _, pp := range p.plugins {
		if pp.plugin == i {
			return pp
		}
	}
	return nil
}

func (p *Plugins) PluginByName(n string) *Plugin {
	for _, pp := range p.plugins {
		if pp.plugin.Name() == n {
			return pp
		}
	}
	return nil
}
