package manga

// Recommendation represents the main model for MyAnimeList manga recommendation.
type Recommendation struct {
	ID    int       `json:"id"`
	Title string    `json:"title"`
	Image string    `json:"image"`
	Users []UserRec `json:"users"`
}

// UserRec represents user recommendation simple model.
type UserRec struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}
