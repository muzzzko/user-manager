package user

import (
	"context"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	"github/user-manager/internal/entity"
	errorpkg "github/user-manager/internal/error"
	"github/user-manager/pkg/client/client/user"
	"github/user-manager/pkg/client/models"
	"github/user-manager/test/tools/postgres"
	"github/user-manager/tools/password"
)

func TestCreate(t *testing.T) {
	if err := postgres.TruncateUserProfile(); err != nil {
		log.Fatalf("fail to truncate user profile table")
	}

	t.Run("successful creation", func(tt *testing.T) {
		pswd := "42adfAfLK"
		pswdHash, _ := password.Hash(pswd)

		userEntity := entity.User{
			FirstName:    "Egor",
			LastName:     "Shestakov",
			Nickname:     "muzzzko",
			Email:        "johnsmith@gmail.com",
			PasswordHash: pswdHash,
			Country: entity.Country{
				ID:   1,
				Code: "UK",
			},
		}

		params := user.NewCreateUserParams()
		params.Body = &models.UserInfo{
			FirstName: userEntity.FirstName,
			LastName:  userEntity.LastName,
			Nickname:  userEntity.Nickname,
			Email:     userEntity.Email,
			Password:  pswd,
			CountryID: userEntity.Country.ID,
		}

		res, err := umClient.User.CreateUser(params)
		assert.Nil(tt, err)

		userEntity.ID = res.Payload.ID
		userFromDB, err := postgres.GetUserByID(res.Payload.ID)
		assert.Nil(tt, err)
		assert.Equal(tt, userEntity, userFromDB)
	})

	t.Run("create user with existed email", func(tt *testing.T) {
		params := user.NewCreateUserParams()
		params.Body = &models.UserInfo{
			FirstName: "Egor",
			LastName:  "Shestakov",
			Nickname:  "muzzzko",
			Email:     "existedemail@gmail.com",
			Password:  "42adfAfLK",
			CountryID: 1,
		}

		_, err := umClient.User.CreateUser(params)
		assert.Nil(tt, err)

		_, err = umClient.User.CreateUser(params)
		errBody := err.(*user.CreateUserUnprocessableEntity)
		assert.Equal(tt, errorpkg.GetDomainErrCode(context.Background(), errorpkg.UserAlreadyExists), errBody.Payload.Code)
	})

	t.Run("create user with invalid email", func(tt *testing.T) {
		params := user.NewCreateUserParams()
		params.Body = &models.UserInfo{
			FirstName: "Egor",
			LastName:  "Shestakov",
			Nickname:  "muzzzko",
			Email:     "invalid",
			Password:  "42adfAfLK",
			CountryID: 1,
		}

		_, err := umClient.User.CreateUser(params)
		errBody := err.(*user.CreateUserUnprocessableEntity)
		assert.Equal(tt, errorpkg.GetValidationErrCode(), errBody.Payload.Code)
	})

	t.Run("create user with not existed country", func(tt *testing.T) {
		params := user.NewCreateUserParams()
		params.Body = &models.UserInfo{
			FirstName: "Egor",
			LastName:  "Shestakov",
			Nickname:  "muzzzko",
			Email:     "johnsmith@gmail.com",
			Password:  "42adfAfLK",
			CountryID: 435345345,
		}

		_, err := umClient.User.CreateUser(params)
		errBody := err.(*user.CreateUserUnprocessableEntity)
		assert.Equal(tt, errorpkg.GetDomainErrCode(context.Background(), errorpkg.CountryNotFound), errBody.Payload.Code)
	})
}
