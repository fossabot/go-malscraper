package top

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/pkg/malscraper/config"
	"github.com/rl404/go-malscraper/pkg/malscraper/constant"
	model "github.com/rl404/go-malscraper/pkg/malscraper/model/top"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser"
	"github.com/rl404/go-malscraper/pkg/malscraper/utils"
)

// PeopleParser is parser for MyAnimeList top people list.
// Example: https://myanimelist.net/people.php
type PeopleParser struct {
	parser.BaseParser
	Page int
	Data []model.People
}

// InitPeopleParser to initiate all fields and data of PeopleParser.
func InitPeopleParser(config config.Config, page ...int) (people PeopleParser, err error) {
	people.Page = 0
	people.Config = config

	if len(page) > 0 {
		people.Page = 50 * (page[0] - 1)
	}

	// Checking to redis if using redis in config.
	// Redis key's pattern is `top-people:{page}`.
	redisKey := constant.RedisGetTopPeople + ":" + strconv.Itoa(people.Page)
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
	err = people.InitParser("/people.php?limit="+strconv.Itoa(people.Page), "#content")
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

// setAllDetail to set all top people detail.
func (pp *PeopleParser) setAllDetail() {
	var topList []model.People
	area := pp.Parser.Find(".people-favorites-ranking-table")
	area.Find("tr.ranking-list").Each(func(i int, eachPeople *goquery.Selection) {
		nameArea := eachPeople.Find(".people")

		topList = append(topList, model.People{
			Rank:         pp.getRank(eachPeople),
			ID:           pp.getID(nameArea),
			Name:         pp.getName(nameArea),
			JapaneseName: pp.getJapaneseName(nameArea),
			Image:        pp.getImage(nameArea),
			Birthday:     pp.getBirthday(eachPeople),
			Favorite:     pp.getFavorite(eachPeople),
		})
	})

	pp.Data = topList
}

// getRank to get people rank.
func (pp *PeopleParser) getRank(eachPeople *goquery.Selection) int {
	rank := eachPeople.Find("td").First().Find("span").Text()
	return utils.StrToNum(rank)
}

//getID to get people id.
func (pp *PeopleParser) getID(nameArea *goquery.Selection) int {
	id, _ := nameArea.Find("a").First().Attr("href")
	id = utils.GetValueFromSplit(id, "/", 4)
	return utils.StrToNum(id)
}

// getName to get people name.
func (pp *PeopleParser) getName(nameArea *goquery.Selection) string {
	return nameArea.Find(".information a").Text()
}

// getJapaneseName to get people japanese name.
func (pp *PeopleParser) getJapaneseName(nameArea *goquery.Selection) string {
	japName := nameArea.Find(".information span").Text()

	if japName != "" {
		japName = japName[1 : len(japName)-1]
	}

	return japName
}

// getImage to get people image.
func (pp *PeopleParser) getImage(nameArea *goquery.Selection) string {
	image, _ := nameArea.Find("img").First().Attr("data-src")
	return utils.URLCleaner(image, "image", pp.Config.CleanImageURL)
}

// getBirthday to get people birthday.
func (pp *PeopleParser) getBirthday(eachPeople *goquery.Selection) string {
	day := eachPeople.Find(".birthday").Text()
	day = strings.TrimSpace(day)

	r := regexp.MustCompile(`\s+`)
	day = r.ReplaceAllString(day, " ")

	if day == "Unknown" {
		day = ""
	}

	return day
}

// getFavorite to get people number favorite.
func (pp *PeopleParser) getFavorite(eachPeople *goquery.Selection) int {
	fav := eachPeople.Find(".favorites").Text()
	return utils.StrToNum(fav)
}
