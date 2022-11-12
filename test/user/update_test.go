package user

import (
	"context"
	"log"
	"testing"

	"github.com/go-openapi/strfmt"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"

	"github/user-manager/internal/entity"
	errorpkg "github/user-manager/internal/error"
	"github/user-manager/pkg/client/client/user"
	"github/user-manager/pkg/client/models"
	"github/user-manager/test/tools/postgres"
	"github/user-manager/tools/password"
)

func TestUpdate(t *testing.T) {
	if err := postgres.TruncateUserProfile(); err != nil {
		log.Fatalf("fail to truncate user profile table")
	}

	t.Run("successful updating", func(tt *testing.T) {
		paramsForCreating := user.NewCreateUserParams()
		paramsForCreating.Body = &models.UserInfo{
			FirstName: "John",
			LastName:  "Smith",
			Nickname:  "John1337",
			Email:     "johnsmith@gmail.com",
			Password:  "42adfAfLK",
			CountryID: 1,
		}

		res, err := umClient.User.CreateUser(paramsForCreating)
		assert.Nil(tt, err)

		pswd := "53aedAfLK"
		pswdHash, _ := password.Hash(pswd)

		userEntity := entity.User{
			FirstName:    "Egor",
			LastName:     "Shestakov",
			Nickname:     "muzzzko",
			Email:        "egorshestakov@gmail.com",
			PasswordHash: pswdHash,
			Country: entity.Country{
				ID:   2,
				Code: "FR",
			},
		}

		paramsForUpdating := user.NewUpdateUserParams()
		paramsForUpdating.Body = &models.User{
			UserInfo: models.UserInfo{
				FirstName: userEntity.FirstName,
				LastName:  userEntity.LastName,
				Nickname:  userEntity.Nickname,
				Email:     userEntity.Email,
				Password:  pswd,
				CountryID: userEntity.Country.ID,
			},
			ID: res.Payload.ID,
		}

		_, err = umClient.User.UpdateUser(paramsForUpdating)
		assert.Nil(tt, err)

		userEntity.ID = res.Payload.ID
		userFromDB, err := postgres.GetUserByID(res.Payload.ID)
		assert.Nil(tt, err)
		assert.Equal(tt, userEntity, userFromDB)
	})

	t.Run("update user with email which already belongs to another user", func(tt *testing.T) {
		model := models.UserInfo{
			FirstName: "Egor",
			LastName:  "Shestakov",
			Nickname:  "muzzzko",
			Email:     "existedemail@gmail.com",
			Password:  "42adfAfLK",
			CountryID: 1,
		}

		paramsForCreating := user.NewCreateUserParams()
		paramsForCreating.Body = &model

		_, err := umClient.User.CreateUser(paramsForCreating)
		assert.Nil(tt, err)

		model.Email = "anotheremail@gmail.com"

		res, err := umClient.User.CreateUser(paramsForCreating)
		assert.Nil(tt, err)

		model.Email = "existedemail@gmail.com"

		paramsForUpdating := user.NewUpdateUserParams()
		paramsForUpdating.Body = &models.User{
			UserInfo: model,
			ID:       res.Payload.ID,
		}

		_, err = umClient.User.UpdateUser(paramsForUpdating)
		errBody := err.(*user.UpdateUserUnprocessableEntity)
		assert.Equal(tt, errorpkg.GetDomainErrCode(context.Background(), errorpkg.UserAlreadyExists), errBody.Payload.Code)
	})

	t.Run("update user with invalid email", func(tt *testing.T) {
		params := user.NewUpdateUserParams()
		params.Body = &models.User{
			UserInfo: models.UserInfo{
				FirstName: "Egor",
				LastName:  "Shestakov",
				Nickname:  "muzzzko",
				Email:     "invalid",
				Password:  "42adfAfLK",
				CountryID: 1,
			},
			ID: strfmt.UUID(uuid.New().String()),
		}

		_, err := umClient.User.UpdateUser(params)
		errBody := err.(*user.UpdateUserUnprocessableEntity)
		assert.Equal(tt, errorpkg.GetValidationErrCode(), errBody.Payload.Code)
	})

	t.Run("update user with not existed country", func(tt *testing.T) {
		params := user.NewUpdateUserParams()
		params.Body = &models.User{
			UserInfo: models.UserInfo{
				FirstName: "Egor",
				LastName:  "Shestakov",
				Nickname:  "muzzzko",
				Email:     "johnsmith@gmail.com",
				Password:  "42adfAfLK",
				CountryID: 3544586,
			},
			ID: strfmt.UUID(uuid.New().String()),
		}

		_, err := umClient.User.UpdateUser(params)
		errBody := err.(*user.UpdateUserUnprocessableEntity)
		assert.Equal(tt, errorpkg.GetDomainErrCode(context.Background(), errorpkg.CountryNotFound), errBody.Payload.Code)
	})

	t.Run("update user which doesn't exist", func(tt *testing.T) {
		params := user.NewUpdateUserParams()
		params.Body = &models.User{
			UserInfo: models.UserInfo{
				FirstName: "Egor",
				LastName:  "Shestakov",
				Nickname:  "muzzzko",
				Email:     "johnsmith@gmail.com",
				Password:  "42adfAfLK",
				CountryID: 1,
			},
			ID: strfmt.UUID(uuid.New().String()),
		}

		_, err := umClient.User.UpdateUser(params)
		errBody := err.(*user.UpdateUserUnprocessableEntity)
		assert.Equal(tt, errorpkg.GetDomainErrCode(context.Background(), errorpkg.UserNotFound), errBody.Payload.Code)
	})
}
