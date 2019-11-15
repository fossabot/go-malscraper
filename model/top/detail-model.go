package top

// TopAnimeMangaData for main top anime/manga data model.
type TopAnimeMangaData struct {
	Rank      string `json:"rank"`
	Image     string `json:"image"`
	Id        string `json:"id"`
	Title     string `json:"title"`
	Type      string `json:"type"`
	Episode   string `json:"episode"`
	Volume    string `json:"volume"`
	StartDate string `json:"start_date"`
	EndDate   string `json:"end_date"`
	Member    string `json:"member"`
	Score     string `json:"score"`
}

// TopCharacterData for main top character data model.
type TopCharacterData struct {
	Rank         string    `json:"rank"`
	Id           string    `json:"id"`
	Name         string    `json:"name"`
	JapaneseName string    `json:"japanese_name"`
	Image        string    `json:"image"`
	Animeography []IdTitle `json:"animeography"`
	Mangaography []IdTitle `json:"mangaography"`
	Favorite     string    `json:"favorite"`
}

// IdTitle is common model contain id and title.
type IdTitle struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}
