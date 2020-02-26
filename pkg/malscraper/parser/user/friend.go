package user

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/pkg/malscraper/config"
	"github.com/rl404/go-malscraper/pkg/malscraper/constant"
	model "github.com/rl404/go-malscraper/pkg/malscraper/model/user"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser"
	"github.com/rl404/go-malscraper/pkg/malscraper/utils"
)

// UserFriendParser is parser for MyAnimeList user friend list.
// Example: https://myanimelist.net/profile/rl404/friends
type UserFriendParser struct {
	parser.BaseParser
	Username string
	Page     int
	Data     []model.Friend
}

// InitUserFriendParser to initiate all fields and data of UserFriendParser.
func InitUserFriendParser(config config.Config, username string, page ...int) (userFriend UserFriendParser, err error) {
	userFriend.Username = username
	userFriend.Page = 0
	userFriend.Config = config

	if len(page) > 0 {
		userFriend.Page = 100 * (page[0] - 1)
	}

	// Checking to redis if using redis in config.
	// Redis key's pattern is `user-friend:{name},{page}`.
	redisKey := constant.RedisGetUserFriend + ":" + userFriend.Username + "," + strconv.Itoa(userFriend.Page)
	if config.RedisClient != nil {
		found, err := utils.UnmarshalFromRedis(config.RedisClient, redisKey, &userFriend.Data)
		if err != nil {
			userFriend.SetResponse(constant.InternalErrorCode, err.Error())
			return userFriend, err
		}

		if found {
			userFriend.SetResponse(constant.SuccessCode, constant.SuccessMessage)
			return userFriend, nil
		}
	}

	// Get MyAnimeList HTML source page and initiate the parser.
	err = userFriend.InitParser("/profile/"+userFriend.Username+"/friends?offset="+strconv.Itoa(userFriend.Page), "#content")
	if err != nil {
		return userFriend, err
	}

	// Fill in data field.
	userFriend.setAllDetail()

	// Save data field to redis if using redis in config.
	if config.RedisClient != nil {
		go utils.SaveToRedis(config.RedisClient, redisKey, userFriend.Data, config.CacheTime)
	}

	return userFriend, nil
}

// setAllDetail to set all user friend detail information.
func (user *UserFriendParser) setAllDetail() {
	var friendList []model.Friend
	area := user.Parser.Find(".majorPad")
	area.Find(".friendHolder").Each(func(i int, friend *goquery.Selection) {
		friendArea := friend.Find(".friendBlock")
		friendList = append(friendList, model.Friend{
			Name:        user.getName(friendArea),
			Image:       user.getImage(friendArea),
			LastOnline:  user.getLastOnline(friendArea),
			FriendSince: user.getFriendSince(friendArea),
		})
	})
	user.Data = friendList
}

// getName to get user's friend name.
func (user *UserFriendParser) getName(friendArea *goquery.Selection) string {
	name, _ := friendArea.Find("a").Attr("href")
	return utils.GetValueFromSplit(name, "/", 4)
}

// getImage to get user's friend image.
func (user *UserFriendParser) getImage(friendArea *goquery.Selection) string {
	image, _ := friendArea.Find("a img").Attr("src")
	return utils.URLCleaner(image, "image", user.Config.CleanImageURL)
}

// getLastOnline to get user's friend last online date.
func (user *UserFriendParser) getLastOnline(friendArea *goquery.Selection) string {
	lastOnline := friendArea.Find(".friendBlock div:nth-of-type(3)").Text()
	return strings.TrimSpace(lastOnline)
}

// getFriendSince to get user's date start become friend.
func (user *UserFriendParser) getFriendSince(friendArea *goquery.Selection) string {
	friendSince := friendArea.Find(".friendBlock div:nth-of-type(4)").Text()
	friendSince = strings.Replace(friendSince, "Friends since", "", -1)
	return strings.TrimSpace(friendSince)
}
