package resolver

import (
	"context"
	"github.com/hugobally/mimiko_api/internal/authorization"
	"github.com/hugobally/mimiko_api/internal/db/models"
)

type MapResolver struct{ *Root }

func (r *MapResolver) Author(ctx context.Context, m *models.Map) (*models.User, error) {
	err := authorization.ReadUserData(ctx, m.AuthorID)
	if err != nil {
		return nil, err
	}

	author, err := r.Database.FindUser(m.AuthorID)
	if err != nil {
		return nil, err
	}

	return author, nil
}

func (r *MapResolver) Knots(ctx context.Context, m *models.Map) ([]*models.Knot, error) {
	var knots []*models.Knot

	err := r.Database.Model(m).Association("Knots").Find(&knots)
	if err != nil {
		return nil, err
	}

	return knots, nil
}

func (r *MapResolver) Links(ctx context.Context, m *models.Map) ([]*models.Link, error) {
	var links []*models.Link

	err := r.Database.Model(m).Association("Links").Find(&links)
	if err != nil {
		return nil, err
	}

	return links, nil
}
