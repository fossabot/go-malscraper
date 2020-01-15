package utils

import (
	"net/url"
	"reflect"
	"testing"
	"time"

	searchModel "github.com/rl404/go-malscraper/pkg/malscraper/model/search"
)

// TestImageURLCleaner to test cleaning image URL from MyAnimeList.
func TestImageURLCleaner(t *testing.T) {
	urls := [][]string{
		// Anime Image
		{"https://cdn.myanimelist.net/r/23x32/images/anime/8/65409.webp?s=5a37d57b31e0e3948166fcea8ca89625", "https://cdn.myanimelist.net/images/anime/8/65409.webp"},
		// Manga image
		{"https://cdn.myanimelist.net/r/80x120/images/manga/3/214566.jpg?s=48212bcd0396d503a01166149a29c67e", "https://cdn.myanimelist.net/images/manga/3/214566.jpg"},
		// User image
		{"https://cdn.myanimelist.net/r/76x120/images/userimages/6098374.jpg?s=4b8e4f091fbb3ecda6b9833efab5bd9b", "https://cdn.myanimelist.net/images/userimages/6098374.jpg"},
		// Empty user image
		{"https://cdn.myanimelist.net/r/76x120/images/questionmark_50.gif?s=8e0400788aa6af2a2f569649493e2b0f", ""},
	}

	for i, url := range urls {
		if !reflect.DeepEqual(url[1], ImageURLCleaner(url[0])) {
			t.Errorf("ImageURLCleaner(%v) failed: expected %v got %v", i, url[1], url[0])
		}
	}
}

// TestIVideoURLCleaner to test cleaning video URL from MyAnimeList.
func TestVideoURLCleaner(t *testing.T) {
	urls := [][]string{
		{"https://www.youtube.com/embed/qig4KOK2R2g?enablejsapi=1&wmode=opaque&autoplay=1", "https://www.youtube.com/watch?v=qig4KOK2R2g"},
		{"https://www.youtube.com/embed/j2hiC9BmJlQ?enablejsapi=1&wmode=opaque&autoplay=1", "https://www.youtube.com/watch?v=j2hiC9BmJlQ"},
	}

	for i, url := range urls {
		if !reflect.DeepEqual(url[1], VideoURLCleaner(url[0])) {
			t.Errorf("VideoURLCleaner(%v) failed: expected %v got %v", i, url[1], url[0])
		}
	}
}

// TestArrayFilter to test removing empty element in array string.
func TestArrayFilter(t *testing.T) {
	arrays := [][]string{
		{"1", "2", "", "3"},
		{"", "1", "", "2"},
		{""},
	}

	results := [][]string{
		{"1", "2", "3"},
		{"1", "2"},
		nil,
	}

	for i, array := range arrays {
		if !reflect.DeepEqual(results[i], ArrayFilter(array)) {
			t.Errorf("ArrayFilter(%v) failed: expected %v got %v", i, results[i], array)
		}
	}
}

// InArrayTest is a simple struct for InArray test.
type InArrayTest struct {
	Arrays []string
	Value  string
	Result bool
}

// TestInArray to test if value exist in array.
func TestInArray(t *testing.T) {
	arrays := []InArrayTest{
		{[]string{"1", "2", "3"}, "2", true},
		{[]string{"1", "2", "3"}, "4", false},
	}

	for i, array := range arrays {
		if !reflect.DeepEqual(array.Result, InArray(array.Arrays, array.Value)) {
			t.Errorf("InArray(%v) failed: expected %v got %v", i, array.Result, !array.Result)
		}
	}
}

// TestGetSeasonName to test season name.
func TestGetSeasonName(t *testing.T) {
	months := map[int]string{
		1:  "winter",
		2:  "winter",
		3:  "winter",
		4:  "spring",
		5:  "spring",
		6:  "spring",
		7:  "summer",
		8:  "summer",
		9:  "summer",
		10: "fall",
		11: "fall",
		12: "fall",
	}

	for m, s := range months {
		if !reflect.DeepEqual(s, GetSeasonName(m)) {
			t.Errorf("GetSeasonName(%v) failed: expected %v got %v", m, s, GetSeasonName(m))
		}
	}
}

// TestGetCurrentSeason to test current season.
func TestGetCurrentSeason(t *testing.T) {
	seasons := []string{"winter", "spring", "summer", "fall"}

	if !InArray(seasons, GetCurrentSeason()) {
		t.Errorf("GetCurrentSeason() failed: expected valid season name got %v", GetCurrentSeason())
	}
}

// StrToNumTest is a simple struct for StrToNum test.
type StrToNumTest struct {
	NumStr string
	NumInt int
}

// TestStrToNum to test string conversion to int.
func TestStrToNum(t *testing.T) {
	strList := []StrToNumTest{
		{"1", 1},
		{"2,234 ", 2234},
		{" 3,345,456", 3345456},
		{"-9234", -9234},
		{"asd", 0},
	}

	for i, str := range strList {
		numInt := StrToNum(str.NumStr)
		if !reflect.DeepEqual(str.NumInt, numInt) {
			t.Errorf("StrToNum(%v) failed: expected %v got %v", i, str.NumInt, numInt)
		}
	}
}

// StrToFloatTest is a simple struct for StrToFloat test.
type StrToFloatTest struct {
	NumStr   string
	NumFloat float64
}

// TestStrToFloat to test string conversion to int.
func TestStrToFloat(t *testing.T) {
	strList := []StrToFloatTest{
		{"1", 1.0},
		{"2,234.5 ", 2234.5},
		{" 3,345,456.123", 3345456.123},
		{"-9234.43", -9234.43},
		{"asd", 0},
	}

	for i, str := range strList {
		numFloat := StrToFloat(str.NumStr)
		if !reflect.DeepEqual(str.NumFloat, numFloat) {
			t.Errorf("StrToFloat(%v) failed: expected %v got %v", i, str.NumFloat, numFloat)
		}
	}
}

// ValueSplitTest is a simple struct for GetValueFromSpit test.
type ValueSplitTest struct {
	Str       string
	Separator string
	Index     int
	Result    string
}

// TestGetValueFromSplit to test value from splitted string.
func TestGetValueFromSplit(t *testing.T) {
	testList := []ValueSplitTest{
		{"https://myanimelist.net/anime/39701/Nanatsu_no_Taizai__Kamigami_no_Gekirin", "/", 3, "anime"},
		{"/anime/genre/2/Adventure", "/", 3, "2"},
		{"/anime/genre/2/Adventure", "/", 6, ""},
		{"Completed 333/333 · Score 9", " · ", 1, "Score 9"},
	}

	for i, str := range testList {
		value := GetValueFromSplit(str.Str, str.Separator, str.Index)
		if !reflect.DeepEqual(str.Result, value) {
			t.Errorf("GetValueFromSplit(%v) failed: expected %v got %v", i, str.Result, value)
		}
	}
}

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
