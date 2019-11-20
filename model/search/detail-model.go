package search

// SearchAnimeMangaData for main anime/manga search data model.
type SearchAnimeMangaData struct {
	Id      int     `json:"id"`
	Title   string  `json:"title"`
	Image   string  `json:"image"`
	Type    string  `json:"type"`
	Episode int     `json:"episode"`
	Volume  int     `json:"volume"`
	Score   float64 `json:"score"`
	Summary string  `json:"summary"`
}

// SearchCharPeopleData for main character/people search data model.
type SearchCharPeopleData struct {
	Id       int       `json:"id"`
	Name     string    `json:"name"`
	Nickname string    `json:"nickname"`
	Image    string    `json:"image"`
	Anime    []IdTitle `json:"anime"`
	Manga    []IdTitle `json:"manga"`
}

// IdTitle is common model contain id and title.
type IdTitle struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

// SearchUserData for main user search data model.
type SearchUserData struct {
	Name       string `json:"name"`
	Image      string `json:"image"`
	LastOnline string `json:"last_online"`
}
