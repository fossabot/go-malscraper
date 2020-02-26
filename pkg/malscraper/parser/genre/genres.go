package genre

import (
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/pkg/malscraper/config"
	"github.com/rl404/go-malscraper/pkg/malscraper/constant"
	"github.com/rl404/go-malscraper/pkg/malscraper/model/common"
	model "github.com/rl404/go-malscraper/pkg/malscraper/model/genre"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser"
	"github.com/rl404/go-malscraper/pkg/malscraper/utils"
)

// GenresParser is parser for MyAnimeList all anime & manga genres.
// Example: https://myanimelist.net/anime.php
//          https://myanimelist.net/manga.php
type GenresParser struct {
	parser.BaseParser
	Type string
	Data []model.Genre
}

// InitGenresParser to initiate all fields and data of GenresParser.
func InitGenresParser(config config.Config, gType string) (genres GenresParser, err error) {
	genres.Type = gType
	genres.Config = config

	if !utils.InArray(constant.MainType, genres.Type) {
		genres.ResponseCode = constant.BadRequestCode
		return genres, common.ErrInvalidMainType
	}

	// Checking to redis if using redis in config.
	// Redis key's pattern is `genres:{type}`.
	redisKey := constant.RedisGetGenres + ":" + genres.Type
	if config.RedisClient != nil {
		found, err := utils.UnmarshalFromRedis(config.RedisClient, redisKey, &genres.Data)
		if err != nil {
			genres.SetResponse(constant.InternalErrorCode, err.Error())
			return genres, err
		}

		if found {
			genres.SetResponse(constant.SuccessCode, constant.SuccessMessage)
			return genres, nil
		}
	}

	// Get MyAnimeList HTML source page and initiate the parser.
	err = genres.InitParser("/"+genres.Type+".php", ".anime-manga-search .genre-link")
	if err != nil {
		return genres, err
	}

	// Fill in data field.
	genres.setAllDetail()

	// Save data field to redis if using redis in config.
	if config.RedisClient != nil {
		go utils.SaveToRedis(config.RedisClient, redisKey, genres.Data, config.CacheTime)
	}

	return genres, nil
}

// setAllDetail to set all genre detail information.
func (gp *GenresParser) setAllDetail() {
	var genres []model.Genre
	gp.Parser.Find(".genre-list a").Each(func(i int, area *goquery.Selection) {
		link, _ := area.Attr("href")
		id := utils.GetValueFromSplit(link, "/", 3)
		name := utils.GetValueFromSplit(link, "/", 4)
		genres = append(genres, model.Genre{
			ID:    utils.StrToNum(id),
			Name:  strings.Replace(name, "_", "", -1),
			Count: gp.getCount(area),
		})
	})
	gp.Data = genres
}

// getCount to get genre count.
func (gp *GenresParser) getCount(area *goquery.Selection) int {
	count := area.Text()
	r, _ := regexp.Compile(`\([0-9,]+\)`)
	count = r.FindString(count)
	count = count[1 : len(count)-1]
	return utils.StrToNum(count)
}
