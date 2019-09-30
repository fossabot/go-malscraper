package general

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/helper"
	"github.com/rl404/go-malscraper/model"
)

// RecommendationModel is an extended model from MainModel for user.
type RecommendationModel struct {
	model.MainModel
	Type string
	Id1  int
	Id2  int
	Data RecommendationData
}

// InitRecommendationModel to initiate fields in parent (MainModel) model.
func (i *RecommendationModel) InitRecommendationModel(t string, id1 int, id2 int) (RecommendationData, int, string) {
	i.Type = t
	i.Id1 = id1
	i.Id2 = id2

	i.InitModel("/recommendations/"+i.Type+"/"+strconv.Itoa(i.Id1)+"-"+strconv.Itoa(i.Id2), "#content")

	i.Parser = i.GetParser(i.Url)

	if i.ResponseCode != 200 {
		return i.Data, i.ResponseCode, i.ErrorMessage
	}

	i.SetAllDetail()

	return i.Data, i.ResponseCode, i.ErrorMessage
}

// SetAllDetail to fill all recommendationd data.
func (i *RecommendationModel) SetAllDetail() {
	i.SetSource()
	i.SetRecommendation()
}

// SetSource to set recommendation source.
func (i *RecommendationModel) SetSource() {
	var source RecomSource

	sourceArea := i.Parser.Find(".borderDark table tr")

	likedArea := sourceArea.Find("td").First()

	source.Liked = IdTitleTypeImage{
		Id:    i.GetSourceId(likedArea),
		Title: i.GetSourceTitle(likedArea),
		Type:  i.Type,
		Image: i.GetSourceImage(likedArea),
	}

	likedArea.Remove()
	recommendationArea := sourceArea.Find("td").First()

	source.Recommendation = IdTitleTypeImage{
		Id:    i.GetSourceId(recommendationArea),
		Title: i.GetSourceTitle(recommendationArea),
		Type:  i.Type,
		Image: i.GetSourceImage(recommendationArea),
	}

	i.Data.Source = source
}

// GetSourceId to get source id.
func (i *RecommendationModel) GetSourceId(sourceArea *goquery.Selection) string {
	id, _ := sourceArea.Find("a").Attr("href")
	splitId := strings.Split(id, "/")

	return splitId[4]
}

// GetSourceTitle to get source title.
func (i *RecommendationModel) GetSourceTitle(sourceArea *goquery.Selection) string {
	return sourceArea.Find("strong").Text()
}

// GetSourceImage to get source image.
func (i *RecommendationModel) GetSourceImage(sourceArea *goquery.Selection) string {
	image, _ := sourceArea.Find("img").Attr("src")
	return helper.ImageUrlCleaner(image)
}

// SetRecommendation to set recommendation list.
func (i *RecommendationModel) SetRecommendation() {
	var recomList []RecomList

	sourceArea := i.Parser.Find(".borderDark")
	sourceArea.Find(".borderClass").Each(func(j int, eachRecom *goquery.Selection) {
		recomList = append(recomList, RecomList{
			Username:       i.GetRecomUser(eachRecom),
			Recommendation: i.GetRecomText(eachRecom),
		})
	})

	i.Data.Recommendation = recomList
}

// GetRecomUser to get user who recommend the anime/manga.
func (i *RecommendationModel) GetRecomUser(eachRecom *goquery.Selection) string {
	return eachRecom.Find("a[href*=\"/profile/\"]").Text()
}

// GetRecomText to get recommendation content.
func (i *RecommendationModel) GetRecomText(eachRecom *goquery.Selection) string {
	eachRecom.Find("a").Remove()
	return strings.TrimSpace(eachRecom.Find("span").Text())
}
