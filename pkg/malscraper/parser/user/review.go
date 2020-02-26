package user

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/pkg/malscraper/config"
	"github.com/rl404/go-malscraper/pkg/malscraper/constant"
	"github.com/rl404/go-malscraper/pkg/malscraper/model/common"
	model "github.com/rl404/go-malscraper/pkg/malscraper/model/user"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser"
	"github.com/rl404/go-malscraper/pkg/malscraper/utils"
)

// ReviewParser is parser for MyAnimeList user review list.
// Example: https://myanimelist.net/profile/rl404/reviews
type ReviewParser struct {
	parser.BaseParser
	Username string
	Type     string
	Page     int
	Data     []model.Review
}

// InitReviewParser to initiate all fields and data of ReviewParser.
func InitReviewParser(config config.Config, username string, page ...int) (review ReviewParser, err error) {
	review.Username = username
	review.Page = 1
	review.Config = config

	if len(page) > 0 {
		review.Page = page[0]
	}

	// Checking to redis if using redis in config.
	// Redis key's pattern is `user-review:{name},{page}`.
	redisKey := constant.RedisGetUserReview + ":" + review.Username + "," + strconv.Itoa(review.Page)
	if config.RedisClient != nil {
		found, err := utils.UnmarshalFromRedis(config.RedisClient, redisKey, &review.Data)
		if err != nil {
			review.SetResponse(constant.InternalErrorCode, err.Error())
			return review, err
		}

		if found {
			review.SetResponse(constant.SuccessCode, constant.SuccessMessage)
			return review, nil
		}
	}

	// Get MyAnimeList HTML source page and initiate the parser.
	err = review.InitParser("/profile/"+review.Username+"/reviews?p="+strconv.Itoa(review.Page), "#content")
	if err != nil {
		return review, err
	}

	// Fill in data field.
	review.setAllDetail()

	// Save data field to redis if using redis in config.
	if config.RedisClient != nil {
		go utils.SaveToRedis(config.RedisClient, redisKey, review.Data, config.CacheTime)
	}

	return review, nil
}

// setAllDetail to set all review detail information.
func (rp *ReviewParser) setAllDetail() {
	var reviews []model.Review

	rp.Parser.Find(".borderDark").Each(func(i int, area *goquery.Selection) {

		topArea := area.Find(".spaceit").First()
		bottomArea := topArea.Next()
		veryBottomArea := bottomArea.Next()

		reviews = append(reviews, model.Review{
			ID:      rp.getID(veryBottomArea),
			Source:  rp.getSource(topArea, bottomArea),
			Helpful: rp.getHelpful(topArea),
			Date:    rp.getDate(topArea),
			Episode: rp.getProgress(topArea, "anime"),
			Chapter: rp.getProgress(topArea, "manga"),
			Score:   rp.getScore(bottomArea),
			Review:  rp.getReview(bottomArea),
		})
	})

	rp.Data = reviews
}

// getID to get review id.
func (rp *ReviewParser) getID(veryBottomArea *goquery.Selection) int {
	id, _ := veryBottomArea.Find("a").First().Attr("href")
	id = utils.GetValueFromSplit(id, "?id=", 1)
	return utils.StrToNum(id)
}

// getSource to get review source.
func (rp *ReviewParser) getSource(topArea *goquery.Selection, bottomArea *goquery.Selection) model.Source {
	sourceArea := topArea.Find("div:nth-of-type(2)")
	return model.Source{
		ID:    rp.getSourceID(sourceArea),
		Type:  rp.getSourceType(sourceArea),
		Title: rp.getSourceTitle(sourceArea),
		Image: rp.getSourceImage(bottomArea),
	}
}

// getSourceID to get review source id.
func (rp *ReviewParser) getSourceID(sourceArea *goquery.Selection) int {
	id, _ := sourceArea.Find("strong a").First().Attr("href")
	id = utils.GetValueFromSplit(id, "/", 4)
	return utils.StrToNum(id)
}

// getSourceType to get review source type.
func (rp *ReviewParser) getSourceType(sourceArea *goquery.Selection) string {
	typ := sourceArea.Find("small").First().Text()
	typ = strings.Replace(typ, "(", "", -1)
	typ = strings.Replace(typ, ")", "", -1)
	rp.Type = strings.ToLower(typ)
	return rp.Type
}

// getSourceTitle to get review source title.
func (rp *ReviewParser) getSourceTitle(sourceArea *goquery.Selection) string {
	title := sourceArea.Find("strong").First().Text()
	return strings.TrimSpace(title)
}

// getSourceImage to get review source image.
func (rp *ReviewParser) getSourceImage(bottomArea *goquery.Selection) string {
	image, _ := bottomArea.Find(".picSurround img").First().Attr("data-src")
	return utils.URLCleaner(image, "image", rp.Config.CleanImageURL)
}

// getHelpful to get review helpful number.
func (rp *ReviewParser) getHelpful(topArea *goquery.Selection) int {
	helpful := topArea.Find("span[id^=\"rhelp\"]").Text()
	return utils.StrToNum(helpful)
}

// getDate to get review date and time.
func (rp *ReviewParser) getDate(topArea *goquery.Selection) common.DateTime {
	area := topArea.Find("div").First().Find("div").First()
	date := area.Text()
	time, _ := area.Attr("title")
	return common.DateTime{
		Date: date,
		Time: time,
	}
}

// getProgress to get review episode/chapter.
func (rp *ReviewParser) getProgress(topArea *goquery.Selection, t string) string {
	if rp.Type != t {
		return ""
	}

	area := topArea.Find("div").First().Find("div:nth-of-type(2)").Text()
	value := strings.Replace(area, "episodes seen", "", -1)
	value = strings.Replace(value, "chapters read", "", -1)
	return strings.TrimSpace(value)
}

// getScore to get review score.
func (rp *ReviewParser) getScore(bottomArea *goquery.Selection) map[string]int {
	score := make(map[string]int)
	area := bottomArea.Find("table").First()
	area.Find("tr").Each(func(i int, eachScore *goquery.Selection) {
		scoreType := strings.ToLower(eachScore.Find("td:nth-of-type(1)").Text())
		score[scoreType] = utils.StrToNum(eachScore.Find("td:nth-of-type(2)").Text())
	})
	return score
}

// getReview to get review content.
func (rp *ReviewParser) getReview(bottomArea *goquery.Selection) string {
	bottomArea.Find("div").Remove()
	bottomArea.Find("a").Remove()

	r := regexp.MustCompile(`[^\S\r\n]+`)
	reviewContent := r.ReplaceAllString(bottomArea.Text(), " ")
	r = regexp.MustCompile(`(\n\n \n)`)
	reviewContent = r.ReplaceAllString(reviewContent, "")

	return strings.TrimSpace(reviewContent)
}
