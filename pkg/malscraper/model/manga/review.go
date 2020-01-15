package manga

import "github.com/rl404/go-malscraper/pkg/malscraper/model/common"

// Review represents the main model for MyAnimeList manga review.
type Review struct {
	ID       int             `json:"id"`
	Username string          `json:"username"`
	Image    string          `json:"image"`
	Helpful  int             `json:"helpful"`
	Date     common.DateTime `json:"date"`
	Chapter  string          `json:"chapter"`
	Score    map[string]int  `json:"score"`
	Review   string          `json:"review"`
}
