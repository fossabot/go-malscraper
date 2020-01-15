package controller

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rl404/go-malscraper/internal/view"
)

// registerMagazineRoutes registers all magazine routes.
func registerMagazineRoutes(router *chi.Mux) {
	router.Get("/magazines", getMagazines)
	router.Get("/magazine/{id}", getMagazine)
}

// getMagazines is magazine route to get MyAnimeList all magazines, and serializations.
// Example: https://myanimelist.net/manga/magazine
func getMagazines(w http.ResponseWriter, r *http.Request) {
	parser, err := MalService.GetMagazines()

	if err != nil {
		view.RespondWithJSON(w, parser.ResponseCode, err.Error(), nil)
	} else {
		view.RespondWithJSON(w, parser.ResponseCode, parser.ResponseMessage.Error(), parser.Data)
	}
}

// getMagazine is magazine route to get MyAnimeList magazine's manga list.
// Example: https://myanimelist.net/manga/magazine/1/Big_Comic_Original
func getMagazine(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))

	if page == 0 {
		page = 1
	}

	parser, err := MalService.GetMagazine(id, page)

	if err != nil {
		view.RespondWithJSON(w, parser.ResponseCode, err.Error(), nil)
	} else {
		view.RespondWithJSON(w, parser.ResponseCode, parser.ResponseMessage.Error(), parser.Data)
	}
}
