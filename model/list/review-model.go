package list

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/helper"
	"github.com/rl404/go-malscraper/model"
)

// ReviewModel is an extended model from MainModel for anime/manga review list.
type ReviewModel struct {
	model.MainModel
	Type string
	Page int
	Data []ReviewData
}

// InitReviewModel to initiate fields in parent (MainModel) model.
func (i *ReviewModel) InitReviewModel(t string, p int) ([]ReviewData, int, string) {
	i.Type = t
	i.Page = p

	if i.Type != "bestvoted" {
		i.InitModel("/reviews.php?t="+i.Type+"&p="+strconv.Itoa(i.Page), "#content")
	} else {
		i.InitModel("/reviews.php?st="+i.Type+"&p="+strconv.Itoa(i.Page), "#content")
	}

	i.Parser = i.GetParser(i.Url)

	if i.ResponseCode != 200 {
		return i.Data, i.ResponseCode, i.ErrorMessage
	}

	i.SetAllDetail()

	return i.Data, i.ResponseCode, i.ErrorMessage
}

// SetAllDetail to fill all anime/manga review list.
func (i *ReviewModel) SetAllDetail() {
	var reviewList []ReviewData

	area := i.Parser.Find("#content").First()
	area.Find(".borderDark").Each(func(j int, eachReview *goquery.Selection) {

		topArea := eachReview.Find(".spaceit").First()
		bottomArea := topArea.Next()
		veryBottomArea := bottomArea.Next()

		reviewList = append(reviewList, ReviewData{
			Id:       i.GetId(veryBottomArea),
			Source:   i.GetSource(topArea, bottomArea),
			Username: i.GetUsername(topArea),
			Image:    i.GetImage(topArea),
			Helpful:  i.GetHelpful(topArea),
			Date:     i.GetDate(topArea),
			Episode:  i.GetProgress(topArea, "anime"),
			Chapter:  i.GetProgress(topArea, "manga"),
			Score:    i.GetScore(bottomArea),
			Review:   i.GetReview(bottomArea),
		})
	})

	i.Data = reviewList
}

// GetId to get review id.
func (i *ReviewModel) GetId(veryBottomArea *goquery.Selection) int {
	id, _ := veryBottomArea.Find("a").First().Attr("href")
	splitId := strings.Split(id, "?id=")
	idInt, _ := strconv.Atoi(splitId[1])
	return idInt
}

// GetSource to get review source.
func (i *ReviewModel) GetSource(topArea *goquery.Selection, bottomArea *goquery.Selection) ReviewSource {
	sourceArea := topArea.Find(".mb8:nth-of-type(2)")
	return ReviewSource{
		Id:    i.GetSourceId(sourceArea),
		Type:  i.GetSourceType(sourceArea),
		Title: i.GetSourceTitle(sourceArea),
		Image: i.GetSourceImage(bottomArea),
	}
}

// GetSourceId to get review source id.
func (i *ReviewModel) GetSourceId(sourceArea *goquery.Selection) int {
	id, _ := sourceArea.Find("strong a").First().Attr("href")
	splitId := strings.Split(id, "/")
	idInt, _ := strconv.Atoi(splitId[4])
	return idInt
}

// GetSourceType to get review source type.
func (i *ReviewModel) GetSourceType(sourceArea *goquery.Selection) string {
	typ := sourceArea.Find("small").First().Text()
	typ = strings.Replace(typ, "(", "", -1)
	typ = strings.Replace(typ, ")", "", -1)
	return strings.ToLower(typ)
}

// GetSourceTitle to get review source title.
func (i *ReviewModel) GetSourceTitle(sourceArea *goquery.Selection) string {
	title := sourceArea.Find("strong").First().Text()
	return strings.TrimSpace(title)
}

// GetSourceImage to get review source image.
func (i *ReviewModel) GetSourceImage(bottomArea *goquery.Selection) string {
	image, _ := bottomArea.Find(".picSurround img").First().Attr("data-src")
	return helper.ImageUrlCleaner(image)
}

// GetUsername to get review username.
func (i *ReviewModel) GetUsername(topArea *goquery.Selection) string {
	user := topArea.Find("table").First().Find("td:nth-of-type(2)").Find("a").First().Text()
	return user
}

// GetImage to get review user image.
func (i *ReviewModel) GetImage(topArea *goquery.Selection) string {
	image, _ := topArea.Find("table td img").First().Attr("src")
	return helper.ImageUrlCleaner(image)
}

// GetHelpful to get review helpful number.
func (i *ReviewModel) GetHelpful(topArea *goquery.Selection) int {
	helpful := topArea.Find("table td:nth-of-type(2) strong").Text()
	helpfulInt, _ := strconv.Atoi(strings.TrimSpace(helpful))
	return helpfulInt
}

// GetDate to get review date and time.
func (i *ReviewModel) GetDate(topArea *goquery.Selection) DateTime {
	area := topArea.Find("div").First().Find("div").First()
	date := area.Text()
	time, _ := area.Attr("title")
	return DateTime{
		Date: date,
		Time: time,
	}
}

// GetProgress to get review episode/chapter.
func (i *ReviewModel) GetProgress(topArea *goquery.Selection, t string) string {
	if i.Type != t {
		return ""
	}

	area := topArea.Find("div").First().Find("div:nth-of-type(2)").Text()
	value := strings.Replace(area, "episodes seen", "", -1)
	value = strings.Replace(value, "chapters read", "", -1)
	return strings.TrimSpace(value)
}

// GetScore to get review score.
func (i *ReviewModel) GetScore(bottomArea *goquery.Selection) map[string]int {
	score := make(map[string]int)
	area := bottomArea.Find("table").First()
	area.Find("tr").Each(func(j int, eachScore *goquery.Selection) {
		scoreType := strings.ToLower(eachScore.Find("td:nth-of-type(1)").Text())
		scoreValue, _ := strconv.Atoi(eachScore.Find("td:nth-of-type(2)").Text())
		score[scoreType] = scoreValue
	})
	return score
}

// GetReview to get review content.
func (i *ReviewModel) GetReview(bottomArea *goquery.Selection) string {
	bottomArea.Find("div").Remove()
	bottomArea.Find("a").Remove()

	r, _ := regexp.Compile(`[^\S\r\n]+`)
	reviewContent := r.ReplaceAllString(bottomArea.Text(), " ")
	r, _ = regexp.Compile(`(\n\n \n)`)
	reviewContent = r.ReplaceAllString(reviewContent, "")

	return strings.TrimSpace(reviewContent)
}
