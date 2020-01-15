package anime

// Staff represents the main model for MyAnimeList anime staff.
type Staff struct {
	ID    int    `json:"id"`
	Name  string `json:"name"`
	Role  string `json:"role"`
	Image string `json:"image"`
}
