package list

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/helper"
	"github.com/rl404/go-malscraper/model"
)

// RecommendationModel is an extended model from MainModel for anime/manga recommendation list.
type RecommendationModel struct {
	model.MainModel
	Type string
	Page int
	Data []RecommendationData
}

// InitReviewModel to initiate fields in parent (MainModel) model.
func (i *RecommendationModel) InitRecommendationModel(t string, p int) ([]RecommendationData, int, string) {
	i.Type = t
	i.Page = 100 * (p - 1)

	i.InitModel("/recommendations.php?s=recentrecs&t="+i.Type+"&show="+strconv.Itoa(i.Page), "#content")

	i.Parser = i.GetParser(i.Url)

	if i.ResponseCode != 200 {
		return i.Data, i.ResponseCode, i.ErrorMessage
	}

	i.SetAllDetail()

	return i.Data, i.ResponseCode, i.ErrorMessage
}

// SetAllDetail to fill all anime/manga recommendation list.
func (i *RecommendationModel) SetAllDetail() {
	var recomList []RecommendationData

	area := i.Parser.Find("#content")
	area.Find("div[class=\"spaceit borderClass\"]").Each(func(j int, eachRecom *goquery.Selection) {
		recomList = append(recomList, RecommendationData{
			Username:       i.GetUsername(eachRecom),
			Date:           i.GetDate(eachRecom),
			Source:         i.GetSource(eachRecom),
			Recommendation: i.GetRecom(eachRecom),
		})
	})

	i.Data = recomList
}

// GetUsername to get recommendation username.
func (i *RecommendationModel) GetUsername(eachRecom *goquery.Selection) string {
	eachRecom.Find("table").First().Next().Next().Find("a").First().Remove()
	return eachRecom.Find("table").First().Next().Next().Find("a").Text()
}

// GetDate to get recommendation date.
func (i *RecommendationModel) GetDate(eachRecom *goquery.Selection) string {
	date := eachRecom.Find("table").First().Next().Next().Text()
	splitDate := strings.Split(date, "-")
	return strings.TrimSpace(splitDate[len(splitDate)-1])
}

// GetSource to get recommendation source anime/manga.
func (i *RecommendationModel) GetSource(eachRecom *goquery.Selection) Source {
	area := eachRecom.Find("table tr")
	return Source{
		Liked:          i.GetSourceDetail(area),
		Recommendation: i.GetSourceDetail(area),
	}
}

// GetSourceDetail to get recommendation source detail.
func (i *RecommendationModel) GetSourceDetail(area *goquery.Selection) IdTitleTypeImage {
	area = area.Find("td").First()

	liked := IdTitleTypeImage{
		Id:    i.GetSourceId(area),
		Title: i.GetSourceTitle(area),
		Type:  i.Type,
		Image: i.GetSourceImage(area),
	}

	area.Remove()
	return liked
}

// GetSourceId to get source id.
func (i *RecommendationModel) GetSourceId(area *goquery.Selection) int {
	id, _ := area.Find("a").First().Attr("href")
	splitId := strings.Split(id, "/")
	idInt, _ := strconv.Atoi(splitId[4])
	return idInt
}

// GetSourceTitle to get source title.
func (i *RecommendationModel) GetSourceTitle(area *goquery.Selection) string {
	return area.Find("strong").First().Text()
}

// GetSourceImage to get source image.
func (i *RecommendationModel) GetSourceImage(area *goquery.Selection) string {
	image, _ := area.Find("img").First().Attr("data-src")
	return helper.ImageUrlCleaner(image)
}

// GetRecom to get reommendation content.
func (i *RecommendationModel) GetRecom(eachRecom *goquery.Selection) string {
	return eachRecom.Find(".recommendations-user-recs-text").Text()
}
