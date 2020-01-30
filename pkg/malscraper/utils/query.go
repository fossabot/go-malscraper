package utils

import (
	"net/url"
	"strconv"

	model "github.com/rl404/go-malscraper/pkg/malscraper/model/search"
)

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
