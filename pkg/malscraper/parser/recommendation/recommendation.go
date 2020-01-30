package recommendation

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/pkg/malscraper/config"
	"github.com/rl404/go-malscraper/pkg/malscraper/constant"
	"github.com/rl404/go-malscraper/pkg/malscraper/model/common"
	model "github.com/rl404/go-malscraper/pkg/malscraper/model/recommendation"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser"
	"github.com/rl404/go-malscraper/pkg/malscraper/utils"
)

// RecommendationParser is parser for MyAnimeList anime & manga recommendation information.
// Example: https://myanimelist.net/recommendations/anime/1-205
//          https://myanimelist.net/recommendations/manga/1-21
type RecommendationParser struct {
	parser.BaseParser
	Type string
	ID1  int
	ID2  int
	Data model.Recommendation
}

// InitRecommendationParser to initiate all fields and data of RecommendationParser.
func InitRecommendationParser(config config.Config, recommendationType string, id1 int, id2 int) (recommendation RecommendationParser, err error) {
	recommendation.Type = recommendationType
	recommendation.ID1 = id1
	recommendation.ID2 = id2
	recommendation.Config = config

	if !utils.InArray(constant.MainType, recommendation.Type) {
		recommendation.ResponseCode = 400
		return recommendation, common.ErrInvalidMainType
	}

	// Checking to redis if using redis in config.
	// Redis key's pattern is `recommendation:{type},{id1},{id2}`.
	redisKey := constant.RedisGetRecommendation + ":" + recommendation.Type + "," + strconv.Itoa(recommendation.ID1) + "," + strconv.Itoa(recommendation.ID2)
	if config.RedisClient != nil {
		found, err := utils.UnmarshalFromRedis(config.RedisClient, redisKey, &recommendation.Data)
		if err != nil {
			recommendation.SetResponse(500, err.Error())
			return recommendation, err
		}

		if found {
			recommendation.SetResponse(200, constant.SuccessMessage)
			return recommendation, nil
		}
	}

	// Get MyAnimeList HTML source page and initiate the parser.
	err = recommendation.InitParser("/recommendations/"+recommendation.Type+"/"+strconv.Itoa(recommendation.ID1)+"-"+strconv.Itoa(recommendation.ID2), "#content")
	if err != nil {
		return recommendation, err
	}

	// Fill in data field.
	recommendation.setAllDetail()

	// Save data field to redis if using redis in config.
	if config.RedisClient != nil {
		go utils.SaveToRedis(config.RedisClient, redisKey, recommendation.Data, config.CacheTime)
	}

	return recommendation, nil
}

// setAllDetail to set all recommendation detail information.
func (rp *RecommendationParser) setAllDetail() {
	rp.setSource()
	rp.setContent()
}

// setSource to set anime & manga's recommendation source.
func (rp *RecommendationParser) setSource() {
	var source model.Source

	sourceArea := rp.Parser.Find(".borderDark table tr")

	likedArea := sourceArea.Find("td").First()

	source.Liked = model.SourceDetail{
		ID:    rp.getSourceID(likedArea),
		Title: rp.getSourceTitle(likedArea),
		Type:  rp.Type,
		Image: rp.getSourceImage(likedArea),
	}

	likedArea.Remove()
	recommendationArea := sourceArea.Find("td").First()

	source.Recommended = model.SourceDetail{
		ID:    rp.getSourceID(recommendationArea),
		Title: rp.getSourceTitle(recommendationArea),
		Type:  rp.Type,
		Image: rp.getSourceImage(recommendationArea),
	}

	rp.Data.Source = source
}

// getSourceID to get anime & manga's recommendation source id.
func (rp *RecommendationParser) getSourceID(sourceArea *goquery.Selection) int {
	id, _ := sourceArea.Find("a").Attr("href")
	id = utils.GetValueFromSplit(id, "/", 4)
	return utils.StrToNum(id)
}

// getSourceTitle to get anime & manga's recommendation source title.
func (rp *RecommendationParser) getSourceTitle(sourceArea *goquery.Selection) string {
	return sourceArea.Find("strong").Text()
}

// getSourceImage to get anime & manga's recommendation source image.
func (rp *RecommendationParser) getSourceImage(sourceArea *goquery.Selection) string {
	image, _ := sourceArea.Find("img").Attr("src")
	return utils.URLCleaner(image, "image", rp.Config.CleanImageURL)
}

// setContent to set anime & manga's recommendation content.
func (rp *RecommendationParser) setContent() {
	var contents []model.User

	sourceArea := rp.Parser.Find(".borderDark")
	sourceArea.Find(".borderClass").Each(func(i int, eachRecom *goquery.Selection) {
		contents = append(contents, model.User{
			Username: rp.getRecomUser(eachRecom),
			Content:  rp.getRecomText(eachRecom),
		})
	})

	rp.Data.Users = contents
}

// getRecomUser to get user who recommend the anime/manga.
func (rp *RecommendationParser) getRecomUser(eachRecom *goquery.Selection) string {
	return eachRecom.Find("a[href*=\"/profile/\"]").Text()
}

// getRecomText to get anime & manga's recommendation content.
func (rp *RecommendationParser) getRecomText(eachRecom *goquery.Selection) string {
	eachRecom.Find("a").Remove()
	content := strings.TrimSpace(eachRecom.Find("div").First().Text())

	r := regexp.MustCompile(`(\n\n\s*)`)
	content = r.ReplaceAllString(content, " ")

	return content
}
