package parser

import (
	"errors"
	"reflect"
	"testing"
	"time"
)

// ParserTest is a simple struct for InitParser test.
type ParserTest struct {
	URL     string
	Area    string
	IsError bool
}

// TestInitParser to test initiating parser.
func TestInitParser(t *testing.T) {
	testList := []ParserTest{
		{URL: "/anime/1", Area: "#content", IsError: false},
		{URL: "/anime/1", Area: "#random", IsError: true},
		{URL: "/random", Area: "#content", IsError: true},
		{URL: "/random random", Area: "#content", IsError: true},
	}

	for _, p := range testList {
		var baseParser BaseParser
		baseParser.InitParser(p.URL, p.Area)
		if !p.IsError && baseParser.ResponseCode != 200 {
			t.Errorf("InitParser(\"%v\", \"%v\") failed: expected to return 200", p.URL, p.Area)
		}
		time.Sleep(1 * time.Second)
	}
}

// TestSetResponse to test set reponse code and message.
func TestSetResponse(t *testing.T) {
	var baseParser BaseParser
	baseParser.SetResponse(200, "success")

	if !reflect.DeepEqual(baseParser.ResponseCode, 200) {
		t.Errorf("Expected response code %v got %v", 200, baseParser.ResponseCode)
	}

	if !reflect.DeepEqual(baseParser.ResponseMessage, errors.New("success")) {
		t.Errorf("Expected response message %v got %v", errors.New("success"), baseParser.ResponseMessage)
	}
}
