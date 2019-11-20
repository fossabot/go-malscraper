package top

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/helper"
	"github.com/rl404/go-malscraper/model"
)

// TopCharacterModel is an extended model from MainModel for top character list.
type TopCharacterModel struct {
	model.MainModel
	Page int
	Data []TopCharacterData
}

// InitTopCharacterModel to initiate fields in parent (MainModel) model.
func (i *TopCharacterModel) InitTopCharacterModel(p int) ([]TopCharacterData, int, string) {
	i.Page = 50 * (p - 1)
	i.InitModel("/character.php?limit="+strconv.Itoa(i.Page), "#content")

	i.Parser = i.GetParser(i.Url)

	if i.ResponseCode != 200 {
		return i.Data, i.ResponseCode, i.ErrorMessage
	}

	i.SetAllDetail()

	return i.Data, i.ResponseCode, i.ErrorMessage
}

// SetAllDetail to set all top character detail.
func (i *TopCharacterModel) SetAllDetail() {
	var topList []TopCharacterData
	area := i.Parser.Find("#content .characters-favorites-ranking-table")
	area.Find("tr.ranking-list").Each(func(j int, eachChar *goquery.Selection) {
		nameArea := eachChar.Find(".people")

		topList = append(topList, TopCharacterData{
			Rank:         i.GetRank(eachChar),
			Id:           i.GetId(nameArea),
			Name:         i.GetName(nameArea),
			JapaneseName: i.GetJapaneseName(nameArea),
			Image:        i.GetImage(nameArea),
			Animeography: i.GetRole(eachChar, ".animeography"),
			Mangaography: i.GetRole(eachChar, ".mangaography"),
			Favorite:     i.GetFavorite(eachChar),
		})
	})

	i.Data = topList
}

// GetRank to get character rank.
func (i *TopCharacterModel) GetRank(eachChar *goquery.Selection) int {
	rank := eachChar.Find("td").First().Find("span").Text()
	rankInt, _ := strconv.Atoi(strings.TrimSpace(rank))
	return rankInt
}

// GetId to get character id.
func (i *TopCharacterModel) GetId(nameArea *goquery.Selection) int {
	id, _ := nameArea.Find("a").First().Attr("href")
	splitId := strings.Split(id, "/")
	idInt, _ := strconv.Atoi(splitId[4])
	return idInt
}

// GetName to get character name.
func (i *TopCharacterModel) GetName(nameArea *goquery.Selection) string {
	return nameArea.Find(".information a").Text()
}

// GetJapaneseName to get character japanese name.
func (i *TopCharacterModel) GetJapaneseName(nameArea *goquery.Selection) string {
	japName := nameArea.Find(".information span").Text()

	if japName != "" {
		japName = japName[1 : len(japName)-1]
	}

	return japName
}

// GetImage to get character image.
func (i *TopCharacterModel) GetImage(nameArea *goquery.Selection) string {
	image, _ := nameArea.Find("img").First().Attr("data-src")
	return helper.ImageUrlCleaner(image)
}

// GetRole to get character role in anime + manga.
func (i *TopCharacterModel) GetRole(eachChar *goquery.Selection, areaClass string) []IdTitle {
	var roleList []IdTitle
	area := eachChar.Find(areaClass)
	area.Find(".title").Each(func(j int, eachRole *goquery.Selection) {
		linkA := eachRole.Find("a").First()
		link, _ := linkA.Attr("href")
		splitId := strings.Split(link, "/")
		idInt, _ := strconv.Atoi(splitId[4])
		roleList = append(roleList, IdTitle{
			Id:    idInt,
			Title: linkA.Text(),
		})
	})

	return roleList
}

// GetFavorite to get character favorite number.
func (i *TopCharacterModel) GetFavorite(eachChar *goquery.Selection) int {
	fav := eachChar.Find(".favorites").Text()
	fav = strings.TrimSpace(strings.Replace(fav, ",", "", -1))
	favInt, _ := strconv.Atoi(fav)
	return favInt
}
