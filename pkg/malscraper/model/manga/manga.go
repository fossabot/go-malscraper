package manga

import "github.com/rl404/go-malscraper/pkg/malscraper/model/common"

// Manga represent the main model for MyAnimeList manga information.
type Manga struct {
	ID                int                   `json:"id"`
	Cover             string                `json:"cover"`
	Title             string                `json:"title"`
	AlternativeTitles AlternativeTitle      `json:"alternativeTitles"`
	Synopsis          string                `json:"synopsis"`
	Score             float64               `json:"score"`
	Voter             int                   `json:"voter"`
	Rank              int                   `json:"rank"`
	Popularity        int                   `json:"popularity"`
	Member            int                   `json:"member"`
	Favorite          int                   `json:"favorite"`
	Type              string                `json:"type"`
	Volume            int                   `json:"volume"`
	Chapter           int                   `json:"chapter"`
	Status            string                `json:"status"`
	StartDate         StartEndDate          `json:"startDate"`
	Genres            []common.IDName       `json:"genres"`
	Authors           []common.IDName       `json:"authors"`
	Serializations    []common.IDName       `json:"serializations"`
	Related           map[string][]Related  `json:"related"`
	Characters        []Character           `json:"characters"`
	Reviews           []Review              `json:"reviews"`
	Recommendations   []MangaRecommendation `json:"recommendations"`
}

// AlternativeTitle represents alternative english, synonym, and japanese title of manga.
type AlternativeTitle struct {
	English  string `json:"english"`
	Synonym  string `json:"synonym"`
	Japanese string `json:"japanese"`
}

// StartEndDate represents manga start and end of publishing date.
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

// MangaRecommendation represents simple manga recommendation.
type MangaRecommendation struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Image string `json:"image"`
	Count int    `json:"count"`
}
