package service

import "github.com/rl404/go-malscraper/model/additional"

// AdditionalService for all additional methods.
type AdditionalService struct{}

// GetStat to get anime & manga statistics.
func (g *AdditionalService) GetStat(t string, id int) (additional.StatData, int, string) {
	var StatModel additional.StatModel
	return StatModel.InitStatModel(t, id)
}