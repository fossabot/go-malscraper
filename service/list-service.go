package service

import "github.com/rl404/go-malscraper/model/list"

// ListsService for all list-related methods.
type ListsService struct{}

// GetAllAnimeGenre to get all anime genre.
func (l *ListsService) GetAllAnimeGenre() ([]list.GenreData, int, string) {
	var genreModel list.GenreModel
	return genreModel.InitGenreModel("anime")
}

// GetAllMangaGenre to get all anime genre.
func (l *ListsService) GetAllMangaGenre() ([]list.GenreData, int, string) {
	var genreModel list.GenreModel
	return genreModel.InitGenreModel("manga")
}

// GetAllStudioProducer to get all anime studio producer.
func (l *ListsService) GetAllStudioProducer() ([]list.ProducerData, int, string) {
	var studioProducerModel list.ProducerModel
	return studioProducerModel.InitProducerModel("anime")
}

// GetAllMagazine to get all manga magazine.
func (l *ListsService) GetAllMagazine() ([]list.ProducerData, int, string) {
	var magazineModel list.ProducerModel
	return magazineModel.InitProducerModel("manga")
}

// GetAllReview to get all anime/manga review.
func (l *ListsService) GetAllReview(t string, p int) ([]list.ReviewData, int, string) {
	var reviewModel list.ReviewModel
	return reviewModel.InitReviewModel(t, p)
}

// GetAllRecommendation to get all anime/manga recommendation.
func (l *ListsService) GetAllRecommendation(t string, p int) ([]list.RecommendationData, int, string) {
	var recommendationModel list.RecommendationModel
	return recommendationModel.InitRecommendationModel(t, p)
}
