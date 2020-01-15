package recommendation

// Recommendation represents the main model for MyAnimeList recommendation.
type Recommendation struct {
	Source Source `json:"source"`
	Users  []User `json:"users"`
}

// Source represents recommendation source (anime/manga) model.
type Source struct {
	Liked       SourceDetail `json:"liked"`
	Recommended SourceDetail `json:"recommended"`
}

// SourceDetail represents anime/manga source detail model.
type SourceDetail struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Type  string `json:"type"`
	Image string `json:"image"`
}

// User represents simple user model with their recommendation content.
type User struct {
	Username string `json:"username"`
	Content  string `json:"content"`
}
