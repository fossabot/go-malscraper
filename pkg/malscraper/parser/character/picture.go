package character

import (
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/pkg/malscraper/config"
	"github.com/rl404/go-malscraper/pkg/malscraper/constant"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser"
	"github.com/rl404/go-malscraper/pkg/malscraper/utils"
)

// PictureParser is parser for MyAnimeList character picture list.
// Example: https://myanimelist.net/character/1/Spike_Spiegel/pictures
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
	// Redis key's pattern is `character-picture:{id}`.
	redisKey := constant.RedisGetCharacterPicture + ":" + strconv.Itoa(picture.ID)
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
	err = picture.InitParser("/character/"+strconv.Itoa(picture.ID)+"/a/pictures", "#content table tr td")
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
func (cp *PictureParser) setAllDetail() {
	var pictures []string

	area := cp.Parser.Next().Find("table").First()
	area.Find("img").Each(func(i int, eachPic *goquery.Selection) {
		link, _ := eachPic.Attr("data-src")
		pictures = append(pictures, link)
	})

	cp.Data = pictures
}
