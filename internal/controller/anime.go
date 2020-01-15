package controller

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rl404/go-malscraper/internal/view"
)

// registerAnimeRoutes registers all anime routes.
func registerAnimeRoutes(router *chi.Mux) {
	router.Get("/anime/{id}", getAnime)
	router.Get("/anime/{id}/videos", getAnimeVideo)
	router.Get("/anime/{id}/episodes", getAnimeEpisode)
	router.Get("/anime/{id}/reviews", getAnimeReview)
	router.Get("/anime/{id}/recommendations", getAnimeRecommendation)
	router.Get("/anime/{id}/stats", getAnimeStats)
	router.Get("/anime/{id}/characters", getAnimeCharacter)
	router.Get("/anime/{id}/staff", getAnimeStaff)
	router.Get("/anime/{id}/pictures", getAnimePicture)
}

// getAnime is anime route to get MyAnimeList anime information.
// Example: https://myanimelist.net/anime/1
func getAnime(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	parser, err := MalService.GetAnime(id)

	if err != nil {
		view.RespondWithJSON(w, parser.ResponseCode, err.Error(), nil)
	} else {
		view.RespondWithJSON(w, parser.ResponseCode, parser.ResponseMessage.Error(), parser.Data)
	}
}

// getAnimeVideo is anime route to get MyAnimeList anime's video list.
// Example: https://myanimelist.net/anime/1/Cowboy_Bebop/video
func getAnimeVideo(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))

	if page == 0 {
		page = 1
	}

	parser, err := MalService.GetAnimeVideo(id, page)

	if err != nil {
		view.RespondWithJSON(w, parser.ResponseCode, err.Error(), nil)
	} else {
		view.RespondWithJSON(w, parser.ResponseCode, parser.ResponseMessage.Error(), parser.Data)
	}
}

// getAnimeEpisode is anime route to get MyAnimeList anime's episode list.
// Example: https://myanimelist.net/anime/1/Cowboy_Bebop/episode
func getAnimeEpisode(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))

	if page == 0 {
		page = 1
	}

	parser, err := MalService.GetAnimeEpisode(id, page)

	if err != nil {
		view.RespondWithJSON(w, parser.ResponseCode, err.Error(), nil)
	} else {
		view.RespondWithJSON(w, parser.ResponseCode, parser.ResponseMessage.Error(), parser.Data)
	}
}

// getAnimeReview is anime route to get MyAnimeList anime's review list.
// Example: https://myanimelist.net/anime/1/Cowboy_Bebop/reviews
func getAnimeReview(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))

	if page == 0 {
		page = 1
	}

	parser, err := MalService.GetAnimeReview(id, page)

	if err != nil {
		view.RespondWithJSON(w, parser.ResponseCode, err.Error(), nil)
	} else {
		view.RespondWithJSON(w, parser.ResponseCode, parser.ResponseMessage.Error(), parser.Data)
	}
}

// getAnimeRecommendation is anime route to get MyAnimeList anime's recommendation list.
// Example: https://myanimelist.net/anime/1/Cowboy_Bebop/userrecs
func getAnimeRecommendation(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	parser, err := MalService.GetAnimeRecommendation(id)

	if err != nil {
		view.RespondWithJSON(w, parser.ResponseCode, err.Error(), nil)
	} else {
		view.RespondWithJSON(w, parser.ResponseCode, parser.ResponseMessage.Error(), parser.Data)
	}
}

// getAnimeStats is anime route to get MyAnimeList anime's stats.
// Example: https://myanimelist.net/anime/1/Cowboy_Bebop/stats
func getAnimeStats(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))

	if page == 0 {
		page = 1
	}

	parser, err := MalService.GetAnimeStats(id, page)

	if err != nil {
		view.RespondWithJSON(w, parser.ResponseCode, err.Error(), nil)
	} else {
		view.RespondWithJSON(w, parser.ResponseCode, parser.ResponseMessage.Error(), parser.Data)
	}
}

// getAnimeCharacter is anime route to get MyAnimeList anime's character list.
// Example: https://myanimelist.net/anime/1/Cowboy_Bebop/characters
func getAnimeCharacter(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	parser, err := MalService.GetAnimeCharacter(id)

	if err != nil {
		view.RespondWithJSON(w, parser.ResponseCode, err.Error(), nil)
	} else {
		view.RespondWithJSON(w, parser.ResponseCode, parser.ResponseMessage.Error(), parser.Data)
	}
}

// getAnimeStaff is anime route to get MyAnimeList anime's staff list.
// Example: https://myanimelist.net/anime/1/Cowboy_Bebop/characters
func getAnimeStaff(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	parser, err := MalService.GetAnimeStaff(id)

	if err != nil {
		view.RespondWithJSON(w, parser.ResponseCode, err.Error(), nil)
	} else {
		view.RespondWithJSON(w, parser.ResponseCode, parser.ResponseMessage.Error(), parser.Data)
	}
}

// getAnimePicture is anime route to get MyAnimeList anime's picture list.
// Example: https://myanimelist.net/anime/1/Cowboy_Bebop/pics
func getAnimePicture(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	parser, err := MalService.GetAnimePicture(id)

	if err != nil {
		view.RespondWithJSON(w, parser.ResponseCode, err.Error(), nil)
	} else {
		view.RespondWithJSON(w, parser.ResponseCode, parser.ResponseMessage.Error(), parser.Data)
	}
}
