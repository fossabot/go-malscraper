package user

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
)

// UserHistoryModel is an extended model from MainModel for user friend list.
type UserHistoryModel struct {
	model.MainModel
	Username string
	Type     string
	Data     []UserHistoryData
}

// InitUserHistoryModel to initiate fields in parent (MainModel) model.
func (i *UserHistoryModel) InitUserHistoryModel(u string, t string) ([]UserHistoryData, int, string) {
	i.Username = u
	i.Type = t

	if i.Type == "" {
		i.InitModel("/history/"+i.Username, "#content")
	} else {
		i.InitModel("/history/"+i.Username+"/"+i.Type, "#content")
	}

	i.Parser = i.GetParser(i.Url)

	if i.ResponseCode != 200 {
		return i.Data, i.ResponseCode, i.ErrorMessage
	}

	i.SetAllDetail()

	return i.Data, i.ResponseCode, i.ErrorMessage
}

// SetAllDetail to set user history detail data.
func (i *UserHistoryModel) SetAllDetail() {
	var historyList []UserHistoryData
	area := i.Parser.Find("#content table")
	area.Find("tr").EachWithBreak(func(j int, eachHistory *goquery.Selection) bool {

		historyClass, _ := eachHistory.Find("td").First().Attr("class")
		if historyClass != "borderClass" {
			return true
		}

		nameArea := eachHistory.Find("td").First()

		historyList = append(historyList, UserHistoryData{
			Id:       i.GetId(nameArea),
			Title:    i.GetTitle(nameArea),
			Type:     i.GetType(nameArea),
			Progress: i.GetProgress(nameArea),
			Date:     i.GetDate(eachHistory),
		})

		return true
	})

	i.Data = historyList
}

// GetId to get anime/manga id.
func (i *UserHistoryModel) GetId(nameArea *goquery.Selection) string {
	id, _ := nameArea.Find("a").First().Attr("href")
	splitId := strings.Split(id, "=")

	return splitId[1]
}

// GetTitle to get anime/manga title.
func (i *UserHistoryModel) GetTitle(nameArea *goquery.Selection) string {
	return nameArea.Find("a").First().Text()
}

// GetType to get anime/manga type.
func (i *UserHistoryModel) GetType(nameArea *goquery.Selection) string {
	t, _ := nameArea.Find("a").First().Attr("href")
	splitType := strings.Split(t, ".php")

	return splitType[0][1:]
}

// GetProgress to get anime/manga progress.
func (i *UserHistoryModel) GetProgress(nameArea *goquery.Selection) string {
	return nameArea.Find("strong").First().Text()
}

// GetDate to get anime/manga history update date.
func (i *UserHistoryModel) GetDate(eachHistory *goquery.Selection) string {
	date := eachHistory.Find("td:nth-of-type(2)")
	date.Find("a").Remove()

	return strings.TrimSpace(date.Text())
}
