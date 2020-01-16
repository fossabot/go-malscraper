package user

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/pkg/malscraper/constant"
	"github.com/rl404/go-malscraper/pkg/malscraper/model/common"
	model "github.com/rl404/go-malscraper/pkg/malscraper/model/user"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser"
	"github.com/rl404/go-malscraper/pkg/malscraper/utils"
)

// UserHistoryParser is parser for MyAnimeList user history list.
// Example: https://myanimelist.net/history/rl404
type UserHistoryParser struct {
	parser.BaseParser
	Username string
	Type     string
	Data     []model.UserHistory
}

// InitUserHistoryParser to initiate all fields and data of UserHistoryParser.
func InitUserHistoryParser(username string, historyType ...string) (userHistory UserHistoryParser, err error) {
	userHistory.Username = username
	userHistory.Type = ""

	if len(historyType) > 0 {
		userHistory.Type = historyType[0]

		if userHistory.Type != "" && !utils.InArray(constant.MainType, userHistory.Type) {
			userHistory.ResponseCode = 400
			return userHistory, common.ErrInvalidMainType
		}
	}

	if userHistory.Type == "" {
		err = userHistory.InitParser("/history/"+userHistory.Username, "#content")
	} else {
		err = userHistory.InitParser("/history/"+userHistory.Username+"/"+userHistory.Type, "#content")
	}

	if err != nil {
		return userHistory, err
	}

	userHistory.setAllDetail()
	return userHistory, nil
}

// setAllDetail to set all user history detail information.
func (user *UserHistoryParser) setAllDetail() {
	var historyList []model.UserHistory
	area := user.Parser.Find("table")
	area.Find("tr").EachWithBreak(func(i int, history *goquery.Selection) bool {

		historyClass, _ := history.Find("td").First().Attr("class")
		if historyClass != "borderClass" {
			return true
		}

		nameArea := history.Find("td").First()

		historyList = append(historyList, model.UserHistory{
			ID:       user.getId(nameArea),
			Title:    user.getTitle(nameArea),
			Type:     user.getType(nameArea),
			Progress: user.getProgress(nameArea),
			Date:     user.getDate(history),
		})

		return true
	})

	user.Data = historyList
}

// getId to get user's anime/manga history id.
func (user *UserHistoryParser) getId(nameArea *goquery.Selection) int {
	idLink, _ := nameArea.Find("a").First().Attr("href")
	id := utils.GetValueFromSplit(idLink, "=", 1)
	return utils.StrToNum(id)
}

// getTitle to get user's anime/manga history title.
func (user *UserHistoryParser) getTitle(nameArea *goquery.Selection) string {
	return nameArea.Find("a").First().Text()
}

// getType to get user's anime/manga history type.
func (user *UserHistoryParser) getType(nameArea *goquery.Selection) string {
	typeLink, _ := nameArea.Find("a").First().Attr("href")
	return utils.GetValueFromSplit(typeLink, ".php", 0)[1:]
}

// getProgress to get user's anime/manga history progress.
func (user *UserHistoryParser) getProgress(nameArea *goquery.Selection) int {
	return utils.StrToNum(nameArea.Find("strong").First().Text())
}

// getDate to get user's anime/manga history update date.
func (user *UserHistoryParser) getDate(history *goquery.Selection) string {
	date := history.Find("td:nth-of-type(2)")
	date.Find("a").Remove()
	return strings.TrimSpace(date.Text())
}
