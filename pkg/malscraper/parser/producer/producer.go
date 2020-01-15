package producer

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/pkg/malscraper/model/common"
	model "github.com/rl404/go-malscraper/pkg/malscraper/model/producer"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser"
	"github.com/rl404/go-malscraper/pkg/malscraper/utils"
)

// ProducerParser is parser for MyAnimeList producer/studio/licensor's anime list.
// Example: https://myanimelist.net/anime/producer/1/Studio_Pierrot
type ProducerParser struct {
	parser.BaseParser
	ID   int
	Page int
	Data []model.Anime
}

// InitProducerParser to initiate all fields and data of ProducerParser.
func InitProducerParser(id int, page ...int) (producer ProducerParser, err error) {
	producer.ID = id
	producer.Page = 1

	if len(page) > 0 {
		producer.Page = page[0]
	}

	err = producer.InitParser("/anime/producer/"+strconv.Itoa(producer.ID)+"/?page="+strconv.Itoa(producer.Page), "#content .js-categories-seasonal")
	if err != nil {
		return producer, err
	}

	producer.setAllDetail()
	return producer, nil
}

// setAllDetail to set all producer detail information.
func (pp *ProducerParser) setAllDetail() {
	var producers []model.Anime

	pp.Parser.Find("div[class=\"seasonal-anime js-seasonal-anime\"]").Each(func(i int, eachArea *goquery.Selection) {
		nameArea := eachArea.Find("div.title")
		topArea := eachArea.Find("div.prodsrc")
		infoArea := eachArea.Find(".information")

		producers = append(producers, model.Anime{
			ID:        pp.getID(nameArea),
			Image:     pp.getImage(eachArea),
			Title:     pp.getTitle(nameArea),
			Genres:    pp.getGenres(eachArea),
			Synopsis:  pp.getSynopsis(eachArea),
			Source:    pp.getSource(topArea),
			Producers: pp.getProducer(topArea),
			Episode:   pp.getProgress(topArea),
			Licensors: pp.getLicensors(eachArea),
			Type:      pp.getType(infoArea),
			StartDate: pp.getStartDate(infoArea),
			Member:    pp.getMember(infoArea),
			Score:     pp.getScore(infoArea),
		})
	})

	pp.Data = producers
}

// getID to get anime's id.
func (pp *ProducerParser) getID(nameArea *goquery.Selection) int {
	id, _ := nameArea.Find("p a").Attr("href")
	id = utils.GetValueFromSplit(id, "/", 4)
	return utils.StrToNum(id)
}

// getImage to get anime's image.
func (pp *ProducerParser) getImage(eachArea *goquery.Selection) string {
	image, _ := eachArea.Find("div.image img").Attr("data-src")
	return utils.ImageURLCleaner(image)
}

// getTitle to get anime's title.
func (pp *ProducerParser) getTitle(nameArea *goquery.Selection) string {
	return nameArea.Find("p a").Text()
}

// getGenre to get anime's genres.
func (pp *ProducerParser) getGenres(eachArea *goquery.Selection) []model.Genre {
	var genres []model.Genre
	genreArea := eachArea.Find("div[class=\"genres js-genre\"]")
	genreArea.Find("a").Each(func(i int, genre *goquery.Selection) {
		genreLink, _ := genre.Attr("href")
		splitLink := strings.Split(genreLink, "/")
		genres = append(genres, model.Genre{
			ID:   utils.StrToNum(splitLink[3]),
			Type: splitLink[1],
			Name: genre.Text(),
		})
	})
	return genres
}

// getSynopsis to get anime's synopsis.
func (pp *ProducerParser) getSynopsis(eachArea *goquery.Selection) string {
	return strings.TrimSpace(eachArea.Find("div[class=\"synopsis js-synopsis\"]").Text())
}

// getSource to get anime's source.
func (pp *ProducerParser) getSource(topArea *goquery.Selection) string {
	return strings.TrimSpace(topArea.Find("span.source").Text())
}

// getProducer to get anime's producer.
func (pp *ProducerParser) getProducer(area *goquery.Selection) []common.IDName {
	var producers []common.IDName
	area = area.Find("span.producer")
	area.Find("a").Each(func(i int, each *goquery.Selection) {
		producers = append(producers, common.IDName{
			ID:   pp.getProducerID(each),
			Name: pp.getProducerName(each),
		})
	})
	return producers
}

// getProducerID to get anime's producer id.
func (pp *ProducerParser) getProducerID(area *goquery.Selection) int {
	link, _ := area.Attr("href")
	id := utils.GetValueFromSplit(link, "/", 3)
	return utils.StrToNum(id)
}

// getProducerName to get anime's producer name.
func (pp *ProducerParser) getProducerName(area *goquery.Selection) string {
	return area.Text()
}

// getProgress to get anime's episode.
func (pp *ProducerParser) getProgress(area *goquery.Selection) int {
	progress := area.Find("div.eps").Text()
	replacer := strings.NewReplacer("eps", "", "ep", "", "vols", "", "vol", "")
	progress = replacer.Replace(progress)

	if progress == "?" {
		return 0
	}

	return utils.StrToNum(progress)
}

// getLicensors to get anime's licensor.
func (pp *ProducerParser) getLicensors(eachArea *goquery.Selection) []string {
	licensor, _ := eachArea.Find("div[class=\"synopsis js-synopsis\"] .licensors").Attr("data-licensors")
	return utils.ArrayFilter(strings.Split(licensor, ","))
}

// getType to get anime's type (TV, movie, ova, etc).
func (pp *ProducerParser) getType(area *goquery.Selection) string {
	typeArea := area.Find(".info").Text()
	return utils.GetValueFromSplit(typeArea, "-", 0)
}

// getStartDate to get anime's start airing date.
func (pp *ProducerParser) getStartDate(area *goquery.Selection) string {
	airing := area.Find(".info .remain-time").Text()
	return strings.TrimSpace(airing)
}

// getMember to get anime's member number.
func (pp *ProducerParser) getMember(area *goquery.Selection) int {
	member := area.Find(".scormem span[class^=member]").Text()
	return utils.StrToNum(member)
}

// getScore to get anime's score.
func (pp *ProducerParser) getScore(area *goquery.Selection) float64 {
	score := area.Find(".scormem .score").Text()
	score = strings.TrimSpace(score)

	if score == "N/A" {
		return 0
	}

	return utils.StrToFloat(score)
}
