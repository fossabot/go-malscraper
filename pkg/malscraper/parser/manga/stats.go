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

// StatsParser is parser for MyAnimeList manga stats information.
// Example: https://myanimelist.net/manga/1/Monster/stats
type StatsParser struct {
	parser.BaseParser
	ID   int
	Page int
	Data model.Stats
}

// InitStatsParser to initiate all fields and data of StatsParser.
func InitStatsParser(config config.Config, id int, page ...int) (stats StatsParser, err error) {
	stats.ID = id
	stats.Page = 0
	stats.Config = config

	if len(page) > 0 {
		stats.Page = 75 * (page[0] - 1)
	}

	// Checking to redis if using redis in config.
	// Redis key's pattern is `manga-stats:{id},{page}`.
	redisKey := constant.RedisGetMangaStats + ":" + strconv.Itoa(stats.ID) + "," + strconv.Itoa(stats.Page)
	if config.RedisClient != nil {
		found, err := utils.UnmarshalFromRedis(config.RedisClient, redisKey, &stats.Data)
		if err != nil {
			stats.SetResponse(500, err.Error())
			return stats, err
		}

		if found {
			stats.SetResponse(200, constant.SuccessMessage)
			return stats, nil
		}
	}

	// Get MyAnimeList HTML source page and initiate the parser.
	err = stats.InitParser("/manga/"+strconv.Itoa(stats.ID)+"/a/stats?m=all&show="+strconv.Itoa(stats.Page), ".js-scrollfix-bottom-rel")
	if err != nil {
		return stats, err
	}

	// Fill in data field.
	stats.setAllDetail()

	// Save data field to redis if using redis in config.
	if config.RedisClient != nil {
		go utils.SaveToRedis(config.RedisClient, redisKey, stats.Data, config.CacheTime)
	}

	return stats, nil
}

// setAllDetail to set all stats detail information.
func (sp *StatsParser) setAllDetail() {
	sp.setSummary()
	sp.setScore()
	sp.setUser()
}

// setSummary to set stats' summary data.
func (sp *StatsParser) setSummary() {
	summaryMap := make(map[string]int)

	area := sp.Parser.Find("h2").First().Next()

	for goquery.NodeName(area) == "div" {

		summaryType := area.Find("span").Text()
		summaryType = strings.Replace(summaryType, ":", "", -1)
		summaryType = strings.Replace(summaryType, " ", "_", -1)
		summaryType = strings.ToLower(summaryType)

		area.Find("span").Remove()
		summaryValue := area.Text()

		summaryMap[summaryType] = utils.StrToNum(summaryValue)

		area = area.Next()
	}

	sp.Data.Summary = summaryMap
}

// setScore to set stats' score.
func (sp *StatsParser) setScore() {
	var scores []model.Score

	sp.Parser.Find("h2").First().Remove()

	area := sp.Parser.Find("h2").First().Next()

	if goquery.NodeName(area) == "table" {
		area.Find("tr").Each(func(i int, eachScore *goquery.Selection) {
			scores = append(scores, model.Score{
				Type:    sp.getScoreType(eachScore),
				Vote:    sp.getScoreVote(eachScore),
				Percent: sp.getScorePercent(eachScore),
			})
		})
	}

	sp.Data.Score = scores
}

// getScoreType to get stats' score type.
func (sp *StatsParser) getScoreType(eachScore *goquery.Selection) int {
	return utils.StrToNum(eachScore.Find("td").First().Text())
}

// getScoreVote to get stats' score number of vote.
func (sp *StatsParser) getScoreVote(eachScore *goquery.Selection) int {
	vote := eachScore.Find("td:nth-of-type(2) span small").Text()
	vote = strings.Replace(vote, " votes", "", -1)
	return utils.StrToNum(vote[1 : len(vote)-1])
}

// getScorePercent to get stats' score percent.
func (sp *StatsParser) getScorePercent(eachScore *goquery.Selection) float64 {
	eachScore.Find("td:nth-of-type(2) span small").Remove()
	percent := eachScore.Find("td:nth-of-type(2) span").Text()
	percent = strings.Replace(percent, "%", "", -1)
	return utils.StrToFloat(percent)
}

// setUser to get stats' user who vote the score.
func (sp *StatsParser) setUser() {
	var users []model.UserStats

	area := sp.Parser.Find(".table-recently-updated")

	area.Find("tr").EachWithBreak(func(i int, eachUser *goquery.Selection) bool {
		if eachUser.Find("td div").Text() == "" {
			return true
		}

		usernameArea := eachUser.Find("td").First()

		users = append(users, model.UserStats{
			Image:    sp.getUserImage(usernameArea),
			Username: sp.getUsername(usernameArea),
			Score:    sp.getUserScore(eachUser),
			Status:   sp.getUserStatus(eachUser),
			Volume:   sp.getUserProgress(eachUser, "4"),
			Chapter:  sp.getUserProgress(eachUser, "5"),
			Date:     sp.getUserDate(eachUser),
		})

		return true
	})

	sp.Data.Users = users
}

// getUserImage to get user image.
func (sp *StatsParser) getUserImage(usernameArea *goquery.Selection) string {
	image, _ := usernameArea.Find("a").Attr("style")
	return utils.URLCleaner(image[21:len(image)-1], "image", sp.Config.CleanImageURL)
}

// getUsername to get user username.
func (sp *StatsParser) getUsername(usernameArea *goquery.Selection) string {
	return strings.TrimSpace(usernameArea.Text())
}

// getUserScore to get user score.
func (sp *StatsParser) getUserScore(eachUser *goquery.Selection) int {
	return utils.StrToNum(eachUser.Find("td:nth-of-type(2)").Text())
}

// getUserStatus to get user watching/reading status.
func (sp *StatsParser) getUserStatus(eachUser *goquery.Selection) string {
	return strings.ToLower(eachUser.Find("td:nth-of-type(3)").Text())
}

// getUserProgress to get user progress.
func (sp *StatsParser) getUserProgress(eachUser *goquery.Selection, cnt string) string {
	progress := eachUser.Find("td:nth-of-type(" + cnt + ")").Text()
	return strings.TrimSpace(progress)
}

// getUserDate to get user date.
func (sp *StatsParser) getUserDate(eachUser *goquery.Selection) string {
	return eachUser.Find("td:nth-of-type(6)").Text()
}
