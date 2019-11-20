package list

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
)

// GenreModel is an extended model from MainModel for anime/manga genre list.
type GenreModel struct {
	model.MainModel
	Type string
	Data []GenreData
}

// InitGenreModel to initiate fields in parent (MainModel) model.
func (i *GenreModel) InitGenreModel(t string) ([]GenreData, int, string) {
	i.Type = t
	i.InitModel("/"+i.Type+".php", ".anime-manga-search .genre-link")

	i.Parser = i.GetParser(i.Url)

	if i.ResponseCode != 200 {
		return i.Data, i.ResponseCode, i.ErrorMessage
	}

	i.SetAllDetail()

	return i.Data, i.ResponseCode, i.ErrorMessage
}

// SetAllDetail to fill all anime/manga genre list.
func (i *GenreModel) SetAllDetail() {
	var genreList []GenreData

	area := i.Parser.Find(".anime-manga-search .genre-link").First()
	area.Find(".genre-list a").Each(func(j int, eachGenre *goquery.Selection) {

		link, _ := eachGenre.Attr("href")
		splitLink := strings.Split(link, "/")
		name := strings.Replace(splitLink[4], "_", "", -1)
		idInt, _ := strconv.Atoi(splitLink[3])

		count := i.GetCount(eachGenre)

		genreList = append(genreList, GenreData{
			Id:    idInt,
			Name:  name,
			Count: count,
		})
	})

	i.Data = genreList
}

// GetCount to get genre count.
func (i *GenreModel) GetCount(eachGenre *goquery.Selection) int {
	count := eachGenre.Text()

	r, _ := regexp.Compile(`\([0-9,]+\)`)
	count = r.FindString(count)
	count = count[1 : len(count)-1]
	count = strings.Replace(count, ",", "", -1)
	countInt, _ := strconv.Atoi(count)

	return countInt
}
