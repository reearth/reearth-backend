package mongodoc

import (
	"context"

	"github.com/reearth/reearth-backend/internal/usecase"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ClientCollection struct {
	Client         *Client
	CollectionName string
}

func (c *ClientCollection) Collection() *mongo.Collection {
	return c.Client.Collection(c.CollectionName)
}

func (c *ClientCollection) FindOne(ctx context.Context, filter interface{}, consumer Consumer) error {
	return c.Client.FindOne(ctx, c.CollectionName, filter, consumer)
}

func (c *ClientCollection) Find(ctx context.Context, filter interface{}, consumer Consumer) error {
	return c.Client.Find(ctx, c.CollectionName, filter, consumer)
}

func (c *ClientCollection) Count(ctx context.Context, filter interface{}) (int64, error) {
	return c.Client.Count(ctx, c.CollectionName, filter)
}

func (c *ClientCollection) Paginate(ctx context.Context, filter interface{}, findOptions *options.FindOptions, p *usecase.Pagination, consumer Consumer) (*usecase.PageInfo, error) {
	return c.Client.Paginate(ctx, c.CollectionName, filter, findOptions, p, consumer)
}

func (c *ClientCollection) SaveOne(ctx context.Context, id string, replacement interface{}) error {
	return c.Client.SaveOne(ctx, c.CollectionName, id, replacement)
}

func (c *ClientCollection) SaveAll(ctx context.Context, ids []string, updates []interface{}) error {
	return c.Client.SaveAll(ctx, c.CollectionName, ids, updates)
}

func (c *ClientCollection) RemoveOne(ctx context.Context, id string) error {
	return c.Client.RemoveOne(ctx, c.CollectionName, id)
}

func (c *ClientCollection) RemoveAll(ctx context.Context, ids []string) error {
	return c.Client.RemoveAll(ctx, c.CollectionName, ids)
}

func (c *ClientCollection) CreateIndex(ctx context.Context, keys []string) []string {
	return c.Client.CreateIndex(ctx, c.CollectionName, keys)
}
