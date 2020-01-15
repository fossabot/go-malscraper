package review

import "github.com/rl404/go-malscraper/pkg/malscraper/model/common"

// Review represents the main model for MyAnimeList review information.
type Review struct {
	ID       int             `json:"id"`
	Username string          `json:"username"`
	Image    string          `json:"image"`
	Source   Source          `json:"source"`
	Helpful  int             `json:"helpful"`
	Date     common.DateTime `json:"date"`
	Episode  string          `json:"episode"`
	Chapter  string          `json:"chapter"`
	Score    map[string]int  `json:"score"`
	Review   string          `json:"review"`
}

// Source represents a simple review source (anime/manga) model.
type Source struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Type  string `json:"type"`
	Image string `json:"image"`
}
