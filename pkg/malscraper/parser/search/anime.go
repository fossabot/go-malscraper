package search

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/pkg/malscraper/constant"
	"github.com/rl404/go-malscraper/pkg/malscraper/model/common"
	model "github.com/rl404/go-malscraper/pkg/malscraper/model/search"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser"
	"github.com/rl404/go-malscraper/pkg/malscraper/utils"
)

// AnimeParser is parser for MyAnimeList anime search result list.
// Example: https://myanimelist.net/anime.php?q=naruto
type AnimeParser struct {
	parser.BaseParser
	Query model.Query
	Data  []model.Anime
}

// InitAdvAnimeParser to initiate all fields and data of AnimeParser.
func InitAdvAnimeParser(queryObj model.Query) (anime AnimeParser, err error) {
	anime.Query = queryObj

	if len(anime.Query.Query) < 3 {
		anime.ResponseCode = constant.BadRequestCode
		return anime, common.Err3LettersSearch
	}

	u, _ := url.Parse("/anime.php")
	q := utils.SetSearchParams(u, queryObj)
	u.RawQuery = q.Encode()

	err = anime.InitParser(u.String(), "div.js-categories-seasonal")
	if err != nil {
		return anime, err
	}

	anime.setAllDetail()
	return anime, nil
}

// InitAnimeParser to initiate basic fields for InitAdvAnimeParser.
func InitAnimeParser(query string, page ...int) (anime AnimeParser, err error) {
	var queryObj model.Query
	queryObj.Query = query

	if len(page) > 0 {
		queryObj.Page = page[0]
	}

	return InitAdvAnimeParser(queryObj)
}

// setAllDetail to fill all anime/manga search list.
func (ap *AnimeParser) setAllDetail() {
	var searchList []model.Anime
	area := ap.Parser.Find("table")
	area.Find("tr").EachWithBreak(func(i int, eachSearch *goquery.Selection) bool {
		if eachSearch.Find(".picSurround").Text() == "" {
			return true
		}

		nameArea := eachSearch.Find("td:nth-of-type(2)")

		searchList = append(searchList, model.Anime{
			Image:     ap.getImage(eachSearch),
			ID:        ap.getID(nameArea),
			Title:     ap.getTitle(nameArea),
			Summary:   ap.getSummary(nameArea),
			Type:      ap.getType(eachSearch),
			Episode:   ap.getProgress(eachSearch),
			Score:     ap.getScore(eachSearch),
			StartDate: ap.getStartDate(eachSearch),
			EndDate:   ap.getEndDate(eachSearch),
			Member:    ap.getMember(eachSearch),
			Rated:     ap.getRated(eachSearch),
		})

		return true
	})

	ap.Data = searchList
}

// getImage to get anime/manga image.
func (ap *AnimeParser) getImage(eachSearch *goquery.Selection) string {
	image, _ := eachSearch.Find("td a img").Attr("data-src")
	return utils.ImageURLCleaner(image)
}

// getID to get anime/manga id.
func (ap *AnimeParser) getID(nameArea *goquery.Selection) int {
	id, _ := nameArea.Find("div[id^=sarea]").Attr("id")
	id = strings.Replace(id, "sarea", "", -1)
	return utils.StrToNum(id)
}

// getTitle to get anime/manga title.
func (ap *AnimeParser) getTitle(nameArea *goquery.Selection) string {
	return nameArea.Find("strong").First().Text()
}

// getSummary to get anime/manga summary.
func (ap *AnimeParser) getSummary(nameArea *goquery.Selection) string {
	summary := nameArea.Find(".pt4").Text()
	return strings.Replace(summary, "read more.", "", -1)
}

// getType to get anime/manga type.
func (ap *AnimeParser) getType(eachSearch *goquery.Selection) string {
	t := eachSearch.Find("td:nth-of-type(3)").Text()
	return strings.TrimSpace(t)
}

// getProgress to get anime/manga episode/volume.
func (ap *AnimeParser) getProgress(eachSearch *goquery.Selection) int {
	progress := eachSearch.Find("td:nth-of-type(4)").Text()
	progress = strings.TrimSpace(progress)

	if progress == "-" {
		return 0
	}

	return utils.StrToNum(progress)
}

// getScore to get anime/manga score.
func (ap *AnimeParser) getScore(eachSearch *goquery.Selection) float64 {
	score := eachSearch.Find("td:nth-of-type(5)").Text()
	score = strings.TrimSpace(score)

	if score == "N/A" {
		return 0.0
	}

	return utils.StrToFloat(score)
}

// getStartDate to get anime/manga start date.
func (ap *AnimeParser) getStartDate(eachSearch *goquery.Selection) string {
	sDate := eachSearch.Find("td:nth-of-type(6)").Text()
	sDate = strings.TrimSpace(sDate)

	if sDate == "-" {
		return ""
	}

	return sDate
}

// getEndDate to get anime/manga end date.
func (ap *AnimeParser) getEndDate(eachSearch *goquery.Selection) string {
	eDate := eachSearch.Find("td:nth-of-type(7)").Text()
	eDate = strings.TrimSpace(eDate)

	if eDate == "-" {
		return ""
	}

	return eDate
}

// getMember to get anime/manga member.
func (ap *AnimeParser) getMember(eachSearch *goquery.Selection) int {
	member := eachSearch.Find("td:nth-of-type(8)").Text()
	return utils.StrToNum(member)
}

// getRated to get anime/manga rated.
func (ap *AnimeParser) getRated(eachSearch *goquery.Selection) string {
	rated := eachSearch.Find("td:nth-of-type(9)").Text()
	rated = strings.TrimSpace(rated)

	if rated == "-" {
		return ""
	}

	return rated
}
