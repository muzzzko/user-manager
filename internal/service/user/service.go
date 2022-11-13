package user

import (
	"context"
	"fmt"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"

	"github/user-manager/internal/constant"
	"github/user-manager/internal/generated/server/models"
	"github/user-manager/pkg/event"
	"github/user-manager/tools/logger"
	"github/user-manager/tools/password"
)

type Service struct {
	userRepo          userRepository
	countryRepo       countryRepository
	userEventProducer userEventProducer
}

func NewService(
	userRepo userRepository,
	countryRepo countryRepository,
	userEventProducer userEventProducer,
) *Service {
	return &Service{
		userRepo:          userRepo,
		countryRepo:       countryRepo,
		userEventProducer: userEventProducer,
	}
}

func (s *Service) CreateUser(ctx context.Context, creatingUser models.CreatingUser) (strfmt.UUID, error) {
	ctxLogger := logger.GetFromContext(ctx)

	country, err := s.countryRepo.GetCountryByCode(ctx, creatingUser.CountryCode)
	if err != nil {
		return "", err
	}

	passwordHash, err := password.Hash(creatingUser.Password)
	if err != nil {
		return "", err
	}

	user := mapModelUserInfoToEntityUser(creatingUser.UserInfo)
	user.ID = strfmt.UUID(uuid.New().String())
	user.Country = country
	user.PasswordHash = passwordHash

	if err := s.userRepo.Save(ctx, user); err != nil {
		return "", err
	}

	if err := s.userEventProducer.Produce(ctx, user.ID, event.CreateAction); err != nil {
		ctxLogger.
			WithError(err).
			WithNickname(user.Nickname).
			Error("fail produce event while creating user")
	}

	return user.ID, nil
}

func (s *Service) DeleteUser(ctx context.Context, userID strfmt.UUID) error {
	ctxLogger := logger.GetFromContext(ctx)

	if err := s.userRepo.Delete(ctx, userID); err != nil {
		return err
	}

	if err := s.userEventProducer.Produce(ctx, userID, event.DeleteAction); err != nil {
		ctxLogger.
			WithError(err).
			WithUserID(userID).
			Error("fail produce event while deleting user")
	}

	return nil
}

func (s *Service) UpdateUser(ctx context.Context, updatingUser models.UpdatingUser) (models.User, error) {
	ctxLogger := logger.GetFromContext(ctx)

	country, err := s.countryRepo.GetCountryByCode(ctx, updatingUser.CountryCode)
	if err != nil {
		return models.User{}, err
	}

	passwordHash, err := password.Hash(updatingUser.Password)
	if err != nil {
		return models.User{}, err
	}

	user := mapModelUserInfoToEntityUser(updatingUser.UserInfo)
	user.ID = updatingUser.ID
	user.Country = country
	user.PasswordHash = passwordHash

	if err := s.userRepo.Update(ctx, user); err != nil {
		return models.User{}, err
	}

	if err := s.userEventProducer.Produce(ctx, updatingUser.ID, event.UpdateAction); err != nil {
		ctxLogger.
			WithError(err).
			WithUserID(updatingUser.ID).
			Error("fail produce event while updating user")
	}

	return mapEntityUserToModelUser(user), nil
}

func (s *Service) GetUsersByFilters(
	ctx context.Context,
	filters models.Filters,
	limit int64,
	next *string,
) ([]*models.User, error) {
	filtersMap := make(map[string]string)
	if filters.CountryCode != nil {
		filtersMap[constant.CountryCodeFilter] = fmt.Sprintf("%s", *filters.CountryCode)
	}

	users, err := s.userRepo.GetUsersByFilters(ctx, filtersMap, limit, next)
	if err != nil {
		return nil, err
	}

	return mapEntityUsersToModelUsers(users), nil
}
