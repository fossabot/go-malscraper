package genre

// Genre represents the main model for MyAnimeList anime & manga genres.
type Genre struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Count int    `json:"count"`
}
