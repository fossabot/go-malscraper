package genre

import "github.com/rl404/go-malscraper/pkg/malscraper/model/common"

// Manga represents the main model for MyAnimeList genre's manga list.
type Manga struct {
	ID             int             `json:"id"`
	Image          string          `json:"image"`
	Title          string          `json:"title"`
	Genres         []common.Genre  `json:"genres"`
	Synopsis       string          `json:"synopsis"`
	Authors        []common.IDName `json:"authors"`
	Volume         int             `json:"volume"`
	Serializations []string        `json:"serializations"`
	Type           string          `json:"type"`
	StartDate      string          `json:"startDate"`
	Member         int             `json:"member"`
	Score          float64         `json:"score"`
}
