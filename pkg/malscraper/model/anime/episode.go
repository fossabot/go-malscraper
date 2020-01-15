package anime

// Episode represents the main model for MyAnimeList anime episode information.
type Episode struct {
	Episode       int    `json:"episode"`
	Title         string `json:"title"`
	JapaneseTitle string `json:"japaneseTitle"`
	AiredDate     string `json:"airedDate"`
	Link          string `json:"link"`
	Tag           string `json:"tag"`
}
