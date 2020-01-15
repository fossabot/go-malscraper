package controller

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rl404/go-malscraper/internal/view"
)

// registerRecommendationRoutes registers all recommendation routes.
func registerRecommendationRoutes(router *chi.Mux) {
	router.Get("/recommendations/{t}", getRecommendations)
	router.Get("/recommendation/{t}/{id1}-{id2}", getRecommendation)
	router.Get("/recommendation/{t}/{id1}/{id2}", getRecommendation)
}

// getRecommendations is recommendation route to get MyAnimeList recommendation list.
// Example: https://myanimelist.net/recommendations.php?s=recentrecs&t=anime
func getRecommendations(w http.ResponseWriter, r *http.Request) {
	rType := chi.URLParam(r, "t")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))

	if page == 0 {
		page = 1
	}

	parser, err := MalService.GetRecommendations(rType, page)

	if err != nil {
		view.RespondWithJSON(w, parser.ResponseCode, err.Error(), nil)
	} else {
		view.RespondWithJSON(w, parser.ResponseCode, parser.ResponseMessage.Error(), parser.Data)
	}
}

// getRecommendation is recommendation route to get MyAnimeList recommendation.
// Example: https://myanimelist.net/recommendations/anime/1-205
//          https://myanimelist.net/recommendations/manga/1-21
func getRecommendation(w http.ResponseWriter, r *http.Request) {
	rType := chi.URLParam(r, "t")
	id1, _ := strconv.Atoi(chi.URLParam(r, "id1"))
	id2, _ := strconv.Atoi(chi.URLParam(r, "id2"))

	parser, err := MalService.GetRecommendation(rType, id1, id2)

	if err != nil {
		view.RespondWithJSON(w, parser.ResponseCode, err.Error(), nil)
	} else {
		view.RespondWithJSON(w, parser.ResponseCode, parser.ResponseMessage.Error(), parser.Data)
	}
}
