package character

// Character represents the main model for MyAnimeList character information.
type Character struct {
	ID           int          `json:"id"`
	Image        string       `json:"image"`
	Nickname     string       `json:"nickname"`
	Name         string       `json:"name"`
	KanjiName    string       `json:"kanjiName"`
	Favorite     int          `json:"favorite"`
	About        string       `json:"about"`
	Animeography []Ography    `json:"animeography"`
	Mangaography []Ography    `json:"mangaography"`
	VoiceActors  []VoiceActor `json:"voiceActors"`
}

// Ography represents simple animeography and mangaopgrahy.
type Ography struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Image string `json:"image"`
	Role  string `json:"role"`
}

// VoiceActor represents simple model of character's voice actor.
type VoiceActor struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Role  string `json:"role"`
	Image string `json:"image"`
}
