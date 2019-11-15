package top

// TopAnimeMangaData for main top anime/manga data model.
type TopAnimeMangaData struct {
	Rank      string `json:"rank"`
	Image     string `json:"image"`
	Id        string `json:"id"`
	Title     string `json:"title"`
	Type      string `json:"type"`
	Episode   string `json:"episode"`
	Volume    string `json:"volume"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Member    string `json:"member"`
	Score     string `json:"score"`
}
