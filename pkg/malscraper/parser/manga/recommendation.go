package manga

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/pkg/malscraper/config"
	"github.com/rl404/go-malscraper/pkg/malscraper/constant"
	model "github.com/rl404/go-malscraper/pkg/malscraper/model/manga"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser"
	"github.com/rl404/go-malscraper/pkg/malscraper/utils"
)

// RecommendationParser is parser for MyAnimeList manga recommendation information.
// Example: https://mymangalist.net/manga/1/Monster/userrecs
type RecommendationParser struct {
	parser.BaseParser
	ID   int
	Data []model.Recommendation
}

// InitRecommendationParser to initiate all fields and data of RecommendationParser.
func InitRecommendationParser(config config.Config, id int) (recommendation RecommendationParser, err error) {
	recommendation.ID = id
	recommendation.Config = config

	// Checking to redis if using redis in config.
	// Redis key's pattern is `manga-recommendation:{id}`.
	redisKey := constant.RedisGetMangaRecommendation + ":" + strconv.Itoa(recommendation.ID)
	if config.RedisClient != nil {
		found, err := utils.UnmarshalFromRedis(config.RedisClient, redisKey, &recommendation.Data)
		if err != nil {
			recommendation.SetResponse(constant.InternalErrorCode, err.Error())
			return recommendation, err
		}

		if found {
			recommendation.SetResponse(constant.SuccessCode, constant.SuccessMessage)
			return recommendation, nil
		}
	}

	// Get MyAnimeList HTML source page and initiate the parser.
	err = recommendation.InitParser("/manga/"+strconv.Itoa(recommendation.ID)+"/a/userrecs", "#content")
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
	var recommendations []model.Recommendation

	area := rp.Parser.Find(".js-scrollfix-bottom-rel").First()
	area.Find("div.borderClass table").Each(func(i int, eachRecommendation *goquery.Selection) {
		contentArea := eachRecommendation.Find("td:nth-of-type(2)")

		recommendations = append(recommendations, model.Recommendation{
			ID:    rp.getID(contentArea),
			Title: rp.getTitle(contentArea),
			Image: rp.getImage(eachRecommendation),
			Users: rp.getUsers(eachRecommendation),
		})
	})

	rp.Data = recommendations
}

// getID to get recommended manga id.
func (rp *RecommendationParser) getID(contentArea *goquery.Selection) int {
	id, _ := contentArea.Find("a").First().Attr("href")
	id = utils.GetValueFromSplit(id, "/", 4)
	return utils.StrToNum(id)
}

// getTitle to get recommended manga title.
func (rp *RecommendationParser) getTitle(contentArea *goquery.Selection) string {
	return contentArea.Find("strong").First().Text()
}

// getImage to get manga image.
func (rp *RecommendationParser) getImage(eachRecommendation *goquery.Selection) string {
	image, _ := eachRecommendation.Find("img").Attr("data-src")
	return utils.URLCleaner(image, "image", rp.Config.CleanImageURL)
}

// getUsers to get recommendation content.
func (rp *RecommendationParser) getUsers(eachRecommendation *goquery.Selection) []model.UserRec {
	var contents []model.UserRec

	contentArea := eachRecommendation.Find("td:nth-of-type(2)")
	contents = append(contents, model.UserRec{
		Username: rp.getUsername(contentArea),
		Content:  rp.getContent(contentArea),
	})

	otherArea := eachRecommendation.Find("div[id^=simaid]")
	otherArea.Find(".borderClass").Each(func(i int, eachOther *goquery.Selection) {
		contents = append(contents, model.UserRec{
			Username: rp.getUsername(eachOther),
			Content:  rp.getContent(eachOther),
		})
	})

	return contents
}

// getUsername to get username who wrote the recommendation.
func (rp *RecommendationParser) getUsername(contentArea *goquery.Selection) string {
	userArea := contentArea.Find(".borderClass .spaceit_pad:nth-of-type(2)").First()
	userArea.Find("a").First().Remove()
	return userArea.Find("a").Text()
}

// getContent to get recommendation content.
func (rp *RecommendationParser) getContent(contentArea *goquery.Selection) string {
	content := contentArea.Find(".borderClass .spaceit_pad").First()
	content.Find("a").Remove()
	return strings.TrimSpace(content.Text())
}
