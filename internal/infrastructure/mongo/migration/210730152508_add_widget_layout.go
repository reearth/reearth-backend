package migration

import (
	"context"

	"github.com/labstack/gommon/log"
	"github.com/reearth/reearth-backend/internal/infrastructure/mongo/mongodoc"
	"github.com/reearth/reearth-backend/pkg/builtin"
	"github.com/reearth/reearth-backend/pkg/id"
	"go.mongodb.org/mongo-driver/bson"
)

func AddWidgetLayout(ctx context.Context, c DBClient) error {
	col := c.WithCollection("scene")

	return col.Find(ctx, bson.D{}, &mongodoc.BatchConsumer{
		Size: 50,
		Callback: func(rows []bson.Raw) error {

			ids := make([]string, 0, len(rows))
			newRows := make([]interface{}, 0, len(rows))

			log.Infof("migration: AddWidgetLayout: hit scenes: %d\n", len(rows))

			for _, row := range rows {
				var doc mongodoc.SceneDocument
				if err := bson.Unmarshal(row, &doc); err != nil {
					return err
				}

				widgets := make([]mongodoc.SceneWidgetDocument, 0, len(doc.Widgets))
				for _, w := range doc.Widgets {
					if w.WidgetLayout == nil {
						pid, _ := id.PluginIDFrom(w.Plugin)

						wl := builtin.GetPlugin(pid).
							Extension(id.PluginExtensionID(w.Extension)).
							Layout()

						wldoc := mongodoc.WidgetLayout{
							Extendable: wl.Extendable,
							Extended:   false,
							Floating:   wl.Floating,
							DefaultLocation: &mongodoc.WidgetLocation{
								Zone:    wl.DefaultLocation.Zone,
								Section: wl.DefaultLocation.Section,
								Area:    wl.DefaultLocation.Area,
							},
						}

						w.WidgetLayout = &wldoc
					}
					widgets = append(widgets, w)
				}
				doc.Widgets = widgets

				ids = append(ids, doc.ID)
				newRows = append(newRows, doc)
			}

			return col.SaveAll(ctx, ids, newRows)
		},
	})
}
