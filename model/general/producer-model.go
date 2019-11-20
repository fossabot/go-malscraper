package general

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/helper"
	"github.com/rl404/go-malscraper/model"
)

// ProducerModel is an extended model from MainModel for studio/producer.
type ProducerModel struct {
	model.MainModel
	T1   string
	T2   string
	Id   int
	Page int
	Data []ProducerData
}

// InitProducerModel to initiate fields in parent (MainModel) model.
func (i *ProducerModel) InitProducerModel(t1 string, t2 string, id int, page int) ([]ProducerData, int, string) {
	i.T1 = t1
	i.T2 = t2
	i.Id = id
	i.Page = page

	if i.T2 == "producer" {
		if i.T1 == "anime" {
			i.InitModel("/anime/producer/"+strconv.Itoa(i.Id)+"/?page="+strconv.Itoa(i.Page), "#content .js-categories-seasonal")
		} else {
			i.InitModel("/manga/magazine/"+strconv.Itoa(i.Id)+"/?page="+strconv.Itoa(i.Page), "#content .js-categories-seasonal")
		}
	} else {
		i.InitModel("/"+i.T1+"/genre/"+strconv.Itoa(i.Id)+"/?page="+strconv.Itoa(i.Page), "#contentWrapper")
	}

	i.Parser = i.GetParser(i.Url)

	if i.ResponseCode != 200 {
		return i.Data, i.ResponseCode, i.ErrorMessage
	}

	i.SetAllDetail()

	return i.Data, i.ResponseCode, i.ErrorMessage
}

// SetAllDetail to fill all studio/producer/genre data.
func (i *ProducerModel) SetAllDetail() {
	var producerDataList []ProducerData

	i.Parser.Find("div[class=\"seasonal-anime js-seasonal-anime\"]").Each(func(j int, eachArea *goquery.Selection) {
		nameArea := eachArea.Find("div.title")
		producerArea := eachArea.Find("div.prodsrc")
		infoArea := eachArea.Find(".information")

		producerDataList = append(producerDataList, ProducerData{
			Id:            i.GetId(nameArea),
			Image:         i.GetImage(eachArea),
			Title:         i.GetTitle(nameArea),
			Genre:         i.GetGenre(eachArea),
			Synopsis:      i.GetSynopsis(eachArea),
			Source:        i.GetSource(producerArea),
			Producer:      i.GetProducer("anime", producerArea),
			Author:        i.GetProducer("manga", producerArea),
			Episode:       i.GetEpisode("anime", producerArea),
			Volumes:       i.GetEpisode("manga", producerArea),
			Licensor:      i.GetLicensor(eachArea),
			Serialization: i.GetSerialization(eachArea),
			Type:          i.GetType(infoArea),
			AiringStart:   i.GetAiring(infoArea),
			Member:        i.GetMember(infoArea),
			Score:         i.GetScore(infoArea),
		})
	})

	i.Data = producerDataList
}

// GetId to get id of anime/manga from the list.
func (i *ProducerModel) GetId(nameArea *goquery.Selection) int {
	id, _ := nameArea.Find("p a").Attr("href")
	splitId := strings.Split(id, "/")
	idInt, _ := strconv.Atoi(splitId[4])
	return idInt
}

// GetImage to get anime/manga image from the list.
func (i *ProducerModel) GetImage(eachArea *goquery.Selection) string {
	image, _ := eachArea.Find("div.image img").Attr("data-src")
	return helper.ImageUrlCleaner(image)
}

// GetTitle to get anime/manga title from the list.
func (i *ProducerModel) GetTitle(nameArea *goquery.Selection) string {
	return nameArea.Find("p a").Text()
}

// GetGenre to get anime/manga genre from the list.
func (i *ProducerModel) GetGenre(eachArea *goquery.Selection) []IdTypeName {
	var genreList []IdTypeName

	genreArea := eachArea.Find("div[class=\"genres js-genre\"]")
	genreArea.Find("a").Each(func(j int, eachGenre *goquery.Selection) {
		genreLink, _ := eachGenre.Attr("href")
		splitLink := strings.Split(genreLink, "/")
		idInt, _ := strconv.Atoi(splitLink[3])
		genreList = append(genreList, IdTypeName{
			Id:   idInt,
			Type: splitLink[1],
			Name: eachGenre.Text(),
		})
	})

	return genreList
}

// GetSynopsis to get anime/manga synopsis from the list.
func (i *ProducerModel) GetSynopsis(eachArea *goquery.Selection) string {
	return strings.TrimSpace(eachArea.Find("div[class=\"synopsis js-synopsis\"]").Text())
}

// GetSource to get anime/manga source from the list.
func (i *ProducerModel) GetSource(producerArea *goquery.Selection) string {
	return strings.TrimSpace(producerArea.Find("span.source").Text())
}

// GetProducer to get anime producer from the list.
func (i *ProducerModel) GetProducer(t string, producerArea *goquery.Selection) []IdName {
	var idNameList []IdName

	if t != i.T1 {
		return nil
	}

	producerArea = producerArea.Find("span.producer")
	producerArea.Find("a").Each(func(j int, eachProducer *goquery.Selection) {
		idNameList = append(idNameList, IdName{
			Id:   i.GetProducerId(eachProducer),
			Name: i.GetProducerName(eachProducer),
		})
	})

	return idNameList
}

// GetProducerId to get producer id.
func (i *ProducerModel) GetProducerId(eachProducer *goquery.Selection) int {
	link, _ := eachProducer.Attr("href")
	splitLink := strings.Split(link, "/")

	if i.T1 == "anime" {
		idInt, _ := strconv.Atoi(splitLink[3])
		return idInt
	}

	idInt, _ := strconv.Atoi(splitLink[4])
	return idInt
}

// GetProducerName to get producer name.
func (i *ProducerModel) GetProducerName(eachProducer *goquery.Selection) string {
	return eachProducer.Text()
}

// GetEpisode to get anime/manga episode/volume.
func (i *ProducerModel) GetEpisode(t string, producerArea *goquery.Selection) int {
	if i.T1 != t {
		return 0
	}

	episodeArea := producerArea.Find("div.eps").Text()
	episodeArea = strings.Replace(episodeArea, "eps", "", -1)
	episodeArea = strings.Replace(episodeArea, "ep", "", -1)
	episodeArea = strings.Replace(episodeArea, "vols", "", -1)
	episodeArea = strings.Replace(episodeArea, "vol", "", -1)
	episodeArea = strings.TrimSpace(episodeArea)

	if episodeArea == "?" {
		return 0
	}

	epInt, _ := strconv.Atoi(episodeArea)
	return epInt
}

// GetLicensor to get anime licensor.
func (i *ProducerModel) GetLicensor(eachArea *goquery.Selection) []string {
	if i.T1 != "anime" {
		return nil
	}

	licensor, _ := eachArea.Find("div[class=\"synopsis js-synopsis\"] .licensors").Attr("data-licensors")
	return helper.ArrayFilter(strings.Split(licensor, ","))
}

// GetSerialization to get manga serialization.
func (i *ProducerModel) GetSerialization(eachArea *goquery.Selection) string {
	if i.T1 != "manga" {
		return ""
	}

	return eachArea.Find("div[class=\"synopsis js-synopsis\"] .serialization a").Text()
}

// GetType to get anime type.
func (i *ProducerModel) GetType(infoArea *goquery.Selection) string {
	if i.T1 != "anime" {
		return ""
	}

	typeArea := infoArea.Find(".info").Text()
	splitType := strings.Split(typeArea, "-")
	return strings.TrimSpace(splitType[0])
}

// GetAiring to get anime/manga start airing/publishing date.
func (i *ProducerModel) GetAiring(infoArea *goquery.Selection) string {
	airing := infoArea.Find(".info .remain-time").Text()
	return strings.TrimSpace(airing)
}

// GetMember to get anime/manga member number.
func (i *ProducerModel) GetMember(infoArea *goquery.Selection) int {
	member := infoArea.Find(".scormem span[class^=member]").Text()
	member = strings.Replace(member, ",", "", -1)
	memberInt, _ := strconv.Atoi(strings.TrimSpace(member))
	return memberInt
}

// GetScore to get anime/manga score.
func (i *ProducerModel) GetScore(infoArea *goquery.Selection) float64 {
	score := infoArea.Find(".scormem .score").Text()
	score = strings.TrimSpace(score)

	if score == "N/A" {
		return 0
	}

	scoreFloat, _ := strconv.ParseFloat(strings.TrimSpace(score), 64)
	return scoreFloat
}
