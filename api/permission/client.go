package permission

import (
	"context"
	"errors"
	"github.com/hugobally/mimiko/backend/auth"
	"github.com/hugobally/mimiko/backend/prisma"
)

type Client struct {
	Prisma *prisma.Client
}

func NewClient(p *prisma.Client) *Client {
	return &Client{p}
}

func (c *Client) ReadUserPrivate(ctx context.Context, user *prisma.User) error {
	ctxUser, err := auth.GetUserFromContext(ctx)
	if err != nil {
		return err
	}

	if ctxUser.ID != user.ID {
		return errors.New("unauthorized")
	}

	return nil
}

func (c *Client) GetMapOwner(ctx context.Context, mapId string) (*prisma.User, error) {
	return c.Prisma.Map(prisma.MapWhereUniqueInput{
		ID: &mapId,
	}).Author().Exec(ctx)
}

func (c *Client) ReadMapData(ctx context.Context, m *prisma.Map) error {
	if !m.Public {
		user, err := auth.GetUserFromContext(ctx)
		if err != nil {
			return err
		}

		owner, err := c.GetMapOwner(ctx, m.ID)
		if err != nil {
			return err
		}

		if user.ID != owner.ID {
			return errors.New("unauthorized")
		}
	}

	return nil
}

func (c *Client) ModifyMap(ctx context.Context, mapId string) error {
	user, err := auth.GetUserFromContext(ctx)
	if err != nil {
		return err
	}

	owner, err := c.GetMapOwner(ctx, mapId)
	if err != nil {
		return err
	}

	if user.ID != owner.ID {
		return errors.New("unauthorized")
	}

	return nil
}
