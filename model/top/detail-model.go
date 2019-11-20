package top

// TopAnimeMangaData for main top anime/manga data model.
type TopAnimeMangaData struct {
	Rank      int     `json:"rank"`
	Image     string  `json:"image"`
	Id        int     `json:"id"`
	Title     string  `json:"title"`
	Type      string  `json:"type"`
	Episode   int     `json:"episode"`
	Volume    int     `json:"volume"` // manga
	StartDate string  `json:"start_date"`
	EndDate   string  `json:"end_date"`
	Member    int     `json:"member"`
	Score     float64 `json:"score"`
}

// TopCharacterData for main top character data model.
type TopCharacterData struct {
	Rank         int       `json:"rank"`
	Id           int       `json:"id"`
	Name         string    `json:"name"`
	JapaneseName string    `json:"japanese_name"`
	Image        string    `json:"image"`
	Favorite     int       `json:"favorite"`
	Animeography []IdTitle `json:"animeography"`
	Mangaography []IdTitle `json:"mangaography"`
}

// IdTitle is common model contain id and title.
type IdTitle struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
}

// TopPeopleData for main top people data model.
type TopPeopleData struct {
	Rank         int    `json:"rank"`
	Id           int    `json:"id"`
	Name         string `json:"name"`
	JapaneseName string `json:"japanese_name"`
	Image        string `json:"image"`
	Birthday     string `json:"birthday"`
	Favorite     int    `json:"favorite"`
}
