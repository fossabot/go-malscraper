package anime

// Character represents the main model for MyAnimeList anime character.
type Character struct {
	ID          int          `json:"id"`
	Name        string       `json:"name"`
	Image       string       `json:"image"`
	Role        string       `json:"role"`
	VoiceActors []VoiceActor `json:"voiceActors"`
}

// VoiceActor represents simple model of character's voice actor.
type VoiceActor struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Role  string `json:"role"`
	Image string `json:"image"`
}
