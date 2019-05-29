package malscraper

import (
	"github.com/rl404/go-malscraper/model/user"
)

func MalScraper() *MalService {
	m := &MalService{}
	return m
}

type MalService struct {
	User *UserService
	// General		*GeneralService
	// Additional	*AdditionalService
	// Lists		*ListsService
	// Search		*SearchService
	// Seasonal		*SeasonalService
	// Top			*TopService
}

type UserService struct{}

func (u *UserService) GetUser(username string) (user.UserData, int, string) {
	var userModel user.UserModel
	return userModel.InitUserModel(username)
}
