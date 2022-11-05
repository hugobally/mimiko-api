package mutation

import (
	"context"
	"github.com/hugobally/mimiko_api/internal/authorization"
	"github.com/hugobally/mimiko_api/internal/db/models"
	"github.com/hugobally/mimiko_api/internal/graph/gqltypes"
)

func (r *MutationResolver) CreateKnot(ctx context.Context, mapID uint, input gqltypes.KnotInput) (*models.Knot, error) {
	m, err := r.Database.FindMap(mapID)
	if err != nil {
		return nil, err
	}

	err = authorization.ModifyMap(ctx, m)
	if err != nil {
		return nil, err
	}

	// TODO Transaction

	k := models.Knot{
		TrackID: *input.TrackID,
		Level:   *input.Level,
		Visited: input.Visited,
		MapID:   m.ID,
	}
	res := r.Database.Create(&k)
	if res.Error != nil {
		return nil, res.Error
	}

	l := models.Link{
		SourceID: *input.SourceID,
		TargetID: k.ID,
		MapID:    m.ID,
	}
	res = r.Database.Create(&l)
	if res.Error != nil {
		return nil, res.Error
	}

	k.ParentLinks = append(k.ParentLinks, l)
	return &k, nil
}

func (r *MutationResolver) UpdateKnot(ctx context.Context, knotID uint, input gqltypes.KnotInput) (*models.Knot, error) {
	//
	//	user, err := authentication.UserIdFromContext(ctx)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	owner, err := r.Prisma.Knot(prisma.KnotWhereUniqueInput{
	//		ID: &knotId,
	//	}).Map().Author().Exec(ctx)
	//
	//	if owner.ID != user.ID {
	//		return nil, errors.New("unauthorized")
	//	}
	//
	//	return r.Prisma.UpdateKnot(prisma.KnotUpdateParams{
	//		Data: prisma.KnotUpdateInput{
	//			Visited: input.Visited,
	//			TrackId: input.TrackId,
	//		},
	//		Where: prisma.KnotWhereUniqueInput{
	//			ID: &knotId,
	//		},
	//	}).Exec(ctx)
	return nil, nil
}

func (r *MutationResolver) DeleteKnots(ctx context.Context, mapID uint, knotIDs []uint) (*gqltypes.MutationResult, error) {
	//
	//	err := r.Permission.ModifyMap(ctx, mapId)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	res, err := r.Prisma.DeleteManyKnots(&prisma.KnotWhereInput{
	//		IDIn: knotIds,
	//		Map:  &prisma.MapWhereInput{ID: &mapId},
	//	}).Exec(ctx)
	//
	//	if err != nil {
	//		return &gqltypes.MutationResult{Success: false, Count: 0}, err
	//	}
	//
	//	return &gqltypes.MutationResult{Success: true, Count: int(res.Count)}, nil
	return nil, nil
}
