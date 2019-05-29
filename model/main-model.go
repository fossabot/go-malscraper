package model

import (
  "net/http"

  "github.com/PuerkitoBio/goquery"
)

const MyAnimeListUrl string = "https://myanimelist.net"

type MainModel struct {
	MyAnimeListUrl 	string
	Parser 			*goquery.Document
	ParserArea 		string
	Url 			string
	ResponseCode 	int
	ErrorMessage 	string
}

func (c *MainModel) InitModel() {
	c.MyAnimeListUrl = MyAnimeListUrl
	c.Url = c.MyAnimeListUrl + c.Url
	c.Parser = c.GetParser(c.Url)
}

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

func (c *MainModel) SetMessage(responseCode int, errorMessage string) {
	c.ResponseCode = responseCode
	c.ErrorMessage = errorMessage
}