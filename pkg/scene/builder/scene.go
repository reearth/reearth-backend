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

type widgetJSON struct {
	PluginID    string       `json:"pluginId"`
	ExtensionID string       `json:"extensionId"`
	Property    propertyJSON `json:"property"`
	Extended    bool         `json:"extended"`
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
		WidgetAlignSystem: b.widgetAlignSystem(ctx, s),
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
		res = append(res, &widgetJSON{
			PluginID:    w.Plugin().String(),
			ExtensionID: string(w.Extension()),
			Property:    b.property(ctx, findProperty(p, w.Property())),
			Extended:    w.WidgetLayout().Extended,
		})
	}
	return res
}

func (b *Builder) widgetAlignSystem(ctx context.Context, s *scene.Scene) *widgetAlignSystemJSON {
	sas := s.WidgetAlignSystem()

	res := widgetAlignSystemJSON{Inner: buildWidgetZone(sas), Outer: buildWidgetZone(sas)}
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

func toString(wids []*id.WidgetID) []string {
	if wids == nil {
		return nil
	}
	docids := make([]string, 0, len(wids))
	for _, wid := range wids {
		if wid == nil {
			continue
		}
		docids = append(docids, wid.String())
	}
	return docids
}

func buildWidgetZone(sas *scene.WidgetAlignSystem) widgetZone {
	return widgetZone{
		Left:   buildWidgetSection(sas),
		Center: buildWidgetSection(sas),
		Right:  buildWidgetSection(sas),
	}
}

func buildWidgetSection(sas *scene.WidgetAlignSystem) widgetSection {
	return widgetSection{
		Top: widgetArea{
			WidgetIDs: toString(sas.WidgetIDs("outer", "right", "top")),
			Align:     *sas.Alignment("outer", "right", "top"),
		},
		Middle: widgetArea{
			WidgetIDs: toString(sas.WidgetIDs("outer", "right", "middle")),
			Align:     *sas.Alignment("outer", "right", "middle"),
		},
		Bottom: widgetArea{
			WidgetIDs: toString(sas.WidgetIDs("outer", "right", "bottom")),
			Align:     *sas.Alignment("outer", "right", "bottom"),
		},
	}
}
