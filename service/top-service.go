package service

import "github.com/rl404/go-malscraper/model/top"

// TopService for all top-list-related methods.
type TopService struct{}

// GetTopAnime to get top anime list.
func (t *TopService) GetTopAnime(typ int, page int) ([]top.TopAnimeMangaData, int, string) {
	var topModel top.TopAnimeMangaModel
	return topModel.InitTopAnimeMangaModel("anime", typ, page)
}

// GetTopManga to get top manga list.
func (t *TopService) GetTopManga(typ int, page int) ([]top.TopAnimeMangaData, int, string) {
	var topModel top.TopAnimeMangaModel
	return topModel.InitTopAnimeMangaModel("manga", typ, page)
}
