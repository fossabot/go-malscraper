package user

// User represents the main model for MyAnimeList user profile.
type User struct {
	Username       string     `json:"username"`
	Image          string     `json:"image"`
	LastOnline     string     `json:"lastOnline"`
	Gender         string     `json:"gender"`
	Birthday       string     `json:"birthday"`
	Location       string     `json:"location"`
	JoinedDate     string     `json:"joinedDate"`
	ForumPost      int        `json:"forumPost"`
	Review         int        `json:"review"`
	Recommendation int        `json:"recommendation"`
	BlogPost       int        `json:"blogPost"`
	Club           int        `json:"club"`
	Sns            []string   `json:"sns"`
	Friend         UserFriend `json:"friend"`
	About          string     `json:"about"`
	AnimeStats     AnimeStats `json:"animeStats"`
	MangaStats     MangaStats `json:"mangaStats"`
	Favorite       Favorite   `json:"favorite"`
}

// UserFriend represents a simple user's friends list.
type UserFriend struct {
	Count   int            `json:"count"`
	Friends []SimpleFriend `json:"friends"`
}

// SimpleFriend represents a simple friend model.
type SimpleFriend struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

// AnimeStats represents detail anime stats model.
type AnimeStats struct {
	Days      float64         `json:"days"`
	MeanScore float64         `json:"meanScore"`
	Status    AnimeStatsCount `json:"status"`
	History   []History       `json:"history"`
}

// MangaStats represents detail manga stats model.
type MangaStats struct {
	Days      float64         `json:"days"`
	MeanScore float64         `json:"meanScore"`
	Status    MangaStatsCount `json:"status"`
	History   []History       `json:"history"`
}

// AnimeStatsCount represents anime status and its count.
type AnimeStatsCount struct {
	Watching    int `json:"watching"`
	Completed   int `json:"completed"`
	OnHold      int `json:"onHold"`
	Dropped     int `json:"dropped"`
	PlanToWatch int `json:"planToWatch"`
	Total       int `json:"total"`
	Rewatched   int `json:"rewatched"`
	Episode     int `json:"episode"`
}

// MangaStatsCount represents manga status and its count.
type MangaStatsCount struct {
	Reading    int `json:"reading"`
	Completed  int `json:"completed"`
	OnHold     int `json:"onHold"`
	Dropped    int `json:"dropped"`
	PlanToRead int `json:"planToRead"`
	Total      int `json:"total"`
	Reread     int `json:"reread"`
	Chapter    int `json:"chapter"`
	Volume     int `json:"volume"`
}

// History represents user's anime/manga watched/read history.
type History struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Image    string `json:"image"`
	Progress string `json:"progress"`
	Score    int    `json:"score"`
	Status   string `json:"status"`
	Date     string `json:"date"`
}

// Favorite represents user's favorite anime, manga, character, and people.
type Favorite struct {
	Anime     []FavAnimeManga `json:"anime"`
	Manga     []FavAnimeManga `json:"manga"`
	Character []FavCharacter  `json:"character"`
	People    []FavPeople     `json:"people"`
}

// FavAnimeManga represents user's favorite anime and manga detail information.
type FavAnimeManga struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Image string `json:"image"`
	Type  string `json:"type"`
	Year  int    `json:"year"`
}

// FavCharacter represents user's favorite character detail information.
type FavCharacter struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Image       string `json:"image"`
	SourceID    int    `json:"sourceId"`
	SourceTitle string `json:"sourceTitle"`
	SourceType  string `json:"sourceType"`
}

// FavPeople represents user's favorite people detail information.
type FavPeople struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}
