package people

import (
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser"
)

// PeoplePictureParser is parser for MyAnimeList people picture list.
// Example: https://myanimelist.net/people/1/Tomokazu_Seki/pictures
type PeoplePictureParser struct {
	parser.BaseParser
	Id   int
	Data []string
}

// InitPeoplePictureParser to initiate all fields and data of PeoplePictureParser.
func InitPeoplePictureParser(id int) (picture PeoplePictureParser, err error) {
	picture.Id = id

	err = picture.InitParser("/people/"+strconv.Itoa(picture.Id)+"/a/pictures", "#content table tr td")
	if err != nil {
		return picture, err
	}

	picture.setAllDetail()
	return picture, nil
}

// setAllDetail to set pictures list.
func (cp *PeoplePictureParser) setAllDetail() {
	var pictures []string

	area := cp.Parser.Next().Find("table").First()
	area.Find("img").Each(func(i int, eachPic *goquery.Selection) {
		link, _ := eachPic.Attr("data-src")
		pictures = append(pictures, link)
	})

	cp.Data = pictures
}
