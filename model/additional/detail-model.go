package additional

// StatData for main anime & manga statistics data model.
type StatData struct {
	Summary 	map[string]string	`json:"summary"`
	Score 		[]Score 			`json:"score"`
	User 		[]User 				`json:"user"`
}

// Score is a model for anime/manga score.
type Score struct {
	Type 	 	string 		`json:"type"`
	Vote 		string 		`json:"vote"`
	Percent 	string 		`json:"percent"`
}

// User is a model for user in StatData model.
type User struct {
 	Image 		string 		`json:"image"`
 	Username 	string 		`json:"username"`
 	Score 		string 		`json:"score"`
 	Status 		string 		`json:"status"`
 	Episode 	string 		`json:"episode"`
 	Volume 		string 		`json:"volume"` 	// manga
 	Chapter 	string 		`json:"chapter"` 	// manga
 	Date 		string 		`json:"date"`
}
