package user

import "github.com/rl404/go-malscraper/pkg/malscraper/model/common"

// Review represents the main model for MyAnimeList user review.
type Review struct {
	ID      int             `json:"id"`
	Source  Source          `json:"source"`
	Helpful int             `json:"helpful"`
	Date    common.DateTime `json:"date"`
	Episode string          `json:"episode"`
	Chapter string          `json:"chapter"` // manga
	Score   map[string]int  `json:"score"`
	Review  string          `json:"review"`
}

// Source represents simple source (anime/manga) review model.
type Source struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Type  string `json:"type"`
	Image string `json:"image"`
}
