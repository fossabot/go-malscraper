package search

// User represents the main model for MyAnimeList user search list.
type User struct {
	Name       string `json:"name"`
	Image      string `json:"image"`
	LastOnline string `json:"lastOnline"`
}
