package manga

// Character represents the main model for MyAnimeList manga character.
type Character struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
	Role  string `json:"role"`
}
