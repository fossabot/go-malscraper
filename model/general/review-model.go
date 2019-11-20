package general

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/helper"
	"github.com/rl404/go-malscraper/model"
)

// ReviewModel is an extended model from MainModel for user.
type ReviewModel struct {
	model.MainModel
	Type string
	Id   int
	Data ReviewData
}

// InitReviewModel to initiate fields in parent (MainModel) model.
func (i *ReviewModel) InitReviewModel(id int) (ReviewData, int, string) {
	i.Id = id
	i.InitModel("/reviews.php?id="+strconv.Itoa(i.Id), "#content")

	i.Parser = i.GetParser(i.Url)

	if i.ResponseCode != 200 {
		return i.Data, i.ResponseCode, i.ErrorMessage
	}

	i.SetAllDetail()

	return i.Data, i.ResponseCode, i.ErrorMessage
}

// SetAllDetail to fill all anime & manga data.
func (i *ReviewModel) SetAllDetail() {
	i.SetId()
	i.SetSource()
	i.SetUsername()
	i.SetImage()
	i.SetHelpful()
	i.SetDate()
	i.SetEpisode()
	i.SetScore()
	i.SetReview()
}

// SetId to set review id.
func (i *ReviewModel) SetId() {
	i.Data.Id = i.Id
}

// SetSource to set review source.
func (i *ReviewModel) SetSource() {
	sourceArea := i.Parser.Find(".borderDark .spaceit")

	topArea := sourceArea.Find(".mb8")
	bottomArea := sourceArea.Next()

	i.Data.Source = IdTitleTypeImage{
		Id:    i.GetSourceId(topArea),
		Type:  i.GetSourceType(topArea),
		Title: i.GetSourceTitle(topArea),
		Image: i.GetSourceImage(bottomArea),
	}
}

// GetSourceId to get source id.
func (i *ReviewModel) GetSourceId(topArea *goquery.Selection) int {
	id, _ := topArea.Find("strong a").Attr("href")
	splitId := strings.Split(id, "/")
	idInt, _ := strconv.Atoi(splitId[4])
	return idInt
}

// GetSourceType to get source type.
func (i *ReviewModel) GetSourceType(topArea *goquery.Selection) string {
	typeStr := topArea.Find("small").First().Text()
	typeStr = strings.Replace(typeStr, "(", "", -1)
	typeStr = strings.Replace(typeStr, ")", "", -1)
	i.Type = strings.ToLower(typeStr)
	return i.Type
}

// GetSourceTitle to get source title.
func (i *ReviewModel) GetSourceTitle(topArea *goquery.Selection) string {
	return strings.TrimSpace(topArea.Find("strong").Text())
}

// GetSourceImage to get source image.
func (i *ReviewModel) GetSourceImage(bottomArea *goquery.Selection) string {
	image, _ := bottomArea.Find(".picSurround img").Attr("data-src")
	return helper.ImageUrlCleaner(image)
}

// SetUsername to set user who write the review.
func (i *ReviewModel) SetUsername() {
	i.Data.Username = i.Parser.Find(".borderDark .spaceit table td:nth-of-type(2) a").First().Text()
}

// SetImage to set user image.
func (i *ReviewModel) SetImage() {
	image, _ := i.Parser.Find(".borderDark .spaceit table td img").Attr("src")
	i.Data.Image = helper.ImageUrlCleaner(image)
}

// SetHelpful to set user number who say helpful.
func (i *ReviewModel) SetHelpful() {
	helpful := i.Parser.Find(".borderDark .spaceit table td:nth-of-type(2) strong").First().Text()
	i.Data.Helpful, _ = strconv.Atoi(strings.TrimSpace(helpful))
}

// SetDate to set date and time of the review.
func (i *ReviewModel) SetDate() {
	dateArea := i.Parser.Find(".borderDark .spaceit div div").First()

	date := dateArea.Text()
	time, _ := dateArea.Attr("title")

	i.Data.Date = ReviewDate{
		Date: date,
		Time: time,
	}
}

// SetEpisode to set anime/manga episode/chapter.
func (i *ReviewModel) SetEpisode() {
	episode := i.Parser.Find(".borderDark .spaceit div div:nth-of-type(2)").First().Text()

	episode = strings.Replace(episode, "episodes seen", "", -1)
	episode = strings.Replace(episode, "chapters read", "", -1)
	episode = strings.TrimSpace(episode)

	if i.Type == "anime" {
		i.Data.Episode = episode
	} else {
		i.Data.Chapter = episode
	}
}

// SetScore to set review score.
func (i *ReviewModel) SetScore() {
	scoreMap := make(map[string]int)

	scoreArea := i.Parser.Find(".borderDark .spaceit").Next().Find("table")
	scoreArea.Find("tr").Each(func(j int, eachScore *goquery.Selection) {
		scoreType := strings.ToLower(eachScore.Find("td").First().Text())
		scoreValue, _ := strconv.Atoi(eachScore.Find("td:nth-of-type(2)").Text())
		scoreMap[scoreType] = scoreValue
	})

	i.Data.Score = scoreMap
}

// SetReview to set review content.
func (i *ReviewModel) SetReview() {
	reviewArea := i.Parser.Find(".borderDark .spaceit").First().Next()
	reviewArea.Find("div").Remove()

	i.Data.Review = strings.TrimSpace(reviewArea.Text())
}
