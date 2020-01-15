package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/rl404/go-malscraper/internal/view"
	model "github.com/rl404/go-malscraper/pkg/malscraper/model/search"
)

// registerSearchRoutes registers all search routes.
func registerSearchRoutes(router *chi.Mux) {
	router.Get("/search/anime", searchAnime)
	router.Get("/search/manga", searchManga)
	router.Get("/search/character", searchCharacter)
	router.Get("/search/people", searchPeople)
	router.Get("/search/user", searchUser)
}

// searchAnime is search route to get MyAnimeList anime search result list.
// Example: https://myanimelist.net/anime.php?q=naruto
func searchAnime(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))

	if page == 0 {
		page = 1
	}

	var queryObj model.Query
	queryObj.Query = query
	queryObj.Page = page
	queryObj.Type, _ = strconv.Atoi(r.URL.Query().Get("type"))
	queryObj.Score, _ = strconv.Atoi(r.URL.Query().Get("score"))
	queryObj.Status, _ = strconv.Atoi(r.URL.Query().Get("status"))
	queryObj.Producer, _ = strconv.Atoi(r.URL.Query().Get("producer"))
	queryObj.Rating, _ = strconv.Atoi(r.URL.Query().Get("rating"))
	queryObj.StartDate, _ = time.Parse("2006-01-02", r.URL.Query().Get("start"))
	queryObj.EndDate, _ = time.Parse("2006-01-02", r.URL.Query().Get("end"))
	queryObj.IsExcludeGenre, _ = strconv.Atoi(r.URL.Query().Get("exclude"))
	// todo: genre
	queryObj.Letter = r.URL.Query().Get("letter")

	parser, err := MalService.AdvSearchAnime(queryObj)

	if err != nil {
		view.RespondWithJSON(w, parser.ResponseCode, err.Error(), nil)
	} else {
		view.RespondWithJSON(w, parser.ResponseCode, parser.ResponseMessage.Error(), parser.Data)
	}
}

// searchManga is search route to get MyAnimeList manga search result list.
// Example: https://myanimelist.net/manga.php?q=naruto
func searchManga(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))

	if page == 0 {
		page = 1
	}

	var queryObj model.Query
	queryObj.Query = query
	queryObj.Page = page
	queryObj.Type, _ = strconv.Atoi(r.URL.Query().Get("type"))
	queryObj.Score, _ = strconv.Atoi(r.URL.Query().Get("score"))
	queryObj.Status, _ = strconv.Atoi(r.URL.Query().Get("status"))
	queryObj.Magazine, _ = strconv.Atoi(r.URL.Query().Get("magazine"))
	queryObj.StartDate, _ = time.Parse("2006-01-02", r.URL.Query().Get("start"))
	queryObj.EndDate, _ = time.Parse("2006-01-02", r.URL.Query().Get("end"))
	queryObj.IsExcludeGenre, _ = strconv.Atoi(r.URL.Query().Get("exclude"))
	// todo: genre
	queryObj.Letter = r.URL.Query().Get("letter")

	parser, err := MalService.AdvSearchManga(queryObj)

	if err != nil {
		view.RespondWithJSON(w, parser.ResponseCode, err.Error(), nil)
	} else {
		view.RespondWithJSON(w, parser.ResponseCode, parser.ResponseMessage.Error(), parser.Data)
	}
}

// searchCharacter is search route to get MyAnimeList character search result list.
// Example: https://myanimelist.net/character.php?q=luffy
func searchCharacter(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))

	if page == 0 {
		page = 1
	}

	parser, err := MalService.SearchCharacter(query, page)

	if err != nil {
		view.RespondWithJSON(w, parser.ResponseCode, err.Error(), nil)
	} else {
		view.RespondWithJSON(w, parser.ResponseCode, parser.ResponseMessage.Error(), parser.Data)
	}
}

// searchPeople is search route to get MyAnimeList people search result list.
// Example: https://myanimelist.net/people.php?q=kana
func searchPeople(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))

	if page == 0 {
		page = 1
	}

	parser, err := MalService.SearchPeople(query, page)

	if err != nil {
		view.RespondWithJSON(w, parser.ResponseCode, err.Error(), nil)
	} else {
		view.RespondWithJSON(w, parser.ResponseCode, parser.ResponseMessage.Error(), parser.Data)
	}
}

// searchUser is search route to get MyAnimeList user search result list.
// Example: https://myanimelist.net/people.php?q=kana
func searchUser(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("query")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))

	if page == 0 {
		page = 1
	}

	parser, err := MalService.SearchUser(query, page)

	if err != nil {
		view.RespondWithJSON(w, parser.ResponseCode, err.Error(), nil)
	} else {
		view.RespondWithJSON(w, parser.ResponseCode, parser.ResponseMessage.Error(), parser.Data)
	}
}
