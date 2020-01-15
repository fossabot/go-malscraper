package recommendation

// Recommend represents the main model for MyAnimeList recommendation list.
type Recommend struct {
	Username string `json:"username"`
	Date     string `json:"date"`
	Source   Source `json:"source"`
	Content  string `json:"content"`
}
