package user

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/helper"
	"github.com/rl404/go-malscraper/model"
)

// UserFriendModel is an extended model from MainModel for user friend list.
type UserFriendModel struct {
	model.MainModel
	Username string
	Page     int
	Data     []UserFriendData
}

// InitUserFriendModel to initiate fields in parent (MainModel) model.
func (i *UserFriendModel) InitUserFriendModel(u string, p int) ([]UserFriendData, int, string) {
	i.Username = u
	i.Page = 100 * (p - 1)

	i.InitModel("/profile/"+i.Username+"/friends?offset="+strconv.Itoa(i.Page), "#content")

	i.Parser = i.GetParser(i.Url)

	if i.ResponseCode != 200 {
		return i.Data, i.ResponseCode, i.ErrorMessage
	}

	i.SetAllDetail()

	return i.Data, i.ResponseCode, i.ErrorMessage
}

// SetAllDetail to set user friend detail data.
func (i *UserFriendModel) SetAllDetail() {
	var friendList []UserFriendData
	area := i.Parser.Find("#content .majorPad")
	area.Find(".friendHolder").Each(func(j int, eachFriend *goquery.Selection) {
		friend := eachFriend.Find(".friendBlock")

		friendList = append(friendList, UserFriendData{
			Image:       i.GetImage(friend),
			Name:        i.GetName(friend),
			LastOnline:  i.GetLastOnline(friend),
			FriendSince: i.GetFriendSince(friend),
		})
	})

	i.Data = friendList
}

// GetImage to get friend image.
func (i *UserFriendModel) GetImage(friend *goquery.Selection) string {
	image, _ := friend.Find("a img").Attr("src")
	return helper.ImageUrlCleaner(image)
}

// GetName to get friend name.
func (i *UserFriendModel) GetName(friend *goquery.Selection) string {
	name, _ := friend.Find("a").Attr("href")
	splitName := strings.Split(name, "/")
	return splitName[4]
}

// GetLastOnline to get friend last online date.
func (i *UserFriendModel) GetLastOnline(friend *goquery.Selection) string {
	lastOnline := friend.Find(".friendBlock div:nth-of-type(3)").Text()
	return strings.TrimSpace(lastOnline)
}

// GetFriendSince to get date start befriending.
func (i *UserFriendModel) GetFriendSince(friend *goquery.Selection) string {
	friendSince := friend.Find(".friendBlock div:nth-of-type(4)").Text()
	friendSince = strings.Replace(friendSince, "Friends since", "", -1)
	return strings.TrimSpace(friendSince)
}
