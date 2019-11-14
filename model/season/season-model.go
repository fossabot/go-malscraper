package season

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/helper"
	"github.com/rl404/go-malscraper/model"
)

var SeasonList = []string{"winter", "spring", "summer", "fall"}

// SeasonModel is an extended model from MainModel for anime season list.
type SeasonModel struct {
	model.MainModel
	Season string
	Year   int
	Data   []SeasonData
}

// InitSeasonModel to initiate fields in parent (MainModel) model.
func (i *SeasonModel) InitSeasonModel(year int, season string) ([]SeasonData, int, string) {
	i.Season = season
	i.Year = year

	if helper.InArray(SeasonList, i.Season) == false && i.Season != "" {
		i.SetMessage(400, "Invalid season")
		return i.Data, i.ResponseCode, i.ErrorMessage
	}

	i.InitModel("/anime/season/"+strconv.Itoa(i.Year)+"/"+i.Season, "#content")

	i.Parser = i.GetParser(i.Url)

	if i.ResponseCode != 200 {
		return i.Data, i.ResponseCode, i.ErrorMessage
	}

	i.SetAllDetail()

	return i.Data, i.ResponseCode, i.ErrorMessage
}

// SetAllDetail to fill all anime season detail.
func (i *SeasonModel) SetAllDetail() {
	var seasonList []SeasonData
	area := i.Parser.Find("#content")
	area.Find("div.seasonal-anime.js-seasonal-anime").Each(func(j int, eachAnime *goquery.Selection) {
		nameArea := eachAnime.Find("div.title")
		producerArea := eachAnime.Find("div.prodsrc")
		infoArea := eachAnime.Find(".information")

		seasonList = append(seasonList, SeasonData{
			Image:       i.GetImage(eachAnime),
			Id:          i.GetId(nameArea),
			Title:       i.GetTitle(nameArea),
			Producer:    i.GetProducer(producerArea),
			Episode:     i.GetEpisode(producerArea),
			Source:      i.GetSource(producerArea),
			Genre:       i.GetGenre(eachAnime),
			Synopsis:    i.GetSynopsis(eachAnime),
			Licensor:    i.GetLicensor(eachAnime),
			Type:        i.GetType(infoArea),
			AiringStart: i.GetAiring(infoArea),
			Member:      i.GetMember(infoArea),
			Score:       i.GetScore(infoArea),
		})
	})

	i.Data = seasonList
}

// GetImage to get anime image.
func (i *SeasonModel) GetImage(eachAnime *goquery.Selection) string {
	image, _ := eachAnime.Find("div.image img").Attr("src")

	if image == "" {
		image, _ = eachAnime.Find("div.image img").Attr("data-src")
	}

	return helper.ImageUrlCleaner(image)
}

// GetId to get anime id.
func (i *SeasonModel) GetId(nameArea *goquery.Selection) string {
	id, _ := nameArea.Find("p a").Attr("href")
	splitId := strings.Split(id, "/")

	return splitId[4]
}

// GetTitle to get anime title.
func (i *SeasonModel) GetTitle(nameArea *goquery.Selection) string {
	return nameArea.Find("p a").Text()
}

// GetProducer to get producer list.
func (i *SeasonModel) GetProducer(producerArea *goquery.Selection) []IdName {
	var producerList []IdName
	area := producerArea.Find("span.producer")
	area.Find("a").Each(func(j int, eachProducer *goquery.Selection) {
		producerList = append(producerList, IdName{
			Id:   i.GetProducerId(eachProducer),
			Name: i.GetProducerName(eachProducer),
		})
	})

	return producerList
}

// GetProducerId to get producer id.
func (i *SeasonModel) GetProducerId(eachProducer *goquery.Selection) string {
	id, _ := eachProducer.Attr("href")
	splitId := strings.Split(id, "/")

	return splitId[3]
}

// GetProducerName to get producer name.
func (i *SeasonModel) GetProducerName(eachProducer *goquery.Selection) string {
	return eachProducer.Text()
}

// GetEpisode to get anime episode.
func (i *SeasonModel) GetEpisode(producerArea *goquery.Selection) string {
	episode := producerArea.Find("div.eps").Text()
	episode = strings.Replace(episode, "eps", "", -1)
	episode = strings.Replace(episode, "ep", "", -1)
	episode = strings.TrimSpace(episode)

	if episode == "?" {
		return ""
	}

	return episode
}

// GetSource to get anime source.
func (i *SeasonModel) GetSource(producerArea *goquery.Selection) string {
	return strings.TrimSpace(producerArea.Find("span.source").Text())
}

// GetGenre to get anime genre list.
func (i *SeasonModel) GetGenre(eachAnime *goquery.Selection) []IdName {
	var genreList []IdName
	area := eachAnime.Find("div[class=\"genres js-genre\"]")
	area.Find("a").Each(func(j int, eachGenre *goquery.Selection) {
		id, _ := eachGenre.Attr("href")
		splitId := strings.Split(id, "/")

		genreList = append(genreList, IdName{
			Id:   splitId[3],
			Name: splitId[4],
		})

	})
	return genreList
}

// GetSynopsis to get anime synopsis.
func (i *SeasonModel) GetSynopsis(eachAnime *goquery.Selection) string {
	synopsis := eachAnime.Find("div[class=\"synopsis js-synopsis\"]").Text()
	synopsis = strings.TrimSpace(synopsis)

	if synopsis == "(No synopsis yet.)" {
		synopsis = ""
	}

	return synopsis
}

// GetLicensor to get anime licensor.
func (i *SeasonModel) GetLicensor(eachAnime *goquery.Selection) []string {
	licensors, _ := eachAnime.Find("div[class=\"synopsis js-synopsis\"] .licensors").Attr("data-licensors")
	licensorsList := strings.Split(licensors, ",")

	return helper.ArrayFilter(licensorsList)
}

// GetType to get anime type.
func (i *SeasonModel) GetType(infoArea *goquery.Selection) string {
	t := infoArea.Find(".info").Text()
	splitType := strings.Split(t, "-")

	return strings.TrimSpace(splitType[0])
}

// GetAiring to get anime airing date.
func (i *SeasonModel) GetAiring(infoArea *goquery.Selection) string {
	airing := infoArea.Find(".info .remain-time").Text()
	return strings.TrimSpace(airing)
}

// GetMember to get anime member count.
func (i *SeasonModel) GetMember(infoArea *goquery.Selection) string {
	member := infoArea.Find(".scormem span[class^=member]").Text()
	member = strings.Replace(member, ",", "", -1)

	return strings.TrimSpace(member)
}

// GetScore to get anime score.
func (i *SeasonModel) GetScore(infoArea *goquery.Selection) string {
	score := infoArea.Find(".scormem .score").Text()
	score = strings.Replace(score, "N/A", "", -1)

	return strings.TrimSpace(score)
}
