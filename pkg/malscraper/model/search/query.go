package search

import "time"

// Query represents anime & manga search query model.
type Query struct {
	Query          string
	Page           int
	Type           int
	Score          int
	Status         int
	Producer       int
	Magazine       int
	Rating         int
	StartDate      time.Time
	EndDate        time.Time
	IsExcludeGenre int
	Genre          []int
	Letter         string
}
