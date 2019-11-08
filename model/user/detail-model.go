package user

// UserData for main user data model.
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

// AnimeStat for anime statistic model.
type AnimeStat struct {
	Days      string      `json:"days"`
	History   []History   `json:"history"`
	MeanScore string      `json:"mean_score"`
	Status    AnimeStatus `json:"status"`
}

// MangaStat for manga statistic model.
type MangaStat struct {
	Days      string      `json:"days"`
	History   []History   `json:"history"`
	MeanScore string      `json:"mean_score"`
	Status    MangaStatus `json:"status"`
}

// History for anime & manga history model.
type History struct {
	Date     string `json:"date"`
	Id       string `json:"id"`
	Image    string `json:"image"`
	Progress string `json:"progress"`
	Score    string `json:"score"`
	Status   string `json:"status"`
	Title    string `json:"title"`
}

// AnimeStatus for anime progress status model.
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

// MangaStatus for manga progress status model.
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

// Friend for friend data model.
type Friend struct {
	Count string       `json:"count"`
	Data  []FriendData `json:"data"`
}

// FriendData for friend detail data model.
type FriendData struct {
	Image string `json:"image"`
	Name  string `json:"name"`
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
	Id    string `json:"id"`
	Image string `json:"image"`
	Title string `json:"title"`
	Type  string `json:"type"`
	Year  string `json:"year"`
}

// FavCharacter for character favorite detail data model.
type FavCharacter struct {
	Id         string `json:"id"`
	Image      string `json:"image"`
	MediaId    string `json:"media_id"`
	MediaTitle string `json:"media_title"`
	MediaType  string `json:"media_type"`
	Name       string `json:"name"`
}

// FavPeople for people favorite detail data model.
type FavPeople struct {
	Id    string `json:"id"`
	Image string `json:"image"`
	Name  string `json:"name"`
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
