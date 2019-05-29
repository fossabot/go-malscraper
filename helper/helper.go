package helper

import "regexp"
import "strings"

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

func VideoUrlCleaner(str string) string {
	r, _ := regexp.Compile(`\?.+`)
	str = r.ReplaceAllString(str, "")

	return str
}

func MapMerge(map1 map[string]string, map2 map[string]string) map[string]string {
	for i := range map2 {
		map1[i] = map2[i]
	}
	return map1
}
