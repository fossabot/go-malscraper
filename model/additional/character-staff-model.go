package additional

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/helper"
	"github.com/rl404/go-malscraper/model"
)

// CharacterStaffModel is an extended model from MainModel for anime/manga additional character+staff list.
type CharacterStaffModel struct {
	model.MainModel
	Type string
	Id   int
	Data CharacterStaffData
}

// InitVideoModel to initiate fields in parent (MainModel) model.
func (i *CharacterStaffModel) InitCharacterStaffModel(t string, id int) (CharacterStaffData, int, string) {
	i.Type = t
	i.Id = id
	i.InitModel("/"+i.Type+"/"+strconv.Itoa(i.Id)+"/a/characters", ".js-scrollfix-bottom-rel")

	i.Parser = i.GetParser(i.Url)

	if i.ResponseCode != 200 {
		return i.Data, i.ResponseCode, i.ErrorMessage
	}

	i.SetAllDetail()

	return i.Data, i.ResponseCode, i.ErrorMessage
}

// SetAllDetail to fill all additional character + staff data.
func (i *CharacterStaffModel) SetAllDetail() {
	i.SetCharacter()
	i.SetStaff()
}

// SetCharacter to set character data.
func (i *CharacterStaffModel) SetCharacter() {
	var charList []Character

	i.Parser.Find(".js-scrollfix-bottom-rel").Find("article").Remove()

	charArea := i.Parser.Find(".js-scrollfix-bottom-rel h2").First().Next()

	for goquery.NodeName(charArea) == "table" {

		charNameArea := charArea.Find("td:nth-of-type(2)")
		vaArea := charArea.Find("td:nth-of-type(3)")

		charList = append(charList, Character{
			Id:    i.GetCharId(charNameArea),
			Image: i.GetCharImage(charArea),
			Name:  i.GetCharName(charNameArea),
			Role:  i.GetCharRole(charNameArea),
			Va:    i.GetVa(vaArea),
		})

		charArea = charArea.Next()
	}

	i.Data.Character = charList
}

// GetCharId to get character id.
func (i *CharacterStaffModel) GetCharId(charNameArea *goquery.Selection) int {
	id, _ := charNameArea.Find("a").First().Attr("href")
	splitId := strings.Split(id, "/")
	idInt, _ := strconv.Atoi(splitId[4])
	return idInt
}

// GetCharImage to get character image.
func (i *CharacterStaffModel) GetCharImage(charArea *goquery.Selection) string {
	image, _ := charArea.Find("td .picSurround img").Attr("data-src")
	return helper.ImageUrlCleaner(image)
}

// GetCharName to get character name.
func (i *CharacterStaffModel) GetCharName(charNameArea *goquery.Selection) string {
	return charNameArea.Find("a").First().Text()
}

// GetCharRole to get character role.
func (i *CharacterStaffModel) GetCharRole(charNameArea *goquery.Selection) string {
	return charNameArea.Find("small").First().Text()
}

// GetVa to get va list.
func (i *CharacterStaffModel) GetVa(vaArea *goquery.Selection) []Staff {
	var vaList []Staff

	vaArea = vaArea.Find("table")
	vaArea.Find("tr").Each(func(j int, eachVa *goquery.Selection) {

		vaNameArea := eachVa.Find("td").First()

		vaList = append(vaList, Staff{
			Id:    i.GetCharId(vaNameArea),
			Image: i.GetCharImage(eachVa),
			Name:  i.GetCharName(vaNameArea),
			Role:  i.GetCharRole(vaNameArea),
		})
	})

	return vaList
}

// SetStaff to set staff data.
func (i *CharacterStaffModel) SetStaff() {
	var staffList []Staff

	i.Parser.Find(".js-scrollfix-bottom-rel h2").First().Remove()

	staffArea := i.Parser.Find(".js-scrollfix-bottom-rel h2").First().Next()

	for goquery.NodeName(staffArea) == "table" {

		staffNameArea := staffArea.Find("td:nth-of-type(2)")

		staffList = append(staffList, Staff{
			Id:    i.GetCharId(staffNameArea),
			Image: i.GetCharImage(staffArea),
			Name:  i.GetCharName(staffNameArea),
			Role:  i.GetCharRole(staffNameArea),
		})

		staffArea = staffArea.Next()
	}

	i.Data.Staff = staffList
}
