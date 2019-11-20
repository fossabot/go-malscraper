package list

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/model"
)

// ProducerModel is an extended model from MainModel for anime/manga producer/magazine list.
type ProducerModel struct {
	model.MainModel
	Type string
	Data []ProducerData
}

// InitProducerModel to initiate fields in parent (MainModel) model.
func (i *ProducerModel) InitProducerModel(t string) ([]ProducerData, int, string) {
	i.Type = t
	if i.Type == "anime" {
		i.InitModel("/anime/producer", ".anime-manga-search")
	} else {
		i.InitModel("/manga/magazine", ".anime-manga-search")
	}

	i.Parser = i.GetParser(i.Url)

	if i.ResponseCode != 200 {
		return i.Data, i.ResponseCode, i.ErrorMessage
	}

	i.SetAllDetail()

	return i.Data, i.ResponseCode, i.ErrorMessage
}

// SetAllDetail to fill all anime/manga producer/magazine list.
func (i *ProducerModel) SetAllDetail() {
	var producerList []ProducerData

	area := i.Parser.Find(".anime-manga-search").First()
	area.Find(".genre-list a").Each(func(j int, eachProducer *goquery.Selection) {
		producerList = append(producerList, ProducerData{
			Id:    i.GetId(eachProducer),
			Name:  i.GetName(eachProducer),
			Count: i.GetCount(eachProducer),
		})
	})

	i.Data = producerList
}

// GetId to get producer id.
func (i *ProducerModel) GetId(eachProducer *goquery.Selection) int {
	link, _ := eachProducer.Attr("href")
	splitLink := strings.Split(link, "/")
	idInt, _ := strconv.Atoi(splitLink[3])
	return idInt
}

// GetName to get producer name.
func (i *ProducerModel) GetName(eachProducer *goquery.Selection) string {
	name := eachProducer.Text()

	r, _ := regexp.Compile(`\([0-9,]+\)`)
	name = r.ReplaceAllString(name, "")

	return strings.TrimSpace(name)
}

// GetCount to get producer count.
func (i *ProducerModel) GetCount(eachProducer *goquery.Selection) int {
	count := eachProducer.Text()

	r, _ := regexp.Compile(`\([0-9,]+\)`)
	count = r.FindString(count)
	count = count[1 : len(count)-1]
	count = strings.Replace(count, ",", "", -1)
	countInt, _ := strconv.Atoi(count)
	return countInt
}
