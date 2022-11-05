package resolver

import (
	"context"
	"github.com/hugobally/mimiko/backend/prisma"
)

type mapResolver struct{ *Resolver }

func (r *mapResolver) Author(ctx context.Context, obj *prisma.Map) (*prisma.User, error) {
	err := r.Permission.ReadMapData(ctx, obj)
	if err != nil {
		return nil, err
	}

	return r.Prisma.Map(prisma.MapWhereUniqueInput{
		ID: &obj.ID,
	}).Author().Exec(ctx)
}

func (r *mapResolver) Knots(ctx context.Context, obj *prisma.Map) ([]prisma.Knot, error) {
	err := r.Permission.ReadMapData(ctx, obj)
	if err != nil {
		return nil, err
	}

	return r.Prisma.Map(prisma.MapWhereUniqueInput{
		ID: &obj.ID,
	}).Knots(nil).Exec(ctx)
}

func (r *mapResolver) Links(ctx context.Context, obj *prisma.Map) ([]prisma.Link, error) {
	err := r.Permission.ReadMapData(ctx, obj)
	if err != nil {
		return nil, err
	}

	return r.Prisma.Map(prisma.MapWhereUniqueInput{
		ID: &obj.ID,
	}).Links(nil).Exec(ctx)
}
