package search

import (
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/pkg/malscraper/model/common"
	model "github.com/rl404/go-malscraper/pkg/malscraper/model/search"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser"
	"github.com/rl404/go-malscraper/pkg/malscraper/utils"
)

// MangaParser is parser for MyAnimeList manga search result list.
// Example: https://mymangalist.net/manga.php?q=naruto
type MangaParser struct {
	parser.BaseParser
	Query model.Query
	Data  []model.Manga
}

// InitAdvMangaParser to initiate all fields and data of MangaParser.
func InitAdvMangaParser(queryObj model.Query) (manga MangaParser, err error) {
	manga.Query = queryObj

	if len(manga.Query.Query) < 3 {
		manga.ResponseCode = 400
		return manga, common.Err3LettersSearch
	}

	u, _ := url.Parse("/manga.php")
	q := utils.SetSearchParams(u, queryObj)
	u.RawQuery = q.Encode()

	err = manga.InitParser(u.String(), "div.js-categories-seasonal")
	if err != nil {
		return manga, err
	}

	manga.setAllDetail()
	return manga, nil
}

// InitMangaParser to initiate basic fields for InitAdvMangaParser.
func InitMangaParser(query string, page ...int) (manga MangaParser, err error) {
	var queryObj model.Query
	queryObj.Query = query

	if len(page) > 0 {
		queryObj.Page = page[0]
	}

	return InitAdvMangaParser(queryObj)
}

// setAllDetail to fill all manga search list.
func (mp *MangaParser) setAllDetail() {
	var searchList []model.Manga
	area := mp.Parser.Find("table")
	area.Find("tr").EachWithBreak(func(i int, eachSearch *goquery.Selection) bool {
		if eachSearch.Find(".picSurround").Text() == "" {
			return true
		}

		nameArea := eachSearch.Find("td:nth-of-type(2)")

		searchList = append(searchList, model.Manga{
			Image:     mp.getImage(eachSearch),
			ID:        mp.getID(nameArea),
			Title:     mp.getTitle(nameArea),
			Summary:   mp.getSummary(nameArea),
			Type:      mp.getType(eachSearch),
			Volume:    mp.getVolume(eachSearch),
			Chapter:   mp.getChapter(eachSearch),
			Score:     mp.getScore(eachSearch),
			StartDate: mp.getStartDate(eachSearch),
			EndDate:   mp.getEndDate(eachSearch),
			Member:    mp.getMember(eachSearch),
		})

		return true
	})

	mp.Data = searchList
}

// getImage to get manga image.
func (mp *MangaParser) getImage(eachSearch *goquery.Selection) string {
	image, _ := eachSearch.Find("td a img").Attr("data-src")
	return utils.ImageURLCleaner(image)
}

// getID to get manga id.
func (mp *MangaParser) getID(nameArea *goquery.Selection) int {
	id, _ := nameArea.Find("div[id^=sarea]").Attr("id")
	id = strings.Replace(id, "sarea", "", -1)
	return utils.StrToNum(id)
}

// getTitle to get manga title.
func (mp *MangaParser) getTitle(nameArea *goquery.Selection) string {
	return nameArea.Find("strong").First().Text()
}

// getSummary to get manga summary.
func (mp *MangaParser) getSummary(nameArea *goquery.Selection) string {
	summary := nameArea.Find(".pt4").Text()
	return strings.Replace(summary, "read more.", "", -1)
}

// getType to get manga type.
func (mp *MangaParser) getType(eachSearch *goquery.Selection) string {
	t := eachSearch.Find("td:nth-of-type(3)").Text()
	return strings.TrimSpace(t)
}

// getVolume to get manga volume.
func (mp *MangaParser) getVolume(eachSearch *goquery.Selection) int {
	progress := eachSearch.Find("td:nth-of-type(4)").Text()
	progress = strings.TrimSpace(progress)

	if progress == "-" {
		return 0
	}

	return utils.StrToNum(progress)
}

// getChapter to get manga chapter.
func (mp *MangaParser) getChapter(eachSearch *goquery.Selection) int {
	progress := eachSearch.Find("td:nth-of-type(5)").Text()
	progress = strings.TrimSpace(progress)

	if progress == "-" {
		return 0
	}

	return utils.StrToNum(progress)
}

// getScore to get manga score.
func (mp *MangaParser) getScore(eachSearch *goquery.Selection) float64 {
	score := eachSearch.Find("td:nth-of-type(6)").Text()
	score = strings.TrimSpace(score)

	if score == "N/A" {
		return 0.0
	}

	return utils.StrToFloat(score)
}

// getStartDate to get manga start date.
func (mp *MangaParser) getStartDate(eachSearch *goquery.Selection) string {
	sDate := eachSearch.Find("td:nth-of-type(7)").Text()
	sDate = strings.TrimSpace(sDate)

	if sDate == "-" {
		return ""
	}

	return sDate
}

// getEndDate to get manga end date.
func (mp *MangaParser) getEndDate(eachSearch *goquery.Selection) string {
	eDate := eachSearch.Find("td:nth-of-type(8)").Text()
	eDate = strings.TrimSpace(eDate)

	if eDate == "-" {
		return ""
	}

	return eDate
}

// getMember to get manga member.
func (mp *MangaParser) getMember(eachSearch *goquery.Selection) int {
	member := eachSearch.Find("td:nth-of-type(9)").Text()
	return utils.StrToNum(member)
}
