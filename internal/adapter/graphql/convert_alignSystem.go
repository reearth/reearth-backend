package graphql

import (
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/scene"
)

func toAlignSystem(sas *scene.WidgetAlignSystem) *WidgetAlignSystem {
	widgetAlignDoc := WidgetAlignSystem{Inner: &WidgetZone{
		Left: &WidgetSection{
			Top: &WidgetArea{
				WidgetIds: IDsFrom(sas.WidgetIDs("inner", "left", "top")),
				Align:     sas.Alignment("inner", "left", "top"),
			},
			Middle: &WidgetArea{
				WidgetIds: IDsFrom(sas.WidgetIDs("inner", "left", "middle")),
				Align:     sas.Alignment("inner", "left", "middle"),
			},
			Bottom: &WidgetArea{
				WidgetIds: IDsFrom(sas.WidgetIDs("inner", "left", "bottom")),
				Align:     sas.Alignment("inner", "left", "bottom"),
			},
		},
		Center: &WidgetSection{
			Top: &WidgetArea{
				WidgetIds: IDsFrom(sas.WidgetIDs("inner", "center", "top")),
				Align:     sas.Alignment("inner", "center", "top"),
			},
			Middle: &WidgetArea{
				WidgetIds: IDsFrom(sas.WidgetIDs("inner", "center", "middle")),
				Align:     sas.Alignment("inner", "center", "middle"),
			},
			Bottom: &WidgetArea{
				WidgetIds: IDsFrom(sas.WidgetIDs("inner", "center", "bottom")),
				Align:     sas.Alignment("inner", "center", "bottom"),
			},
		},
		Right: &WidgetSection{
			Top: &WidgetArea{
				WidgetIds: IDsFrom(sas.WidgetIDs("inner", "right", "top")),
				Align:     sas.Alignment("inner", "right", "top"),
			},
			Middle: &WidgetArea{
				WidgetIds: IDsFrom(sas.WidgetIDs("inner", "right", "middle")),
				Align:     sas.Alignment("inner", "right", "middle"),
			},
			Bottom: &WidgetArea{
				WidgetIds: IDsFrom(sas.WidgetIDs("inner", "right", "bottom")),
				Align:     sas.Alignment("inner", "right", "bottom"),
			},
		},
	}, Outer: &WidgetZone{
		Left: &WidgetSection{
			Top: &WidgetArea{
				WidgetIds: IDsFrom(sas.WidgetIDs("outer", "left", "top")),
				Align:     sas.Alignment("outer", "left", "top"),
			},
			Middle: &WidgetArea{
				WidgetIds: IDsFrom(sas.WidgetIDs("outer", "left", "middle")),
				Align:     sas.Alignment("outer", "left", "middle"),
			},
			Bottom: &WidgetArea{
				WidgetIds: IDsFrom(sas.WidgetIDs("outer", "left", "bottom")),
				Align:     sas.Alignment("outer", "left", "bottom"),
			},
		},
		Center: &WidgetSection{
			Top: &WidgetArea{
				WidgetIds: IDsFrom(sas.WidgetIDs("outer", "center", "top")),
				Align:     sas.Alignment("outer", "center", "top"),
			},
			Middle: &WidgetArea{
				WidgetIds: IDsFrom(sas.WidgetIDs("outer", "center", "middle")),
				Align:     sas.Alignment("outer", "center", "middle"),
			},
			Bottom: &WidgetArea{
				WidgetIds: IDsFrom(sas.WidgetIDs("outer", "center", "bottom")),
				Align:     sas.Alignment("outer", "center", "bottom"),
			},
		},
		Right: &WidgetSection{
			Top: &WidgetArea{
				WidgetIds: IDsFrom(sas.WidgetIDs("outer", "right", "top")),
				Align:     sas.Alignment("outer", "right", "top"),
			},
			Middle: &WidgetArea{
				WidgetIds: IDsFrom(sas.WidgetIDs("outer", "right", "middle")),
				Align:     sas.Alignment("outer", "right", "middle"),
			},
			Bottom: &WidgetArea{
				WidgetIds: IDsFrom(sas.WidgetIDs("outer", "right", "bottom")),
				Align:     sas.Alignment("outer", "right", "bottom"),
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
