package search

// People represents the main model for MyAnimeList people search result.
type People struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Image    string `json:"image"`
}
