package people

import (
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/pkg/malscraper/config"
	"github.com/rl404/go-malscraper/pkg/malscraper/constant"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser"
	"github.com/rl404/go-malscraper/pkg/malscraper/utils"
)

// PeoplePictureParser is parser for MyAnimeList people picture list.
// Example: https://myanimelist.net/people/1/Tomokazu_Seki/pictures
type PeoplePictureParser struct {
	parser.BaseParser
	ID   int
	Data []string
}

// InitPeoplePictureParser to initiate all fields and data of PeoplePictureParser.
func InitPeoplePictureParser(config config.Config, id int) (picture PeoplePictureParser, err error) {
	picture.ID = id
	picture.Config = config

	// Checking to redis if using redis in config.
	// Redis key's pattern is `people-picture:{id}`.
	redisKey := constant.RedisGetPeoplePicture + ":" + strconv.Itoa(picture.ID)
	if config.RedisClient != nil {
		found, err := utils.UnmarshalFromRedis(config.RedisClient, redisKey, &picture.Data)
		if err != nil {
			picture.SetResponse(500, err.Error())
			return picture, err
		}

		if found {
			picture.SetResponse(200, constant.SuccessMessage)
			return picture, nil
		}
	}

	// Get MyAnimeList HTML source page and initiate the parser.
	err = picture.InitParser("/people/"+strconv.Itoa(picture.ID)+"/a/pictures", "#content table tr td")
	if err != nil {
		return picture, err
	}

	// Fill in data field.
	picture.setAllDetail()

	// Save data field to redis if using redis in config.
	if config.RedisClient != nil {
		go utils.SaveToRedis(config.RedisClient, redisKey, picture.Data, config.CacheTime)
	}

	return picture, nil
}

// setAllDetail to set pictures list.
func (cp *PeoplePictureParser) setAllDetail() {
	var pictures []string

	area := cp.Parser.Next().Find("table").First()
	area.Find("img").Each(func(i int, eachPic *goquery.Selection) {
		link, _ := eachPic.Attr("data-src")
		pictures = append(pictures, link)
	})

	cp.Data = pictures
}
