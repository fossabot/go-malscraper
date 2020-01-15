package manga

import (
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser"
)

// PictureParser is parser for MyAnimeList manga picture list.
// Example: https://myanimelist.net/manga/1/Monster/pics
type PictureParser struct {
	parser.BaseParser
	Id   int
	Data []string
}

// InitPictureParser to initiate all fields and data of PictureParser.
func InitPictureParser(id int) (picture PictureParser, err error) {
	picture.Id = id

	err = picture.InitParser("/manga/"+strconv.Itoa(picture.Id)+"/a/pics", ".js-scrollfix-bottom-rel")
	if err != nil {
		return picture, err
	}

	picture.setAllDetail()
	return picture, nil
}

// setAllDetail to set pictures list.
func (pp *PictureParser) setAllDetail() {
	var pictures []string

	area := pp.Parser.Find("table").First()
	area.Find("img").Each(func(i int, eachPic *goquery.Selection) {
		link, _ := eachPic.Attr("data-src")
		pictures = append(pictures, link)
	})

	pp.Data = pictures
}
