package anime

import (
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/pkg/malscraper/config"
	"github.com/rl404/go-malscraper/pkg/malscraper/constant"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser"
	"github.com/rl404/go-malscraper/pkg/malscraper/utils"
)

// PictureParser is parser for MyAnimeList anime picture list.
// Example: https://myanimelist.net/anime/1/Cowboy_Bebop/pics
type PictureParser struct {
	parser.BaseParser
	ID   int
	Data []string
}

// InitPictureParser to initiate all fields and data of PictureParser.
func InitPictureParser(config config.Config, id int) (picture PictureParser, err error) {
	picture.ID = id
	picture.Config = config

	// Checking to redis if using redis in config.
	// Redis key's pattern is `anime-picture:{id}`.
	redisKey := constant.RedisGetAnimePicture + ":" + strconv.Itoa(picture.ID)
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
	err = picture.InitParser("/anime/"+strconv.Itoa(picture.ID)+"/a/pics", ".js-scrollfix-bottom-rel")
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
func (pp *PictureParser) setAllDetail() {
	var pictures []string

	area := pp.Parser.Find("table").First()
	area.Find("img").Each(func(i int, eachPic *goquery.Selection) {
		link, _ := eachPic.Attr("data-src")
		pictures = append(pictures, link)
	})

	pp.Data = pictures
}
