package season

// SeasonData for main anime season data model.
type SeasonData struct {
	Image       string   `json:"image"`
	Id          string   `json:"id"`
	Title       string   `json:"title"`
	Producer    []IdName `json:"producer"`
	Episode     string   `json:"episode"`
	Source      string   `json:"source"`
	Genre       []IdName `json:"genre"`
	Synopsis    string   `json:"synopsis"`
	Licensor    []string `json:"licensor"`
	Type        string   `json:"type"`
	AiringStart string   `json:"airing_start"`
	Member      string   `json:"member"`
	Score       string   `json:"score"`
}

// IdName is common model contain id and name.
type IdName struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}
