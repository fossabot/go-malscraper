package search

import (
	"regexp"
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/pkg/malscraper/config"
	"github.com/rl404/go-malscraper/pkg/malscraper/constant"
	"github.com/rl404/go-malscraper/pkg/malscraper/model/common"
	model "github.com/rl404/go-malscraper/pkg/malscraper/model/search"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser"
	"github.com/rl404/go-malscraper/pkg/malscraper/utils"
)

// PeopleParser is parser for MyAnimeList people search result list.
// Example: https://myanimelist.net/people.php?q=kana
type PeopleParser struct {
	parser.BaseParser
	Query string
	Page  int
	Data  []model.People
}

// InitPeopleParser to initiate all fields and data of PeopleParser.
func InitPeopleParser(config config.Config, query string, page ...int) (people PeopleParser, err error) {
	people.Query = query
	people.Page = 0
	people.Config = config

	if len(page) > 0 {
		people.Page = 50 * (page[0] - 1)
	}

	if len(people.Query) < 3 {
		people.ResponseCode = constant.BadRequestCode
		return people, common.Err3LettersSearch
	}

	// Checking to redis if using redis in config.
	// Redis key's pattern is `search-people:{query},{page}`.
	redisKey := constant.RedisSearchPeople + ":" + people.Query + "," + strconv.Itoa(people.Page)
	if config.RedisClient != nil {
		found, err := utils.UnmarshalFromRedis(config.RedisClient, redisKey, &people.Data)
		if err != nil {
			people.SetResponse(constant.InternalErrorCode, err.Error())
			return people, err
		}

		if found {
			people.SetResponse(constant.SuccessCode, constant.SuccessMessage)
			return people, nil
		}
	}

	// Get MyAnimeList HTML source page and initiate the parser.
	err = people.InitParser("/people.php?q="+people.Query+"&show="+strconv.Itoa(people.Page), "#content")
	if err != nil {
		return people, err
	}

	// Fill in data field.
	people.setAllDetail()

	// Save data field to redis if using redis in config.
	if config.RedisClient != nil {
		go utils.SaveToRedis(config.RedisClient, redisKey, people.Data, config.CacheTime)
	}

	return people, nil
}

// setAllDetail to set all people detail information.
func (cp *PeopleParser) setAllDetail() {
	var peoples []model.People

	area := cp.Parser.Find("table")
	area.Find("tr").EachWithBreak(func(i int, eachSearch *goquery.Selection) bool {
		if i == 0 {
			return true
		}

		nameArea := eachSearch.Find("td:nth-of-type(2)")

		peoples = append(peoples, model.People{
			Image:    cp.getImage(eachSearch),
			ID:       cp.getID(nameArea),
			Name:     cp.getName(nameArea),
			Nickname: cp.getNickname(nameArea),
		})

		return true
	})

	cp.Data = peoples
}

// getImage to get people image.
func (cp *PeopleParser) getImage(eachSearch *goquery.Selection) string {
	image, _ := eachSearch.Find("td a img").Attr("data-src")
	return utils.URLCleaner(image, "image", cp.Config.CleanImageURL)
}

// getID to get people id.
func (cp *PeopleParser) getID(nameArea *goquery.Selection) int {
	id, _ := nameArea.Find("a").First().Attr("href")
	id = utils.GetValueFromSplit(id, "/", 2)
	return utils.StrToNum(id)
}

// getName to get people name.
func (cp *PeopleParser) getName(nameArea *goquery.Selection) string {
	return nameArea.Find("a").First().Text()
}

// getNickname to get people nickname.
func (cp *PeopleParser) getNickname(nameArea *goquery.Selection) string {
	nick := nameArea.Find("small").First().Text()
	if nick != "" {
		r := regexp.MustCompile(`\s+`)
		nick = r.ReplaceAllString(nick, " ")
		return nick[1 : len(nick)-1]
	}
	return nick
}
