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
}

type countryRepository interface {
	GetCountryByID(ctx context.Context, id int64) (entity.Country, error)
}

type userEventProducer interface {
	Produce(
		ctx context.Context,
		userID strfmt.UUID,
		action string,
	) error
}
