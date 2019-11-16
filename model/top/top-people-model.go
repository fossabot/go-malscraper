package top

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/helper"
	"github.com/rl404/go-malscraper/model"
)

// TopPeopleModel is an extended model from MainModel for top people list.
type TopPeopleModel struct {
	model.MainModel
	Page int
	Data []TopPeopleData
}

// InitTopPeopleModel to initiate fields in parent (MainModel) model.
func (i *TopPeopleModel) InitTopPeopleModel(p int) ([]TopPeopleData, int, string) {
	i.Page = 50 * (p - 1)
	i.InitModel("/people.php?limit="+strconv.Itoa(i.Page), "#content")

	i.Parser = i.GetParser(i.Url)

	if i.ResponseCode != 200 {
		return i.Data, i.ResponseCode, i.ErrorMessage
	}

	i.SetAllDetail()

	return i.Data, i.ResponseCode, i.ErrorMessage
}

// SetAllDetail to set all top character detail.
func (i *TopPeopleModel) SetAllDetail() {
	var topList []TopPeopleData
	area := i.Parser.Find("#content .people-favorites-ranking-table")
	area.Find("tr.ranking-list").Each(func(j int, eachPeople *goquery.Selection) {
		nameArea := eachPeople.Find(".people")

		topList = append(topList, TopPeopleData{
			Rank:         i.GetRank(eachPeople),
			Id:           i.GetId(nameArea),
			Name:         i.GetName(nameArea),
			JapaneseName: i.GetJapaneseName(nameArea),
			Image:        i.GetImage(nameArea),
			Birthday:     i.GetBirthday(eachPeople),
			Favorite:     i.GetFavorite(eachPeople),
		})
	})

	i.Data = topList
}

// GetRank to get people rank.
func (i *TopPeopleModel) GetRank(eachPeople *goquery.Selection) string {
	rank := eachPeople.Find("td").First().Find("span").Text()
	return strings.TrimSpace(rank)
}

//GetId to get people id.
func (i *TopPeopleModel) GetId(nameArea *goquery.Selection) string {
	id, _ := nameArea.Find("a").First().Attr("href")
	splitId := strings.Split(id, "/")

	return splitId[4]
}

// GetName to get people name.
func (i *TopPeopleModel) GetName(nameArea *goquery.Selection) string {
	return nameArea.Find(".information a").Text()
}

// GetJapaneseName to get people japanese name.
func (i *TopPeopleModel) GetJapaneseName(nameArea *goquery.Selection) string {
	japName := nameArea.Find(".information span").Text()

	if japName != "" {
		japName = japName[1 : len(japName)-1]
	}

	return japName
}

// GetImage to get people image.
func (i *TopPeopleModel) GetImage(nameArea *goquery.Selection) string {
	image, _ := nameArea.Find("img").First().Attr("data-src")
	return helper.ImageUrlCleaner(image)
}

// GetBirthday to get people birthday.
func (i *TopPeopleModel) GetBirthday(eachPeople *goquery.Selection) string {
	day := eachPeople.Find(".birthday").Text()
	day = strings.TrimSpace(day)

	r, _ := regexp.Compile(`\s+`)
	day = r.ReplaceAllString(day, " ")

	if day == "Unknown" {
		day = ""
	}

	return day
}

// GetFavorite to get people number favorite.
func (i *TopPeopleModel) GetFavorite(eachPeople *goquery.Selection) string {
	fav := eachPeople.Find(".favorites").Text()
	return strings.TrimSpace(strings.Replace(fav, ",", "", -1))
}
