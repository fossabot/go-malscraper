package service

import "github.com/rl404/go-malscraper/model/user"

// UserService for all user-related methods.
type UserService struct{}

// GetUser to get user profile information.
func (u *UserService) GetUser(username string) (user.UserData, int, string) {
	var userModel user.UserModel
	return userModel.InitUserModel(username)
}

// GetUserList to get user anime/manga list.
func (u *UserService) GetUserList(username string, typ string, status int) ([]user.UserListData, int, string) {
	var userListModel user.UserListModel
	return userListModel.InitUserListModel(username, typ, status)
}

// GetUserCover to get user cover of anime/manga list.
func (u *UserService) GetUserCover(username string, typ string, query string) (string, int, string) {
	var userCoverModel user.UserCoverModel
	return userCoverModel.InitUserCoverModel(username, typ, query)
}
