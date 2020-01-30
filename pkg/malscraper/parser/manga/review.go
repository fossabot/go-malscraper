package manga

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/pkg/malscraper/config"
	"github.com/rl404/go-malscraper/pkg/malscraper/constant"
	"github.com/rl404/go-malscraper/pkg/malscraper/model/common"
	model "github.com/rl404/go-malscraper/pkg/malscraper/model/manga"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser"
	"github.com/rl404/go-malscraper/pkg/malscraper/utils"
)

// ReviewParser is parser for MyAnimeList manga review list.
// Example: https://myanimelist.net/manga/1/Monster/reviews
type ReviewParser struct {
	parser.BaseParser
	ID   int
	Page int
	Data []model.Review
}

// InitReviewParser to initiate all fields and data of ReviewParser.
func InitReviewParser(config config.Config, id int, page ...int) (review ReviewParser, err error) {
	review.ID = id
	review.Page = 1
	review.Config = config

	if len(page) > 0 {
		review.Page = page[0]
	}

	// Checking to redis if using redis in config.
	// Redis key's pattern is `manga-review:{id},{page{`.
	redisKey := constant.RedisGetMangaReview + ":" + strconv.Itoa(review.ID) + "," + strconv.Itoa(review.Page)
	if config.RedisClient != nil {
		found, err := utils.UnmarshalFromRedis(config.RedisClient, redisKey, &review.Data)
		if err != nil {
			review.SetResponse(500, err.Error())
			return review, err
		}

		if found {
			review.SetResponse(200, constant.SuccessMessage)
			return review, nil
		}
	}

	// Get MyAnimeList HTML source page and initiate the parser.
	err = review.InitParser("/manga/"+strconv.Itoa(review.ID)+"/a/reviews?p="+strconv.Itoa(review.Page), "#content")
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

	area := rp.Parser.Find(".js-scrollfix-bottom-rel").First()
	area.Find(".borderDark").Each(func(i int, review *goquery.Selection) {

		topArea := review.Find(".spaceit").First()
		bottomArea := topArea.Next()
		veryBottomArea := bottomArea.Next()

		reviews = append(reviews, model.Review{
			ID:       rp.getID(veryBottomArea),
			Username: rp.getUsername(topArea),
			Image:    rp.getImage(topArea),
			Helpful:  rp.getHelpful(topArea),
			Date:     rp.getDate(topArea),
			Chapter:  rp.getProgress(topArea),
			Score:    rp.getScore(bottomArea),
			Review:   rp.getReview(bottomArea),
		})
	})

	rp.Data = reviews
}

// getID to get review's id.
func (rp *ReviewParser) getID(veryBottomArea *goquery.Selection) int {
	id, _ := veryBottomArea.Find("a").First().Attr("href")
	id = utils.GetValueFromSplit(id, "?id=", 1)
	return utils.StrToNum(id)
}

// getUsername to get username who wrote the review.
func (rp *ReviewParser) getUsername(topArea *goquery.Selection) string {
	return topArea.Find("table td:nth-of-type(2)").Find("a").First().Text()
}

// getImage to get user image.
func (rp *ReviewParser) getImage(topArea *goquery.Selection) string {
	image, _ := topArea.Find("table td:nth-of-type(1)").Find("img").Attr("src")
	return utils.URLCleaner(image, "image", rp.Config.CleanImageURL)
}

// getHelpful to get number of helpful.
func (rp *ReviewParser) getHelpful(topArea *goquery.Selection) int {
	helpful := topArea.Find("table td:nth-of-type(2) strong").First().Text()
	return utils.StrToNum(helpful)
}

// getDate to get review date.
func (rp *ReviewParser) getDate(topArea *goquery.Selection) common.DateTime {
	dateArea := topArea.Find("div").Find("div").First()
	date := dateArea.Text()
	time, _ := dateArea.Attr("title")
	return common.DateTime{
		Date: date,
		Time: time,
	}
}

// getProgress to get review chapter.
func (rp *ReviewParser) getProgress(topArea *goquery.Selection) string {
	progress := topArea.Find("div").First().Find("div:nth-of-type(2)").Text()
	progress = strings.Replace(progress, "episodes seen", "", -1)
	progress = strings.Replace(progress, "chapters read", "", -1)
	return strings.TrimSpace(progress)
}

// getScore to get review score.
func (rp *ReviewParser) getScore(bottomArea *goquery.Selection) map[string]int {
	scoreMap := make(map[string]int)

	area := bottomArea.Find("table")
	area.Find("tr").Each(func(i int, eachScore *goquery.Selection) {
		scoreType := strings.ToLower(eachScore.Find("td:nth-of-type(1)").Text())
		scoreMap[scoreType] = utils.StrToNum(eachScore.Find("td:nth-of-type(2)").Text())
	})

	return scoreMap
}

// getReview to get review content.
func (rp *ReviewParser) getReview(bottomArea *goquery.Selection) string {
	bottomArea.Find("a").Remove()
	bottomArea.Find("div[id^=score]").Remove()

	r := regexp.MustCompile(`[^\S\r\n]+`)
	reviewContent := r.ReplaceAllString(bottomArea.Text(), " ")

	r = regexp.MustCompile(`(\n\n \n)`)
	reviewContent = r.ReplaceAllString(reviewContent, "")

	return strings.TrimSpace(reviewContent)
}
