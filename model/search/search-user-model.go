package search

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/helper"
	"github.com/rl404/go-malscraper/model"
)

// SearchUserModel is an extended model from MainModel for user search list.
type SearchUserModel struct {
	model.MainModel
	Type  string
	Query string
	Page  int
	Data  []SearchUserData
}

// InitSearchUserModel to initiate fields in parent (MainModel) model.
func (i *SearchUserModel) InitSearchUserModel(query string, page int) ([]SearchUserData, int, string) {
	i.Query = query
	i.Page = 24 * (page - 1)

	if len(i.Query) < 3 {
		i.SetMessage(400, "Search query needs at least 3 letters")
		return i.Data, i.ResponseCode, i.ErrorMessage
	}

	i.InitModel("/users.php?q="+i.Query+"&show="+strconv.Itoa(i.Page), "#content")

	i.Parser = i.GetParser(i.Url)

	if i.ResponseCode != 200 {
		return i.Data, i.ResponseCode, i.ErrorMessage
	}

	i.SetAllDetail()

	return i.Data, i.ResponseCode, i.ErrorMessage
}

// SetAllDetail to fill all user search list.
func (i *SearchUserModel) SetAllDetail() {
	var userList []SearchUserData
	area := i.Parser.Find("#content")
	area.Find("td.borderClass").Each(func(j int, eachUser *goquery.Selection) {
		userList = append(userList, SearchUserData{
			Name:       i.GetName(eachUser),
			Image:      i.GetImage(eachUser),
			LastOnline: i.GetLastOnline(eachUser),
		})
	})

	i.Data = userList
}

// GetName to get user name.
func (i *SearchUserModel) GetName(eachUser *goquery.Selection) string {
	return eachUser.Find("a").First().Text()
}

// GetImage to get user image.
func (i *SearchUserModel) GetImage(eachUser *goquery.Selection) string {
	image, _ := eachUser.Find("img").First().Attr("src")
	return helper.ImageUrlCleaner(image)
}

// GetLastOnline to get user last online date.
func (i *SearchUserModel) GetLastOnline(eachUser *goquery.Selection) string {
	return strings.TrimSpace(eachUser.Find("small").First().Text())
}
