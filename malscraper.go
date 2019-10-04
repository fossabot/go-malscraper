package malscraper

import "github.com/rl404/go-malscraper/service"

// MalScraper is the main function to call all methods in go-malscraper.
func MalScraper() *MalService {
	m := &MalService{}
	return m
}

// MalService for all go-malscraper service.
type MalService struct {
	User       *service.UserService
	General    *service.GeneralService
	Additional *service.AdditionalService
	// Lists		*ListsService
	// Search		*SearchService
	// Seasonal		*SeasonalService
	// Top			*TopService
}
