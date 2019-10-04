package additional

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/helper"
	"github.com/rl404/go-malscraper/model"
)

// StatModel is an extended model from MainModel for anime/manga statistics.
type StatModel struct {
	model.MainModel
	Id   int
	Type string
	Data StatData
}

// InitStatModel to initiate fields in parent (MainModel) model.
func (i *StatModel) InitStatModel(t string, id int) (StatData, int, string) {
	i.Id = id
	i.Type = t
	i.InitModel("/"+i.Type+"/"+strconv.Itoa(i.Id)+"/a/stats", ".js-scrollfix-bottom-rel")

	i.Parser = i.GetParser(i.Url)

	if i.ResponseCode != 200 {
		return i.Data, i.ResponseCode, i.ErrorMessage
	}

	i.SetAllDetail()

	return i.Data, i.ResponseCode, i.ErrorMessage
}

// SetAllDetail to fill all stat data.
func (i *StatModel) SetAllDetail() {
	i.SetSummary()
	i.SetScore()
	i.SetUser()
}

// SetSummary to set stat summary data.
func (i *StatModel) SetSummary() {
	summaryMap := make(map[string]string)

	summaryArea := i.Parser.Find(".js-scrollfix-bottom-rel h2").First().Next()

	for goquery.NodeName(summaryArea) == "div" {

		summaryType := summaryArea.Find("span").Text()
		summaryType = strings.Replace(summaryType, ":", "", -1)
		summaryType = strings.ToLower(summaryType)

		summaryArea.Find("span").Remove()
		summaryValue := summaryArea.Text()
		summaryValue = strings.Replace(summaryValue, ",", "", -1)

		summaryMap[summaryType] = strings.TrimSpace(summaryValue)

		summaryArea = summaryArea.Next()
	}

	i.Data.Summary = summaryMap
}

// SetScore to set stat score.
func (i *StatModel) SetScore() {
	var scoreList []Score

	i.Parser.Find(".js-scrollfix-bottom-rel h2").First().Remove()

	scoreArea := i.Parser.Find(".js-scrollfix-bottom-rel h2").First().Next()

	if goquery.NodeName(scoreArea) == "table" {
		scoreArea.Find("tr").Each(func(j int, eachScore *goquery.Selection) {
			scoreList = append(scoreList, Score{
				Type:    i.GetScoreType(eachScore),
				Vote:    i.GetScoreVote(eachScore),
				Percent: i.GetScorePercent(eachScore),
			})
		})
	}

	i.Data.Score = scoreList
}

// GetScoreType to get score type.
func (i *StatModel) GetScoreType(eachScore *goquery.Selection) string {
	return eachScore.Find("td").First().Text()
}

// GetScoreVote to get score number of vote.
func (i *StatModel) GetScoreVote(eachScore *goquery.Selection) string {
	vote := eachScore.Find("td:nth-of-type(2) span small").Text()
	vote = strings.Replace(vote, " votes", "", -1)

	return vote[1 : len(vote)-1]
}

// GetScorePercent to get score percent.
func (i *StatModel) GetScorePercent(eachScore *goquery.Selection) string {
	eachScore.Find("td:nth-of-type(2) span small").Remove()
	percent := eachScore.Find("td:nth-of-type(2) span").Text()
	percent = strings.Replace(percent, "%", "", -1)

	return strings.TrimSpace(percent)
}

// SetUser to get user who vote the score.
func (i *StatModel) SetUser() {
	var userList []User

	userArea := i.Parser.Find(".table-recently-updated")

	userArea.Find("tr").EachWithBreak(func(j int, eachUser *goquery.Selection) bool {
		if eachUser.Find("td div").Text() == "" {
			return true
		}

		usernameArea := eachUser.Find("td").First()

		userList = append(userList, User{
			Image:    i.GetUserImage(usernameArea),
			Username: i.GetUsername(usernameArea),
			Score:    i.GetUserScore(eachUser),
			Status:   i.GetUserStatus(eachUser),
			Episode:  i.GetUserProgress(eachUser, "anime", "4"),
			Volume:   i.GetUserProgress(eachUser, "manga", "4"),
			Chapter:  i.GetUserProgress(eachUser, "manga", "5"),
			Date:     i.GetUserDate(eachUser),
		})

		return true
	})

	i.Data.User = userList
}

// GetUserImage to get user image.
func (i *StatModel) GetUserImage(usernameArea *goquery.Selection) string {
	image, _ := usernameArea.Find("a").Attr("style")
	return helper.ImageUrlCleaner(image[21 : len(image)-1])
}

// GetUesrname to get user username.
func (i *StatModel) GetUsername(usernameArea *goquery.Selection) string {
	return strings.TrimSpace(usernameArea.Text())
}

// GetUserScore to get user score.
func (i *StatModel) GetUserScore(eachUser *goquery.Selection) string {
	return eachUser.Find("td:nth-of-type(2)").Text()
}

// GetUserStatus to get user watching/reading status.
func (i *StatModel) GetUserStatus(eachUser *goquery.Selection) string {
	return strings.ToLower(eachUser.Find("td:nth-of-type(3)").Text())
}

// GetUserProgress to get user progress.
func (i *StatModel) GetUserProgress(eachUser *goquery.Selection, t string, cnt string) string {
	if i.Type != t {
		return ""
	}

	progress := eachUser.Find("td:nth-of-type(" + cnt + ")").Text()
	return strings.TrimSpace(progress)
}

// GetUserDate to get user date.
func (i *StatModel) GetUserDate(eachUser *goquery.Selection) string {
	if i.Type == "anime" {
		return eachUser.Find("td:nth-of-type(5)").Text()
	}

	return eachUser.Find("td:nth-of-type(6)").Text()
}
