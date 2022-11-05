package mutation

import (
	"context"
	"errors"
	"github.com/hugobally/mimiko/backend/api/models"
	"github.com/hugobally/mimiko/backend/auth"
	"github.com/hugobally/mimiko/backend/prisma"
)

func (r *Resolver) CreateKnots(ctx context.Context, mapId string, inputs []models.KnotInput) ([]prisma.Knot, error) {
	err := r.Permission.ModifyMap(ctx, mapId)
	if err != nil {
		return nil, err
	}

	var newKnots []prisma.Knot
	for _, input := range inputs {
		if input.TrackId == nil || input.Level == nil {
			continue
		}

		knot, err := r.Prisma.CreateKnot(prisma.KnotCreateInput{
			Map: prisma.MapCreateOneWithoutKnotsInput{
				Connect: &prisma.MapWhereUniqueInput{
					ID: &mapId,
				},
			},
			TrackId: *input.TrackId,
			Level:   *input.Level,
			Visited: input.Visited,
		}).Exec(ctx)
		if err != nil {
			return newKnots, err
		}
		newKnots = append(newKnots, *knot)
	}
	return newKnots, nil
}

func (r *Resolver) UpdateKnot(ctx context.Context, knotId string, input models.KnotInput) (*prisma.Knot, error) {

	user, err := auth.GetUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	owner, err := r.Prisma.Knot(prisma.KnotWhereUniqueInput{
		ID: &knotId,
	}).Map().Author().Exec(ctx)

	if owner.ID != user.ID {
		return nil, errors.New("unauthorized")
	}

	return r.Prisma.UpdateKnot(prisma.KnotUpdateParams{
		Data: prisma.KnotUpdateInput{
			Visited: input.Visited,
			TrackId: input.TrackId,
		},
		Where: prisma.KnotWhereUniqueInput{
			ID: &knotId,
		},
	}).Exec(ctx)
}

func (r *Resolver) DeleteKnots(ctx context.Context, mapId string, knotIds []string) (*models.MutationResult, error) {

	err := r.Permission.ModifyMap(ctx, mapId)
	if err != nil {
		return nil, err
	}

	res, err := r.Prisma.DeleteManyKnots(&prisma.KnotWhereInput{
		IDIn: knotIds,
		Map:  &prisma.MapWhereInput{ID: &mapId},
	}).Exec(ctx)

	if err != nil {
		return &models.MutationResult{Success: false, Count: 0}, err
	}

	return &models.MutationResult{Success: true, Count: int(res.Count)}, nil
}
