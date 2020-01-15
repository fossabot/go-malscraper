package anime

import "github.com/rl404/go-malscraper/pkg/malscraper/model/common"

// Anime represent the main model for MyAnimeList anime information.
type Anime struct {
	ID                int                   `json:"id"`
	Cover             string                `json:"cover"`
	Title             string                `json:"title"`
	AlternativeTitles AlternativeTitle      `json:"alternativeTitles"`
	Video             string                `json:"video"`
	Synopsis          string                `json:"synopsis"`
	Score             float64               `json:"score"`
	Voter             int                   `json:"voter"`
	Rank              int                   `json:"rank"`
	Popularity        int                   `json:"popularity"`
	Member            int                   `json:"member"`
	Favorite          int                   `json:"favorite"`
	Type              string                `json:"type"`
	Episode           int                   `json:"episode"`
	Status            string                `json:"status"`
	StartDate         StartEndDate          `json:"startDate"`
	Premiered         string                `json:"premiered"`
	Broadcast         string                `json:"broadcast"`
	Producers         []common.IDName       `json:"producers"`
	Licensors         []common.IDName       `json:"licensors"`
	Studios           []common.IDName       `json:"studios"`
	Source            string                `json:"source"`
	Genres            []common.IDName       `json:"genres"`
	Duration          string                `json:"duration"`
	Rating            string                `json:"rating"`
	Related           map[string][]Related  `json:"related"`
	Characters        []AnimeCharacter      `json:"characters"`
	Staff             []Staff               `json:"staff"`
	Song              Song                  `json:"song"`
	Reviews           []Review              `json:"reviews"`
	Recommendations   []AnimeRecommendation `json:"recommendations"`
}

// AlternativeTitle represents alternative english, synonym, and japanese title of anime.
type AlternativeTitle struct {
	English  string `json:"english"`
	Synonym  string `json:"synonym"`
	Japanese string `json:"japanese"`
}

// StartEndDate represents anime start and end of airing date.
type StartEndDate struct {
	Start string `json:"start"`
	End   string `json:"end"`
}

// Related represents related anime & manga model.
type Related struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Type  string `json:"type"`
}

// AnimeCharacter represents simple character model with its voice actor information.
type AnimeCharacter struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Role    string `json:"role"`
	Image   string `json:"image"`
	VaId    int    `json:"vaId"`
	VaName  string `json:"vaName"`
	VaImage string `json:"vaImage"`
	VaRole  string `json:"vaRole"`
}

// Song represents list of opening and ending anime songs.
type Song struct {
	Opening []string `json:"opening"`
	Closing []string `json:"closing"`
}

// AnimeRecommendation represents simple anime recommendation.
type AnimeRecommendation struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Image string `json:"image"`
	Count int    `json:"count"`
}
