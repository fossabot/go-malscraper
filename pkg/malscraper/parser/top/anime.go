package top

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/pkg/malscraper/constant"
	"github.com/rl404/go-malscraper/pkg/malscraper/model/common"
	model "github.com/rl404/go-malscraper/pkg/malscraper/model/top"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser"
	"github.com/rl404/go-malscraper/pkg/malscraper/utils"
)

// AnimeParser is parser for MyAnimeList top anime list.
// Example: https://myanimelist.net/topanime.php
type AnimeParser struct {
	parser.BaseParser
	StringType string
	Type       int
	Page       int
	Data       []model.Anime
}

// InitAnimeParser to initiate all fields and data of AnimeParser.
func InitAnimeParser(params ...int) (anime AnimeParser, err error) {
	anime.StringType = ""
	anime.Type = 0
	anime.Page = 0

	for i, param := range params {
		switch i {
		case 0:
			if param > len(constant.TopAnimeTypes)-1 || param < 0 {
				anime.ResponseCode = 400
				return anime, common.ErrInvalidMainType
			}

			anime.Type = param
			anime.StringType = constant.TopAnimeTypes[param]
		case 1:
			anime.Page = 50 * (param - 1)
		}
	}

	err = anime.InitParser("/topanime.php?type="+anime.StringType+"&limit="+strconv.Itoa(anime.Page), "#content")
	if err != nil {
		return anime, err
	}

	anime.setAllDetail()
	return anime, nil
}

// setAllDetail to set all top anime detail.
func (ap *AnimeParser) setAllDetail() {
	var topList []model.Anime
	area := ap.Parser.Find("table")
	area.Find("tr.ranking-list").Each(func(j int, eachTop *goquery.Selection) {
		nameArea := eachTop.Find("td .detail")
		infoArea, _ := nameArea.Find("div.information").Html()
		parsedInfo := strings.Split(infoArea, "<br/>")

		topList = append(topList, model.Anime{
			Rank:      ap.getRank(eachTop),
			Image:     ap.getImage(eachTop),
			ID:        ap.getID(nameArea),
			Title:     ap.getTitle(nameArea),
			Type:      ap.getType(parsedInfo),
			Episode:   ap.getEpCh(parsedInfo),
			StartDate: ap.getDate(parsedInfo, 0),
			EndDate:   ap.getDate(parsedInfo, 1),
			Member:    ap.getMember(parsedInfo),
			Score:     ap.getScore(eachTop),
		})
	})

	ap.Data = topList
}

// getRank to get anime rank.
func (ap *AnimeParser) getRank(eachTop *goquery.Selection) int {
	rank := eachTop.Find("td").First().Find("span").Text()
	return utils.StrToNum(rank)
}

// getImage to get anime image.
func (ap *AnimeParser) getImage(eachTop *goquery.Selection) string {
	image, _ := eachTop.Find("td:nth-of-type(2) a img").Attr("data-src")
	return utils.ImageURLCleaner(image)
}

// getID to get anime id.
func (ap *AnimeParser) getID(nameArea *goquery.Selection) int {
	id, _ := nameArea.Find("div").First().Attr("id")
	idInt, _ := strconv.Atoi(strings.Replace(id, "area", "", -1))
	return idInt
}

// getTitle to get anime title.
func (ap *AnimeParser) getTitle(nameArea *goquery.Selection) string {
	return nameArea.Find("a").First().Text()
}

// getType to get anime type.
func (ap *AnimeParser) getType(parsedInfo []string) string {
	splitType := strings.Split(strings.TrimSpace(parsedInfo[0]), " ")
	return splitType[0]
}

// getEpCh to get anime episode/chapter.
func (ap *AnimeParser) getEpCh(parsedInfo []string) int {
	splitEpCh := strings.Split(strings.TrimSpace(parsedInfo[0]), " ")
	if splitEpCh[1][1:] == "?" {
		return 0
	}
	return utils.StrToNum(splitEpCh[1][1:])
}

// getDate to get anime start/end date.
func (ap *AnimeParser) getDate(parsedInfo []string, t int) string {
	splitDate := strings.Split(strings.TrimSpace(parsedInfo[1]), "-")
	return strings.TrimSpace(splitDate[t])
}

// getMember to get anime member number.
func (ap *AnimeParser) getMember(parsedInfo []string) int {
	member := strings.TrimSpace(parsedInfo[2])
	member = strings.Replace(member, "members", "", -1)
	member = strings.Replace(member, "favorites", "", -1)
	member = strings.Replace(member, ",", "", -1)
	return utils.StrToNum(member)
}

// getScore to get anime score.
func (ap *AnimeParser) getScore(eachTop *goquery.Selection) float64 {
	score := eachTop.Find("td:nth-of-type(3)").Text()
	score = strings.TrimSpace(strings.Replace(score, "N/A", "", -1))
	return utils.StrToFloat(score)
}
