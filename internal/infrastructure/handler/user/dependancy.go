package user

import (
	"context"
	"github.com/go-openapi/strfmt"
	"github/user-manager/internal/generated/server/models"
)

type userService interface {
	CreateUser(ctx context.Context, userInfo models.CreatingUser) (strfmt.UUID, error)
	DeleteUser(ctx context.Context, userID strfmt.UUID) error
	UpdateUser(ctx context.Context, user models.UpdatingUser) (models.User, error)

	GetUsersByFilters(
		ctx context.Context,
		filters models.Filters,
		limit int64,
		next *string,
	) ([]*models.User, error)
}
