package query

import (
	"context"
	"github.com/hugobally/mimiko_api/internal/authorization"
	"github.com/hugobally/mimiko_api/internal/db/models"
	"github.com/hugobally/mimiko_api/internal/graph/gqltypes"
)

func (r *QueryResolver) Map(ctx context.Context, mapID uint) (*models.Map, error) {
	m, err := r.Database.FindMap(mapID)
	if err != nil {
		return nil, err
	}

	err = authorization.ReadMapData(ctx, m)

	if err != nil {
		return nil, err
	}

	return m, nil
}

func (r *QueryResolver) Maps(ctx context.Context, params *gqltypes.MapsFilter) ([]*models.Map, error) {
	var maps []*models.Map
	var filter *models.Map

	if params != nil && params.Author != nil {
		err := authorization.ReadUserData(ctx, *params.Author)
		if err != nil {
			return nil, err
		}

		filter = &models.Map{AuthorID: *params.Author}
	} else {
		filter = &models.Map{Public: true}
	}

	res := r.Database.Where(filter).Order("created_at desc").Find(&maps)
	if res.Error != nil {
		return nil, res.Error
	}

	return maps, nil
}
