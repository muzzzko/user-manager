package user

import (
	"context"
	"github.com/go-openapi/strfmt"

	"github/user-manager/internal/generated/server/models"
)

type userService interface {
	CreateUser(ctx context.Context, userInfo models.UserInfo) (strfmt.UUID, error)
	DeleteUser(ctx context.Context, userID strfmt.UUID) error
	UpdateUser(ctx context.Context, user models.User) (models.User, error)
}
