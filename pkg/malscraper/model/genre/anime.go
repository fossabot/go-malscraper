package genre

import "github.com/rl404/go-malscraper/pkg/malscraper/model/common"

// Anime represents the main model for MyAnimeList genre's anime list.
type Anime struct {
	ID        int             `json:"id"`
	Image     string          `json:"image"`
	Title     string          `json:"title"`
	Genres    []common.Genre  `json:"genres"`
	Synopsis  string          `json:"synopsis"`
	Source    string          `json:"source"`
	Producers []common.IDName `json:"producers"`
	Episode   int             `json:"episode"`
	Licensors []string        `json:"licensors"`
	Type      string          `json:"type"`
	StartDate string          `json:"startDate"`
	Member    int             `json:"member"`
	Score     float64         `json:"score"`
}
