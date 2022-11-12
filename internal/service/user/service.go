package user

import (
	"context"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"

	"github/user-manager/internal/entity"
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

func (s *Service) CreateUser(ctx context.Context, userInfo models.UserInfo) (strfmt.UUID, error) {
	ctxLogger := logger.GetFromContext(ctx)

	country, err := s.countryRepo.GetCountryByID(ctx, userInfo.CountryID)
	if err != nil {
		return "", err
	}

	passwordHash, err := password.Hash(userInfo.Password)
	if err != nil {
		return "", err
	}

	user := entity.User{
		ID:           strfmt.UUID(uuid.New().String()),
		FirstName:    userInfo.FirstName,
		LastName:     userInfo.LastName,
		Nickname:     userInfo.Nickname,
		PasswordHash: passwordHash,
		Email:        userInfo.Email,
		Country:      country,
	}

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

func (s *Service) UpdateUser(ctx context.Context, user models.User) (models.User, error) {
	ctxLogger := logger.GetFromContext(ctx)

	country, err := s.countryRepo.GetCountryByID(ctx, user.CountryID)
	if err != nil {
		return models.User{}, err
	}

	passwordHash, err := password.Hash(user.Password)
	if err != nil {
		return models.User{}, err
	}

	newUser := entity.User{
		ID:           user.ID,
		FirstName:    user.FirstName,
		LastName:     user.LastName,
		Nickname:     user.Nickname,
		PasswordHash: passwordHash,
		Email:        user.Email,
		Country:      country,
	}

	if err := s.userRepo.Update(ctx, newUser); err != nil {
		return models.User{}, err
	}

	if err := s.userEventProducer.Produce(ctx, user.ID, event.UpdateAction); err != nil {
		ctxLogger.
			WithError(err).
			WithUserID(user.ID).
			Error("fail produce event while updating user")
	}

	return user, nil
}
