package graphql

import (
	"io"

	"github.com/99designs/gqlgen/graphql"
	"github.com/reearth/reearth-backend/internal/usecase"
	"github.com/reearth/reearth-backend/internal/usecase/interfaces"
	"github.com/reearth/reearth-backend/pkg/file"
	"github.com/reearth/reearth-backend/pkg/scene"
	"github.com/reearth/reearth-backend/pkg/visualizer"
)

// func refToBool(s *bool) bool {
// 	if s == nil {
// 		return false
// 	}
// 	return *s
// }

// func refToInt(s *int) int {
// 	if s == nil {
// 		return 0
// 	}
// 	return *s
// }

func refToIndex(s *int) int {
	if s == nil {
		return -1
	}
	return *s
}

func refToString(s *string) string {
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

func toPageInfo(p *usecase.PageInfo) *PageInfo {
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

func toWidgetLayout(w *scene.WidgetLayout) *WidgetLayout {
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
	var e *WidgetExtendable
	var wl WidgetLayout
	if w.Extendable != nil {
		e = &WidgetExtendable{
			Vertically:   w.Extendable.Vertically,
			Horizontally: w.Extendable.Horizontally,
		}
		wl = WidgetLayout{
			Extendable:      e,
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

func toVisualizer(v visualizer.Visualizer) Visualizer {
	switch v {
	case visualizer.VisualizerCesium:
		return VisualizerCesium
	}
	return Visualizer("")
}

func fromFile(f *graphql.Upload) *file.File {
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

func fromListOperation(op ListOperation) interfaces.ListOperation {
	if op == ListOperationAdd {
		return interfaces.ListOperationAdd
	} else if op == ListOperationMove {
		return interfaces.ListOperationMove
	} else if op == ListOperationRemove {
		return interfaces.ListOperationRemove
	}
	return interfaces.ListOperation("")
}
