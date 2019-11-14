package service

import "github.com/rl404/go-malscraper/model/season"

// SeasonService for all season-related methods.
type SeasonService struct{}

// GetSeason to get list anime in the season.
func (u *SeasonService) GetSeason(year int, s string) ([]season.SeasonData, int, string) {
	var seasonModel season.SeasonModel
	return seasonModel.InitSeasonModel(year, s)
}
