package user

import (
	"testing"

	"github.com/jackc/pgx/v5"
	"github.com/stretchr/testify/assert"

	"github/user-manager/pkg/client/client/user"
	"github/user-manager/pkg/client/models"
	"github/user-manager/test/tools/postgres"
)

func TestDelete(t *testing.T) {
	t.Run("successful deleting", func(tt *testing.T) {
		t.Run("delete user", func(tt *testing.T) {
			paramsForCreating := user.NewCreateUserParams()
			paramsForCreating.Body = &models.UserInfo{
				FirstName: "Egor",
				LastName:  "Shestakov",
				Nickname:  "muzzzko",
				Email:     "userfordeleting@gmail.com",
				Password:  "42adfAfLK",
				CountryID: 1,
			}

			res, err := umClient.User.CreateUser(paramsForCreating)
			assert.Nil(tt, err)

			paramsForDeleting := user.NewDeleteUserParams()
			paramsForDeleting.UserID = res.Payload.ID

			_, err = umClient.User.DeleteUser(paramsForDeleting)
			assert.Nil(tt, err)

			_, err = postgres.GetUserByID(res.Payload.ID)
			assert.Equal(tt, pgx.ErrNoRows, err)
		})
	})
}
