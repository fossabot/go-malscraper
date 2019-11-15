package top

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/helper"
	"github.com/rl404/go-malscraper/model"
)

// TopAnimeMangaModel is an extended model from MainModel for top anime/manga list.
type TopAnimeMangaModel struct {
	model.MainModel
	SuperType string
	Type      string
	Page      int
	Data      []TopAnimeMangaData
}

// InitTopAnimeMangaModel to initiate fields in parent (MainModel) model.
func (i *TopAnimeMangaModel) InitTopAnimeMangaModel(superType string, t int, p int) ([]TopAnimeMangaData, int, string) {
	i.SuperType = superType
	i.Page = 50 * (p - 1)

	if i.SuperType == "anime" {
		i.Type = helper.GetTopAnimeType()[t]
		i.InitModel("/topanime.php?type="+i.Type+"&limit="+strconv.Itoa(i.Page), "#content")
	} else {
		i.Type = helper.GetTopMangaType()[t]
		i.InitModel("/topmanga.php?type="+i.Type+"&limit="+strconv.Itoa(i.Page), "#content")
	}

	i.Parser = i.GetParser(i.Url)

	if i.ResponseCode != 200 {
		return i.Data, i.ResponseCode, i.ErrorMessage
	}

	i.SetAllDetail()

	return i.Data, i.ResponseCode, i.ErrorMessage
}

// SetAllDetail to set all top anime/manga detail.
func (i *TopAnimeMangaModel) SetAllDetail() {
	var topList []TopAnimeMangaData
	area := i.Parser.Find("table")
	area.Find("tr.ranking-list").Each(func(j int, eachTop *goquery.Selection) {
		nameArea := eachTop.Find("td .detail")
		infoArea, _ := nameArea.Find("div.information").Html()
		parsedInfo := strings.Split(infoArea, "<br/>")

		topList = append(topList, TopAnimeMangaData{
			Rank:      i.GetRank(eachTop),
			Image:     i.GetImage(eachTop),
			Id:        i.GetId(nameArea),
			Title:     i.GetTitle(nameArea),
			Type:      i.GetType(parsedInfo),
			Episode:   i.GetEpCh("anime", parsedInfo),
			Volume:    i.GetEpCh("manga", parsedInfo),
			StartDate: i.GetDate(parsedInfo, 0),
			EndDate:   i.GetDate(parsedInfo, 1),
			Member:    i.GetMember(parsedInfo),
			Score:     i.GetScore(eachTop),
		})
	})

	i.Data = topList
}

// GetRank to get anime/manga rank.
func (i *TopAnimeMangaModel) GetRank(eachTop *goquery.Selection) string {
	rank := eachTop.Find("td").First().Find("span").Text()
	return strings.TrimSpace(rank)
}

// GetImage to get anime/manga image.
func (i *TopAnimeMangaModel) GetImage(eachTop *goquery.Selection) string {
	image, _ := eachTop.Find("td:nth-of-type(2) a img").Attr("data-src")
	return helper.ImageUrlCleaner(image)
}

// GetId to get anime/manga id.
func (i *TopAnimeMangaModel) GetId(nameArea *goquery.Selection) string {
	id, _ := nameArea.Find("div").First().Attr("id")
	return strings.Replace(id, "area", "", -1)
}

// GetTitle to get anime/manga title.
func (i *TopAnimeMangaModel) GetTitle(nameArea *goquery.Selection) string {
	return nameArea.Find("a").First().Text()
}

// GetType to get anime/manga type.
func (i *TopAnimeMangaModel) GetType(parsedInfo []string) string {
	splitType := strings.Split(strings.TrimSpace(parsedInfo[0]), " ")
	return splitType[0]
}

// GetEpCh to get anime/manga episode/chapter.
func (i *TopAnimeMangaModel) GetEpCh(t string, parsedInfo []string) string {
	if i.SuperType != t {
		return ""
	}

	splitEpCh := strings.Split(strings.TrimSpace(parsedInfo[0]), " ")
	if splitEpCh[1][1:] == "?" {
		return ""
	}

	return splitEpCh[1][1:]
}

// GetDate to get anime/manga start/end date.
func (i *TopAnimeMangaModel) GetDate(parsedInfo []string, t int) string {
	splitDate := strings.Split(strings.TrimSpace(parsedInfo[1]), "-")
	return strings.TrimSpace(splitDate[t])
}

// GetMember to get anime/manga member number.
func (i *TopAnimeMangaModel) GetMember(parsedInfo []string) string {
	member := strings.TrimSpace(parsedInfo[2])
	member = strings.Replace(member, "members", "", -1)
	member = strings.Replace(member, "favorites", "", -1)
	member = strings.Replace(member, ",", "", -1)

	return strings.TrimSpace(member)
}

// GetScore to get anime/manga score.
func (i *TopAnimeMangaModel) GetScore(eachTop *goquery.Selection) string {
	score := eachTop.Find("td:nth-of-type(3)").Text()
	return strings.TrimSpace(strings.Replace(score, "N/A", "", -1))
}
