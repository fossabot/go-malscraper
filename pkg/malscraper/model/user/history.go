package user

// UserHistory represents the main model for MyAnimelist user's history list.
type UserHistory struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Type     string `json:"type"`
	Progress int    `json:"progress"`
	Date     string `json:"date"`
}
