package query

import (
	"context"
	"github.com/hugobally/mimiko/backend/api/models"
	"github.com/hugobally/mimiko/backend/auth"
	"github.com/hugobally/mimiko/backend/prisma"
)

func (r *Resolver) Map(ctx context.Context, mapID string) (*prisma.Map, error) {
	fetchedMap, err := r.Prisma.Map(prisma.MapWhereUniqueInput{
		ID: &mapID,
	}).Exec(ctx)

	if err != nil {
		return nil, err
	}

	err = r.Permission.ReadMapData(ctx, fetchedMap)

	if err != nil {
		return nil, err
	}

	return fetchedMap, nil
}

// TODO Separate user maps endpoint and general search endpoint
func (r *Resolver) Maps(ctx context.Context, filter *models.MapsFilter) ([]prisma.Map, error) {
	user, err := auth.GetUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	var params *prisma.MapsParams

	order := prisma.MapOrderByInputCreatedAtDesc

	if filter != nil && filter.UserId != nil {
		params = &prisma.MapsParams{
			Where: &prisma.MapWhereInput{
				Author: &prisma.UserWhereInput{
					ID: filter.UserId,
				},
			},
			OrderBy: &order,
		}
		if user.ID != *filter.UserId {
			onlyPublic := true
			params.Where.Public = &onlyPublic
		}
	} else {
		params = MapsDefaultFilter(order)
	}

	if filter != nil {
		if filter.Offset > 0 {
			params.Skip = &filter.Offset
		}
		if filter.Limit > 0 && filter.Limit <= 100 {
			params.First = &filter.Limit
		}
	}

	result, err := r.Prisma.Maps(params).Exec(ctx)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func MapsDefaultFilter(order prisma.MapOrderByInput) *prisma.MapsParams {
	onlyPublic := true
	defaultPageLength := int32(100)

	return &prisma.MapsParams{
		Where: &prisma.MapWhereInput{
			Public: &onlyPublic,
		},
		OrderBy: &order,
		First: &defaultPageLength,
	}
}
