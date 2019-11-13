package search

// SearchAnimeMangaData for main anime/manga search data model.
type SearchAnimeMangaData struct {
	Image   string `json:"image"`
	Id      string `json:"id"`
	Title   string `json:"title"`
	Summary string `json:"summary"`
	Type    string `json:"type"`
	Episode string `json:"episode"`
	Volume  string `json:"volume"`
	Score   string `json:"score"`
}

// SearchCharPeopleData for main character/people search data model.
type SearchCharPeopleData struct {
	Image    string    `json:"image"`
	Id       string    `json:"id"`
	Name     string    `json:"name"`
	Nickname string    `json:"nickname"`
	Anime    []IdTitle `json:"anime"`
	Manga    []IdTitle `json:"manga"`
}

// IdTitle is common model contain id and title.
type IdTitle struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

// SearchUserData for main user search data model.
type SearchUserData struct {
	Name       string `json:"name"`
	Image      string `json:"image"`
	LastOnline string `json:"last_online"`
}
