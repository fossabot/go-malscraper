package user

// UserData for main user data model.
type UserData struct {
	Username       string    `json:"username"`
	Image          string    `json:"image"`
	LastOnline     string    `json:"last_online"`
	Gender         string    `json:"gender"`
	Birthday       string    `json:"birthday"`
	Location       string    `json:"location"`
	JoinedDate     string    `json:"joined_date"`
	ForumPost      int       `json:"forum_post"`
	Review         int       `json:"review"`
	Recommendation int       `json:"recommendation"`
	BlogPost       int       `json:"blog_post"`
	Club           int       `json:"club"`
	Sns            []string  `json:"sns"`
	Friend         Friend    `json:"friend"`
	About          string    `json:"about"`
	AnimeStat      AnimeStat `json:"anime_stat"`
	MangaStat      MangaStat `json:"manga_stat"`
	Favorite       Favorite  `json:"favorite"`
}

// AnimeStat for anime statistic model.
type AnimeStat struct {
	Days      float64     `json:"days"`
	MeanScore float64     `json:"mean_score"`
	Status    AnimeStatus `json:"status"`
	History   []History   `json:"history"`
}

// MangaStat for manga statistic model.
type MangaStat struct {
	Days      float64     `json:"days"`
	MeanScore float64     `json:"mean_score"`
	Status    MangaStatus `json:"status"`
	History   []History   `json:"history"`
}

// History for anime & manga history model.
type History struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Image    string `json:"image"`
	Progress string `json:"progress"`
	Score    int    `json:"score"`
	Status   string `json:"status"`
	Date     string `json:"date"`
}

// AnimeStatus for anime progress status model.
type AnimeStatus struct {
	Completed   int `json:"completed"`
	Dropped     int `json:"dropped"`
	Episode     int `json:"episode"`
	OnHold      int `json:"on_hold"`
	PlanToWatch int `json:"plan_to_watch"`
	Rewatched   int `json:"rewatched"`
	Total       int `json:"total"`
	Watching    int `json:"watching"`
}

// MangaStatus for manga progress status model.
type MangaStatus struct {
	Chapter    int `json:"chapter"`
	Completed  int `json:"completed"`
	Dropped    int `json:"dropped"`
	OnHold     int `json:"on_hold"`
	PlanToRead int `json:"plan_to_read"`
	Reading    int `json:"reading"`
	Reread     int `json:"reread"`
	Total      int `json:"total"`
	Volume     int `json:"volume"`
}

// Friend for friend data model.
type Friend struct {
	Count int          `json:"count"`
	Data  []FriendData `json:"data"`
}

// FriendData for friend detail data model.
type FriendData struct {
	Name  string `json:"name"`
	Image string `json:"image"`
}

// Favorite for favorite data model.
type Favorite struct {
	Anime     []FavAnimeManga `json:"anime"`
	Character []FavCharacter  `json:"character"`
	Manga     []FavAnimeManga `json:"manga"`
	People    []FavPeople     `json:"people"`
}

// FavAnimeManga for anime & manga favorite detail data model.
type FavAnimeManga struct {
	Id    int    `json:"id"`
	Title string `json:"title"`
	Image string `json:"image"`
	Type  string `json:"type"`
	Year  int    `json:"year"`
}

// FavCharacter for character favorite detail data model.
type FavCharacter struct {
	Id         int    `json:"id"`
	Image      string `json:"image"`
	Name       string `json:"name"`
	MediaId    int    `json:"media_id"`
	MediaTitle string `json:"media_title"`
	MediaType  string `json:"media_type"`
}

// FavPeople for people favorite detail data model.
type FavPeople struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Image string `json:"image"`
}

// UserFriendData for main user friend list model.
type UserFriendData struct {
	Name        string `json:"name"`
	Image       string `json:"image"`
	LastOnline  string `json:"last_online"`
	FriendSince string `json:"friend_since"`
}

// UserHistoryData for main user history list model.
type UserHistoryData struct {
	Id       int    `json:"id"`
	Title    string `json:"title"`
	Type     string `json:"type"`
	Progress int    `json:"progress"`
	Date     string `json:"date"`
}

// UserListData for main user anime/manga list model.
// TODO: Handle dynamic (dirty) data type
type UserListData struct {
	Status            int         `json:"status"`
	Score             int         `json:"score"`
	Tags              interface{} `json:"tags"`
	IsRewatching      interface{} `json:"is_rewatching"`
	WatchedEpisode    int         `json:"num_watched_episodes"`
	AnimeTitle        interface{} `json:"anime_title"`
	AnimeEpisode      int         `json:"anime_num_episodes"`
	AnimeAiringStatus int         `json:"anime_airing_status"`
	AnimeId           int         `json:"anime_id"`
	AnimeStudio       []IdName    `json:"anime_studios"`
	AnimeLicensor     []IdName    `json:"anime_licensors"`
	AnimeSeason       Season      `json:"anime_season"`
	HasEpisodeVideo   bool        `json:"has_episode_video"`
	HasPromotionVideo bool        `json:"has_promotion_video"`
	HasVideo          bool        `json:"has_video"`
	VideoUrl          string      `json:"video_url"`
	AnimeUrl          string      `json:"anime_url"`
	AnimeImage        string      `json:"anime_image_path"`
	AnimeType         string      `json:"anime_media_type_string"`
	AnimeRating       string      `json:"anime_mpaa_rating_string"`
	AnimeStartDate    string      `json:"anime_start_date_string"`
	AnimeEndDate      string      `json:"anime_end_date_string"`
	IsRereading       interface{} `json:"is_rereading"`            // manga
	ReadChapter       int         `json:"num_read_chapters"`       // manga
	ReadVolume        int         `json:"num_read_volumes"`        // manga
	MangaTitle        interface{} `json:"manga_title"`             // manga
	MangaChapter      int         `json:"manga_num_chapters"`      // manga
	MangaVolume       int         `json:"manga_num_volume"`        // manga
	MangaStatus       int         `json:"manga_publishing_status"` // manga
	MangaId           int         `json:"manga_id"`                // manga
	MangaMagazine     []IdName    `json:"manga_magazines"`         // manga
	MangaUrl          string      `json:"manga_url"`               // manga
	MangaImage        string      `json:"manga_image_path"`        // manga
	MangaType         string      `json:"manga_media_type_string"` // manga
	MangaStartDate    string      `json:"manga_start_date_string"` // manga
	MangaEndDate      string      `json:"manga_end_date_string"`   // manga
	Retail            string      `json:"retail_string"`           // manga
	IsAddedToList     bool        `json:"is_added_to_list"`
	StartDate         string      `json:"start_date_string"`
	FinishDate        string      `json:"finish_date_string"`
	Day               int         `json:"days_string"`
	Storage           string      `json:"storage_string"`
	Priority          string      `json:"priority_string"`
}

// IdName is common model contains id and name.
type IdName struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// Season is season model contains year and season.
type Season struct {
	Year   int    `json:"year"`
	Season string `json:"season"`
}
