package search

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/helper"
	"github.com/rl404/go-malscraper/model"
)

// SearchAnimeMangaModel is an extended model from MainModel for anime/manga search list.
type SearchAnimeMangaModel struct {
	model.MainModel
	Type  string
	Query string
	Page  int
	Data  []SearchAnimeMangaData
}

// InitSearchAnimeMangaModel to initiate fields in parent (MainModel) model.
func (i *SearchAnimeMangaModel) InitSearchAnimeMangaModel(t string, query string, page int) ([]SearchAnimeMangaData, int, string) {
	i.Type = t
	i.Query = query
	i.Page = 50 * (page - 1)

	if len(i.Query) < 3 {
		i.SetMessage(400, "Search query needs at least 3 letters")
		return i.Data, i.ResponseCode, i.ErrorMessage
	}

	i.InitModel("/"+i.Type+".php?q="+i.Query+"&show="+strconv.Itoa(i.Page), "div[class^=js-categories-seasonal]")

	i.Parser = i.GetParser(i.Url)

	if i.ResponseCode != 200 {
		return i.Data, i.ResponseCode, i.ErrorMessage
	}

	i.SetAllDetail()

	return i.Data, i.ResponseCode, i.ErrorMessage
}

// SetAllDetail to fill all anime/manga search list.
func (i *SearchAnimeMangaModel) SetAllDetail() {
	var searchList []SearchAnimeMangaData
	area := i.Parser.Find("div[class^=js-categories-seasonal] table")
	area.Find("tr").EachWithBreak(func(j int, eachSearch *goquery.Selection) bool {
		if eachSearch.Find(".picSurround").Text() == "" {
			return true
		}

		nameArea := eachSearch.Find("td:nth-of-type(2)")

		searchList = append(searchList, SearchAnimeMangaData{
			Image:   i.GetImage(eachSearch),
			Id:      i.GetId(nameArea),
			Title:   i.GetTitle(nameArea),
			Summary: i.GetSummary(nameArea),
			Type:    i.GetType(eachSearch),
			Episode: i.GetProgress("anime", eachSearch),
			Volume:  i.GetProgress("manga", eachSearch),
			Score:   i.GetScore(eachSearch),
		})

		return true
	})

	i.Data = searchList
}

// GetImage to get anime/manga image.
func (i *SearchAnimeMangaModel) GetImage(eachSearch *goquery.Selection) string {
	image, _ := eachSearch.Find("td a img").Attr("data-src")
	return helper.ImageUrlCleaner(image)
}

// GetId to get anime/manga id.
func (i *SearchAnimeMangaModel) GetId(nameArea *goquery.Selection) string {
	id, _ := nameArea.Find("div[id^=sarea]").Attr("id")
	id = strings.Replace(id, "sarea", "", -1)

	return id
}

// GetTitle to get anime/manga title.
func (i *SearchAnimeMangaModel) GetTitle(nameArea *goquery.Selection) string {
	return nameArea.Find("strong").First().Text()
}

// GetSummary to get anime/manga summary.
func (i *SearchAnimeMangaModel) GetSummary(nameArea *goquery.Selection) string {
	summary := nameArea.Find(".pt4").Text()

	return strings.Replace(summary, "read more.", "", -1)
}

// GetType to get anime/manga type.
func (i *SearchAnimeMangaModel) GetType(eachSearch *goquery.Selection) string {
	t := eachSearch.Find("td:nth-of-type(3)").Text()
	return strings.TrimSpace(t)
}

// GetProgress to get anime/manga episode/volume.
func (i *SearchAnimeMangaModel) GetProgress(t string, eachSearch *goquery.Selection) string {
	if i.Type != t {
		return ""
	}

	progress := eachSearch.Find("td:nth-of-type(4)").Text()
	progress = strings.TrimSpace(progress)

	if progress == "-" {
		return ""
	}
	return progress
}

// GetScore to get anime/manga score.
func (i *SearchAnimeMangaModel) GetScore(eachSearch *goquery.Selection) string {
	score := eachSearch.Find("td:nth-of-type(5)").Text()
	score = strings.TrimSpace(score)

	if score == "N/A" {
		return ""
	}
	return score
}
