package gqlmodel

import (
	"io"

	"github.com/99designs/gqlgen/graphql"
	"github.com/reearth/reearth-backend/internal/usecase"
	"github.com/reearth/reearth-backend/internal/usecase/interfaces"
	"github.com/reearth/reearth-backend/pkg/file"
	"github.com/reearth/reearth-backend/pkg/scene"
	"github.com/reearth/reearth-backend/pkg/visualizer"
)

func RefToIndex(i *int) int {
	if i == nil {
		return -1
	}
	return *i
}

func RefToString(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

func stringToRef(s string) *string {
	if s == "" {
		return nil
	}
	return &s
}

func ToPageInfo(p *usecase.PageInfo) *PageInfo {
	if p == nil {
		return &PageInfo{}
	}
	return &PageInfo{
		StartCursor:     p.StartCursor(),
		EndCursor:       p.EndCursor(),
		HasNextPage:     p.HasNextPage(),
		HasPreviousPage: p.HasPreviousPage(),
	}
}

func ToWidgetLayout(w *scene.WidgetLayout) *WidgetLayout {
	if w == nil {
		return nil
	}
	var l *WidgetLocation
	if w.DefaultLocation != nil {
		l = &WidgetLocation{
			Zone:    (*WidgetZoneType)(&w.DefaultLocation.Zone),
			Section: (*WidgetSectionType)(&w.DefaultLocation.Section),
			Area:    (*WidgetAreaType)(&w.DefaultLocation.Area),
		}
	}

	var wl WidgetLayout
	if w.Extendable != nil {
		wl = WidgetLayout{
			Extendable: &WidgetExtendable{
				Vertically:   w.Extendable.Vertically,
				Horizontally: w.Extendable.Horizontally,
			},
			Extended:        w.Extended,
			Floating:        w.Floating,
			DefaultLocation: l,
		}
	} else {
		wl = WidgetLayout{
			Extended:        w.Extended,
			Floating:        w.Floating,
			DefaultLocation: l,
		}
	}
	return &wl
}

func ToVisualizer(v visualizer.Visualizer) Visualizer {
	switch v {
	case visualizer.VisualizerCesium:
		return VisualizerCesium
	}
	return Visualizer("")
}

func FromFile(f *graphql.Upload) *file.File {
	if f == nil {
		return nil
	}
	return &file.File{
		Content:     io.NopCloser(f.File),
		Path:        f.Filename,
		Size:        f.Size,
		ContentType: f.ContentType,
	}
}

func FromListOperation(op ListOperation) interfaces.ListOperation {
	if op == ListOperationAdd {
		return interfaces.ListOperationAdd
	} else if op == ListOperationMove {
		return interfaces.ListOperationMove
	} else if op == ListOperationRemove {
		return interfaces.ListOperationRemove
	}
	return interfaces.ListOperation("")
}
