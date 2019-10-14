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

// GetPicture to get anime/manga additional picture list.
func (g *AdditionalService) GetPicture(t string, id int) ([]string, int, string) {
	var PictureModel additional.PictureModel
	return PictureModel.InitPictureModel(t, id)
}

// GetCharacterPicture to get character additional picture list.
func (g *AdditionalService) GetCharacterPicture(id int) ([]string, int, string) {
	var CharacterPictureModel additional.CharacterPeoplePictureModel
	return CharacterPictureModel.InitCharacterPeoplePictureModel("character", id)
}

// GetPeoplePicture to get people additional picture list.
func (g *AdditionalService) GetPeoplePicture(id int) ([]string, int, string) {
	var PeoplePictureModel additional.CharacterPeoplePictureModel
	return PeoplePictureModel.InitCharacterPeoplePictureModel("people", id)
}

// GetEpisode to get anime additional episode list.
func (g *AdditionalService) GetEpisode(id int, p int) ([]additional.EpisodeData, int, string) {
	var EpisodeModel additional.EpisodeModel
	return EpisodeModel.InitEpisodeModel(id, p)
}

// GetAnimeReview to get anime additional review list.
func (g *AdditionalService) GetAnimeReview(id int, p int) ([]additional.ReviewData, int, string) {
	var ReviewModel additional.ReviewModel
	return ReviewModel.InitReviewModel("anime", id, p)
}

// GetMangaReview to get manga additional review list.
func (g *AdditionalService) GetMangaReview(id int, p int) ([]additional.ReviewData, int, string) {
	var ReviewModel additional.ReviewModel
	return ReviewModel.InitReviewModel("manga", id, p)
}
