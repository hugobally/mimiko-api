package mutation

import (
	"context"
	"github.com/hugobally/mimiko/backend/auth"
	"github.com/hugobally/mimiko/backend/prisma"
	"github.com/hugobally/mimiko/backend/validation"
)

// TODO Unique usernames
func (r *Resolver) UpdateUsername(ctx context.Context, newUsername string) (*prisma.User, error) {
	err := validation.Length(newUsername, 3, 30)
	if err != nil {
		return nil, err
	}
	err = validation.AlphaNum(newUsername)
	if err != nil {
		return nil, err
	}

	user, err := auth.GetUserFromContext(ctx)
	if err != nil {
		return nil, err
	}

	return r.Prisma.UpdateUser(prisma.UserUpdateParams{
		Data: prisma.UserUpdateInput{
			Username: &newUsername,
		},
		Where: prisma.UserWhereUniqueInput{
			ID: &user.ID,
		},
	}).Exec(ctx)
}
