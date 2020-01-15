package controller

import (
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
	"github.com/rl404/go-malscraper/internal/view"
)

// registerUserRoutes registers all user routes.
func registerUserRoutes(router *chi.Mux) {
	router.Get("/user/{u}", getUser)
	router.Get("/user/{u}/friends", getUserFriend)
	router.Get("/user/{u}/history", getUserHistory)
	router.Get("/user/{u}/reviews", getUserReview)
	router.Get("/user/{u}/recommendations", getUserRecommendation)
	// router.Get("/user-list/{u}", ph.GetUserList)
	// router.Get("/user-list/{u}/{t}", ph.GetUserList)
	// router.Get("/user-list/{u}/{t}/{s}", ph.GetUserList)
	// router.Get("/user-cover/{u}", ph.GetUserCover)
	// router.Get("/user-cover/{u}/{t}", ph.GetUserCover)
	// router.Get("/user-cover/{u}/{t}/{q}", ph.GetUserCover)
}

// getUser is user route to get MyAnimeList user profile.
// Example: https://myanimelist.net/profile/rl404
func getUser(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "u")

	parser, err := MalService.GetUser(username)

	if err != nil {
		view.RespondWithJSON(w, parser.ResponseCode, err.Error(), nil)
	} else {
		view.RespondWithJSON(w, parser.ResponseCode, parser.ResponseMessage.Error(), parser.Data)
	}
}

// getUserFriend is user route to get MyAnimeList user's friend list.
// Example: https://myanimelist.net/profile/rl404/friends
func getUserFriend(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "u")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))

	if page == 0 {
		page = 1
	}

	parser, err := MalService.GetUserFriend(username, page)

	if err != nil {
		view.RespondWithJSON(w, parser.ResponseCode, err.Error(), nil)
	} else {
		view.RespondWithJSON(w, parser.ResponseCode, parser.ResponseMessage.Error(), parser.Data)
	}
}

// getUserHistory is user route to get MyAnimeList user's history list.
// Example: https://myanimelist.net/history/rl404
func getUserHistory(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "u")
	historyType := r.URL.Query().Get("type")

	parser, err := MalService.GetUserHistory(username, historyType)

	if err != nil {
		view.RespondWithJSON(w, parser.ResponseCode, err.Error(), nil)
	} else {
		view.RespondWithJSON(w, parser.ResponseCode, parser.ResponseMessage.Error(), parser.Data)
	}
}

// getUserReview is user route to get MyAnimeList user's review list.
// Example: https://myanimelist.net/profile/rl404/reviews
func getUserReview(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "u")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))

	if page == 0 {
		page = 1
	}

	parser, err := MalService.GetUserReview(username, page)

	if err != nil {
		view.RespondWithJSON(w, parser.ResponseCode, err.Error(), nil)
	} else {
		view.RespondWithJSON(w, parser.ResponseCode, parser.ResponseMessage.Error(), parser.Data)
	}
}

// getUserRecommendation is user route to get MyAnimeList user's recommendation list.
// Example: https://myanimelist.net/profile/rl404/recommendations
func getUserRecommendation(w http.ResponseWriter, r *http.Request) {
	username := chi.URLParam(r, "u")
	page, _ := strconv.Atoi(r.URL.Query().Get("page"))

	if page == 0 {
		page = 1
	}

	parser, err := MalService.GetUserRecommendation(username, page)

	if err != nil {
		view.RespondWithJSON(w, parser.ResponseCode, err.Error(), nil)
	} else {
		view.RespondWithJSON(w, parser.ResponseCode, parser.ResponseMessage.Error(), parser.Data)
	}
}
