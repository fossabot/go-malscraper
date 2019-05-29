package user

// ----- MAIN USER MODEL -----
type UserData struct {
	About          string    `json:"about"`
	AnimeStat      AnimeStat `json:"anime_stat"`
	Birthday       string    `json:"birthday"`
	BlogPost       string    `json:"blog_post"`
	Club           string    `json:"club"`
	Favorite       Favorite  `json:"favorite"`
	ForumPost      string    `json:"forum_post"`
	Friend         Friend    `json:"friend"`
	Gender         string    `json:"gender"`
	Image          string    `json:"image"`
	JoinedDate     string    `json:"joined_date"`
	LastOnline     string    `json:"last_online"`
	Location       string    `json:"location"`
	MangaStat      MangaStat `json:"manga_stat"`
	Recommendation string    `json:"recommendation"`
	Review         string    `json:"review"`
	Sns            []string  `json:"sns"`
	Username       string    `json:"username"`
}

// ----- ANIME & MANGA STATISTIC -----
type AnimeStat struct {
	Days      string      `json:"days"`
	History   []History   `json:"history"`
	MeanScore string      `json:"mean_score"`
	Status    AnimeStatus `json:"status"`
}
type MangaStat struct {
	Days      string      `json:"days"`
	History   []History   `json:"history"`
	MeanScore string      `json:"mean_score"`
	Status    MangaStatus `json:"status"`
}

// ----- HISTORY -----
type History struct {
	Date     string `json:"date"`
	Id       string `json:"id"`
	Image    string `json:"image"`
	Progress string `json:"progress"`
	Score    string `json:"score"`
	Status   string `json:"status"`
	Title    string `json:"title"`
}

// ----- ANIME & MANGA PROGRESS STATUS -----
type AnimeStatus struct {
	Completed   string `json:"completed"`
	Dropped     string `json:"dropped"`
	Episode     string `json:"episode"`
	OnHold      string `json:"on_hold"`
	PlanToWatch string `json:"plan_to_watch"`
	Rewatched   string `json:"rewatched"`
	Total       string `json:"total"`
	Watching    string `json:"watching"`
}
type MangaStatus struct {
	Chapter    string `json:"chapter"`
	Completed  string `json:"completed"`
	Dropped    string `json:"dropped"`
	OnHold     string `json:"on_hold"`
	PlanToRead string `json:"plan_to_read"`
	Reading    string `json:"reading"`
	Reread     string `json:"reread"`
	Total      string `json:"total"`
	Volume     string `json:"volume"`
}

// ---- FRIEND -----
type Friend struct {
	Count string       `json:"count"`
	Data  []FriendData `json:"data"`
}
type FriendData struct {
	Image string `json:"image"`
	Name  string `json:"name"`
}

// ----- FAVORITE -----
type Favorite struct {
	Anime     []FavAnimeManga `json:"anime"`
	Character []FavCharacter  `json:"character"`
	Manga     []FavAnimeManga `json:"manga"`
	People    []FavPeople     `json:"people"`
}
type FavAnimeManga struct {
	Id    string `json:"id"`
	Image string `json:"image"`
	Title string `json:"title"`
	Type  string `json:"type"`
	Year  string `json:"year"`
}
type FavCharacter struct {
	Id         string `json:"id"`
	Image      string `json:"image"`
	MediaId    string `json:"media_id"`
	MediaTitle string `json:"media_title"`
	MediaType  string `json:"media_type"`
	Name       string `json:"name"`
}
type FavPeople struct {
	Id    string `json:"id"`
	Image string `json:"image"`
	Name  string `json:"name"`
}
