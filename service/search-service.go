package service

import "github.com/rl404/go-malscraper/model/search"

// SearchService for all search-related methods.
type SearchService struct{}

// SearchAnime to get anime search result list.
func (u *SearchService) SearchAnime(query string, page int) ([]search.SearchAnimeMangaData, int, string) {
	var searchModel search.SearchAnimeMangaModel
	return searchModel.InitSearchAnimeMangaModel("anime", query, page)
}

// SearchManga to get manga search result list.
func (u *SearchService) SearchManga(query string, page int) ([]search.SearchAnimeMangaData, int, string) {
	var searchModel search.SearchAnimeMangaModel
	return searchModel.InitSearchAnimeMangaModel("manga", query, page)
}

// SearchCharacter to get character search result list.
func (u *SearchService) SearchCharacter(query string, page int) ([]search.SearchCharPeopleData, int, string) {
	var searchModel search.SearchCharPeopleModel
	return searchModel.InitSearchCharPeopleModel("character", query, page)
}

// SearchPeople to get people search result list.
func (u *SearchService) SearchPeople(query string, page int) ([]search.SearchCharPeopleData, int, string) {
	var searchModel search.SearchCharPeopleModel
	return searchModel.InitSearchCharPeopleModel("people", query, page)
}
