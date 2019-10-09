package additional

// StatData for main anime & manga statistics data model.
type StatData struct {
	Summary map[string]string `json:"summary"`
	Score   []Score           `json:"score"`
	User    []User            `json:"user"`
}

// Score is a model for anime/manga score.
type Score struct {
	Type    string `json:"type"`
	Vote    string `json:"vote"`
	Percent string `json:"percent"`
}

// User is a model for user in StatData model.
type User struct {
	Image    string `json:"image"`
	Username string `json:"username"`
	Score    string `json:"score"`
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
	Episode string `json:"episode"`
	Title   string `json:"title"`
	Link    string `json:"link"`
}

// Promotion is a model for promotion video in VideoData.
type Promotion struct {
	Title string `json:"title"`
	Link  string `json:"link"`
}

// CharacterStaffData for main anime/manga additional character+staff list.
type CharacterStaffData struct {
	Character []Character `json:"character"`
	Staff     []Staff     `json:"staff"`
}

// Character is a model for character in CharacterStaffData.
type Character struct {
	Id    string  `json:"id"`
	Image string  `json:"image"`
	Name  string  `json:"name"`
	Role  string  `json:"role"`
	Va    []Staff `json:"va"`
}

// Staff is a common model with id, name, role, and image.
type Staff struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Role  string `json:"role"`
	Image string `json:"image"`
}

// EpisodeData for main anime additional episode list.
type EpisodeData struct {
	Episode       string `json:"episode"`
	Link          string `json:"link"`
	Title         string `json:"title"`
	JapaneseTitle string `json:"japanese_title"`
	Aired         string `json:"aired"`
}
