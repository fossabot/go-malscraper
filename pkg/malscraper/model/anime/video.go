package anime

// Video represents the main model for MyAnimeList anime video information.
type Video struct {
	Episodes   []SimpleEpisode `json:"episodes"`
	Promotions []Promotion     `json:"promotions"`
}

// SimpleEpisode represents anime episode simple model.
type SimpleEpisode struct {
	Episode int    `json:"episode"`
	Title   string `json:"title"`
	Link    string `json:"link"`
}

// Promotion represents promotion video model.
type Promotion struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}
