package builder

import (
	"context"
	"time"

	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/property"
	"github.com/reearth/reearth-backend/pkg/scene"
)

type sceneJSON struct {
	SchemaVersion     int                     `json:"schemaVersion"`
	ID                string                  `json:"id"`
	PublishedAt       time.Time               `json:"publishedAt"`
	Property          propertyJSON            `json:"property"`
	Plugins           map[string]propertyJSON `json:"plugins"`
	Layers            []*layerJSON            `json:"layers"`
	Widgets           []*widgetJSON           `json:"widgets"`
	WidgetAlignSystem *widgetAlignSystemJSON  `json:"widgetAlignSystem"`
}

type extendableJSON struct {
	Vertically   *bool `json:"vertically"`
	Horizontally *bool `json:"horizontally"`
}

type widgetJSON struct {
	ID          string         `json:"id"`
	PluginID    string         `json:"pluginId"`
	ExtensionID string         `json:"extensionId"`
	Property    propertyJSON   `json:"property"`
	Extended    *bool          `json:"extended"`
	Extendable  extendableJSON `json:"extendable"`
	Floating    bool           `json:"floating"`
}

func (b *Builder) scene(ctx context.Context, s *scene.Scene, publishedAt time.Time, l []*layerJSON, p []*property.Property) *sceneJSON {
	return &sceneJSON{
		SchemaVersion:     version,
		ID:                s.ID().String(),
		PublishedAt:       publishedAt,
		Property:          b.property(ctx, findProperty(p, s.Property())),
		Plugins:           b.plugins(ctx, s, p),
		Widgets:           b.widgets(ctx, s, p),
		Layers:            l,
		WidgetAlignSystem: b.widgetAlignSystem(s),
	}
}

func (b *Builder) plugins(ctx context.Context, s *scene.Scene, p []*property.Property) map[string]propertyJSON {
	scenePlugins := s.PluginSystem().Plugins()
	res := map[string]propertyJSON{}
	for _, sp := range scenePlugins {
		if sp == nil {
			continue
		}
		if pp := sp.Property(); pp != nil {
			res[sp.Plugin().String()] = b.property(ctx, findProperty(p, *pp))
		}
	}
	return res
}

func (b *Builder) widgets(ctx context.Context, s *scene.Scene, p []*property.Property) []*widgetJSON {
	sceneWidgets := s.WidgetSystem().Widgets()
	res := make([]*widgetJSON, 0, len(sceneWidgets))
	for _, w := range sceneWidgets {
		if !w.Enabled() {
			continue
		}
		if w.WidgetLayout().Extendable != nil {
			res = append(res, &widgetJSON{
				ID:          w.ID().String(),
				PluginID:    w.Plugin().String(),
				ExtensionID: string(w.Extension()),
				Property:    b.property(ctx, findProperty(p, w.Property())),
				Extended:    w.WidgetLayout().Extended,
				Extendable: extendableJSON{
					Vertically:   w.WidgetLayout().Extendable.Vertically,
					Horizontally: w.WidgetLayout().Extendable.Horizontally,
				},
				Floating: w.WidgetLayout().Floating,
			})
		} else {
			res = append(res, &widgetJSON{
				ID:          w.ID().String(),
				PluginID:    w.Plugin().String(),
				ExtensionID: string(w.Extension()),
				Property:    b.property(ctx, findProperty(p, w.Property())),
				Extended:    w.WidgetLayout().Extended,
				Floating:    w.WidgetLayout().Floating,
			})
		}
	}
	return res
}

func (b *Builder) widgetAlignSystem(s *scene.Scene) *widgetAlignSystemJSON {
	sas := s.WidgetAlignSystem()

	res := widgetAlignSystemJSON{Inner: buildWidgetZone(sas.Zone("inner")), Outer: buildWidgetZone(sas.Zone("outer"))}
	return &res
}

func (b *Builder) property(ctx context.Context, p *property.Property) propertyJSON {
	return property.SealProperty(ctx, p).Interface()
}

func findProperty(pp []*property.Property, i id.PropertyID) *property.Property {
	for _, p := range pp {
		if p.ID() == i {
			return p
		}
	}
	return nil
}

func toString(wids []id.WidgetID) []string {
	if wids == nil {
		return nil
	}
	docids := make([]string, 0, len(wids))
	for _, wid := range wids {
		docids = append(docids, wid.String())
	}
	return docids
}

func buildWidgetZone(z *scene.WidgetZone) widgetZone {
	if z == nil {
		return widgetZone{}
	}
	return widgetZone{
		Left:   buildWidgetSection(*z.Section(scene.WidgetSectionLeft)),
		Center: buildWidgetSection(*z.Section(scene.WidgetSectionCenter)),
		Right:  buildWidgetSection(*z.Section(scene.WidgetSectionRight)),
	}
}

func buildWidgetSection(s scene.WidgetSection) widgetSection {
	return widgetSection{
		Middle: buildWidgetArea(*s.Area(scene.WidgetAreaMiddle)),
		Top:    buildWidgetArea(*s.Area(scene.WidgetAreaTop)),
		Bottom: buildWidgetArea(*s.Area(scene.WidgetAreaBottom)),
	}
}

func buildWidgetArea(a scene.WidgetArea) widgetArea {
	return widgetArea{
		WidgetIDs: toString(a.WidgetIDs()),
		Align:     string(a.Alignment()),
	}
}
