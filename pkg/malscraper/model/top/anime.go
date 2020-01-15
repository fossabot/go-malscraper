package top

// Anime represents the main model for MyAnimeList top anime list.
type Anime struct {
	Rank      int     `json:"rank"`
	Title     string  `json:"title"`
	Image     string  `json:"image"`
	ID        int     `json:"id"`
	Type      string  `json:"type"`
	Episode   int     `json:"episode"`
	StartDate string  `json:"startDate"`
	EndDate   string  `json:"endDate"`
	Member    int     `json:"member"`
	Score     float64 `json:"score"`
}
