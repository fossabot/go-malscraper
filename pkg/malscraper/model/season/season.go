package season

// Anime represents the main model for MyAnimeList seasonal anime list.
type Anime struct {
	ID        int        `json:"id"`
	Image     string     `json:"image"`
	Title     string     `json:"title"`
	Genres    []Genre    `json:"genres"`
	Synopsis  string     `json:"synopsis"`
	Source    string     `json:"source"`
	Producers []Producer `json:"producers"`
	Episode   int        `json:"episode"`
	Licensors []string   `json:"licensors"`
	Type      string     `json:"type"`
	StartDate string     `json:"startDate"`
	Member    int        `json:"member"`
	Score     float64    `json:"score"`
}

// Genre represents genre simple model.
type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Producer represents producer simple model.
type Producer struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
