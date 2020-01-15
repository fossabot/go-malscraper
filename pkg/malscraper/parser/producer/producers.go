package producer

import (
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	model "github.com/rl404/go-malscraper/pkg/malscraper/model/producer"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser"
	"github.com/rl404/go-malscraper/pkg/malscraper/utils"
)

// ProducersParser is parser for MyAnimeList all producers, studios, and licensors.
// Example: https://myanimelist.net/anime/producer/1
type ProducersParser struct {
	parser.BaseParser
	Data []model.Producer
}

// InitProducersParser to initiate all fields and data of ProducersParser.
func InitProducersParser() (producers ProducersParser, err error) {
	err = producers.InitParser("/anime/producer", ".anime-manga-search")
	if err != nil {
		return producers, err
	}

	producers.setAllDetail()
	return producers, nil
}

// setAllDetail to set all producer detail information.
func (pp *ProducersParser) setAllDetail() {
	var producers []model.Producer

	pp.Parser.Find(".genre-list a").Each(func(i int, area *goquery.Selection) {
		producers = append(producers, model.Producer{
			ID:    pp.getID(area),
			Name:  pp.getName(area),
			Count: pp.getCount(area),
		})
	})

	pp.Data = producers
}

// getID to get producer id.
func (pp *ProducersParser) getID(area *goquery.Selection) int {
	link, _ := area.Attr("href")
	id := utils.GetValueFromSplit(link, "/", 3)
	return utils.StrToNum(id)
}

// getName to get producer name.
func (pp *ProducersParser) getName(area *goquery.Selection) string {
	name := area.Text()
	r := regexp.MustCompile(`\([0-9,]+\)`)
	name = r.ReplaceAllString(name, "")
	return strings.TrimSpace(name)
}

// getCount to get producer anime count.
func (pp *ProducersParser) getCount(area *goquery.Selection) int {
	count := area.Text()
	r := regexp.MustCompile(`\([0-9,]+\)`)
	count = r.FindString(count)
	count = count[1 : len(count)-1]
	return utils.StrToNum(count)
}
