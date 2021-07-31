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
		WidgetAlignSystem: b.widgetAlignment(ctx, s),
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

func (b *Builder) widgetAlignment(ctx context.Context, s *scene.Scene) *widgetAlignSystemJSON {
	sas := s.WidgetAlignSystem()
	res := widgetAlignSystemJSON{Inner: widgetZone{
		Left: widgetSection{
			Top: widgetArea{
				WidgetIds: toString(sas.WidgetIds("inner", "left", "top")),
				Align:     *sas.Alignment("inner", "left", "top"),
			},
			Middle: widgetArea{
				WidgetIds: toString(sas.WidgetIds("inner", "left", "middle")),
				Align:     *sas.Alignment("inner", "left", "middle"),
			},
			Bottom: widgetArea{
				WidgetIds: toString(sas.WidgetIds("inner", "left", "bottom")),
				Align:     *sas.Alignment("inner", "left", "bottom"),
			},
		},
		Center: widgetSection{
			Top: widgetArea{
				WidgetIds: toString(sas.WidgetIds("inner", "center", "top")),
				Align:     *sas.Alignment("inner", "center", "top"),
			},
			Middle: widgetArea{
				WidgetIds: toString(sas.WidgetIds("inner", "center", "middle")),
				Align:     *sas.Alignment("inner", "center", "middle"),
			},
			Bottom: widgetArea{
				WidgetIds: toString(sas.WidgetIds("inner", "center", "bottom")),
				Align:     *sas.Alignment("inner", "center", "bottom"),
			},
		},
		Right: widgetSection{
			Top: widgetArea{
				WidgetIds: toString(sas.WidgetIds("inner", "right", "top")),
				Align:     *sas.Alignment("inner", "right", "top"),
			},
			Middle: widgetArea{
				WidgetIds: toString(sas.WidgetIds("inner", "right", "middle")),
				Align:     *sas.Alignment("inner", "right", "middle"),
			},
			Bottom: widgetArea{
				WidgetIds: toString(sas.WidgetIds("inner", "right", "bottom")),
				Align:     *sas.Alignment("inner", "right", "bottom"),
			},
		},
	}, Outer: widgetZone{
		Left: widgetSection{
			Top: widgetArea{
				WidgetIds: toString(sas.WidgetIds("outer", "left", "top")),
				Align:     *sas.Alignment("outer", "left", "top"),
			},
			Middle: widgetArea{
				WidgetIds: toString(sas.WidgetIds("outer", "left", "middle")),
				Align:     *sas.Alignment("outer", "left", "middle"),
			},
			Bottom: widgetArea{
				WidgetIds: toString(sas.WidgetIds("outer", "left", "bottom")),
				Align:     *sas.Alignment("outer", "left", "bottom"),
			},
		},
		Center: widgetSection{
			Top: widgetArea{
				WidgetIds: toString(sas.WidgetIds("outer", "center", "top")),
				Align:     *sas.Alignment("outer", "center", "top"),
			},
			Middle: widgetArea{
				WidgetIds: toString(sas.WidgetIds("outer", "center", "middle")),
				Align:     *sas.Alignment("outer", "center", "middle"),
			},
			Bottom: widgetArea{
				WidgetIds: toString(sas.WidgetIds("outer", "center", "bottom")),
				Align:     *sas.Alignment("outer", "center", "bottom"),
			},
		},
		Right: widgetSection{
			Top: widgetArea{
				WidgetIds: toString(sas.WidgetIds("outer", "right", "top")),
				Align:     *sas.Alignment("outer", "right", "top"),
			},
			Middle: widgetArea{
				WidgetIds: toString(sas.WidgetIds("outer", "right", "middle")),
				Align:     *sas.Alignment("outer", "right", "middle"),
			},
			Bottom: widgetArea{
				WidgetIds: toString(sas.WidgetIds("outer", "right", "bottom")),
				Align:     *sas.Alignment("outer", "right", "bottom"),
			},
		},
	}}
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
