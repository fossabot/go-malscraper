package service

import "github.com/rl404/go-malscraper/model/additional"

// AdditionalService for all additional methods.
type AdditionalService struct{}

// GetStat to get anime & manga statistics.
func (g *AdditionalService) GetStat(t string, id int) (additional.StatData, int, string) {
	var StatModel additional.StatModel
	return StatModel.InitStatModel(t, id)
}

// GetVideo to get anime additional video list.
func (g *AdditionalService) GetVideo(id int, p int) (additional.VideoData, int, string) {
	var VideoModel additional.VideoModel
	return VideoModel.InitVideoModel(id, p)
}

// GetCharacterStaff to get anime/manga additional staff+character list.
func (g *AdditionalService) GetCharacterStaff(t string, id int) (additional.CharacterStaffData, int, string) {
	var CharacterStaffModel additional.CharacterStaffModel
	return CharacterStaffModel.InitCharacterStaffModel(t, id)
}
