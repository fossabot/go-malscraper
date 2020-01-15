package controller

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rl404/go-malscraper/internal/view"
)

// registerCharacterRoutes registers all character routes.
func registerCharacterRoutes(router *chi.Mux) {
	router.Get("/character/{id}", getCharacter)
	router.Get("/character/{id}/pictures", getCharacterPicture)
}

// getCharacter is character route to get MyAnimeList character information.
// Example: https://myanimelist.net/character/1/Spike_Spiegel
func getCharacter(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	parser, err := MalService.GetCharacter(id)

	if err != nil {
		view.RespondWithJSON(w, parser.ResponseCode, err.Error(), nil)
	} else {
		view.RespondWithJSON(w, parser.ResponseCode, parser.ResponseMessage.Error(), parser.Data)
	}
}

// getCharacterPicture is character route to get MyAnimeList character's picture list.
// Example: https://myanimelist.net/character/1/Spike_Spiegel/pictures
func getCharacterPicture(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	parser, err := MalService.GetCharacterPicture(id)

	if err != nil {
		view.RespondWithJSON(w, parser.ResponseCode, err.Error(), nil)
	} else {
		view.RespondWithJSON(w, parser.ResponseCode, parser.ResponseMessage.Error(), parser.Data)
	}
}
