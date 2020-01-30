package utils

import (
	"regexp"
	"strconv"
	"strings"
	"time"
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

// URLCleaner is wrapper for image and video url cleaner for easier call.
func URLCleaner(str string, URLType string, isNeeded ...bool) string {
	if len(isNeeded) > 0 {
		if !isNeeded[0] {
			return str
		}
	}

	switch strings.ToLower(URLType) {
	case "image", "images", "img":
		return ImageURLCleaner(str)
	case "video", "videos":
		return VideoURLCleaner(str)
	default:
		return str
	}
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
