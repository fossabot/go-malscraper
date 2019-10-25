package list

// GenreData for main anime & manga genre data model.
type GenreData struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Count string `json:"count"`
}

// ProducerData for main anime & manga producer/magazine data model.
type ProducerData struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Count string `json:"count"`
}

//
type ReviewData struct {
	Id       string            `json:"id"`
	Source   ReviewSource      `json:"source"`
	Username string            `json:"username"`
	Image    string            `json:"image"`
	Helpful  string            `json:"helpful"`
	Date     DateTime          `json:"date"`
	Episode  string            `json:"episode"`
	Chapter  string            `json:"chapter"` // manga
	Score    map[string]string `json:"score"`
	Review   string            `json:"review"`
}

// ReviewSource is source model for ReviewData.
type ReviewSource struct {
	Id    string `json:"id"`
	Type  string `json:"type"`
	Title string `json:"title"`
	Image string `json:"image"`
}

// DateTime is common model for date and time.
type DateTime struct {
	Date string `json:"date"`
	Time string `json:"time"`
}
