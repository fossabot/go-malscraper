package additional

import (
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/helper"
	"github.com/rl404/go-malscraper/model"
)

// VideoModel is an extended model from MainModel for anime additional video list.
type VideoModel struct {
	model.MainModel
	Id   int
	Page int
	Data VideoData
}

// InitVideoModel to initiate fields in parent (MainModel) model.
func (i *VideoModel) InitVideoModel(id int, p int) (VideoData, int, string) {
	i.Id = id
	i.Page = p
	i.InitModel("/anime/"+strconv.Itoa(i.Id)+"/a/video?p="+strconv.Itoa(i.Page), ".js-scrollfix-bottom-rel")

	i.Parser = i.GetParser(i.Url)

	if i.ResponseCode != 200 {
		return i.Data, i.ResponseCode, i.ErrorMessage
	}

	i.SetAllDetail()

	return i.Data, i.ResponseCode, i.ErrorMessage
}

// SetAllDetail to fill all additional video data.
func (i *VideoModel) SetAllDetail() {
	i.SetEpisode()
	i.SetPromotion()
}

// SetEpisode to set episode video.
func (i *VideoModel) SetEpisode() {
	var episodeList []Episode

	episodeArea := i.Parser.Find(".episode-video")
	episodeArea.Find(".video-list-outer").Each(func(j int, eachEpisode *goquery.Selection) {
		linkArea := eachEpisode.Find("a").First()

		link, _ := linkArea.Attr("href")

		episodeList = append(episodeList, Episode{
			Title:   i.GetEpisodeTitle(linkArea),
			Episode: i.GetEpisodeNo(linkArea),
			Link:    link,
		})
	})

	i.Data.Episode = episodeList
}

// GetEpisodeNo to get episode number.
func (i *VideoModel) GetEpisodeNo(linkArea *goquery.Selection) string {
	linkArea.Find("span.title").Find("span").Remove()
	return linkArea.Find("span.title").Text()
}

// GetEpisodeTitle to get episode title.
func (i *VideoModel) GetEpisodeTitle(linkArea *goquery.Selection) string {
	return linkArea.Find("span.title span").Text()
}

// SetPromotion to get promotion video.
func (i *VideoModel) SetPromotion() {
	var promoList []Promotion

	area := i.Parser.Find(".promotional-video")
	area.Find(".video-list-outer").Each(func(j int, eachPromo *goquery.Selection) {
		linkArea := eachPromo.Find("a").First()

		promoList = append(promoList, Promotion{
			Title: i.GetPromoTitle(linkArea),
			Link:  i.GetPromoLink(linkArea),
		})
	})

	i.Data.Promotion = promoList
}

// GetPromoTitle to get promotion video title.
func (i *VideoModel) GetPromoTitle(linkArea *goquery.Selection) string {
	return linkArea.Find("span.title").Text()
}

// GetPromoLink to get promotion video link.
func (i *VideoModel) GetPromoLink(linkArea *goquery.Selection) string {
	link, _ := linkArea.Attr("href")
	return helper.VideoUrlCleaner(link)
}
