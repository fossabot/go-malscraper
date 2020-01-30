package controller

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/rl404/go-malscraper/internal/config"
	"github.com/rl404/go-malscraper/internal/view"
	"github.com/rl404/go-malscraper/pkg/malscraper"
)

// MalService for all mal service controller.
var MalService *malscraper.MalService

// init to initiate MalService instance.
func init() {
	malConfig, err := config.GetConfig()
	if err != nil {
		// If there is an error in parsing config file, malscraper will be
		// initiated with default mode (no config).
		MalService = malscraper.Default()
	} else {
		MalService = malscraper.New(malConfig)
	}
}

// RegisterRoutesV1 registers all go-malscraper routes version 1.
func RegisterRoutesV1() http.Handler {
	router := chi.NewRouter()
	registerAnimeRoutes(router)
	registerMangaRoutes(router)
	registerCharacterRoutes(router)
	registerPeopleRoutes(router)
	registerProducerRoutes(router)
	registerMagazineRoutes(router)
	registerGenreRoutes(router)
	registerReviewRoutes(router)
	registerRecommendationRoutes(router)
	registerUserRoutes(router)
	registerSearchRoutes(router)
	registerSeasonRoutes(router)
	registerTopRoutes(router)
	return router
}

// RegisterBaseRoutes registers base routes.
func RegisterBaseRoutes(router *chi.Mux) {
	router.NotFound(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		view.RespondWithJSON(w, 404, "page not found", nil)
	}))

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		view.RespondWithJSON(w, 200, "root", nil)
	})

	router.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		view.RespondWithJSON(w, 200, "pong", nil)
	})
}
