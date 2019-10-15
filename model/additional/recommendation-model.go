package additional

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/helper"
	"github.com/rl404/go-malscraper/model"
)

// RecommendationModel is an extended model from MainModel for anime/manga additional recommendation list.
type RecommendationModel struct {
	model.MainModel
	Id   int
	Type string
	Data []RecommendationData
}

// InitRecommendationModel to initiate fields in parent (MainModel) model.
func (i *RecommendationModel) InitRecommendationModel(t string, id int) ([]RecommendationData, int, string) {
	i.Id = id
	i.Type = t
	i.InitModel("/"+i.Type+"/"+strconv.Itoa(i.Id)+"/a/userrecs", ".js-scrollfix-bottom-rel")

	i.Parser = i.GetParser(i.Url)

	if i.ResponseCode != 200 {
		return i.Data, i.ResponseCode, i.ErrorMessage
	}

	i.SetAllDetail()

	return i.Data, i.ResponseCode, i.ErrorMessage
}

// SetAllDetail to set review list.
func (i *RecommendationModel) SetAllDetail() {
	var recommendationList []RecommendationData

	area := i.Parser.Find(".js-scrollfix-bottom-rel").First()
	area.Find("div.borderClass table").Each(func(j int, eachRecommendation *goquery.Selection) {

		contentArea := eachRecommendation.Find("td:nth-of-type(2)")
		otherArea := eachRecommendation.Find("div[id^=simaid]")

		recommendationList = append(recommendationList, RecommendationData{
			Id:             i.GetId(contentArea),
			Title:          i.GetTitle(contentArea),
			Image:          i.GetImage(eachRecommendation),
			Username:       i.GetUsername(contentArea),
			Recommendation: i.GetRecommendation(contentArea),
			Other:          i.GetOther(otherArea),
		})
	})

	i.Data = recommendationList
}

// GetId to get recommendation id.
func (i *RecommendationModel) GetId(contentArea *goquery.Selection) string {
	id, _ := contentArea.Find("a").First().Attr("href")
	splitId := strings.Split(id, "/")

	return splitId[4]
}

// GetTitle to get anime/manga title.
func (i *RecommendationModel) GetTitle(contentArea *goquery.Selection) string {
	return contentArea.Find("strong").First().Text()
}

// GetImage to get anime/manga image.
func (i *RecommendationModel) GetImage(eachRecommendation *goquery.Selection) string {
	image, _ := eachRecommendation.Find("img").Attr("data-src")
	return helper.ImageUrlCleaner(image)
}

// GetUsername to get username who wrote the recommendation.
func (i *RecommendationModel) GetUsername(contentArea *goquery.Selection) string {
	userArea := contentArea.Find(".borderClass .spaceit_pad:nth-of-type(2)").First()
	userArea.Find("a").First().Remove()

	return userArea.Find("a").Text()
}

// GetRecommendation to get recommendation content.
func (i *RecommendationModel) GetRecommendation(contentArea *goquery.Selection) string {
	content := contentArea.Find(".borderClass .spaceit_pad").First()
	content.Find("a").Remove()
	return strings.TrimSpace(content.Text())
}

// GetOther to get other recommendation.
func (i *RecommendationModel) GetOther(otherArea *goquery.Selection) []OtherRecom {
	var otherList []OtherRecom
	otherArea.Find(".borderClass").Each(func(j int, eachOther *goquery.Selection) {
		otherList = append(otherList, OtherRecom{
			Username:       i.GetOtherUser(eachOther),
			Recommendation: i.GetOtherRecom(eachOther),
		})
	})

	return otherList
}

// GetOtherUser to get other username.
func (i *RecommendationModel) GetOtherUser(eachOther *goquery.Selection) string {
	userArea := eachOther.Find(".spaceit_pad:nth-of-type(2)")
	userArea.Find("a").First().Remove()
	return userArea.Find("a").Text()
}

// GetOtherRecom to get other recommendation content.
func (i *RecommendationModel) GetOtherRecom(eachOther *goquery.Selection) string {
	content := eachOther.Find(".spaceit_pad").First()
	content.Find("a").Remove()
	return strings.TrimSpace(content.Text())
}
