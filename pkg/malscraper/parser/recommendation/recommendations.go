package recommendation

import (
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

// RecommendationsParser is parser for MyAnimeList anime & manga recommendation information.
// Example: https://myanimelist.net/recommendations.php?s=recentrecs&t=anime
type RecommendationsParser struct {
	parser.BaseParser
	Type string
	Page int
	Data []model.Recommend
}

// InitRecommendationsParser to initiate all fields and data of RecommendationsParser.
func InitRecommendationsParser(config config.Config, rType string, page ...int) (recommend RecommendationsParser, err error) {
	recommend.Type = rType
	recommend.Page = 0
	recommend.Config = config

	if len(page) > 0 {
		recommend.Page = 100 * (page[0] - 1)
	}

	if !utils.InArray(constant.MainType, recommend.Type) {
		recommend.ResponseCode = constant.BadRequestCode
		return recommend, common.ErrInvalidMainType
	}

	// Checking to redis if using redis in config.
	// Redis key's pattern is `recommendations:{type},{id}`.
	redisKey := constant.RedisGetRecommendations + ":" + recommend.Type + "," + strconv.Itoa(recommend.Page)
	if config.RedisClient != nil {
		found, err := utils.UnmarshalFromRedis(config.RedisClient, redisKey, &recommend.Data)
		if err != nil {
			recommend.SetResponse(constant.InternalErrorCode, err.Error())
			return recommend, err
		}

		if found {
			recommend.SetResponse(constant.SuccessCode, constant.SuccessMessage)
			return recommend, nil
		}
	}

	// Get MyAnimeList HTML source page and initiate the parser.
	err = recommend.InitParser("/recommendations.php?s=recentrecs&t="+recommend.Type+"&show="+strconv.Itoa(recommend.Page), "#content")
	if err != nil {
		return recommend, err
	}

	// Fill in data field.
	recommend.setAllDetail()

	// Save data field to redis if using redis in config.
	if config.RedisClient != nil {
		go utils.SaveToRedis(config.RedisClient, redisKey, recommend.Data, config.CacheTime)
	}

	return recommend, nil
}

// setAllDetail to set all recommendation detail information
func (rp *RecommendationsParser) setAllDetail() {
	var recommends []model.Recommend

	rp.Parser.Find("div[class=\"spaceit borderClass\"]").Each(func(i int, eachRecom *goquery.Selection) {
		recommends = append(recommends, model.Recommend{
			Username: rp.getUsername(eachRecom),
			Date:     rp.getDate(eachRecom),
			Source:   rp.getSource(eachRecom),
			Content:  rp.getContent(eachRecom),
		})
	})

	rp.Data = recommends
}

// getUsername to get recommendation username.
func (rp *RecommendationsParser) getUsername(eachRecom *goquery.Selection) string {
	eachRecom.Find("table").First().Next().Next().Find("a").First().Remove()
	return eachRecom.Find("table").First().Next().Next().Find("a").Text()
}

// getDate to get recommendation date.
func (rp *RecommendationsParser) getDate(eachRecom *goquery.Selection) string {
	date := eachRecom.Find("table").First().Next().Next().Text()
	splitDate := strings.Split(date, "-")
	return strings.TrimSpace(splitDate[len(splitDate)-1])
}

// getSource to get recommendation source anime/manga.
func (rp *RecommendationsParser) getSource(eachRecom *goquery.Selection) model.Source {
	area := eachRecom.Find("table tr")
	return model.Source{
		Liked:       rp.getSourceDetail(area),
		Recommended: rp.getSourceDetail(area),
	}
}

// getSourceDetail to get recommendation source detail.
func (rp *RecommendationsParser) getSourceDetail(area *goquery.Selection) model.SourceDetail {
	area = area.Find("td").First()

	liked := model.SourceDetail{
		ID:    rp.getSourceID(area),
		Title: rp.getSourceTitle(area),
		Type:  rp.Type,
		Image: rp.getSourceImage(area),
	}

	area.Remove()
	return liked
}

// getSourceID to get source id.
func (rp *RecommendationsParser) getSourceID(area *goquery.Selection) int {
	id, _ := area.Find("a").First().Attr("href")
	id = utils.GetValueFromSplit(id, "/", 4)
	return utils.StrToNum(id)
}

// getSourceTitle to get source title.
func (rp *RecommendationsParser) getSourceTitle(area *goquery.Selection) string {
	return area.Find("strong").First().Text()
}

// getSourceImage to get source image.
func (rp *RecommendationsParser) getSourceImage(area *goquery.Selection) string {
	image, _ := area.Find("img").First().Attr("data-src")
	return utils.URLCleaner(image, "image", rp.Config.CleanImageURL)
}

// getContent to get reommendation content.
func (rp *RecommendationsParser) getContent(eachRecom *goquery.Selection) string {
	return eachRecom.Find(".recommendations-user-recs-text").Text()
}
