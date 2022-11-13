package user

import (
	"github/user-manager/internal/entity"
	"github/user-manager/internal/generated/server/models"
)

func mapEntityUsersToModelUsers(eusers []entity.User) []*models.User {
	musers := make([]*models.User, 0, len(eusers))
	for _, eu := range eusers {
		mu := mapEntityUserToModelUser(eu)
		musers = append(musers, &mu)
	}

	return musers
}

func mapModelUserInfoToEntityUser(userInfo models.UserInfo) entity.User {
	return entity.User{
		FirstName: userInfo.FirstName,
		LastName:  userInfo.LastName,
		Nickname:  userInfo.Nickname,
		Email:     userInfo.Email,
	}
}

func mapEntityUserToModelUser(euser entity.User) models.User {
	return models.User{
		UserInfo: models.UserInfo{
			FirstName:   euser.FirstName,
			LastName:    euser.LastName,
			Nickname:    euser.Nickname,
			Email:       euser.Email,
			CountryCode: euser.Country.Code,
		},
		ID: euser.ID,
	}
}
