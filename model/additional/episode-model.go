package additional

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
)

// EpisodeModel is an extended model from MainModel for anime additional episode list.
type EpisodeModel struct {
	model.MainModel
	Id   int
	Page int
	Data []EpisodeData
}

// InitEpisodeModel to initiate fields in parent (MainModel) model.
func (i *EpisodeModel) InitEpisodeModel(id int, p int) ([]EpisodeData, int, string) {
	i.Id = id
	i.Page = 100 * (p - 1)
	i.InitModel("/anime/"+strconv.Itoa(i.Id)+"/a/episode?offset="+strconv.Itoa(i.Page), ".js-scrollfix-bottom-rel")

	i.Parser = i.GetParser(i.Url)

	if i.ResponseCode != 200 {
		return i.Data, i.ResponseCode, i.ErrorMessage
	}

	i.SetAllDetail()

	return i.Data, i.ResponseCode, i.ErrorMessage
}

// SetAllDetail to fill all additional video data.
func (i *EpisodeModel) SetAllDetail() {
	var epList []EpisodeData

	area := i.Parser.Find(".js-scrollfix-bottom-rel table.episode_list").First()
	area.Find(".episode-list-data").Each(func(j int, eachEp *goquery.Selection) {
		epList = append(epList, EpisodeData{
			Episode:       i.GetEpisode(eachEp),
			Link:          i.GetLink(eachEp),
			Title:         i.GetTitle(eachEp),
			JapaneseTitle: i.GetJapaneseTitle(eachEp),
			Aired:         i.GetAired(eachEp),
		})
	})

	i.Data = epList
}

// GetEpisode to get episode number.
func (i *EpisodeModel) GetEpisode(eachEp *goquery.Selection) int {
	epInt, _ := strconv.Atoi(eachEp.Find(".episode-number").Text())
	return epInt
}

// GetLink to get episode link.
func (i *EpisodeModel) GetLink(eachEp *goquery.Selection) string {
	link, _ := eachEp.Find(".episode-video a").First().Attr("href")
	return link
}

// GetTitle to get episode title.
func (i *EpisodeModel) GetTitle(eachEp *goquery.Selection) string {
	title := eachEp.Find(".episode-title").First()
	return title.Find("a").First().Text()
}

// GetJapaneseTitle to get japanese episode title.
func (i *EpisodeModel) GetJapaneseTitle(eachEp *goquery.Selection) string {
	title := eachEp.Find(".episode-title span").First().Text()
	return strings.TrimSpace(title)
}

// GetAired to get episode aired date.
func (i *EpisodeModel) GetAired(eachEp *goquery.Selection) string {
	aired := eachEp.Find(".episode-aired").Text()
	return strings.Replace(aired, "N/A", "", -1)
}
