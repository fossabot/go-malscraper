package anime

// Stats represents the main model for MyAnimeList anime stats information.
type Stats struct {
	Summary map[string]int `json:"summary"`
	Score   []Score        `json:"score"`
	Users   []UserStats    `json:"users"`
}

// Score represents detail score model and its count.
type Score struct {
	Type    int     `json:"type"`
	Vote    int     `json:"vote"`
	Percent float64 `json:"percent"`
}

// UserStats represents simple user's stats model.
type UserStats struct {
	Username string `json:"username"`
	Image    string `json:"image"`
	Score    int    `json:"score"`
	Status   string `json:"status"`
	Episode  string `json:"episode"`
	Date     string `json:"date"`
}
