package user

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/rl404/go-malscraper/model"
)

// UserCoverModel is a model for user cover.
type UserCoverModel struct {
	model.MainModel
	User     string
	Type     string
	Query    string
	Data     string
	UserList []UserListData
}

// InitUserCoverModel to initiate fields.
func (u *UserCoverModel) InitUserCoverModel(user string, typ string, query string) (string, int, string) {
	u.User = user
	u.Type = typ

	if query != "" {
		u.Query, _ = url.QueryUnescape(query)
	} else {
		u.Query = "tr:hover .animetitle[href*='/{id}/']:before{background-image:url({url})}"
	}

	var userListModel UserListModel
	u.UserList, u.ResponseCode, u.ErrorMessage = userListModel.InitUserListModel(u.User, u.Type, 7)

	if u.ResponseCode != 200 {
		return "", u.ResponseCode, u.ErrorMessage
	}

	u.CreateCSS()

	return u.Data, u.ResponseCode, u.ErrorMessage
}

// CreateCSS to create css image cover from user list data.
func (u *UserCoverModel) CreateCSS() {
	coverString := ""

	for _, c := range u.UserList {
		temp := ""
		if u.Type == "anime" {
			temp = strings.Replace(u.Query, "{id}", strconv.Itoa(c.AnimeId), -1)
			temp = strings.Replace(temp, "{url}", c.AnimeImage, -1)
		} else {
			temp = strings.Replace(u.Query, "{id}", strconv.Itoa(c.MangaId), -1)
			temp = strings.Replace(temp, "{url}", c.MangaImage, -1)
		}
		coverString += temp + "\n"
	}

	u.Data = coverString
}
