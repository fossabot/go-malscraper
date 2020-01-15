package search

// Anime represents the main model for MyAnimeList anime search result.
type Anime struct {
	ID        int     `json:"id"`
	Title     string  `json:"title"`
	Image     string  `json:"image"`
	Summary   string  `json:"summary"`
	Type      string  `json:"type"`
	Episode   int     `json:"episode"`
	Score     float64 `json:"score"`
	StartDate string  `json:"startDate"`
	EndDate   string  `json:"endDate"`
	Member    int     `json:"member"`
	Rated     string  `json:"rated"`
}
