package controller

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rl404/go-malscraper/internal/view"
)

// registerTopRoutes registers all top routes.
func registerTopRoutes(router *chi.Mux) {
	router.Get("/top/anime", getTopAnime)
	router.Get("/top/manga", getTopManga)
	router.Get("/top/character", getTopCharacter)
	router.Get("/top/people", getTopPeople)
}

// getTopAnime is top route to get MyAnimeList top anime list.
// Example: https://myanimelist.net/topanime.php
func getTopAnime(w http.ResponseWriter, r *http.Request) {
	topType, _ := strconv.Atoi(r.URL.Query().Get("type"))
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))

	if page == 0 {
		page = 1
	}

	parser, err := MalService.GetTopAnime(topType, page)

	if err != nil {
		view.RespondWithJSON(w, parser.ResponseCode, err.Error(), nil)
	} else {
		view.RespondWithJSON(w, parser.ResponseCode, parser.ResponseMessage.Error(), parser.Data)
	}
}

// getTopManga is top route to get MyAnimeList top manga list.
// Example: https://myanimelist.net/topmanga.php
func getTopManga(w http.ResponseWriter, r *http.Request) {
	topType, _ := strconv.Atoi(r.URL.Query().Get("type"))
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))

	if page == 0 {
		page = 1
	}

	parser, err := MalService.GetTopManga(topType, page)

	if err != nil {
		view.RespondWithJSON(w, parser.ResponseCode, err.Error(), nil)
	} else {
		view.RespondWithJSON(w, parser.ResponseCode, parser.ResponseMessage.Error(), parser.Data)
	}
}

// getTopCharacter is top route to get MyAnimeList top character list.
// Example: https://myanimelist.net/character.php
func getTopCharacter(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))

	if page == 0 {
		page = 1
	}

	parser, err := MalService.GetTopCharacter(page)

	if err != nil {
		view.RespondWithJSON(w, parser.ResponseCode, err.Error(), nil)
	} else {
		view.RespondWithJSON(w, parser.ResponseCode, parser.ResponseMessage.Error(), parser.Data)
	}
}

// getTopPeople is top route to get MyAnimeList top people list.
// Example: https://myanimelist.net/character.php
func getTopPeople(w http.ResponseWriter, r *http.Request) {
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))

	if page == 0 {
		page = 1
	}

	parser, err := MalService.GetTopPeople(page)

	if err != nil {
		view.RespondWithJSON(w, parser.ResponseCode, err.Error(), nil)
	} else {
		view.RespondWithJSON(w, parser.ResponseCode, parser.ResponseMessage.Error(), parser.Data)
	}
}
