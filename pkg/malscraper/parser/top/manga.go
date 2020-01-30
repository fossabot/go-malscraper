package top

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/pkg/malscraper/config"
	"github.com/rl404/go-malscraper/pkg/malscraper/constant"
	"github.com/rl404/go-malscraper/pkg/malscraper/model/common"
	model "github.com/rl404/go-malscraper/pkg/malscraper/model/top"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser"
	"github.com/rl404/go-malscraper/pkg/malscraper/utils"
)

// MangaParser is parser for MyAnimeList top manga list.
// Example: https://mymangalist.net/topmanga.php
type MangaParser struct {
	parser.BaseParser
	StringType string
	Type       int
	Page       int
	Data       []model.Manga
}

// InitMangaParser to initiate all fields and data of MangaParser.
func InitMangaParser(config config.Config, params ...int) (manga MangaParser, err error) {
	manga.StringType = ""
	manga.Type = 0
	manga.Page = 0
	manga.Config = config

	for i, param := range params {
		switch i {
		case 0:
			if param > len(constant.TopMangaTypes)-1 || param < 0 {
				manga.ResponseCode = 400
				return manga, common.ErrInvalidMainType
			}

			manga.Type = param
			manga.StringType = constant.TopMangaTypes[param]
		case 1:
			manga.Page = 50 * (param - 1)
		}
	}

	// Checking to redis if using redis in config.
	// Redis key's pattern is `top-manga:{type},{page}`.
	redisKey := constant.RedisGetTopManga + ":" + strconv.Itoa(manga.Type) + "," + strconv.Itoa(manga.Page)
	if config.RedisClient != nil {
		found, err := utils.UnmarshalFromRedis(config.RedisClient, redisKey, &manga.Data)
		if err != nil {
			manga.SetResponse(500, err.Error())
			return manga, err
		}

		if found {
			manga.SetResponse(200, constant.SuccessMessage)
			return manga, nil
		}
	}

	// Get MyAnimeList HTML source page and initiate the parser.
	err = manga.InitParser("/topmanga.php?type="+manga.StringType+"&limit="+strconv.Itoa(manga.Page), "#content")
	if err != nil {
		return manga, err
	}

	// Fill in data field.
	manga.setAllDetail()

	// Save data field to redis if using redis in config.
	if config.RedisClient != nil {
		go utils.SaveToRedis(config.RedisClient, redisKey, manga.Data, config.CacheTime)
	}

	return manga, nil
}

// setAllDetail to set all top manga detail.
func (ap *MangaParser) setAllDetail() {
	var topList []model.Manga
	area := ap.Parser.Find("table")
	area.Find("tr.ranking-list").Each(func(j int, eachTop *goquery.Selection) {
		nameArea := eachTop.Find("td .detail")
		infoArea, _ := nameArea.Find("div.information").Html()
		parsedInfo := strings.Split(infoArea, "<br/>")

		topList = append(topList, model.Manga{
			Rank:      ap.getRank(eachTop),
			Image:     ap.getImage(eachTop),
			ID:        ap.getID(nameArea),
			Title:     ap.getTitle(nameArea),
			Type:      ap.getType(parsedInfo),
			Volume:    ap.getEpCh(parsedInfo),
			StartDate: ap.getDate(parsedInfo, 0),
			EndDate:   ap.getDate(parsedInfo, 1),
			Member:    ap.getMember(parsedInfo),
			Score:     ap.getScore(eachTop),
		})
	})

	ap.Data = topList
}

// getRank to get manga rank.
func (ap *MangaParser) getRank(eachTop *goquery.Selection) int {
	rank := eachTop.Find("td").First().Find("span").Text()
	return utils.StrToNum(rank)
}

// getImage to get manga image.
func (ap *MangaParser) getImage(eachTop *goquery.Selection) string {
	image, _ := eachTop.Find("td:nth-of-type(2) a img").Attr("data-src")
	return utils.URLCleaner(image, "image", ap.Config.CleanImageURL)
}

// getID to get manga id.
func (ap *MangaParser) getID(nameArea *goquery.Selection) int {
	id, _ := nameArea.Find("div").First().Attr("id")
	idInt, _ := strconv.Atoi(strings.Replace(id, "area", "", -1))
	return idInt
}

// getTitle to get manga title.
func (ap *MangaParser) getTitle(nameArea *goquery.Selection) string {
	return nameArea.Find("a").First().Text()
}

// getType to get manga type.
func (ap *MangaParser) getType(parsedInfo []string) string {
	splitType := strings.Split(strings.TrimSpace(parsedInfo[0]), " ")
	return splitType[0]
}

// getEpCh to get manga episode/chapter.
func (ap *MangaParser) getEpCh(parsedInfo []string) int {
	splitEpCh := strings.Split(strings.TrimSpace(parsedInfo[0]), " ")
	if splitEpCh[1][1:] == "?" {
		return 0
	}
	return utils.StrToNum(splitEpCh[1][1:])
}

// getDate to get manga start/end date.
func (ap *MangaParser) getDate(parsedInfo []string, t int) string {
	splitDate := strings.Split(strings.TrimSpace(parsedInfo[1]), "-")
	return strings.TrimSpace(splitDate[t])
}

// getMember to get manga member number.
func (ap *MangaParser) getMember(parsedInfo []string) int {
	member := strings.TrimSpace(parsedInfo[2])
	member = strings.Replace(member, "members", "", -1)
	member = strings.Replace(member, "favorites", "", -1)
	member = strings.Replace(member, ",", "", -1)
	return utils.StrToNum(member)
}

// getScore to get manga score.
func (ap *MangaParser) getScore(eachTop *goquery.Selection) float64 {
	score := eachTop.Find("td:nth-of-type(3)").Text()
	score = strings.TrimSpace(strings.Replace(score, "N/A", "", -1))
	return utils.StrToFloat(score)
}
