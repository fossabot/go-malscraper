package magazine

import "github.com/rl404/go-malscraper/pkg/malscraper/model/common"

// Manga represents the main model for MyAnimeList magazine/serialization's manga list.
type Manga struct {
	ID             int             `json:"id"`
	Image          string          `json:"image"`
	Title          string          `json:"title"`
	Genres         []Genre         `json:"genres"`
	Synopsis       string          `json:"synopsis"`
	Authors        []common.IDName `json:"authors"`
	Volume         int             `json:"volume"`
	Serializations []string        `json:"serializations"`
	Type           string          `json:"type"`
	StartDate      string          `json:"startDate"`
	Member         int             `json:"member"`
	Score          float64         `json:"score"`
}

// Genre represents genre simple model.
type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// Magazine represents magazine/serialization simple model with its manga count.
type Magazine struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}
