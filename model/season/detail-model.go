package season

// SeasonData for main anime season data model.
type SeasonData struct {
	Id          int      `json:"id"`
	Title       string   `json:"title"`
	Image       string   `json:"image"`
	Producer    []IdName `json:"producer"`
	Episode     int      `json:"episode"`
	Source      string   `json:"source"`
	Genre       []IdName `json:"genre"`
	Synopsis    string   `json:"synopsis"`
	Licensor    []string `json:"licensor"`
	Type        string   `json:"type"`
	AiringStart string   `json:"airing_start"`
	Member      int      `json:"member"`
	Score       float64  `json:"score"`
}

// IdName is common model contain id and name.
type IdName struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
