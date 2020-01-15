package top

// Character represents the main model for MyAnimeList top character list.
type Character struct {
	Rank         int       `json:"rank"`
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	JapaneseName string    `json:"japaneseName"`
	Image        string    `json:"image"`
	Favorite     int       `json:"favorite"`
	Animeography []Ography `json:"animeography"`
	Mangaography []Ography `json:"mangaography"`
}

// Ography represents simple anime & manga ography model.
type Ography struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}
