package service

import "github.com/rl404/go-malscraper/model/user"

// UserService for all user-related methods.
type UserService struct{}

// GetUser to get user profile information.
func (u *UserService) GetUser(username string) (user.UserData, int, string) {
	var userModel user.UserModel
	return userModel.InitUserModel(username)
}
