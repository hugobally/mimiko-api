package authorization

import (
	"context"
	"errors"
	"github.com/hugobally/mimiko_api/internal/authentication"
	"github.com/hugobally/mimiko_api/internal/db"
	"github.com/hugobally/mimiko_api/internal/db/models"
)

func ReadUserData(ctx context.Context, userID uint) error {
	ctxUserID, err := authentication.UserIdFromContext(ctx)
	if err != nil {
		return err
	}

	if ctxUserID != userID {
		return errors.New("unauthorized")
	}

	return nil
}

func ReadMapData(ctx context.Context, m *models.Map) error {
	if !m.Public {
		id, err := authentication.UserIdFromContext(ctx)
		if err != nil {
			return err
		}

		if id != m.AuthorID {
			return errors.New("unauthorized")
		}
	}

	return nil
}

func CreatePublicMaps(c *db.Client, userID uint) error {
	u, err := c.FindUser(userID)
	if err != nil {
		return err
	}
	if !u.Admin {
		return errors.New("unauthorized")
	}
	return nil
}

func ModifyMap(ctx context.Context, m *models.Map) error {
	u, err := authentication.UserIdFromContext(ctx)
	if err != nil {
		return err
	}

	if u != m.AuthorID {
		return errors.New("unauthorized")
	}

	return nil
}
