package utils

import (
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"

	model "github.com/rl404/go-malscraper/pkg/malscraper/model/search"
)

// ImageURLCleaner to clean dirty image url.
// Example:
// https://cdn.myanimelist.net/r/80x120/images/manga/3/214566.jpg?s=48212bcd0396d503a01166149a29c67e => https://cdn.myanimelist.net/images/manga/3/214566.jpg
// https://cdn.myanimelist.net/r/76x120/images/userimages/6098374.jpg?s=4b8e4f091fbb3ecda6b9833efab5bd9b => https://cdn.myanimelist.net/images/userimages/6098374.jpg
// https://cdn.myanimelist.net/r/76x120/images/questionmark_50.gif?s=8e0400788aa6af2a2f569649493e2b0f => empty string
func ImageURLCleaner(str string) string {
	match, _ := regexp.MatchString("(questionmark)|(qm_50)", str)

	if match {
		return ""
	}

	str = strings.Replace(str, "v.jpg", ".jpg", -1)
	str = strings.Replace(str, "t.jpg", ".jpg", -1)
	str = strings.Replace(str, "_thumb.jpg", ".jpg", -1)
	str = strings.Replace(str, "userimages/thumbs", "userimages", -1)

	r := regexp.MustCompile(`r\/\d{1,3}x\d{1,3}\/`)
	str = r.ReplaceAllString(str, "")
	r = regexp.MustCompile(`\?.+`)
	str = r.ReplaceAllString(str, "")

	return str
}

// VideoURLCleaner to clean dirty video url.
// Example:
// https://www.youtube.com/embed/qig4KOK2R2g?enablejsapi=1&wmode=opaque&autoplay=1 => https://www.youtube.com/watch?v=qig4KOK2R2g
// https://www.youtube.com/embed/j2hiC9BmJlQ?enablejsapi=1&wmode=opaque&autoplay=1 => https://www.youtube.com/watch?v=j2hiC9BmJlQ
func VideoURLCleaner(str string) string {
	r := regexp.MustCompile(`\?.+`)
	str = r.ReplaceAllString(str, "")
	str = strings.Replace(str, "embed/", "watch?v=", -1)

	return str
}

// ArrayFilter to remove empty string from slice.
func ArrayFilter(s []string) []string {
	var r []string
	for _, str := range s {
		if str != "" {
			r = append(r, str)
		}
	}
	return r
}

// InArray to get if value is in array.
func InArray(arr []string, v string) bool {
	for _, a := range arr {
		if a == v {
			return true
		}
	}
	return false
}

// GetCurrentSeason to get current season (spring, summer, fall, winter).
func GetCurrentSeason() string {
	currentMonth, _ := strconv.Atoi(time.Now().Format("1"))
	return GetSeasonName(currentMonth)
}

// GetSeasonName to get season name (spring, summer, fall, winter).
func GetSeasonName(m int) string {
	switch {
	case m >= 1 && m < 4:
		return "winter"
	case m >= 4 && m < 7:
		return "spring"
	case m >= 7 && m < 10:
		return "summer"
	default:
		return "fall"
	}
}

// StrToNum to convert string number to integer including comma removal (1,234 -> 1234).
func StrToNum(strNum string) int {
	strNum = strings.TrimSpace(strNum)
	strNum = strings.Replace(strNum, ",", "", -1)
	intNum, _ := strconv.Atoi(strNum)
	return intNum
}

// StrToFloat to convert string number to float64.
func StrToFloat(strNum string) float64 {
	strNum = strings.TrimSpace(strNum)
	strNum = strings.Replace(strNum, ",", "", -1)
	floatNum, _ := strconv.ParseFloat(strNum, 64)
	return floatNum
}

// GetValueFromSplit to get value from splitted string.
func GetValueFromSplit(str string, separator string, index int) string {
	splitStr := strings.Split(str, separator)
	if len(splitStr) <= index {
		return ""
	}
	return strings.TrimSpace(splitStr[index])
}

// SetSearchParams to set search params for search anime and manga.
func SetSearchParams(u *url.URL, queryObj model.Query) url.Values {
	q := u.Query()

	defaultColums := []string{"a", "b", "c", "d", "e", "f", "g"}
	for _, c := range defaultColums {
		q.Add("c[]", c)
	}

	q.Add("q", queryObj.Query)

	if queryObj.Page > 0 {
		q.Add("show", strconv.Itoa(50*(queryObj.Page-1)))
	}

	q.Add("type", strconv.Itoa(queryObj.Type))
	q.Add("score", strconv.Itoa(queryObj.Score))
	q.Add("status", strconv.Itoa(queryObj.Status))
	q.Add("p", strconv.Itoa(queryObj.Producer))
	q.Add("mid", strconv.Itoa(queryObj.Magazine))

	if !queryObj.StartDate.IsZero() {
		sDay, sMonth, sYear := queryObj.StartDate.Date()
		q.Add("sd", strconv.Itoa(sDay))
		q.Add("sm", strconv.Itoa(int(sMonth)))
		q.Add("sy", strconv.Itoa(sYear))
	}

	if !queryObj.EndDate.IsZero() {
		eDay, eMonth, eYear := queryObj.EndDate.Date()
		q.Add("ed", strconv.Itoa(eDay))
		q.Add("em", strconv.Itoa(int(eMonth)))
		q.Add("ey", strconv.Itoa(eYear))
	}

	q.Add("gx", strconv.Itoa(queryObj.IsExcludeGenre))

	for _, g := range queryObj.Genre {
		q.Add("genre[]", strconv.Itoa(g))
	}

	return q
}
