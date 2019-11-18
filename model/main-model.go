package model

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// MyAnimeListUrl for MyAnimeList base url.
const MyAnimeListUrl string = "https://myanimelist.net"

// MainModel is base model for all model in go-malscraper.
type MainModel struct {
	Parser       *goquery.Document
	ParserArea   string
	Url          string
	ResponseCode int
	ErrorMessage string
}

// InitModel to initiate fields of MainModel.
func (c *MainModel) InitModel(url string, parserArea string) {
	c.Url = MyAnimeListUrl + url
	c.ParserArea = parserArea
}

// GetParser to get goquest parser to parse html.
func (c *MainModel) GetParser(url string) *goquery.Document {
	res, err := http.Get(url)
	if err != nil {
		c.SetMessage(500, err.Error())
		return nil
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		c.SetMessage(res.StatusCode, res.Status)
		return nil
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		c.SetMessage(502, err.Error())
		return nil
	}

	c.SetMessage(200, "Success")
	return doc
}

// SetMessage to set (error) message for models.
func (c *MainModel) SetMessage(responseCode int, errorMessage string) {
	c.ResponseCode = responseCode
	c.ErrorMessage = errorMessage
}
