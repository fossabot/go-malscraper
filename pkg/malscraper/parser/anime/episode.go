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

// EpisodeParser is parser for MyAnimeList anime episode list.
// Example: https://myanimelist.net/anime/1/Cowboy_Bebop/episode
type EpisodeParser struct {
	parser.BaseParser
	ID   int
	Page int
	Data []model.Episode
}

// InitEpisodeParser to initiate all fields and data of EpisodeParser.
func InitEpisodeParser(config config.Config, id int, page ...int) (episode EpisodeParser, err error) {
	episode.ID = id
	episode.Page = 1

	if len(page) > 0 {
		episode.Page = 100 * (page[0] - 1)
	}

	// Checking to redis if using redis in config.
	// Redis key's pattern is `anime-episode:{id},{page}`.
	redisKey := constant.RedisGetAnimeEpisode + ":" + strconv.Itoa(episode.ID) + "," + strconv.Itoa(episode.Page)
	if config.RedisClient != nil {
		found, err := utils.UnmarshalFromRedis(config.RedisClient, redisKey, &episode.Data)
		if err != nil {
			episode.SetResponse(500, err.Error())
			return episode, err
		}

		if found {
			episode.SetResponse(200, constant.SuccessMessage)
			return episode, nil
		}
	}

	// Get MyAnimeList HTML source page and initiate the parser.
	err = episode.InitParser("/anime/"+strconv.Itoa(episode.ID)+"/a/episode?offset="+strconv.Itoa(episode.Page), ".js-scrollfix-bottom-rel")
	if err != nil {
		return episode, err
	}

	// Fill in data field.
	episode.setAllDetail()

	// Save data field to redis if using redis in config.
	if config.RedisClient != nil {
		go utils.SaveToRedis(config.RedisClient, redisKey, episode.Data, config.CacheTime)
	}

	return episode, nil
}

// setAllDetail to set episode detail information.
func (ep *EpisodeParser) setAllDetail() {
	var epList []model.Episode

	area := ep.Parser.Find("table.episode_list").First()
	area.Find(".episode-list-data").Each(func(i int, eachEp *goquery.Selection) {
		epList = append(epList, model.Episode{
			Episode:       ep.getEpisode(eachEp),
			Link:          ep.getLink(eachEp),
			Title:         ep.getTitle(eachEp),
			JapaneseTitle: ep.getJapaneseTitle(eachEp),
			AiredDate:     ep.getAired(eachEp),
			Tag:           ep.getTag(eachEp),
		})
	})

	ep.Data = epList
}

// getEpisode to get episode number.
func (ep *EpisodeParser) getEpisode(eachEp *goquery.Selection) int {
	return utils.StrToNum(eachEp.Find(".episode-number").Text())
}

// getLink to get episode link.
func (ep *EpisodeParser) getLink(eachEp *goquery.Selection) string {
	link, _ := eachEp.Find(".episode-video a").First().Attr("href")
	return link
}

// getTitle to get episode title.
func (ep *EpisodeParser) getTitle(eachEp *goquery.Selection) string {
	title := eachEp.Find(".episode-title").First()
	return title.Find("a").First().Text()
}

// getJapaneseTitle to get japanese episode title.
func (ep *EpisodeParser) getJapaneseTitle(eachEp *goquery.Selection) string {
	title := eachEp.Find(".episode-title span:last-child").Text()
	return strings.TrimSpace(title)
}

// getAired to get episode aired date.
func (ep *EpisodeParser) getAired(eachEp *goquery.Selection) string {
	aired := eachEp.Find(".episode-aired").Text()
	return strings.Replace(aired, "N/A", "", -1)
}

// getTag to get episode tag.
func (ep *EpisodeParser) getTag(eachEp *goquery.Selection) string {
	return eachEp.Find("span.icon-episode-type-bg").Text()
}
