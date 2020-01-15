package user

// Friend represents the main model for MyAnimeList user's friend list.
type Friend struct {
	Name        string `json:"name"`
	Image       string `json:"image"`
	LastOnline  string `json:"lastOnline"`
	FriendSince string `json:"friendSince"`
}
