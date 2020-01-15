package search

// Manga represents the main model for MyAnimeList manga search result.
type Manga struct {
	ID        int     `json:"id"`
	Title     string  `json:"title"`
	Image     string  `json:"image"`
	Summary   string  `json:"summary"`
	Type      string  `json:"type"`
	Volume    int     `json:"volume"`
	Chapter   int     `json:"chapter"`
	Score     float64 `json:"score"`
	StartDate string  `json:"startDate"`
	EndDate   string  `json:"endDate"`
	Member    int     `json:"member"`
}
