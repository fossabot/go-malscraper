package top

// Manga represents the main model for MyAnimeList top manga list.
type Manga struct {
	Rank      int     `json:"rank"`
	Title     string  `json:"title"`
	Image     string  `json:"image"`
	ID        int     `json:"id"`
	Type      string  `json:"type"`
	Volume    int     `json:"volume"`
	StartDate string  `json:"startDate"`
	EndDate   string  `json:"endDate"`
	Member    int     `json:"member"`
	Score     float64 `json:"score"`
}
