package search

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/helper"
	"github.com/rl404/go-malscraper/model"
)

// SearchCharPeopleModel is an extended model from MainModel for character/people search list.
type SearchCharPeopleModel struct {
	model.MainModel
	Type  string
	Query string
	Page  int
	Data  []SearchCharPeopleData
}

// InitSearchCharPeopleModel to initiate fields in parent (MainModel) model.
func (i *SearchCharPeopleModel) InitSearchCharPeopleModel(t string, query string, page int) ([]SearchCharPeopleData, int, string) {
	i.Type = t
	i.Query = query
	i.Page = 50 * (page - 1)

	if len(i.Query) < 3 {
		i.SetMessage(400, "Search query needs at least 3 letters")
		return i.Data, i.ResponseCode, i.ErrorMessage
	}

	i.InitModel("/"+i.Type+".php?q="+i.Query+"&show="+strconv.Itoa(i.Page), "#content")

	i.Parser = i.GetParser(i.Url)

	if i.ResponseCode != 200 {
		return i.Data, i.ResponseCode, i.ErrorMessage
	}

	i.SetAllDetail()

	return i.Data, i.ResponseCode, i.ErrorMessage
}

// SetAllDetail to fill all character/people search list.
func (i *SearchCharPeopleModel) SetAllDetail() {
	var searchList []SearchCharPeopleData
	area := i.Parser.Find("#content table")
	area.Find("tr").EachWithBreak(func(j int, eachSearch *goquery.Selection) bool {

		if j == 0 {
			return true
		}

		nameArea := eachSearch.Find("td:nth-of-type(2)")

		searchList = append(searchList, SearchCharPeopleData{
			Image:    i.GetImage(eachSearch),
			Id:       i.GetId(nameArea),
			Name:     i.GetName(nameArea),
			Nickname: i.GetNickname(nameArea),
			Anime:    i.GetRole("anime", eachSearch),
			Manga:    i.GetRole("manga", eachSearch),
		})

		return true
	})

	i.Data = searchList
}

// GetImage to get character/people image.
func (i *SearchCharPeopleModel) GetImage(eachSearch *goquery.Selection) string {
	image, _ := eachSearch.Find("td a img").Attr("src")
	return helper.ImageUrlCleaner(image)
}

// GetId to get character/people id.
func (i *SearchCharPeopleModel) GetId(nameArea *goquery.Selection) string {
	id, _ := nameArea.Find("a").First().Attr("href")
	splitId := strings.Split(id, "/")

	if i.Type == "character" {
		return splitId[4]
	}
	return splitId[2]
}

// GetName to get character/people name.
func (i *SearchCharPeopleModel) GetName(nameArea *goquery.Selection) string {
	return nameArea.Find("a").First().Text()
}

// GetNickname to get character/people nickname.
func (i *SearchCharPeopleModel) GetNickname(nameArea *goquery.Selection) string {
	nick := nameArea.Find("small").First().Text()
	if nick != "" {
		return nick[1 : len(nick)-1]
	}
	return nick
}

// GetRole to get character anime/manga role.
func (i *SearchCharPeopleModel) GetRole(t string, eachSearch *goquery.Selection) []IdTitle {
	var roleList []IdTitle
	area := eachSearch.Find("td:nth-of-type(3) small")
	area.Find("a").Each(func(j int, eachRole *goquery.Selection) {
		id, _ := eachRole.Attr("href")
		splitId := strings.Split(id, "/")

		roleType := splitId[1]

		if t == roleType && splitId[2] != "" {
			roleList = append(roleList, IdTitle{
				Id:    splitId[2],
				Title: eachRole.Text(),
			})
		}
	})

	return roleList
}
