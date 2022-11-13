package user

import (
	"github.com/stretchr/testify/assert"
	"log"
	"sort"
	"testing"

	"github/user-manager/pkg/client/client/user"
	"github/user-manager/pkg/client/models"
	"github/user-manager/test/tools/postgres"
)

const (
	userAliasEgor = "Egor"
	userAliasJohn = "John"
)

func TestGetUsers(t *testing.T) {
	if err := postgres.TruncateUserProfile(); err != nil {
		log.Fatalf("fail to truncate user profile table")
	}

	userMap := map[string]*models.User{
		userAliasEgor: {
			UserInfo: models.UserInfo{
				FirstName:   "Egor",
				LastName:    "Shestakov",
				Nickname:    "muzzzko",
				Email:       "egor@gmail.com",
				CountryCode: "UK",
			},
		},
		userAliasJohn: {
			UserInfo: models.UserInfo{
				FirstName:   "John",
				LastName:    "Smith",
				Nickname:    "js",
				Email:       "johnsmith@gmail.com",
				CountryCode: "UK",
			},
		},
	}

	createUser(userMap[userAliasEgor], "42adfAfLK")
	createUser(userMap[userAliasJohn], "53adfAfLK")

	t.Run("get users without next, limit and filters", func(tt *testing.T) {
		params := user.NewGetUsersByFiltersParams()

		res, err := umClient.User.GetUsersByFilters(params)
		assert.Nil(tt, err)

		expected := []*models.User{userMap[userAliasEgor], userMap[userAliasJohn]}
		sort.SliceStable(expected, func(i, j int) bool {
			return expected[i].ID > expected[j].ID
		})
		assert.Equal(tt, expected, res.Payload.Users)
		assert.Equal(tt, "", res.Payload.Next)
	})

	t.Run("get users without next, limit with filters", func(tt *testing.T) {
		countryCode := "UK"
		params := user.NewGetUsersByFiltersParams()
		params.Body.Filters = models.Filters{
			CountryCode: &countryCode,
		}

		res, err := umClient.User.GetUsersByFilters(params)
		assert.Nil(tt, err)

		expected := []*models.User{userMap[userAliasEgor], userMap[userAliasJohn]}
		sort.SliceStable(expected, func(i, j int) bool {
			return expected[i].ID > expected[j].ID
		})
		assert.Equal(tt, expected, res.Payload.Users)
		assert.Equal(tt, "", res.Payload.Next)
	})

	t.Run("get users without next with filters and limit", func(tt *testing.T) {
		countryCode := "UK"
		params := user.NewGetUsersByFiltersParams()
		params.Body.Filters = models.Filters{
			CountryCode: &countryCode,
		}
		params.Body.Limit = 1

		res, err := umClient.User.GetUsersByFilters(params)
		assert.Nil(tt, err)
		assert.Equal(tt, len(res.Payload.Users), 1)

		expected := []*models.User{userMap[userAliasEgor], userMap[userAliasJohn]}
		sort.SliceStable(expected, func(i, j int) bool {
			return expected[i].ID > expected[j].ID
		})

		assert.Equal(tt, []*models.User{expected[0]}, res.Payload.Users)
		assert.Equal(tt, expected[0].ID.String(), res.Payload.Next)
	})

	t.Run("get users with next, filters and limit", func(tt *testing.T) {
		expected := []*models.User{userMap[userAliasEgor], userMap[userAliasJohn]}
		sort.SliceStable(expected, func(i, j int) bool {
			return expected[i].ID > expected[j].ID
		})
		next := expected[0].ID.String()

		countryCode := "UK"
		params := user.NewGetUsersByFiltersParams()
		params.Body.Filters = models.Filters{
			CountryCode: &countryCode,
		}
		params.Body.Limit = 2
		params.Body.Next = &next

		res, err := umClient.User.GetUsersByFilters(params)
		assert.Nil(tt, err)
		assert.Equal(tt, len(res.Payload.Users), 1)

		assert.Equal(tt, []*models.User{expected[1]}, res.Payload.Users)
		assert.Equal(tt, "", res.Payload.Next)
	})

	t.Run("get users filter by country", func(tt *testing.T) {
		countryCode := "FR"
		params := user.NewGetUsersByFiltersParams()
		params.Body.Filters = models.Filters{
			CountryCode: &countryCode,
		}

		res, err := umClient.User.GetUsersByFilters(params)
		assert.Nil(tt, err)

		assert.Equal(tt, []*models.User{}, res.Payload.Users)
		assert.Equal(tt, "", res.Payload.Next)
	})
}

func createUser(u *models.User, password string) {
	creatingParams := user.NewCreateUserParams()
	creatingParams.Body = &models.CreatingUser{
		UserInfo: u.UserInfo,
		Password: password,
	}

	res, err := umClient.User.CreateUser(creatingParams)
	if err != nil {
		log.Fatalf("fail create user: %s", err.Error())
	}
	u.ID = res.Payload.ID
}
