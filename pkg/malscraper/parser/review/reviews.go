package review

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/pkg/malscraper/constant"
	"github.com/rl404/go-malscraper/pkg/malscraper/model/common"
	model "github.com/rl404/go-malscraper/pkg/malscraper/model/review"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser"
	"github.com/rl404/go-malscraper/pkg/malscraper/utils"
)

// ReviewsParser is parser for MyAnimeList anime/manga review list.
// Example: https://myanimelist.net/reviews.php
type ReviewsParser struct {
	parser.BaseParser
	Type string
	Page int
	Data []model.Review
}

// InitReviewsParser to initiate all fields and data of ReviewsParser.
func InitReviewsParser(params ...interface{}) (reviews ReviewsParser, err error) {
	reviews.Type = ""
	reviews.Page = 1

	for i, param := range params {
		switch i {
		case 0:
			if v, ok := param.(string); ok {
				reviews.Type = v
			}
		case 1:
			if v, ok := param.(int); ok {
				reviews.Page = v
			}
		}
	}

	if !utils.InArray(constant.ReviewTypes, reviews.Type) {
		reviews.ResponseCode = 400
		return reviews, common.ErrInvalidMainType
	}

	if reviews.Type != "bestvoted" {
		err = reviews.InitParser("/reviews.php?t="+reviews.Type+"&p="+strconv.Itoa(reviews.Page), "#content")
	} else {
		err = reviews.InitParser("/reviews.php?st="+reviews.Type+"&p="+strconv.Itoa(reviews.Page), "#content")
	}

	if err != nil {
		return reviews, err
	}

	reviews.setAllDetail()
	return reviews, nil
}

// setAllDetail to fill all anime/manga review detail information.
func (rp *ReviewsParser) setAllDetail() {
	var reviews []model.Review

	rp.Parser.Find(".borderDark").Each(func(i int, area *goquery.Selection) {

		topArea := area.Find(".spaceit").First()
		bottomArea := topArea.Next()
		veryBottomArea := bottomArea.Next()

		reviews = append(reviews, model.Review{
			ID:       rp.getID(veryBottomArea),
			Source:   rp.getSource(topArea, bottomArea),
			Username: rp.getUsername(topArea),
			Image:    rp.getImage(topArea),
			Helpful:  rp.getHelpful(topArea),
			Date:     rp.getDate(topArea),
			Episode:  rp.getProgress(topArea, "anime"),
			Chapter:  rp.getProgress(topArea, "manga"),
			Score:    rp.getScore(bottomArea),
			Review:   rp.getReview(bottomArea),
		})
	})

	rp.Data = reviews
}

// getID to get review id.
func (rp *ReviewsParser) getID(veryBottomArea *goquery.Selection) int {
	id, _ := veryBottomArea.Find("a").First().Attr("href")
	id = utils.GetValueFromSplit(id, "?id=", 1)
	return utils.StrToNum(id)
}

// getSource to get review source.
func (rp *ReviewsParser) getSource(topArea *goquery.Selection, bottomArea *goquery.Selection) model.Source {
	sourceArea := topArea.Find(".mb8:nth-of-type(2)")
	return model.Source{
		ID:    rp.getSourceID(sourceArea),
		Type:  rp.getSourceType(sourceArea),
		Title: rp.getSourceTitle(sourceArea),
		Image: rp.getSourceImage(bottomArea),
	}
}

// getSourceID to get review source id.
func (rp *ReviewsParser) getSourceID(sourceArea *goquery.Selection) int {
	id, _ := sourceArea.Find("strong a").First().Attr("href")
	id = utils.GetValueFromSplit(id, "/", 4)
	return utils.StrToNum(id)
}

// getSourceType to get review source type.
func (rp *ReviewsParser) getSourceType(sourceArea *goquery.Selection) string {
	typ := sourceArea.Find("small").First().Text()
	typ = strings.Replace(typ, "(", "", -1)
	typ = strings.Replace(typ, ")", "", -1)
	rp.Type = strings.ToLower(typ)
	return rp.Type
}

// getSourceTitle to get review source title.
func (rp *ReviewsParser) getSourceTitle(sourceArea *goquery.Selection) string {
	title := sourceArea.Find("strong").First().Text()
	return strings.TrimSpace(title)
}

// getSourceImage to get review source image.
func (rp *ReviewsParser) getSourceImage(bottomArea *goquery.Selection) string {
	image, _ := bottomArea.Find(".picSurround img").First().Attr("data-src")
	return utils.ImageURLCleaner(image)
}

// getUsername to get review username.
func (rp *ReviewsParser) getUsername(topArea *goquery.Selection) string {
	user := topArea.Find("table").First().Find("td:nth-of-type(2)").Find("a").First().Text()
	return user
}

// getImage to get review user image.
func (rp *ReviewsParser) getImage(topArea *goquery.Selection) string {
	image, _ := topArea.Find("table td img").First().Attr("src")
	return utils.ImageURLCleaner(image)
}

// getHelpful to get review helpful number.
func (rp *ReviewsParser) getHelpful(topArea *goquery.Selection) int {
	helpful := topArea.Find("table td:nth-of-type(2) strong").Text()
	return utils.StrToNum(helpful)
}

// getDate to get review date and time.
func (rp *ReviewsParser) getDate(topArea *goquery.Selection) common.DateTime {
	area := topArea.Find("div").First().Find("div").First()
	date := area.Text()
	time, _ := area.Attr("title")
	return common.DateTime{
		Date: date,
		Time: time,
	}
}

// getProgress to get review episode/chapter.
func (rp *ReviewsParser) getProgress(topArea *goquery.Selection, t string) string {
	if rp.Type != t {
		return ""
	}

	area := topArea.Find("div").First().Find("div:nth-of-type(2)").Text()
	value := strings.Replace(area, "episodes seen", "", -1)
	value = strings.Replace(value, "chapters read", "", -1)
	return strings.TrimSpace(value)
}

// getScore to get review score.
func (rp *ReviewsParser) getScore(bottomArea *goquery.Selection) map[string]int {
	score := make(map[string]int)
	area := bottomArea.Find("table").First()
	area.Find("tr").Each(func(i int, eachScore *goquery.Selection) {
		scoreType := strings.ToLower(eachScore.Find("td:nth-of-type(1)").Text())
		score[scoreType] = utils.StrToNum(eachScore.Find("td:nth-of-type(2)").Text())
	})
	return score
}

// getReview to get review content.
func (rp *ReviewsParser) getReview(bottomArea *goquery.Selection) string {
	bottomArea.Find("div").Remove()
	bottomArea.Find("a").Remove()

	r := regexp.MustCompile(`[^\S\r\n]+`)
	reviewContent := r.ReplaceAllString(bottomArea.Text(), " ")
	r = regexp.MustCompile(`(\n\n \n)`)
	reviewContent = r.ReplaceAllString(reviewContent, "")

	return strings.TrimSpace(reviewContent)
}
