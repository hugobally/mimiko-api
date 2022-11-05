package resolver

import (
	"context"
	"github.com/hugobally/mimiko_api/internal/authorization"
	"github.com/hugobally/mimiko_api/internal/db/models"
)

type UserResolver struct{ *Root }

func (r *UserResolver) Maps(ctx context.Context, u *models.User) ([]*models.Map, error) {
	err := authorization.ReadUserData(ctx, u.ID)
	if err != nil {
		return nil, err
	}

	var maps []*models.Map

	err = r.Database.Model(u).Association("Maps").Find(&maps)
	if err != nil {
		return nil, err
	}

	return maps, nil
}
