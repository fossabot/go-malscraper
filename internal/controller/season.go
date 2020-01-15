package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
	"github.com/rl404/go-malscraper/internal/view"
	"github.com/rl404/go-malscraper/pkg/malscraper/utils"
)

// registerSeasonRoutes registers all season routes.
func registerSeasonRoutes(router *chi.Mux) {
	router.Get("/season", getSeason)
}

// getSeason is season route to get MyAnimeList seasonal anime list .
// Example: https://myanimelist.net/anime/season
func getSeason(w http.ResponseWriter, r *http.Request) {
	season := r.URL.Query().Get("season")
	year, _ := strconv.Atoi(r.URL.Query().Get("year"))

	if year == 0 {
		year = time.Now().Year()
	}

	if season == "" {
		season = utils.GetCurrentSeason()
	}

	parser, err := MalService.GetSeason(year, season)

	if err != nil {
		view.RespondWithJSON(w, parser.ResponseCode, err.Error(), nil)
	} else {
		view.RespondWithJSON(w, parser.ResponseCode, parser.ResponseMessage.Error(), parser.Data)
	}
}
