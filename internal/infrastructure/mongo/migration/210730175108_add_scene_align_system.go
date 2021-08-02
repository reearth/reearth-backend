package migration

import (
	"context"

	"github.com/labstack/gommon/log"
	"github.com/reearth/reearth-backend/internal/infrastructure/mongo/mongodoc"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/scene"
	"go.mongodb.org/mongo-driver/bson"
)

func AddSceneAlignSystem(ctx context.Context, c DBClient) error {
	col := c.WithCollection("scene")

	return col.Find(ctx, bson.D{}, &mongodoc.BatchConsumer{
		Size: 50,
		Callback: func(rows []bson.Raw) error {

			ids := make([]string, 0, len(rows))
			newRows := make([]interface{}, 0, len(rows))

			log.Infof("migration: AddSceneAlignSystem: hit scenes: %d\n", len(rows))

			for _, row := range rows {
				var doc mongodoc.SceneDocument
				if err := bson.Unmarshal(row, &doc); err != nil {
					return err
				}

				swas := scene.NewWidgetAlignSystem()

				for _, w := range doc.Widgets {
					dl := w.WidgetLayout.DefaultLocation
					mdl := scene.WidgetLocation{Zone: dl.Zone, Section: dl.Section, Area: dl.Area}
					wid, _ := id.WidgetIDFrom(w.ID)
					swas.Add(wid, &mdl)
				}

				mwas := *mongodoc.NewWidgetAlignSystem(swas)
				doc.AlignSystem = mwas

				ids = append(ids, doc.ID)
				newRows = append(newRows, doc)
			}

			return col.SaveAll(ctx, ids, newRows)
		},
	})
}
