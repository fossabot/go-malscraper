package anime

import (
	"strconv"

	"github.com/PuerkitoBio/goquery"
	model "github.com/rl404/go-malscraper/pkg/malscraper/model/anime"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser"
	"github.com/rl404/go-malscraper/pkg/malscraper/utils"
)

// StaffParser is parser for MyAnimeList anime staff list.
// Example: https://myanimelist.net/anime/1/Cowboy_Bebop/characters
type StaffParser struct {
	parser.BaseParser
	ID   int
	Data []model.Staff
}

// InitStaffParser to initiate all fields and data of StaffParser.
func InitStaffParser(id int) (staff StaffParser, err error) {
	staff.ID = id

	err = staff.InitParser("/anime/"+strconv.Itoa(staff.ID)+"/a/characters", ".js-scrollfix-bottom-rel")
	if err != nil {
		return staff, err
	}

	staff.setAllDetail()
	return staff, nil
}

// setAllDetail to set all character and staff detail information.
func (csp *StaffParser) setAllDetail() {
	var staffList []model.Staff

	csp.Parser.Find("article").Remove()
	csp.Parser.Find(".js-scrollfix-bottom-rel h2").First().Remove()

	staffArea := csp.Parser.Find(".js-scrollfix-bottom-rel h2").First().Next()

	for goquery.NodeName(staffArea) == "table" {

		staffNameArea := staffArea.Find("td:nth-of-type(2)")

		staffList = append(staffList, model.Staff{
			ID:    csp.getStaffID(staffNameArea),
			Image: csp.getStaffImage(staffArea),
			Name:  csp.getStaffName(staffNameArea),
			Role:  csp.getStaffRole(staffNameArea),
		})

		staffArea = staffArea.Next()
	}

	csp.Data = staffList
}

// getStaffID to get character id.
func (csp *StaffParser) getStaffID(charNameArea *goquery.Selection) int {
	id, _ := charNameArea.Find("a").First().Attr("href")
	id = utils.GetValueFromSplit(id, "/", 4)
	return utils.StrToNum(id)
}

// getStaffImage to get character image.
func (csp *StaffParser) getStaffImage(charArea *goquery.Selection) string {
	image, _ := charArea.Find("td .picSurround img").Attr("data-src")
	return utils.ImageURLCleaner(image)
}

// getStaffName to get character name.
func (csp *StaffParser) getStaffName(charNameArea *goquery.Selection) string {
	return charNameArea.Find("a").First().Text()
}

// getStaffRole to get character role.
func (csp *StaffParser) getStaffRole(charNameArea *goquery.Selection) string {
	return charNameArea.Find("small").First().Text()
}
