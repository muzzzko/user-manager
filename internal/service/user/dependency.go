package user

import (
	"context"
	"github.com/go-openapi/strfmt"
	"github/user-manager/internal/entity"
)

type userRepository interface {
	Save(ctx context.Context, user entity.User) error
	Delete(ctx context.Context, userID strfmt.UUID) error
	Update(ctx context.Context, user entity.User) error

	GetByID(ctx context.Context, id strfmt.UUID) (entity.User, error)
	GetUsersByFilters(
		ctx context.Context,
		filters map[string]string,
		limit int64,
		next *string,
	) ([]entity.User, error)
}

type countryRepository interface {
	GetCountryByCode(ctx context.Context, code string) (entity.Country, error)
}

type userEventProducer interface {
	Produce(
		ctx context.Context,
		userID strfmt.UUID,
		action string,
	) error
}
