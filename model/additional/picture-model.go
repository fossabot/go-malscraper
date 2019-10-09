package additional

import (
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
)

// PictureModel is an extended model from MainModel for anime/manga additional picture list.
type PictureModel struct {
	model.MainModel
	Id   int
	Type string
	Data []string
}

// InitPictureModel to initiate fields in parent (MainModel) model.
func (i *PictureModel) InitPictureModel(t string, id int) ([]string, int, string) {
	i.Id = id
	i.Type = t
	i.InitModel("/"+i.Type+"/"+strconv.Itoa(i.Id)+"/a/pics", ".js-scrollfix-bottom-rel")

	i.Parser = i.GetParser(i.Url)

	if i.ResponseCode != 200 {
		return i.Data, i.ResponseCode, i.ErrorMessage
	}

	i.SetAllDetail()

	return i.Data, i.ResponseCode, i.ErrorMessage
}

// SetAllDetail to set picture list.
func (i *PictureModel) SetAllDetail() {
	var picList []string

	area := i.Parser.Find(".js-scrollfix-bottom-rel table").First()
	area.Find("img").Each(func(j int, eachPic *goquery.Selection) {
		link, _ := eachPic.Attr("src")
		picList = append(picList, link)
	})

	i.Data = picList
}
