package authentication

import (
	"context"
	"errors"
)

type contextKey string

const (
	userCtxKey = contextKey("user")
)

func UserIdFromContext(ctx context.Context) (uint, error) {
	userID, ok := ctx.Value(userCtxKey).(uint)
	if !ok {
		return 0, errors.New("internal server error")
	}
	return userID, nil
}
