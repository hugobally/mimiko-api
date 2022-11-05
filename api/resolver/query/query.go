package query

import (
	"context"
	"github.com/hugobally/mimiko/backend/api/permission"
	"github.com/hugobally/mimiko/backend/auth"
	"github.com/hugobally/mimiko/backend/prisma"
	"github.com/hugobally/mimiko/backend/shared"
)

type Resolver struct {
	*shared.Services
	Permission *permission.Client
}

func (r *Resolver) Me(ctx context.Context) (*prisma.User, error) {
	ctxUser, err := auth.GetUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	return r.Prisma.User(prisma.UserWhereUniqueInput{
		ID: &ctxUser.ID,
	}).Exec(ctx)
}

func NewResolver(s *shared.Services, perm *permission.Client) *Resolver {
	return &Resolver{s, perm}
}
