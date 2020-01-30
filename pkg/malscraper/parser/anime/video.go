package anime

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/pkg/malscraper/config"
	"github.com/rl404/go-malscraper/pkg/malscraper/constant"
	model "github.com/rl404/go-malscraper/pkg/malscraper/model/anime"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser"
	"github.com/rl404/go-malscraper/pkg/malscraper/utils"
)

// VideoParser is parser for MyAnimeList anime video list.
// Example: https://myanimelist.net/anime/1/Cowboy_Bebop/video
type VideoParser struct {
	parser.BaseParser
	ID   int
	Page int
	Data model.Video
}

// InitVideoParser to initiate all fields and data of VideoParser.
func InitVideoParser(config config.Config, id int, page ...int) (video VideoParser, err error) {
	video.ID = id
	video.Page = 1
	video.Config = config

	if len(page) > 0 {
		video.Page = page[0]
	}

	// Checking to redis if using redis in config.
	// Redis key's pattern is `anime-video:{id},{page}`.
	redisKey := constant.RedisGetAnimeVideo + ":" + strconv.Itoa(video.ID) + "," + strconv.Itoa(video.Page)
	if config.RedisClient != nil {
		found, err := utils.UnmarshalFromRedis(config.RedisClient, redisKey, &video.Data)
		if err != nil {
			video.SetResponse(500, err.Error())
			return video, err
		}

		if found {
			video.SetResponse(200, constant.SuccessMessage)
			return video, nil
		}
	}

	// Get MyAnimeList HTML source page and initiate the parser.
	err = video.InitParser("/anime/"+strconv.Itoa(video.ID)+"/a/video?p="+strconv.Itoa(video.Page), ".js-scrollfix-bottom-rel")
	if err != nil {
		return video, err
	}

	// Fill in data field.
	video.setAllDetail()

	// Save data field to redis if using redis in config.
	if config.RedisClient != nil {
		go utils.SaveToRedis(config.RedisClient, redisKey, video.Data, config.CacheTime)
	}

	return video, nil
}

// setAllDetail to set all video detail information.
func (vp *VideoParser) setAllDetail() {
	vp.setEpisode()
	vp.setPromotion()
}

// setEpisode to set episode video.
func (vp *VideoParser) setEpisode() {
	var episodeList []model.SimpleEpisode

	episodeArea := vp.Parser.Find(".episode-video")
	episodeArea.Find(".video-list-outer").Each(func(i int, eachEpisode *goquery.Selection) {
		linkArea := eachEpisode.Find("a").First()

		link, _ := linkArea.Attr("href")

		episodeList = append(episodeList, model.SimpleEpisode{
			Title:   vp.getEpisodeTitle(linkArea),
			Episode: vp.getEpisodeNo(linkArea),
			Link:    link,
		})
	})

	vp.Data.Episodes = episodeList
}

// getEpisodeNo to get episode number.
func (vp *VideoParser) getEpisodeNo(linkArea *goquery.Selection) int {
	linkArea.Find("span.title").Find("span").Remove()
	episode := linkArea.Find("span.title").Text()
	episode = strings.Replace(episode, "Episode", "", -1)
	return utils.StrToNum(episode)
}

// getEpisodeTitle to get episode title.
func (vp *VideoParser) getEpisodeTitle(linkArea *goquery.Selection) string {
	return linkArea.Find("span.title span").Text()
}

// setPromotion to get promotion video.
func (vp *VideoParser) setPromotion() {
	var promoList []model.Promotion

	area := vp.Parser.Find(".promotional-video")
	area.Find(".video-list-outer").Each(func(i int, eachPromo *goquery.Selection) {
		linkArea := eachPromo.Find("a").First()

		promoList = append(promoList, model.Promotion{
			Title: vp.getPromoTitle(linkArea),
			Link:  vp.getPromoLink(linkArea),
		})
	})

	vp.Data.Promotions = promoList
}

// getPromoTitle to get promotion video title.
func (vp *VideoParser) getPromoTitle(linkArea *goquery.Selection) string {
	return linkArea.Find("span.title").Text()
}

// getPromoLink to get promotion video link.
func (vp *VideoParser) getPromoLink(linkArea *goquery.Selection) string {
	link, _ := linkArea.Attr("href")
	return utils.URLCleaner(link, "video", vp.Config.CleanVideoURL)
}
