package utils

import (
	"net/url"
	"testing"
	"time"

	searchModel "github.com/rl404/go-malscraper/pkg/malscraper/model/search"
)

// SearchParamTest is a simple struct for SetSearchParams test.
type SearchParamTest struct {
	URL    string
	Query  searchModel.Query
	Result string
}

// TestSetSearchParams to test anime & manga search params test.
func TestSetSearchParams(t *testing.T) {
	startDate, _ := time.Parse("2006-01-02", "2019-01-02")
	endDate, _ := time.Parse("2006-01-02", "2019-02-02")

	testList := []SearchParamTest{
		{"/anime.php", searchModel.Query{
			Query: "naruto",
			Page:  2}, "/anime.php?c%5B%5D=a&c%5B%5D=b&c%5B%5D=c&c%5B%5D=d&c%5B%5D=e&c%5B%5D=f&c%5B%5D=g&gx=0&mid=0&p=0&q=naruto&score=0&show=50&status=0&type=0",
		},
		{"/manga.php", searchModel.Query{
			Query: "naruto",
			Page:  2,
			Score: 7}, "/manga.php?c%5B%5D=a&c%5B%5D=b&c%5B%5D=c&c%5B%5D=d&c%5B%5D=e&c%5B%5D=f&c%5B%5D=g&gx=0&mid=0&p=0&q=naruto&score=7&show=50&status=0&type=0",
		},
		{"/anime.php", searchModel.Query{
			Query:     "naruto",
			Page:      2,
			StartDate: startDate,
			EndDate:   endDate}, "/anime.php?c%5B%5D=a&c%5B%5D=b&c%5B%5D=c&c%5B%5D=d&c%5B%5D=e&c%5B%5D=f&c%5B%5D=g&ed=2019&em=2&ey=2&gx=0&mid=0&p=0&q=naruto&score=0&sd=2019&show=50&sm=1&status=0&sy=2&type=0",
		},
		{"/anime.php", searchModel.Query{
			Query: "naruto",
			Genre: []int{
				1,
				4,
				5,
			}}, "/anime.php?c%5B%5D=a&c%5B%5D=b&c%5B%5D=c&c%5B%5D=d&c%5B%5D=e&c%5B%5D=f&c%5B%5D=g&genre%5B%5D=1&genre%5B%5D=4&genre%5B%5D=5&gx=0&mid=0&p=0&q=naruto&score=0&status=0&type=0",
		},
	}

	for _, param := range testList {
		u, _ := url.Parse(param.URL)
		q := SetSearchParams(u, param.Query)
		u.RawQuery = q.Encode()
		if u.String() != param.Result {
			t.Errorf("SetSearchParams() failed: expected %v got %v", param.Result, u.String())
		}
	}
}
