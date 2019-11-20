package additional

// CharacterStaffData for main anime/manga additional character+staff list.
type CharacterStaffData struct {
	Character []Character `json:"character"`
	Staff     []Staff     `json:"staff"`
}

// Character is a model for character in CharacterStaffData.
type Character struct {
	Id    int     `json:"id"`
	Image string  `json:"image"`
	Name  string  `json:"name"`
	Role  string  `json:"role"`
	Va    []Staff `json:"va"`
}

// Staff is a common model with id, name, role, and image.
type Staff struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Role  string `json:"role"`
	Image string `json:"image"`
}

// StatData for main anime & manga statistics data model.
type StatData struct {
	Summary map[string]int `json:"summary"`
	Score   []Score        `json:"score"`
	User    []User         `json:"user"`
}

// Score is a model for anime/manga score.
type Score struct {
	Type    int     `json:"type"`
	Vote    int     `json:"vote"`
	Percent float64 `json:"percent"`
}

// User is a model for user in StatData model.
type User struct {
	Username string `json:"username"`
	Image    string `json:"image"`
	Score    int    `json:"score"`
	Status   string `json:"status"`
	Episode  string `json:"episode"`
	Volume   string `json:"volume"`  // manga
	Chapter  string `json:"chapter"` // manga
	Date     string `json:"date"`
}

// VideoData for main anime additional video data model.
type VideoData struct {
	Episode   []Episode   `json:"episode"`
	Promotion []Promotion `json:"promotion"`
}

// Episode is a model for episode video in VideoData.
type Episode struct {
	Title   string `json:"title"`
	Episode string `json:"episode"`
	Link    string `json:"link"`
}

// Promotion is a model for promotion video in VideoData.
type Promotion struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}

// EpisodeData for main anime additional episode list.
type EpisodeData struct {
	Episode       int    `json:"episode"`
	Title         string `json:"title"`
	JapaneseTitle string `json:"japanese_title"`
	Aired         string `json:"aired"`
	Link          string `json:"link"`
}

// ReviewData for main anime/manga additional review list.
type ReviewData struct {
	Id       int            `json:"id"`
	Username string         `json:"username"`
	Image    string         `json:"image"`
	Helpful  int            `json:"helpful"`
	Date     DateTime       `json:"date"`
	Episode  string         `json:"episode"`
	Chapter  string         `json:"chapter"` // manga
	Score    map[string]int `json:"score"`
	Review   string         `json:"review"`
}

// DateTime is common model for date and time.
type DateTime struct {
	Date string `json:"date"`
	Time string `json:"time"`
}

// RecommendationData for main anime/manga additional recommendation list.
type RecommendationData struct {
	Id             int          `json:"id"`
	Title          string       `json:"title"`
	Image          string       `json:"image"`
	Username       string       `json:"username"`
	Recommendation string       `json:"recommendation"`
	Other          []OtherRecom `json:"other"`
}

// OtherRecom is simple modal for other recommandation.
type OtherRecom struct {
	Username       string `json:"username"`
	Recommendation string `json:"recommendation"`
}
