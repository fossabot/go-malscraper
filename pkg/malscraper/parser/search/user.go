package search

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/pkg/malscraper/config"
	"github.com/rl404/go-malscraper/pkg/malscraper/constant"
	"github.com/rl404/go-malscraper/pkg/malscraper/model/common"
	model "github.com/rl404/go-malscraper/pkg/malscraper/model/search"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser"
	"github.com/rl404/go-malscraper/pkg/malscraper/utils"
)

// UserParser is parser for MyAnimeList user search result list.
// Example: https://myanimelist.net/users.php?q=rl404
type UserParser struct {
	parser.BaseParser
	Query string
	Page  int
	Data  []model.User
}

// InitUserParser to initiate all fields and data of UserParser.
func InitUserParser(config config.Config, query string, page ...int) (user UserParser, err error) {
	user.Query = query
	user.Page = 0
	user.Config = config

	if len(page) > 0 {
		user.Page = 24 * (page[0] - 1)
	}

	if len(user.Query) < 3 {
		user.ResponseCode = constant.BadRequestCode
		return user, common.Err3LettersSearch
	}

	// Checking to redis if using redis in config.
	// Redis key's pattern is `search-user:{query},{page}`.
	redisKey := constant.RedisSearchUser + ":" + user.Query + "," + strconv.Itoa(user.Page)
	if config.RedisClient != nil {
		found, err := utils.UnmarshalFromRedis(config.RedisClient, redisKey, &user.Data)
		if err != nil {
			user.SetResponse(constant.InternalErrorCode, err.Error())
			return user, err
		}

		if found {
			user.SetResponse(constant.SuccessCode, constant.SuccessMessage)
			return user, nil
		}
	}

	// Get MyAnimeList HTML source page and initiate the parser.
	err = user.InitParser("/users.php?q="+user.Query+"&show="+strconv.Itoa(user.Page), "#content")
	if err != nil {
		return user, err
	}

	// Fill in data field.
	user.setAllDetail()

	// Save data field to redis if using redis in config.
	if config.RedisClient != nil {
		go utils.SaveToRedis(config.RedisClient, redisKey, user.Data, config.CacheTime)
	}

	return user, nil
}

// setAllDetail to set a;; user detail information.
func (up *UserParser) setAllDetail() {
	var users []model.User
	up.Parser.Find("td.borderClass").Each(func(i int, eachUser *goquery.Selection) {
		users = append(users, model.User{
			Name:       up.getName(eachUser),
			Image:      up.getImage(eachUser),
			LastOnline: up.getLastOnline(eachUser),
		})
	})

	up.Data = users
}

// getName to get user name.
func (up *UserParser) getName(eachUser *goquery.Selection) string {
	return eachUser.Find("a").First().Text()
}

// getImage to get user image.
func (up *UserParser) getImage(eachUser *goquery.Selection) string {
	image, _ := eachUser.Find("img").First().Attr("data-src")
	return utils.URLCleaner(image, "image", up.Config.CleanImageURL)
}

// getLastOnline to get user last online date.
func (up *UserParser) getLastOnline(eachUser *goquery.Selection) string {
	return strings.TrimSpace(eachUser.Find("small").First().Text())
}
