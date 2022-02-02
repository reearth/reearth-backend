package mongo

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"

	"github.com/reearth/reearth-backend/internal/infrastructure/mongo/mongodoc"
	"github.com/reearth/reearth-backend/internal/usecase"
	"github.com/reearth/reearth-backend/internal/usecase/repo"
	"github.com/reearth/reearth-backend/pkg/asset"
	"github.com/reearth/reearth-backend/pkg/id"
	"github.com/reearth/reearth-backend/pkg/log"
	"github.com/reearth/reearth-backend/pkg/rerror"
)

type assetRepo struct {
	client *mongodoc.ClientCollection
	filter *id.TeamIDSet
}

func NewAsset(client *mongodoc.Client) repo.Asset {
	r := &assetRepo{client: client.WithCollection("asset")}
	r.init()
	return r
}

func (r *assetRepo) init() {
	i := r.client.CreateIndex(context.Background(), []string{"team"})
	if len(i) > 0 {
		log.Infof("mongo: %s: index created: %s", "asset", i)
	}
}

func (r *assetRepo) Filtered(filter []id.TeamID) repo.Asset {
	return &assetRepo{
		client: r.client,
		filter: id.NewTeamIDSet(filter...),
	}
}

func (r *assetRepo) FindByID(ctx context.Context, id id.AssetID, teams []id.TeamID) (*asset.Asset, error) {
	return r.findOne(ctx, bson.M{
		"id": id.String(),
	}, teams)
}

func (r *assetRepo) FindByIDs(ctx context.Context, ids []id.AssetID, teams []id.TeamID) ([]*asset.Asset, error) {
	dst := make([]*asset.Asset, 0, len(ids))
	res, err := r.find(ctx, dst, bson.M{
		"id": bson.M{"$in": id.AssetIDsToStrings(ids)},
	}, teams)
	if err != nil {
		return nil, err
	}
	return mapAssets(ids, res), nil
}

func (r *assetRepo) FindByTeam(ctx context.Context, id id.TeamID, pagination *usecase.Pagination) ([]*asset.Asset, *usecase.PageInfo, error) {
	if r.filter != nil && !r.filter.Has(id) {
		return nil, nil, nil
	}
	filter := bson.M{
		"team": id.String(),
	}
	return r.paginate(ctx, filter, pagination, nil)
}

func (r *assetRepo) Save(ctx context.Context, asset *asset.Asset) error {
	if err := r.ok(asset); err != nil {
		return err
	}
	doc, id := mongodoc.NewAsset(asset)
	return r.client.SaveOne(ctx, id, doc)
}

func (r *assetRepo) Remove(ctx context.Context, id id.AssetID) error {
	return r.client.RemoveOneOf(ctx, id.String(), r.applyFilter(nil, nil))
}

func (r *assetRepo) paginate(ctx context.Context, filter bson.M, pagination *usecase.Pagination, teams []id.TeamID) ([]*asset.Asset, *usecase.PageInfo, error) {
	var c mongodoc.AssetConsumer
	pageInfo, err2 := r.client.Paginate(ctx, r.applyFilter(filter, teams), pagination, &c)
	if err2 != nil {
		return nil, nil, rerror.ErrInternalBy(err2)
	}
	return c.Rows, pageInfo, nil
}

func (r *assetRepo) find(ctx context.Context, dst []*asset.Asset, filter bson.M, teams []id.TeamID) ([]*asset.Asset, error) {
	c := mongodoc.AssetConsumer{
		Rows: dst,
	}
	if err2 := r.client.Find(ctx, r.applyFilter(filter, teams), &c); err2 != nil {
		return nil, rerror.ErrInternalBy(err2)
	}
	return c.Rows, nil
}

func (r *assetRepo) findOne(ctx context.Context, filter bson.M, teams []id.TeamID) (*asset.Asset, error) {
	dst := make([]*asset.Asset, 0, 1)
	c := mongodoc.AssetConsumer{
		Rows: dst,
	}
	if err := r.client.FindOne(ctx, r.applyFilter(filter, teams), &c); err != nil {
		return nil, err
	}
	return c.Rows[0], nil
}

func (r *assetRepo) ok(a *asset.Asset) error {
	if r.filter == nil || r.filter.Has(a.Team()) {
		return nil
	}
	return repo.ErrOperationDenied
}

func (r *assetRepo) applyFilter(filter bson.M, teams []id.TeamID) bson.M {
	if filter == nil {
		filter = bson.M{}
	}
	s := r.filter.Clone()
	s.Add(teams...)
	filter["team"] = bson.M{"$in": id.TeamIDsToStrings(s.All())}
	return filter
}

func mapAssets(ids []id.AssetID, rows []*asset.Asset) []*asset.Asset {
	res := make([]*asset.Asset, 0, len(ids))
	for _, id := range ids {
		var r2 *asset.Asset
		for _, r := range rows {
			if r.ID() == id {
				r2 = r
				break
			}
		}
		res = append(res, r2)
	}
	return res
}
