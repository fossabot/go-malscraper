package service

import "github.com/rl404/go-malscraper/model/general"

// GeneralService for all general methods.
type GeneralService struct{}

// GetInfo to get anime & manga information.
func (g *GeneralService) GetInfo(t string, id int) (general.InfoData, int, string) {
	var InfoModel general.InfoModel
	return InfoModel.InitInfoModel(t, id)
}

// GetCharacter to get character information.
func (g *GeneralService) GetCharacter(id int) (general.CharacterData, int, string) {
	var CharacterModel general.CharacterModel
	return CharacterModel.InitCharacterModel(id)
}

// GetPeople to get people information.
func (g *GeneralService) GetPeople(id int) (general.PeopleData, int, string) {
	var PeopleModel general.PeopleModel
	return PeopleModel.InitPeopleModel(id)
}

// GetProducer to get studio/producer information.
func (g *GeneralService) GetProducer(id int, page int) ([]general.ProducerData, int, string) {
	var ProducerModel general.ProducerModel
	return ProducerModel.InitProducerModel("anime", "producer", id, page)
}

// GetMagazine to get magazine information.
func (g *GeneralService) GetMagazine(id int, page int) ([]general.ProducerData, int, string) {
	var ProducerModel general.ProducerModel
	return ProducerModel.InitProducerModel("manga", "producer", id, page)
}

// GetGenre to get magazine information.
func (g *GeneralService) GetGenre(t string, id int, page int) ([]general.ProducerData, int, string) {
	var ProducerModel general.ProducerModel
	return ProducerModel.InitProducerModel(t, "genre", id, page)
}

// GetRecommendation to get anime/manga recommendation.
func (g *GeneralService) GetRecommendation(t string, id1 int, id2 int) (general.RecommendationData, int, string) {
	var RecommendationModel general.RecommendationModel
	return RecommendationModel.InitRecommendationModel(t, id1, id2)
}

// GetReview to get anime/manga review.
func (g *GeneralService) GetReview(id int) (general.ReviewData, int, string) {
	var ReviewModel general.ReviewModel
	return ReviewModel.InitReviewModel(id)
}
