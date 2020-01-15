package anime

import "github.com/rl404/go-malscraper/pkg/malscraper/model/common"

// Review represents the main model for MyAnimeList anime review.
type Review struct {
	ID       int             `json:"id"`
	Username string          `json:"username"`
	Image    string          `json:"image"`
	Helpful  int             `json:"helpful"`
	Date     common.DateTime `json:"date"`
	Episode  string          `json:"episode"`
	Score    map[string]int  `json:"score"`
	Review   string          `json:"review"`
}
