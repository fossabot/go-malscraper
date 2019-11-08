package user

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/rl404/go-malscraper/helper"
	"github.com/rl404/go-malscraper/model"
)

// UserListModel is an extended model from MainModel for user anime/manga list.
type UserListModel struct {
	model.MainModel
	User   string
	Type   string
	Status int
	Data   []UserListData
}

// InitUserListModel to initiate fields in parent (MainModel) model.
func (u *UserListModel) InitUserListModel(user string, typ string, status int) ([]UserListData, int, string) {
	u.User = user
	u.Type = typ
	u.Status = status

	u.InitModel("", "")

	u.SetAllDetail()

	return u.Data, u.ResponseCode, u.ErrorMessage
}

// SetAllDetail to get anime/manga detail data.
func (u *UserListModel) SetAllDetail() {
	var dataList []UserListData
	offset := 0

	client := &http.Client{}

	for true {
		url := fmt.Sprintf("%v/%vlist/%v/load.json?offset=%v&status=%v", u.Url, u.Type, u.User, offset, strconv.Itoa(u.Status))

		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			u.SetMessage(500, err.Error())
			return
		}

		resp, err := client.Do(req)
		if err != nil {
			u.SetMessage(resp.StatusCode, err.Error())
			return
		}

		body, _ := ioutil.ReadAll(resp.Body)

		var dataListTmp []UserListData
		err = json.Unmarshal(body, &dataListTmp)
		if err != nil {
			u.SetMessage(500, err.Error())
			return
		}

		resp.Body.Close()

		if len(dataListTmp) > 0 {
			for i, d := range dataListTmp {
				dataListTmp[i].AnimeImage = helper.ImageUrlCleaner(d.AnimeImage)
				dataListTmp[i].MangaImage = helper.ImageUrlCleaner(d.MangaImage)
			}

			dataList = append(dataList, dataListTmp...)

			offset += 300
		} else {
			break
		}
	}

	u.SetMessage(200, "Success")
	u.Data = dataList
}
