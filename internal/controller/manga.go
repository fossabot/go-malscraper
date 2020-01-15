package controller

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rl404/go-malscraper/internal/view"
)

// registerMangaRoutes registers all manga routes.
func registerMangaRoutes(router *chi.Mux) {
	router.Get("/manga/{id}", getManga)
	router.Get("/manga/{id}/reviews", getMangaReview)
	router.Get("/manga/{id}/recommendations", getMangaRecommendation)
	router.Get("/manga/{id}/stats", getMangaStats)
	router.Get("/manga/{id}/characters", getMangaCharacter)
	router.Get("/manga/{id}/pictures", getMangaPicture)
}

// getManga is manga route to get MyAnimeList manga information.
// Example: https://myanimelist.net/manga/1
func getManga(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	parser, err := MalService.GetManga(id)

	if err != nil {
		view.RespondWithJSON(w, parser.ResponseCode, err.Error(), nil)
	} else {
		view.RespondWithJSON(w, parser.ResponseCode, parser.ResponseMessage.Error(), parser.Data)
	}
}

// getMangaReview is manga route to get MyAnimeList manga's review list.
// Example: https://myanimelist.net/manga/1/Monster/reviews
func getMangaReview(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))

	if page == 0 {
		page = 1
	}

	videos, err := MalService.GetMangaReview(id, page)

	if err != nil {
		view.RespondWithJSON(w, videos.ResponseCode, err.Error(), nil)
	} else {
		view.RespondWithJSON(w, videos.ResponseCode, videos.ResponseMessage.Error(), videos.Data)
	}
}

// getMangaRecommendation is manga route to get MyAnimeList manga's recommendation list.
// Example: https://myanimelist.net/manga/1/Monster/userrecs
func getMangaRecommendation(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	parser, err := MalService.GetMangaRecommendation(id)

	if err != nil {
		view.RespondWithJSON(w, parser.ResponseCode, err.Error(), nil)
	} else {
		view.RespondWithJSON(w, parser.ResponseCode, parser.ResponseMessage.Error(), parser.Data)
	}
}

// getMangaStats is manga route to get MyAnimeList manga's stats.
// Example: https://myanimelist.net/manga/1/Monster/stats
func getMangaStats(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))

	if page == 0 {
		page = 1
	}

	parser, err := MalService.GetMangaStats(id, page)

	if err != nil {
		view.RespondWithJSON(w, parser.ResponseCode, err.Error(), nil)
	} else {
		view.RespondWithJSON(w, parser.ResponseCode, parser.ResponseMessage.Error(), parser.Data)
	}
}

// getMangaCharacter is manga route to get MyAnimeList manga's character list.
// Example: https://myanimelist.net/manga/1/Monster/characters
func getMangaCharacter(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	parser, err := MalService.GetMangaCharacter(id)

	if err != nil {
		view.RespondWithJSON(w, parser.ResponseCode, err.Error(), nil)
	} else {
		view.RespondWithJSON(w, parser.ResponseCode, parser.ResponseMessage.Error(), parser.Data)
	}
}

// getMangaPicture is manga route to get MyAnimeList manga's picture list.
// Example: https://myanimelist.net/manga/1/Monster/characters
func getMangaPicture(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	parser, err := MalService.GetMangaPicture(id)

	if err != nil {
		view.RespondWithJSON(w, parser.ResponseCode, err.Error(), nil)
	} else {
		view.RespondWithJSON(w, parser.ResponseCode, parser.ResponseMessage.Error(), parser.Data)
	}
}
