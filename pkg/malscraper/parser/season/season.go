package season

import (
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/pkg/malscraper/constant"
	"github.com/rl404/go-malscraper/pkg/malscraper/model/common"
	model "github.com/rl404/go-malscraper/pkg/malscraper/model/season"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser"
	"github.com/rl404/go-malscraper/pkg/malscraper/utils"
)

// SeasonParser is parser for MyAnimeList seasonal anime list.
// Example: https://myanimelist.net/anime/season
type SeasonParser struct {
	parser.BaseParser
	Season string
	Year   int
	Data   []model.Anime
}

// InitSeasonParser to initiate all fields and data of SeasonParser.
func InitSeasonParser(params ...interface{}) (season SeasonParser, err error) {
	season.Year = time.Now().Year()
	season.Season = utils.GetCurrentSeason()

	for i, param := range params {
		switch i {
		case 0:
			if v, ok := param.(int); ok {
				season.Year = v
			}
		case 1:
			if v, ok := param.(string); ok {
				season.Season = v
			}
		}
	}

	if !utils.InArray(constant.AnimeSeasons, season.Season) {
		season.ResponseCode = 400
		return season, common.ErrInvalidSeason
	}

	err = season.InitParser("/anime/season/"+strconv.Itoa(season.Year)+"/"+season.Season, "#content")
	if err != nil {
		return season, err
	}

	season.setAllDetail()
	return season, nil
}

// setAllDetail to set all seasonal anime detail information.
func (sp *SeasonParser) setAllDetail() {
	var seasonList []model.Anime
	sp.Parser.Find("div.seasonal-anime.js-seasonal-anime").Each(func(i int, eachAnime *goquery.Selection) {
		nameArea := eachAnime.Find("div.title")
		producerArea := eachAnime.Find("div.prodsrc")
		infoArea := eachAnime.Find(".information")

		seasonList = append(seasonList, model.Anime{
			Image:     sp.getImage(eachAnime),
			ID:        sp.getID(nameArea),
			Title:     sp.getTitle(nameArea),
			Producers: sp.getProducer(producerArea),
			Episode:   sp.getEpisode(producerArea),
			Source:    sp.getSource(producerArea),
			Genres:    sp.getGenre(eachAnime),
			Synopsis:  sp.getSynopsis(eachAnime),
			Licensors: sp.getLicensor(eachAnime),
			Type:      sp.getType(infoArea),
			StartDate: sp.getStartDate(infoArea),
			Member:    sp.getMember(infoArea),
			Score:     sp.getScore(infoArea),
		})
	})

	sp.Data = seasonList
}

// getImage to get anime image.
func (sp *SeasonParser) getImage(eachAnime *goquery.Selection) string {
	image, _ := eachAnime.Find("div.image img").Attr("src")

	if image == "" {
		image, _ = eachAnime.Find("div.image img").Attr("data-src")
	}

	return utils.ImageURLCleaner(image)
}

// getID to get anime id.
func (sp *SeasonParser) getID(nameArea *goquery.Selection) int {
	id, _ := nameArea.Find("p a").Attr("href")
	id = utils.GetValueFromSplit(id, "/", 4)
	return utils.StrToNum(id)
}

// getTitle to get anime title.
func (sp *SeasonParser) getTitle(nameArea *goquery.Selection) string {
	return nameArea.Find("p a").Text()
}

// getProducer to get producer list.
func (sp *SeasonParser) getProducer(producerArea *goquery.Selection) []model.Producer {
	var producerList []model.Producer
	area := producerArea.Find("span.producer")
	area.Find("a").Each(func(i int, eachProducer *goquery.Selection) {
		producerList = append(producerList, model.Producer{
			ID:   sp.getProducerID(eachProducer),
			Name: sp.getProducerName(eachProducer),
		})
	})

	return producerList
}

// getProducerID to get producer id.
func (sp *SeasonParser) getProducerID(eachProducer *goquery.Selection) int {
	id, _ := eachProducer.Attr("href")
	id = utils.GetValueFromSplit(id, "/", 3)
	return utils.StrToNum(id)
}

// getProducerName to get producer name.
func (sp *SeasonParser) getProducerName(eachProducer *goquery.Selection) string {
	return eachProducer.Text()
}

// getEpisode to get anime episode.
func (sp *SeasonParser) getEpisode(producerArea *goquery.Selection) int {
	episode := producerArea.Find("div.eps").Text()
	episode = strings.Replace(episode, "eps", "", -1)
	episode = strings.Replace(episode, "ep", "", -1)
	episode = strings.TrimSpace(episode)

	if episode == "?" {
		return 0
	}

	return utils.StrToNum(episode)
}

// getSource to get anime source.
func (sp *SeasonParser) getSource(producerArea *goquery.Selection) string {
	return strings.TrimSpace(producerArea.Find("span.source").Text())
}

// getGenre to get anime genre list.
func (sp *SeasonParser) getGenre(eachAnime *goquery.Selection) []model.Genre {
	var genreList []model.Genre
	area := eachAnime.Find("div[class=\"genres js-genre\"]")
	area.Find("a").Each(func(i int, eachGenre *goquery.Selection) {
		id, _ := eachGenre.Attr("href")
		splitID := strings.Split(id, "/")
		genreList = append(genreList, model.Genre{
			ID:   utils.StrToNum(splitID[3]),
			Name: splitID[4],
		})

	})
	return genreList
}

// getSynopsis to get anime synopsis.
func (sp *SeasonParser) getSynopsis(eachAnime *goquery.Selection) string {
	synopsis := eachAnime.Find("div[class=\"synopsis js-synopsis\"]").Text()
	synopsis = strings.TrimSpace(synopsis)

	if synopsis == "(No synopsis yet.)" {
		synopsis = ""
	}

	return synopsis
}

// getLicensor to get anime licensor.
func (sp *SeasonParser) getLicensor(eachAnime *goquery.Selection) []string {
	licensors, _ := eachAnime.Find("div[class=\"synopsis js-synopsis\"] .licensors").Attr("data-licensors")
	licensorsList := strings.Split(licensors, ",")
	return utils.ArrayFilter(licensorsList)
}

// getType to get anime type.
func (sp *SeasonParser) getType(infoArea *goquery.Selection) string {
	t := infoArea.Find(".info").Text()
	return utils.GetValueFromSplit(t, "-", 0)
}

// getStartDate to get anime airing date.
func (sp *SeasonParser) getStartDate(infoArea *goquery.Selection) string {
	airing := infoArea.Find(".info .remain-time").Text()
	return strings.TrimSpace(airing)
}

// getMember to get anime member count.
func (sp *SeasonParser) getMember(infoArea *goquery.Selection) int {
	member := infoArea.Find(".scormem span[class^=member]").Text()
	return utils.StrToNum(member)
}

// getScore to get anime score.
func (sp *SeasonParser) getScore(infoArea *goquery.Selection) float64 {
	score := infoArea.Find(".scormem .score").Text()
	score = strings.Replace(score, "N/A", "", -1)
	return utils.StrToFloat(score)
}
