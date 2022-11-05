package mutation

import (
	"context"
	"github.com/hugobally/mimiko/backend/api/models"
	"github.com/hugobally/mimiko/backend/prisma"
)

func (r *Resolver) CreateLinks(ctx context.Context, mapId string, sourceID string, targetIds []string) ([]prisma.Link, error) {
	err := r.Permission.ModifyMap(ctx, mapId)
	if err != nil {
		return nil, err
	}

	var newLinks []prisma.Link
	for _, targetId := range targetIds {
		link, err := r.Prisma.CreateLink(prisma.LinkCreateInput{
			Map: prisma.MapCreateOneWithoutLinksInput{
				Connect: &prisma.MapWhereUniqueInput{
					ID: &mapId,
				},
			},
			Source: sourceID,
			Target: targetId,
		}).Exec(ctx)
		if err != nil {
			return newLinks, err
		}
		newLinks = append(newLinks, *link)
	}
	return newLinks, nil
}

func (r *Resolver) DeleteLinks(ctx context.Context, mapId string, linkIds []string) (*models.MutationResult, error) {
	err := r.Permission.ModifyMap(ctx, mapId)
	if err != nil {
		return nil, err
	}

	res, err := r.Prisma.DeleteManyLinks(&prisma.LinkWhereInput{
		IDIn: linkIds,
		Map:  &prisma.MapWhereInput{ID: &mapId},
	}).Exec(ctx)

	if err != nil {
		return &models.MutationResult{Success: false}, err
	}

	return &models.MutationResult{Success: true, Count: int(res.Count)}, nil
}
