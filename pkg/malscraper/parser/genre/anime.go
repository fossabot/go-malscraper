package genre

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/pkg/malscraper/model/common"
	model "github.com/rl404/go-malscraper/pkg/malscraper/model/genre"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser"
	"github.com/rl404/go-malscraper/pkg/malscraper/utils"
)

// AnimeParser is parser for MyAnimeList genre's anime list.
// Example: https://myanimelist.net/anime/genre/1/Action
type AnimeParser struct {
	parser.BaseParser
	ID   int
	Page int
	Data []model.Anime
}

// InitAnimeParser to initiate all fields and data of AnimeParser.
func InitAnimeParser(id int, page ...int) (genre AnimeParser, err error) {
	genre.ID = id
	genre.Page = 1

	if len(page) > 0 {
		genre.Page = page[0]
	}

	err = genre.InitParser("/anime/genre/"+strconv.Itoa(genre.ID)+"/?page="+strconv.Itoa(genre.Page), "#contentWrapper")
	if err != nil {
		return genre, err
	}

	genre.setAllDetail()
	return genre, nil
}

// setAllDetail to set all genre's anime detail information.
func (ap *AnimeParser) setAllDetail() {
	var genres []model.Anime

	ap.Parser.Find("div[class=\"seasonal-anime js-seasonal-anime\"]").Each(func(i int, eachArea *goquery.Selection) {
		nameArea := eachArea.Find("div.title")
		topArea := eachArea.Find("div.prodsrc")
		infoArea := eachArea.Find(".information")

		genres = append(genres, model.Anime{
			ID:        ap.getID(nameArea),
			Image:     ap.getImage(eachArea),
			Title:     ap.getTitle(nameArea),
			Genres:    ap.getGenres(eachArea),
			Synopsis:  ap.getSynopsis(eachArea),
			Source:    ap.getSource(topArea),
			Producers: ap.getProducer(topArea),
			Episode:   ap.getProgress(topArea),
			Licensors: ap.getLicensors(eachArea),
			Type:      ap.getType(infoArea),
			StartDate: ap.getStartDate(infoArea),
			Member:    ap.getMember(infoArea),
			Score:     ap.getScore(infoArea),
		})
	})

	ap.Data = genres
}

// getID to get anime's id.
func (ap *AnimeParser) getID(nameArea *goquery.Selection) int {
	id, _ := nameArea.Find("p a").Attr("href")
	id = utils.GetValueFromSplit(id, "/", 4)
	return utils.StrToNum(id)
}

// getImage to get anime's image.
func (ap *AnimeParser) getImage(eachArea *goquery.Selection) string {
	image, _ := eachArea.Find("div.image img").Attr("data-src")
	return utils.ImageURLCleaner(image)
}

// getTitle to get anime's title.
func (ap *AnimeParser) getTitle(nameArea *goquery.Selection) string {
	return nameArea.Find("p a").Text()
}

// getGenre to get anime's genres.
func (ap *AnimeParser) getGenres(eachArea *goquery.Selection) []common.Genre {
	var genres []common.Genre
	genreArea := eachArea.Find("div[class=\"genres js-genre\"]")
	genreArea.Find("a").Each(func(i int, genre *goquery.Selection) {
		genreLink, _ := genre.Attr("href")
		splitLink := strings.Split(genreLink, "/")
		genres = append(genres, common.Genre{
			ID:   utils.StrToNum(splitLink[3]),
			Type: splitLink[1],
			Name: genre.Text(),
		})
	})
	return genres
}

// getSynopsis to get anime's synopsis.
func (ap *AnimeParser) getSynopsis(eachArea *goquery.Selection) string {
	synopsis := strings.TrimSpace(eachArea.Find("div[class=\"synopsis js-synopsis\"]").Text())
	if synopsis == "(No synopsis yet.)" {
		return ""
	}
	return synopsis
}

// getSource to get anime's source.
func (ap *AnimeParser) getSource(topArea *goquery.Selection) string {
	return strings.TrimSpace(topArea.Find("span.source").Text())
}

// getProducer to get anime's genre.
func (ap *AnimeParser) getProducer(area *goquery.Selection) []common.IDName {
	var genres []common.IDName
	area = area.Find("span.producer")
	area.Find("a").Each(func(i int, each *goquery.Selection) {
		genres = append(genres, common.IDName{
			ID:   ap.getProducerID(each),
			Name: ap.getProducerName(each),
		})
	})
	return genres
}

// getProducerID to get anime's genre id.
func (ap *AnimeParser) getProducerID(area *goquery.Selection) int {
	link, _ := area.Attr("href")
	id := utils.GetValueFromSplit(link, "/", 3)
	return utils.StrToNum(id)
}

// getProducerName to get anime's genre name.
func (ap *AnimeParser) getProducerName(area *goquery.Selection) string {
	return area.Text()
}

// getProgress to get anime's episode.
func (ap *AnimeParser) getProgress(area *goquery.Selection) int {
	progress := area.Find("div.eps").Text()
	replacer := strings.NewReplacer("eps", "", "ep", "", "vols", "", "vol", "")
	progress = replacer.Replace(progress)

	if progress == "?" {
		return 0
	}

	return utils.StrToNum(progress)
}

// getLicensors to get anime's licensor.
func (ap *AnimeParser) getLicensors(eachArea *goquery.Selection) []string {
	licensor, _ := eachArea.Find("div[class=\"synopsis js-synopsis\"] .licensors").Attr("data-licensors")
	return utils.ArrayFilter(strings.Split(licensor, ","))
}

// getType to get anime's type (TV, movie, ova, etc).
func (ap *AnimeParser) getType(area *goquery.Selection) string {
	typeArea := area.Find(".info").Text()
	return utils.GetValueFromSplit(typeArea, "-", 0)
}

// getStartDate to get anime's start airing date.
func (ap *AnimeParser) getStartDate(area *goquery.Selection) string {
	airing := area.Find(".info .remain-time").Text()
	return strings.TrimSpace(airing)
}

// getMember to get anime's member number.
func (ap *AnimeParser) getMember(area *goquery.Selection) int {
	member := area.Find(".scormem span[class^=member]").Text()
	return utils.StrToNum(member)
}

// getScore to get anime's score.
func (ap *AnimeParser) getScore(area *goquery.Selection) float64 {
	score := area.Find(".scormem .score").Text()
	score = strings.TrimSpace(score)

	if score == "N/A" {
		return 0
	}

	return utils.StrToFloat(score)
}
