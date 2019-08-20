package general

// InfoData for main anime & manga data model.
type InfoData struct {
	Id             string                   `json:"id"`
	Cover          string                   `json:"cover"`
	Title          string                   `json:"title"`
	Title2         Title2                   `json:"title2"`
	Video          string                   `json:"video"`
	Synopsis       string                   `json:"synopsis"`
	Score          string                   `json:"score"`
	Voter          string                   `json:"voter"`
	Rank           string                   `json:"rank"`
	Popularity     string                   `json:"popularity"`
	Members        string                   `json:"members"`
	Favorite       string                   `json:"favorite"`
	Type           string                   `json:"type"`
	Episodes       string                   `json:"episodes"`
	Volumes        string                   `json:"volumes"`  // manga
	Chapters       string                   `json:"chapters"` // manga
	Status         string                   `json:"status"`
	Aired          StartEndDate             `json:"aired"`
	Published      StartEndDate             `json:"published"` // manga
	Premiered      string                   `json:"premiered"`
	Broadcast      string                   `json:"broadcast"`
	Producers      []IdName                 `json:"producers"`
	Licensors      []IdName                 `json:"licensors"`
	Studios        []IdName                 `json:"studios"`
	Source         string                   `json:"source"`
	Genres         []IdName                 `json:"genres"`
	Authors        []IdName                 `json:"authors"`       // manga
	Serialization  string                   `json:"serialization"` // manga
	Duration       string                   `json:"duration"`
	Rating         string                   `json:"rating"`
	Related        map[string][]IdTitleType `json:"related"`
	Character      []Character              `json:"character"`
	Staff          []Staff                  `json:"staff"`
	Song           Song                     `json:"song"`
	Review         []Review                 `json:"review"`
	Recommendation []Recommendation         `json:"recommendation"`
}

// Title2 for altenative anime & manga title.
type Title2 struct {
	English  string `json:"english"`
	Synonym  string `json:"synonym"`
	Japanese string `json:"japanese"`
}

// StartEndDate for start and end airing date.
type StartEndDate struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

// IdName is common model for producer, licensors, studios, and genres.
type IdName struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// IdTitleType is common model for related.
type IdTitleType struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Type  string `json:"type"`
}

// Character for character data model.
type Character struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Role    string `json:"role"`
	Image   string `json:"image"`
	VaName  string `json:"va_name"`
	VaId    string `json:"va_id"`
	VaImage string `json:"va_image"`
	VaRole  string `json:"va_role"`
}

// Staff for staff data model.
type Staff struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Role  string `json:"role"`
	Image string `json:"image"`
}

// Song for opening & closing song data model.
type Song struct {
	Opening []string `json:"opening"`
	Closing []string `json:"closing"`
}

// ReviewDate for date & time in review model.
type ReviewDate struct {
	Date string `json:"date"`
	Time string `json:"time"`
}

// Review for review data model.
type Review struct {
	Id       string            `json:"id"`
	Username string            `json:"username"`
	Image    string            `json:"image"`
	Helpful  string            `json:"helpful"`
	Date     ReviewDate        `json:"date"`
	Episode  string            `json:"episode"`
	Chapter  string            `json:"chapter"`
	Score    map[string]string `json:"score"`
	Review   string            `json:"review"`
}

// recommendation for recommendation data model.
type Recommendation struct {
	Id    string `json:"id"`
	Title string `json:"title"`
	Image string `json:"image"`
	User  string `json:"user"`
}
