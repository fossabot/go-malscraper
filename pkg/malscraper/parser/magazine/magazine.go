package magazine

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/pkg/malscraper/model/common"
	model "github.com/rl404/go-malscraper/pkg/malscraper/model/magazine"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser"
	"github.com/rl404/go-malscraper/pkg/malscraper/utils"
)

// MagazineParser is parser for MyAnimeList magazine/serialization's manga list.
// Example: https://myanimelist.net/manga/magazine/1/Big_Comic_Original
type MagazineParser struct {
	parser.BaseParser
	ID   int
	Page int
	Data []model.Manga
}

// InitMagazineParser to initiate all fields and data of MagazineParser.
func InitMagazineParser(id int, page ...int) (magazine MagazineParser, err error) {
	magazine.ID = id
	magazine.Page = 1

	if len(page) > 0 {
		magazine.Page = page[0]
	}

	err = magazine.InitParser("/manga/magazine/"+strconv.Itoa(magazine.ID)+"/?page="+strconv.Itoa(magazine.Page), "#content .js-categories-seasonal")
	if err != nil {
		return magazine, err
	}

	magazine.setAllDetail()
	return magazine, nil
}

// setAllDetail to set all magazine detail information.
func (pp *MagazineParser) setAllDetail() {
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
func (pp *MagazineParser) getID(nameArea *goquery.Selection) int {
	id, _ := nameArea.Find("p a").Attr("href")
	id = utils.GetValueFromSplit(id, "/", 4)
	return utils.StrToNum(id)
}

// getImage to get manga's image.
func (pp *MagazineParser) getImage(eachArea *goquery.Selection) string {
	image, _ := eachArea.Find("div.image img").Attr("data-src")
	return utils.ImageURLCleaner(image)
}

// getTitle to get manga's title.
func (pp *MagazineParser) getTitle(nameArea *goquery.Selection) string {
	return nameArea.Find("p a").Text()
}

// getGenre to get manga's genres.
func (pp *MagazineParser) getGenres(eachArea *goquery.Selection) []model.Genre {
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

// getSynopsis to get manga's synopsis.
func (pp *MagazineParser) getSynopsis(eachArea *goquery.Selection) string {
	synopsis := strings.TrimSpace(eachArea.Find("div[class=\"synopsis js-synopsis\"] .preline").Text())
	if synopsis == "(No synopsis yet.)" {
		return ""
	}
	return synopsis
}

// getType to get manga's type.
func (pp *MagazineParser) getType(topArea *goquery.Selection) string {
	return strings.TrimSpace(topArea.Find("span.source").Text())
}

// getAuthor to get manga's magazine.
func (pp *MagazineParser) getAuthor(area *goquery.Selection) []common.IDName {
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

// getAuthorID to get manga's magazine id.
func (pp *MagazineParser) getAuthorID(area *goquery.Selection) int {
	link, _ := area.Attr("href")
	id := utils.GetValueFromSplit(link, "/", 4)
	return utils.StrToNum(id)
}

// getAuthorName to get manga's magazine name.
func (pp *MagazineParser) getAuthorName(area *goquery.Selection) string {
	return area.Text()
}

// getProgress to get manga's episode.
func (pp *MagazineParser) getProgress(area *goquery.Selection) int {
	progress := area.Find("div.eps").Text()
	replacer := strings.NewReplacer("eps", "", "ep", "", "vols", "", "vol", "")
	progress = replacer.Replace(progress)

	if progress == "?" {
		return 0
	}

	return utils.StrToNum(progress)
}

// getSerialization to get manga's serialization.
func (pp *MagazineParser) getSerialization(eachArea *goquery.Selection) []string {
	serialization := eachArea.Find("div[class=\"synopsis js-synopsis\"] .serialization a").Text()
	return utils.ArrayFilter(strings.Split(serialization, ","))
}

// getStartDate to get manga's start airing date.
func (pp *MagazineParser) getStartDate(area *goquery.Selection) string {
	airing := area.Find(".info .remain-time").Text()
	return strings.TrimSpace(airing)
}

// getMember to get manga's member number.
func (pp *MagazineParser) getMember(area *goquery.Selection) int {
	member := area.Find(".scormem span[class^=member]").Text()
	return utils.StrToNum(member)
}

// getScore to get manga's score.
func (pp *MagazineParser) getScore(area *goquery.Selection) float64 {
	score := area.Find(".scormem .score").Text()
	score = strings.TrimSpace(score)

	if score == "N/A" {
		return 0
	}

	return utils.StrToFloat(score)
}
