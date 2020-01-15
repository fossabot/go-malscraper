package parser

import (
	"errors"
	"net/http"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/pkg/malscraper/constant"
)

// BaseParser is the base parser for all parsers in go-malscarper.
type BaseParser struct {
	Parser          *goquery.Selection
	ParseArea       string
	URL             string
	ResponseCode    int
	ResponseMessage error
}

// InitParser to initiate base fields.
func (base *BaseParser) InitParser(url string, parseArea string) error {
	base.URL = constant.MyAnimeListURL + url
	base.ParseArea = parseArea
	base.Parser, base.ResponseCode, base.ResponseMessage = base.getParser()

	if base.ResponseCode != http.StatusOK {
		return base.ResponseMessage
	}
	return nil
}

// getParser to get parser of requested MyAnimeList URL.
func (base *BaseParser) getParser() (parser *goquery.Selection, responseCode int, err error) {
	resp, err := http.Get(base.URL)
	if err != nil {
		return nil, 500, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, resp.StatusCode, errors.New(resp.Status)
	}

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, 500, err
	}

	return doc.Find(base.ParseArea).First(), 200, errors.New("success")
}
