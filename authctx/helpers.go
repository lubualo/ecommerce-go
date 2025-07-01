package authctx

import (
	"context"
	"errors"

	"github.com/lubualo/ecommerce-go/models"
)

func WithUser(ctx context.Context, u *models.AuthUser) context.Context {
	return context.WithValue(ctx, AuthUserKey(), u)
}

func UserFromContext(ctx context.Context) (*models.AuthUser, bool) {
	u, ok := ctx.Value(AuthUserKey()).(*models.AuthUser)
	return u, ok
}

func UserUUIDFromContext(ctx context.Context) (string, error) {
	u, ok := ctx.Value(AuthUserKey()).(*models.AuthUser)
	if !ok {
		return "", errors.New("cannot retrieve user UUID from context")
	}
	return u.UUID, nil
}

func IsAdmin(ctx context.Context) bool {
	_, ok := UserFromContext(ctx)
	return ok && true
}
