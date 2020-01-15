package producer

import "github.com/rl404/go-malscraper/pkg/malscraper/model/common"

// Anime represents the main model for MyAnimeList producer/studio/licensor's anime list.
type Anime struct {
	ID        int             `json:"id"`
	Image     string          `json:"image"`
	Title     string          `json:"title"`
	Genres    []Genre         `json:"genres"`
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

// Genre represents genre simple model.
type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// Producer represents producer/studio/licensor simple model with its anime count.
type Producer struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}
