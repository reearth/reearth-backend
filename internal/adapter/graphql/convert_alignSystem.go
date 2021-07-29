package graphql

import (
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/scene"
)

func toAlignSystem(sas *scene.WidgetAlignSystem) *WidgetAlignSystem {
	widgetAlignDoc := WidgetAlignSystem{Inner: &WidgetZone{
		Left: &WidgetSection{
			Top: &WidgetArea{
				WidgetIds: IDsFrom(sas.WidgetIds("inner", "left", "top")),
			},
			Middle: &WidgetArea{
				WidgetIds: IDsFrom(sas.WidgetIds("inner", "left", "middle")),
			},
			Bottom: &WidgetArea{
				WidgetIds: IDsFrom(sas.WidgetIds("inner", "left", "bottom")),
			},
		},
		Center: &WidgetSection{
			Top: &WidgetArea{
				WidgetIds: IDsFrom(sas.WidgetIds("inner", "center", "top")),
			},
			Middle: &WidgetArea{
				WidgetIds: IDsFrom(sas.WidgetIds("inner", "center", "middle")),
			},
			Bottom: &WidgetArea{
				WidgetIds: IDsFrom(sas.WidgetIds("inner", "center", "bottom")),
			},
		},
		Right: &WidgetSection{
			Top: &WidgetArea{
				WidgetIds: IDsFrom(sas.WidgetIds("inner", "right", "top")),
			},
			Middle: &WidgetArea{
				WidgetIds: IDsFrom(sas.WidgetIds("inner", "right", "middle")),
			},
			Bottom: &WidgetArea{
				WidgetIds: IDsFrom(sas.WidgetIds("inner", "right", "bottom")),
			},
		},
	}, Outer: &WidgetZone{
		Left: &WidgetSection{
			Top: &WidgetArea{
				WidgetIds: IDsFrom(sas.WidgetIds("outer", "left", "top")),
			},
			Middle: &WidgetArea{
				WidgetIds: IDsFrom(sas.WidgetIds("outer", "left", "middle")),
			},
			Bottom: &WidgetArea{
				WidgetIds: IDsFrom(sas.WidgetIds("outer", "left", "bottom")),
			},
		},
		Center: &WidgetSection{
			Top: &WidgetArea{
				WidgetIds: IDsFrom(sas.WidgetIds("outer", "center", "top")),
			},
			Middle: &WidgetArea{
				WidgetIds: IDsFrom(sas.WidgetIds("outer", "center", "middle")),
			},
			Bottom: &WidgetArea{
				WidgetIds: IDsFrom(sas.WidgetIds("outer", "center", "bottom")),
			},
		},
		Right: &WidgetSection{
			Top: &WidgetArea{
				WidgetIds: IDsFrom(sas.WidgetIds("outer", "right", "top")),
			},
			Middle: &WidgetArea{
				WidgetIds: IDsFrom(sas.WidgetIds("outer", "right", "middle")),
			},
			Bottom: &WidgetArea{
				WidgetIds: IDsFrom(sas.WidgetIds("outer", "right", "bottom")),
			},
		},
	}}
	return &widgetAlignDoc
}

func IDsFrom(wids []*id.WidgetID) []*id.ID {
	var nids []*id.ID
	for _, w := range wids {
		nids = append(nids, w.IDRef())
	}
	return nids
}
