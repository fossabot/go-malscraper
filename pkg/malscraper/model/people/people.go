package people

// People represents the main model for MyAnimeList people information.
type People struct {
	ID               int          `json:"id"`
	Name             string       `json:"name"`
	Image            string       `json:"image"`
	GivenName        string       `json:"givenName"`
	FamilyName       string       `json:"familyName"`
	AlternativeNames []string     `json:"alternativeNames"`
	Birthday         string       `json:"birthday"`
	Website          string       `json:"website"`
	Favorite         int          `json:"favorite"`
	More             string       `json:"more"`
	VoiceActors      []VoiceActor `json:"voiceActors"`
	Staff            []Staff      `json:"staff"`
	PublishedManga   []Staff      `json:"publishedManga"`
}

// VoiceActor represents voice actor model with their anime and character role.
type VoiceActor struct {
	Anime     Anime     `json:"anime"`
	Character Character `json:"character"`
}

// Anime represents simple anime model for voice actor.
type Anime struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Image string `json:"image"`
}

// Character represents simple character model for voice actor.
type Character struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Role  string `json:"role"`
	Image string `json:"image"`
}

// Staff represents simple staff model.
type Staff struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Image string `json:"image"`
	Role  string `json:"role"`
}
