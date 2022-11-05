package mutation

import (
	"context"
	"github.com/hugobally/mimiko_api/internal/db/models"
)

// TODO Unique usernames
func (r *MutationResolver) UpdateUsername(ctx context.Context, newUsername string) (*models.User, error) {
	//	err := validation.Length(newUsername, 3, 30)
	//	if err != nil {
	//		return nil, err
	//	}
	//	err = validation.AlphaNum(newUsername)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	user, err := authentication.UserIdFromContext(ctx)
	//	if err != nil {
	//		return nil, err
	//	}
	//
	//	return r.Prisma.UpdateUser(prisma.UserUpdateParams{
	//		Data: prisma.UserUpdateInput{
	//			Username: &newUsername,
	//		},
	//		Where: prisma.UserWhereUniqueInput{
	//			ID: &user.ID,
	//		},
	//	}).Exec(ctx)
	return nil, nil
}
