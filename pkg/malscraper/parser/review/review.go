package review

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/pkg/malscraper/config"
	"github.com/rl404/go-malscraper/pkg/malscraper/constant"
	"github.com/rl404/go-malscraper/pkg/malscraper/model/common"
	model "github.com/rl404/go-malscraper/pkg/malscraper/model/review"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser"
	"github.com/rl404/go-malscraper/pkg/malscraper/utils"
)

// ReviewParser is parser for MyAnimeList anime & manga review information.
// Example: https://myanimelist.net/reviews.php?id=1
type ReviewParser struct {
	parser.BaseParser
	ID   int
	Type string
	Data model.Review
}

// InitReviewParser to initiate all fields and data of ReviewParser.
func InitReviewParser(config config.Config, id int) (review ReviewParser, err error) {
	review.ID = id
	review.Config = config

	// Checking to redis if using redis in config.
	// Redis key's pattern is `review:{id}`.
	redisKey := constant.RedisGetReview + ":" + strconv.Itoa(review.ID)
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
	err = review.InitParser("/reviews.php?id="+strconv.Itoa(review.ID), "#content")
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
	rp.setID()
	rp.setSource()
	rp.setUsername()
	rp.setImage()
	rp.setHelpful()
	rp.setDate()
	rp.setProgress()
	rp.setScore()
	rp.setReview()
}

// setID to set review's id.
func (rp *ReviewParser) setID() {
	rp.Data.ID = rp.ID
}

// setSource to set review's source.
func (rp *ReviewParser) setSource() {
	sourceArea := rp.Parser.Find(".borderDark .spaceit")

	topArea := sourceArea.Find(".mb8")
	bottomArea := sourceArea.Next()

	rp.Data.Source = model.Source{
		ID:    rp.getSourceID(topArea),
		Type:  rp.getSourceType(topArea),
		Title: rp.getSourceTitle(topArea),
		Image: rp.getSourceImage(bottomArea),
	}
}

// getSourceID to get review's source id.
func (rp *ReviewParser) getSourceID(topArea *goquery.Selection) int {
	id, _ := topArea.Find("strong a").Attr("href")
	id = utils.GetValueFromSplit(id, "/", 4)
	return utils.StrToNum(id)
}

// getSourceType to get review's source type.
func (rp *ReviewParser) getSourceType(topArea *goquery.Selection) string {
	typeStr := topArea.Find("small").First().Text()
	typeStr = strings.Replace(typeStr, "(", "", -1)
	typeStr = strings.Replace(typeStr, ")", "", -1)
	rp.Type = strings.ToLower(typeStr)
	return rp.Type
}

// getSourceTitle to get review's source title.
func (rp *ReviewParser) getSourceTitle(topArea *goquery.Selection) string {
	return strings.TrimSpace(topArea.Find("strong").Text())
}

// getSourceImage to get review's source image.
func (rp *ReviewParser) getSourceImage(bottomArea *goquery.Selection) string {
	image, _ := bottomArea.Find(".picSurround img").Attr("data-src")
	return utils.URLCleaner(image, "image", rp.Config.CleanImageURL)
}

// setUsername to set review's user who write the review.
func (rp *ReviewParser) setUsername() {
	rp.Data.Username = rp.Parser.Find(".borderDark .spaceit table td:nth-of-type(2) a").First().Text()
}

// setImage to set review's user image.
func (rp *ReviewParser) setImage() {
	image, _ := rp.Parser.Find(".borderDark .spaceit table td img").Attr("src")
	rp.Data.Image = utils.URLCleaner(image, "image", rp.Config.CleanImageURL)
}

// setHelpful to set review's user number who vote helpful.
func (rp *ReviewParser) setHelpful() {
	helpful := rp.Parser.Find(".borderDark .spaceit table td:nth-of-type(2) strong").First().Text()
	rp.Data.Helpful = utils.StrToNum(helpful)
}

// setDate to set review's date and time.
func (rp *ReviewParser) setDate() {
	dateArea := rp.Parser.Find(".borderDark .spaceit div div").First()

	date := dateArea.Text()
	time, _ := dateArea.Attr("title")

	rp.Data.Date = common.DateTime{
		Date: date,
		Time: time,
	}
}

// setProgress to set review's anime/manga episode/chapter.
func (rp *ReviewParser) setProgress() {
	progress := rp.Parser.Find(".borderDark .spaceit div div:nth-of-type(2)").First().Text()

	progress = strings.Replace(progress, "episodes seen", "", -1)
	progress = strings.Replace(progress, "chapters read", "", -1)
	progress = strings.TrimSpace(progress)

	if rp.Type == "anime" {
		rp.Data.Episode = progress
	} else {
		rp.Data.Chapter = progress
	}
}

// setScore to set review's score.
func (rp *ReviewParser) setScore() {
	scoreMap := make(map[string]int)

	scoreArea := rp.Parser.Find(".borderDark .spaceit").Next().Find("table")
	scoreArea.Find("tr").Each(func(i int, score *goquery.Selection) {
		scoreType := strings.ToLower(score.Find("td").First().Text())
		scoreMap[scoreType] = utils.StrToNum(score.Find("td:nth-of-type(2)").Text())
	})

	rp.Data.Score = scoreMap
}

// setReview to set review's content.
func (rp *ReviewParser) setReview() {
	reviewArea := rp.Parser.Find(".borderDark .spaceit").First().Next()
	reviewArea.Find("div").Remove()

	rp.Data.Review = strings.TrimSpace(reviewArea.Text())
}
