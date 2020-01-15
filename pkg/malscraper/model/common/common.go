package common

// DateTime represents common struct containing date and time.
type DateTime struct {
	Date string `json:"date"`
	Time string `json:"time"`
}

// Genre represents genre simple model.
type Genre struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Type string `json:"type"`
}

// IDName represents common struct containing id and name.
type IDName struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// IDTitle represents common struct containing id and title.
type IDTitle struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
}

// IDTitleType represents common struct containing id, title, and type.
type IDTitleType struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Type  string `json:"type"`
}

// IDImageTitle represents common struct containing id, image, and title.
type IDImageTitle struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Image string `json:"image"`
}

// IDTypeName represents common struct containing id, type and name.
type IDTypeName struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}

// IDTitleImageRole represents common struct containing id, title, image, and role.
type IDTitleImageRole struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Image string `json:"image"`
	Role  string `json:"role"`
}

// IDTitleTypeImage represents common struct containing id, title, type, and image.
type IDTitleTypeImage struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Type  string `json:"type"`
	Image string `json:"image"`
}
