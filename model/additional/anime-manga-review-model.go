package additional

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/helper"
	"github.com/rl404/go-malscraper/model"
)

// ReviewModel is an extended model from MainModel for anime/manga additional review list.
type ReviewModel struct {
	model.MainModel
	Id   int
	Type string
	Page int
	Data []ReviewData
}

// InitReviewModel to initiate fields in parent (MainModel) model.
func (i *ReviewModel) InitReviewModel(t string, id int, p int) ([]ReviewData, int, string) {
	i.Id = id
	i.Type = t
	i.Page = p
	i.InitModel("/"+i.Type+"/"+strconv.Itoa(i.Id)+"/a/reviews?p="+strconv.Itoa(i.Page), ".js-scrollfix-bottom-rel")

	i.Parser = i.GetParser(i.Url)

	if i.ResponseCode != 200 {
		return i.Data, i.ResponseCode, i.ErrorMessage
	}

	i.SetAllDetail()

	return i.Data, i.ResponseCode, i.ErrorMessage
}

// SetAllDetail to set review list.
func (i *ReviewModel) SetAllDetail() {
	var reviewList []ReviewData

	area := i.Parser.Find(".js-scrollfix-bottom-rel").First()
	area.Find(".borderDark").Each(func(j int, eachReview *goquery.Selection) {

		topArea := eachReview.Find(".spaceit").First()
		bottomArea := topArea.Next()
		veryBottomArea := bottomArea.Next()

		reviewList = append(reviewList, ReviewData{
			Id:       i.GetId(veryBottomArea),
			Username: i.GetUsername(topArea),
			Image:    i.GetImage(topArea),
			Helpful:  i.GetHelpful(topArea),
			Date:     i.GetDate(topArea),
			Episode:  i.GetProgress("anime", topArea),
			Chapter:  i.GetProgress("manga", topArea),
			Score:    i.GetScore(bottomArea),
			Review:   i.GetReview(bottomArea),
		})
	})

	i.Data = reviewList
}

// GetId to get review id.
func (i *ReviewModel) GetId(veryBottomArea *goquery.Selection) string {
	id, _ := veryBottomArea.Find("a").First().Attr("href")
	splitId := strings.Split(id, "?id=")

	return splitId[1]
}

// GetUsername to get username who wrote the review.
func (i *ReviewModel) GetUsername(topArea *goquery.Selection) string {
	return topArea.Find("table td:nth-of-type(2)").Find("a").First().Text()
}

// GetImage to get user image.
func (i *ReviewModel) GetImage(topArea *goquery.Selection) string {
	image, _ := topArea.Find("table td:nth-of-type(1)").Find("img").Attr("src")
	return helper.ImageUrlCleaner(image)
}

// GetHelpful to get number of helpful.
func (i *ReviewModel) GetHelpful(topArea *goquery.Selection) string {
	helpful := topArea.Find("table td:nth-of-type(2) strong").First().Text()
	return strings.TrimSpace(helpful)
}

// GetDate to get review date.
func (i *ReviewModel) GetDate(topArea *goquery.Selection) DateTime {
	dateArea := topArea.Find("div").Find("div").First()
	date := dateArea.Text()
	time, _ := dateArea.Attr("title")

	return DateTime{
		Date: date,
		Time: time,
	}
}

// GetProgress to get review episode/chapter.
func (i *ReviewModel) GetProgress(t string, topArea *goquery.Selection) string {
	if t != i.Type {
		return ""
	}

	progress := topArea.Find("div").First().Find("div:nth-of-type(2)").Text()
	progress = strings.Replace(progress, "episodes seen", "", -1)
	progress = strings.Replace(progress, "chapters read", "", -1)

	return strings.TrimSpace(progress)
}

// GetScore to get review score.
func (i *ReviewModel) GetScore(bottomArea *goquery.Selection) map[string]string {
	scoreMap := make(map[string]string)

	area := bottomArea.Find("table")
	area.Find("tr").Each(func(j int, eachScore *goquery.Selection) {
		scoreType := strings.ToLower(eachScore.Find("td:nth-of-type(1)").Text())
		scoreValue := eachScore.Find("td:nth-of-type(2)").Text()
		scoreMap[scoreType] = scoreValue
	})

	return scoreMap
}

// GetReview to get review content.
func (i *ReviewModel) GetReview(bottomArea *goquery.Selection) string {
	bottomArea.Find("a").Remove()
	bottomArea.Find("div[id^=score]").Remove()
	return strings.TrimSpace(bottomArea.Text())
}
