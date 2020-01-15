package user

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	model "github.com/rl404/go-malscraper/pkg/malscraper/model/user"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser"
	"github.com/rl404/go-malscraper/pkg/malscraper/utils"
)

// RecommendationParser is parser for MyAnimeList user's recommendation list.
// Example: https://myanimelist.net/profile/rl404/recommendations
type RecommendationParser struct {
	parser.BaseParser
	Username string
	Page     int
	Data     []model.Recommendation
}

// InitRecommendationParser to initiate all fields and data of RecommendationParser.
func InitRecommendationParser(username string, page ...int) (recommendation RecommendationParser, err error) {
	recommendation.Username = username
	recommendation.Page = 1

	if len(page) > 0 {
		recommendation.Page = page[0]
	}

	err = recommendation.InitParser("/profile/"+recommendation.Username+"/recommendations?p="+strconv.Itoa(recommendation.Page), ".container-right")
	if err != nil {
		return recommendation, err
	}

	recommendation.setAllDetail()
	return recommendation, nil
}

// setAllDetail to set all user recommendation detail information.
func (user *RecommendationParser) setAllDetail() {
	var recommendations []model.Recommendation

	user.Parser.Find("div[class=\"spaceit borderClass\"]").EachWithBreak(func(i int, eachRecom *goquery.Selection) bool {
		if eachRecom.Find("table").Text() == "" {
			return true
		}

		recommendations = append(recommendations, model.Recommendation{
			Date:    user.getDate(eachRecom),
			Source:  user.getSource(eachRecom),
			Content: user.getContent(eachRecom),
		})

		return true
	})

	user.Data = recommendations
}

// getDate to get recommendation date.
func (user *RecommendationParser) getDate(eachRecom *goquery.Selection) string {
	date := eachRecom.Find("table").First().Next().Next().Text()
	splitDate := strings.Split(date, "-")
	return strings.TrimSpace(splitDate[len(splitDate)-1])
}

// getSource to get recommendation source anime/manga.
func (user *RecommendationParser) getSource(eachRecom *goquery.Selection) model.RecommendationSource {
	area := eachRecom.Find("table tr")
	return model.RecommendationSource{
		Liked:       user.getSourceDetail(area),
		Recommended: user.getSourceDetail(area),
	}
}

// getSourceDetail to get recommendation source detail.
func (user *RecommendationParser) getSourceDetail(area *goquery.Selection) model.Source {
	area = area.Find("td").First()

	liked := model.Source{
		ID:    user.getSourceID(area),
		Title: user.getSourceTitle(area),
		Type:  user.getSourceType(area),
		Image: user.getSourceImage(area),
	}

	area.Remove()
	return liked
}

// getSourceID to get source id.
func (user *RecommendationParser) getSourceID(area *goquery.Selection) int {
	id, _ := area.Find("a").First().Attr("href")
	id = utils.GetValueFromSplit(id, "/", 4)
	return utils.StrToNum(id)
}

// getSourceTitle to get source title.
func (user *RecommendationParser) getSourceTitle(area *goquery.Selection) string {
	return area.Find("strong").First().Text()
}

// getSourceType to get source type.
func (user *RecommendationParser) getSourceType(area *goquery.Selection) string {
	t, _ := area.Find("a").First().Attr("href")
	return utils.GetValueFromSplit(t, "/", 3)
}

// getSourceImage to get source image.
func (user *RecommendationParser) getSourceImage(area *goquery.Selection) string {
	image, _ := area.Find("img").First().Attr("data-src")
	return utils.ImageURLCleaner(image)
}

// getContent to get reommendation content.
func (user *RecommendationParser) getContent(eachRecom *goquery.Selection) string {
	return eachRecom.Find(".profile-user-recs-text").Text()
}
