package additional

import (
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
)

// CharacterPeoplePictureModel is an extended model from MainModel for character/people additional picture list.
type CharacterPeoplePictureModel struct {
	model.MainModel
	Id   int
	Type string
	Data []string
}

// InitCharacterPeoplePictureModel to initiate fields in parent (MainModel) model.
func (i *CharacterPeoplePictureModel) InitCharacterPeoplePictureModel(t string, id int) ([]string, int, string) {
	i.Id = id
	i.Type = t
	i.InitModel("/"+i.Type+"/"+strconv.Itoa(i.Id)+"/a/pictures", "#content table tr td")

	i.Parser = i.GetParser(i.Url)

	if i.ResponseCode != 200 {
		return i.Data, i.ResponseCode, i.ErrorMessage
	}

	i.SetAllDetail()

	return i.Data, i.ResponseCode, i.ErrorMessage
}

// SetAllDetail to set picture list.
func (i *CharacterPeoplePictureModel) SetAllDetail() {
	var picList []string

	area := i.Parser.Find("#content table tr td").Next().Find("table").First()
	area.Find("img").Each(func(j int, eachPic *goquery.Selection) {
		link, _ := eachPic.Attr("src")
		picList = append(picList, link)
	})

	i.Data = picList
}
