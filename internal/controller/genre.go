package controller

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rl404/go-malscraper/internal/view"
)

// registerGenreRoutes registers all genre routes.
func registerGenreRoutes(router *chi.Mux) {
	router.Get("/genres/{t}", getGenres)
	router.Get("/genre/{id}/anime", getAnimeWithGenre)
	router.Get("/genre/{id}/manga", getMangaWithGenre)
}

// getGenres is genre route to get MyAnimeList all anime & manga genre list.
// Example: https://myanimelist.net/anime.php
//          https://myanimelist.net/manga.php
func getGenres(w http.ResponseWriter, r *http.Request) {
	gType := chi.URLParam(r, "t")

	parser, err := MalService.GetGenres(gType)

	if err != nil {
		view.RespondWithJSON(w, parser.ResponseCode, err.Error(), nil)
	} else {
		view.RespondWithJSON(w, parser.ResponseCode, parser.ResponseMessage.Error(), parser.Data)
	}
}

// getAnimeWithGenre is genre route to get MyAnimeList anime list having specific genre.
// Example: https://myanimelist.net/anime/genre/1/Action
func getAnimeWithGenre(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))

	if page == 0 {
		page = 1
	}

	parser, err := MalService.GetAnimeWithGenre(id, page)

	if err != nil {
		view.RespondWithJSON(w, parser.ResponseCode, err.Error(), nil)
	} else {
		view.RespondWithJSON(w, parser.ResponseCode, parser.ResponseMessage.Error(), parser.Data)
	}
}

// getMangaWithGenre is genre route to get MyAnimeList manga list having specific genre.
// Example: https://myanimelist.net/manga/genre/1/Action
func getMangaWithGenre(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))

	if page == 0 {
		page = 1
	}

	parser, err := MalService.GetMangaWithGenre(id, page)

	if err != nil {
		view.RespondWithJSON(w, parser.ResponseCode, err.Error(), nil)
	} else {
		view.RespondWithJSON(w, parser.ResponseCode, parser.ResponseMessage.Error(), parser.Data)
	}
}
