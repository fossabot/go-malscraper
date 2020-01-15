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

// MangaParser is parser for MyAnimeList genre's manga list.
// Example: https://myanimelist.net/manga/genre/1/Action
type MangaParser struct {
	parser.BaseParser
	ID   int
	Page int
	Data []model.Manga
}

// InitMangaParser to initiate all fields and data of MangaParser.
func InitMangaParser(id int, page ...int) (genre MangaParser, err error) {
	genre.ID = id
	genre.Page = 1

	if len(page) > 0 {
		genre.Page = page[0]
	}

	err = genre.InitParser("/manga/genre/"+strconv.Itoa(genre.ID)+"/?page="+strconv.Itoa(genre.Page), "#contentWrapper")
	if err != nil {
		return genre, err
	}

	genre.setAllDetail()
	return genre, nil
}

// setAllDetail to set all genre detail information.
func (pp *MangaParser) setAllDetail() {
	var mangas []model.Manga

	pp.Parser.Find("div[class=\"seasonal-anime js-seasonal-anime\"]").Each(func(i int, eachArea *goquery.Selection) {
		nameArea := eachArea.Find("div.title")
		topArea := eachArea.Find("div.prodsrc")
		infoArea := eachArea.Find(".information")

		mangas = append(mangas, model.Manga{
			ID:             pp.getID(nameArea),
			Image:          pp.getImage(eachArea),
			Title:          pp.getTitle(nameArea),
			Genres:         pp.getGenres(eachArea),
			Synopsis:       pp.getSynopsis(eachArea),
			Authors:        pp.getAuthor(topArea),
			Volume:         pp.getProgress(topArea),
			Serializations: pp.getSerialization(eachArea),
			Type:           pp.getType(topArea),
			StartDate:      pp.getStartDate(infoArea),
			Member:         pp.getMember(infoArea),
			Score:          pp.getScore(infoArea),
		})
	})

	pp.Data = mangas
}

// getID to get manga's id.
func (pp *MangaParser) getID(nameArea *goquery.Selection) int {
	id, _ := nameArea.Find("p a").Attr("href")
	id = utils.GetValueFromSplit(id, "/", 4)
	return utils.StrToNum(id)
}

// getImage to get manga's image.
func (pp *MangaParser) getImage(eachArea *goquery.Selection) string {
	image, _ := eachArea.Find("div.image img").Attr("data-src")
	return utils.ImageURLCleaner(image)
}

// getTitle to get manga's title.
func (pp *MangaParser) getTitle(nameArea *goquery.Selection) string {
	return nameArea.Find("p a").Text()
}

// getGenre to get manga's genres.
func (pp *MangaParser) getGenres(eachArea *goquery.Selection) []common.Genre {
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

// getSynopsis to get manga's synopsis.
func (pp *MangaParser) getSynopsis(eachArea *goquery.Selection) string {
	synopsis := strings.TrimSpace(eachArea.Find("div[class=\"synopsis js-synopsis\"] .preline").Text())
	if synopsis == "(No synopsis yet.)" {
		return ""
	}
	return synopsis
}

// getType to get manga's type.
func (pp *MangaParser) getType(topArea *goquery.Selection) string {
	return strings.TrimSpace(topArea.Find("span.source").Text())
}

// getAuthor to get manga's genre.
func (pp *MangaParser) getAuthor(area *goquery.Selection) []common.IDName {
	var authors []common.IDName
	area = area.Find("span.producer")
	area.Find("a").Each(func(i int, each *goquery.Selection) {
		authors = append(authors, common.IDName{
			ID:   pp.getAuthorID(each),
			Name: pp.getAuthorName(each),
		})
	})
	return authors
}

// getAuthorID to get manga's genre id.
func (pp *MangaParser) getAuthorID(area *goquery.Selection) int {
	link, _ := area.Attr("href")
	id := utils.GetValueFromSplit(link, "/", 4)
	return utils.StrToNum(id)
}

// getAuthorName to get manga's genre name.
func (pp *MangaParser) getAuthorName(area *goquery.Selection) string {
	return area.Text()
}

// getProgress to get manga's episode.
func (pp *MangaParser) getProgress(area *goquery.Selection) int {
	progress := area.Find("div.eps").Text()
	replacer := strings.NewReplacer("eps", "", "ep", "", "vols", "", "vol", "")
	progress = replacer.Replace(progress)

	if progress == "?" {
		return 0
	}

	return utils.StrToNum(progress)
}

// getSerialization to get manga's serialization.
func (pp *MangaParser) getSerialization(eachArea *goquery.Selection) []string {
	serialization := eachArea.Find("div[class=\"synopsis js-synopsis\"] .serialization a").Text()
	return utils.ArrayFilter(strings.Split(serialization, ","))
}

// getStartDate to get manga's start airing date.
func (pp *MangaParser) getStartDate(area *goquery.Selection) string {
	airing := area.Find(".info .remain-time").Text()
	return strings.TrimSpace(airing)
}

// getMember to get manga's member number.
func (pp *MangaParser) getMember(area *goquery.Selection) int {
	member := area.Find(".scormem span[class^=member]").Text()
	return utils.StrToNum(member)
}

// getScore to get manga's score.
func (pp *MangaParser) getScore(area *goquery.Selection) float64 {
	score := area.Find(".scormem .score").Text()
	score = strings.TrimSpace(score)

	if score == "N/A" {
		return 0
	}

	return utils.StrToFloat(score)
}
