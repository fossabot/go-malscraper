package character

import (
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser"
)

// PictureParser is parser for MyAnimeList character picture list.
// Example: https://myanimelist.net/character/1/Spike_Spiegel/pictures
type PictureParser struct {
	parser.BaseParser
	Id   int
	Data []string
}

// InitPictureParser to initiate all fields and data of PictureParser.
func InitPictureParser(id int) (picture PictureParser, err error) {
	picture.Id = id

	err = picture.InitParser("/character/"+strconv.Itoa(picture.Id)+"/a/pictures", "#content table tr td")
	if err != nil {
		return picture, err
	}

	picture.setAllDetail()
	return picture, nil
}

// setAllDetail to set pictures list.
func (cp *PictureParser) setAllDetail() {
	var pictures []string

	area := cp.Parser.Next().Find("table").First()
	area.Find("img").Each(func(i int, eachPic *goquery.Selection) {
		link, _ := eachPic.Attr("data-src")
		pictures = append(pictures, link)
	})

	cp.Data = pictures
}
