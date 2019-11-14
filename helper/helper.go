package helper

import (
	"regexp"
	"strconv"
	"strings"
	"time"
)

// ImageUrlCleaner to clean dirty image url.
func ImageUrlCleaner(str string) string {
	match, _ := regexp.MatchString("(questionmark)|(qm_50)", str)

	if match {
		return ""
	}

	str = strings.Replace(str, "v.jpg", ".jpg", -1)
	str = strings.Replace(str, "t.jpg", ".jpg", -1)
	str = strings.Replace(str, "_thumb.jpg", ".jpg", -1)
	str = strings.Replace(str, "userimages/thumbs", "userimages", -1)

	r, _ := regexp.Compile(`r\/\d{1,3}x\d{1,3}\/`)
	str = r.ReplaceAllString(str, "")
	r2, _ := regexp.Compile(`\?.+`)
	str = r2.ReplaceAllString(str, "")

	return str
}

// VideoUrlCleaner to clean dirty video url.
func VideoUrlCleaner(str string) string {
	r, _ := regexp.Compile(`\?.+`)
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

// GetCurrentSeason to get current season (spring, summer, fall, winter)
func GetCurrentSeason() string {
	currentMonth, _ := strconv.Atoi(time.Now().Format("1"))
	switch {
	case currentMonth >= 1 && currentMonth < 4:
		return "winter"
	case currentMonth >= 4 && currentMonth < 7:
		return "spring"
	case currentMonth >= 7 && currentMonth < 10:
		return "summer"
	default:
		return "fall"
	}
}
