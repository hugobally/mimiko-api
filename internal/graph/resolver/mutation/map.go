package mutation

import (
	"context"
	"github.com/hugobally/mimiko_api/internal/authentication"
	"github.com/hugobally/mimiko_api/internal/authorization"
	"github.com/hugobally/mimiko_api/internal/db/models"
	"github.com/hugobally/mimiko_api/internal/graph/gqltypes"
	"github.com/hugobally/mimiko_api/internal/validation"
)

func validateMapTitle(title *string) error {
	if title == nil {
		return nil
	}

	err := validation.Length(*title, 1, 80)
	if err != nil {
		return err
	}

	return nil
}

func (r *MutationResolver) CreateMap(ctx context.Context, mapInput gqltypes.MapInput) (*models.Map, error) {
	err := validateMapTitle(mapInput.Title)
	if err != nil {
		return nil, err
	}

	u, err := authentication.UserIdFromContext(ctx)
	if err != nil {
		return nil, err
	}

	if mapInput.Public {
		err := authorization.CreatePublicMaps(r.Database, u)
		if err != nil {
			return nil, err
		}
	}

	m := models.Map{
		Title:      *mapInput.Title,
		FlagshipID: *mapInput.FlagshipID,
		AuthorID:   u,
		Public:     mapInput.Public,
		Knots: []models.Knot{
			{TrackID: *mapInput.FlagshipID},
		},
	}

	res := r.Database.Create(&m)
	if res.Error != nil {
		return nil, res.Error
	}

	return &m, nil
}

func (r *MutationResolver) UpdateMap(ctx context.Context, mapId uint, mapInput gqltypes.MapInput) (*models.Map, error) {
	//	err := r.Permission.ModifyMap(ctx, mapId)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	err = validateMapTitle(mapInput.Title)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	return r.Prisma.UpdateMap(prisma.MapUpdateParams{
	//		Data: prisma.MapUpdateInput{
	//			Title:      mapInput.Title,
	//			FlagshipId: mapInput.FlagshipId,
	//			Public:     mapInput.Public,
	//		},
	//		Where: prisma.MapWhereUniqueInput{
	//			ID: &mapId,
	//		},
	//	}).Exec(ctx)
	return nil, nil
}

func (r *MutationResolver) DeleteMap(ctx context.Context, mapId uint) (*gqltypes.MutationResult, error) {
	//	err := r.Permission.ModifyMap(ctx, mapId)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	_, err = r.Prisma.DeleteMap(prisma.MapWhereUniqueInput{
	//		ID: &mapId,
	//	}).Exec(ctx)
	//
	//	if err != nil {
	//		return &gqltypes.MutationResult{Success: false}, err
	//	}
	//
	//	return &gqltypes.MutationResult{Success: true}, nil
	return nil, nil
}
