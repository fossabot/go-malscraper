package search

// Character represents the main model for MyAnimeList character search result.
type Character struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Nickname string `json:"nickname"`
	Image    string `json:"image"`
	Anime    []Role `json:"anime"`
	Manga    []Role `json:"manga"`
}

// Role represents the simple anime & manga role.
type Role struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}
