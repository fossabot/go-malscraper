package controller

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rl404/go-malscraper/internal/view"
)

// registerReviewRoutes registers all review routes.
func registerReviewRoutes(router *chi.Mux) {
	router.Get("/review/{id}", getReview)
	router.Get("/reviews", getReviews)
	router.Get("/reviews/{t}", getReviews)
}

// getReview is review route to get MyAnimeList review.
// Example: https://myanimelist.net/reviews.php?id=1
func getReview(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	parser, err := MalService.GetReview(id)

	if err != nil {
		view.RespondWithJSON(w, parser.ResponseCode, err.Error(), nil)
	} else {
		view.RespondWithJSON(w, parser.ResponseCode, parser.ResponseMessage.Error(), parser.Data)
	}
}

// getReviews is review route to get MyAnimeList review list.
// Example: https://myanimelist.net/reviews.php
func getReviews(w http.ResponseWriter, r *http.Request) {
	// rType := chi.URLParam(r, "t")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))

	if page == 0 {
		page = 1
	}

	parser, err := MalService.GetReviews()

	if err != nil {
		view.RespondWithJSON(w, parser.ResponseCode, err.Error(), nil)
	} else {
		view.RespondWithJSON(w, parser.ResponseCode, parser.ResponseMessage.Error(), parser.Data)
	}
}
