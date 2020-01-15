package controller

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rl404/go-malscraper/internal/view"
)

// registerProducerRoutes registers all producer routes.
func registerProducerRoutes(router *chi.Mux) {
	router.Get("/producers", getProducers)
	router.Get("/producer/{id}", getProducer)
}

// getProducers is producer route to get MyAnimeList all producers, studios, and licensors.
// Example: https://myanimelist.net/anime/producer
func getProducers(w http.ResponseWriter, r *http.Request) {
	parser, err := MalService.GetProducers()

	if err != nil {
		view.RespondWithJSON(w, parser.ResponseCode, err.Error(), nil)
	} else {
		view.RespondWithJSON(w, parser.ResponseCode, parser.ResponseMessage.Error(), parser.Data)
	}
}

// getProducer is producer route to get MyAnimeList producer's anime list.
// Example: https://myanimelist.net/anime/producer/1/Studio_Pierrot
func getProducer(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))

	if page == 0 {
		page = 1
	}

	parser, err := MalService.GetProducer(id, page)

	if err != nil {
		view.RespondWithJSON(w, parser.ResponseCode, err.Error(), nil)
	} else {
		view.RespondWithJSON(w, parser.ResponseCode, parser.ResponseMessage.Error(), parser.Data)
	}
}
