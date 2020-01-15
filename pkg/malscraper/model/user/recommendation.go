package user

// Recommendation represent the main model for MyAnimeList user recommendation list.
type Recommendation struct {
	Source  RecommendationSource `json:"source"`
	Date    string               `json:"date"`
	Content string               `json:"content"`
}

// RecommendationSource represents recommendation source (anime/manga) model.
type RecommendationSource struct {
	Liked       Source `json:"liked"`
	Recommended Source `json:"recommended"`
}
