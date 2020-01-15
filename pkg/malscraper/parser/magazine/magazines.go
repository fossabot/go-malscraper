package magazine

import (
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	model "github.com/rl404/go-malscraper/pkg/malscraper/model/magazine"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser"
	"github.com/rl404/go-malscraper/pkg/malscraper/utils"
)

// MagazinesParser is parser for MyAnimeList all magazines, and serializations.
// Example: https://myanimelist.net/manga/magazine/1
type MagazinesParser struct {
	parser.BaseParser
	Data []model.Magazine
}

// InitMagazinesParser to initiate all fields and data of MagazinesParser.
func InitMagazinesParser() (magazines MagazinesParser, err error) {
	err = magazines.InitParser("/manga/magazine", ".anime-manga-search")
	if err != nil {
		return magazines, err
	}

	magazines.setAllDetail()
	return magazines, nil
}

// setAllDetail to set all magazine detail information.
func (mp *MagazinesParser) setAllDetail() {
	var magazines []model.Magazine

	mp.Parser.Find(".genre-list a").Each(func(i int, area *goquery.Selection) {
		magazines = append(magazines, model.Magazine{
			ID:    mp.getID(area),
			Name:  mp.getName(area),
			Count: mp.getCount(area),
		})
	})

	mp.Data = magazines
}

// getID to get magazine id.
func (mp *MagazinesParser) getID(area *goquery.Selection) int {
	link, _ := area.Attr("href")
	id := utils.GetValueFromSplit(link, "/", 3)
	return utils.StrToNum(id)
}

// getName to get magazine name.
func (mp *MagazinesParser) getName(area *goquery.Selection) string {
	name := area.Text()
	r := regexp.MustCompile(`\([0-9,]+\)`)
	name = r.ReplaceAllString(name, "")
	return strings.TrimSpace(name)
}

// getCount to get magazine anime count.
func (mp *MagazinesParser) getCount(area *goquery.Selection) int {
	count := area.Text()
	r := regexp.MustCompile(`\([0-9,]+\)`)
	count = r.FindString(count)
	count = count[1 : len(count)-1]
	return utils.StrToNum(count)
}
