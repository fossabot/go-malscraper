package controller

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rl404/go-malscraper/internal/view"
)

// registerPeopleRoutes registers all people routes.
func registerPeopleRoutes(router *chi.Mux) {
	router.Get("/people/{id}", getPeople)
	router.Get("/people/{id}/pictures", getPeoplePicture)
}

// getPeople is people route to get MyAnimeList people information.
// Example: https://myanimelist.net/people/1/Tomokazu_Seki
func getPeople(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	parser, err := MalService.GetPeople(id)

	if err != nil {
		view.RespondWithJSON(w, parser.ResponseCode, err.Error(), nil)
	} else {
		view.RespondWithJSON(w, parser.ResponseCode, parser.ResponseMessage.Error(), parser.Data)
	}
}

// getPeoplePicture is people route to get MyAnimeList people's pictures list.
// Example: https://myanimelist.net/people/1/Tomokazu_Seki/pictures
func getPeoplePicture(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(chi.URLParam(r, "id"))

	parser, err := MalService.GetPeoplePicture(id)

	if err != nil {
		view.RespondWithJSON(w, parser.ResponseCode, err.Error(), nil)
	} else {
		view.RespondWithJSON(w, parser.ResponseCode, parser.ResponseMessage.Error(), parser.Data)
	}
}
