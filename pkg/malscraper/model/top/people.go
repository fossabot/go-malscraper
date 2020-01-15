package top

// People represents the main model for MyAnimeList top people list.
type People struct {
	Rank         int    `json:"rank"`
	ID           int    `json:"id"`
	Name         string `json:"name"`
	JapaneseName string `json:"japaneseName"`
	Image        string `json:"image"`
	Birthday     string `json:"birthday"`
	Favorite     int    `json:"favorite"`
}
