package auth

import (
	"context"
	"errors"
	"fmt"
	"github.com/hugobally/mimiko/backend/prisma"
)

type contextKey string

const (
	userCtxKey = contextKey("user")
)

func GetUserFromContext(ctx context.Context) (*prisma.User, error) {
	user, ok := ctx.Value(userCtxKey).(*prisma.User)
	if !ok {
		fmt.Println("auth.GetUserFromContext : internal assertion error")
		return nil, errors.New("internal server error")
	}
	return user, nil
}
