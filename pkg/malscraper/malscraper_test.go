package malscraper

import (
	"testing"
	"time"

	searchModel "github.com/rl404/go-malscraper/pkg/malscraper/model/search"
)

var (
	// Default config for all malscraper testing.
	// Change to your own testing needs.
	malConfigTest = MalConfig{
		UseDb: false,
	}

	// Malscraper service with default config.
	defaultMal *MalService

	// Malscraper service with user-defined config (malConfigTest above).
	mal *MalService

	// Sleep time between each test so we don't look like spamming their web.
	sleepTime = 5 * time.Second

	// infoIdTest is anime/manga id for testing.
	infoIdTest = 1

	// charIdTest is character id for testing.
	charIdTest = 1

	// peopleIdTest is people id for testing.
	peopleIdTest = 1

	// studioIdTest is studio id for testing.
	studioIdTest = 1

	// magazineIdTest is magazine id for testing.
	magazineIdTest = 1

	// animeGenreIdTest is anime genre id for testing.
	animeGenreIdTest = 1

	// mangaGenreIdTest is manga genre id for testing.
	mangaGenreIdTest = 1

	// animeRecommendationIdTest is anime recommendation id for testing.
	animeRecommendationIdTest = []int{1, 6}

	// mangaRecommendationIdTest is manga recommendation id for testing.
	mangaRecommendationIdTest = []int{1, 3}

	// reviewIdTest is review id for testing.
	reviewIdTest = 1

	// searchTest is query string for search testing.
	searchTest = "naruto"

	// Username for all user testing.
	userTest = "rl404"

	// user for user's review and recommendation testing.
	userTest2 = "Archaeon"

	// advSearchTest for anime & manga advanced search.
	advSearchTest = searchModel.Query{
		Query: "naruto",
	}
)

// TestDefaultMal to test creating default malscraper object.
func TestDefaultMal(t *testing.T) {
	defaultMal = Default()
}

// TestNewMal to test creating user-defined config malscraper object.
func TestNewMal(t *testing.T) {
	mal = New(malConfigTest)
}

// TestGetAnime to test parsing MyAnimeList anime main page.
func TestGetAnime(t *testing.T) {
	res, err := mal.GetAnime(infoIdTest)
	if err != nil {
		t.Errorf("GetAnime(%v) failed: %v", infoIdTest, err.Error())
	} else {
		t.Logf("GetAnime(%v) success: parsed anime %v", infoIdTest, res.Data.Title)
	}
	time.Sleep(sleepTime)
}

// TestGetAnimeVideo to test parsing MyAnimeList anime video page.
func TestGetAnimeVideo(t *testing.T) {
	res, err := mal.GetAnimeVideo(infoIdTest)
	if err != nil {
		t.Errorf("GetAnimeVideo(%v) failed: %v", infoIdTest, err.Error())
	} else {
		t.Logf("GetAnimeVideo(%v) success: parsed %v episodes %v promotions", infoIdTest, len(res.Data.Episodes), len(res.Data.Promotions))
	}
	time.Sleep(sleepTime)
}

// TestGetAnimeEpisode to test parsing MyAnimeList anime episode page.
func TestGetAnimeEpisode(t *testing.T) {
	res, err := mal.GetAnimeEpisode(infoIdTest)
	if err != nil {
		t.Errorf("GetAnimeEpisode(%v) failed: %v", infoIdTest, err.Error())
	} else {
		t.Logf("GetAnimeEpisode(%v) success: parsed %v episodes", infoIdTest, len(res.Data))
	}
	time.Sleep(sleepTime)
}

// TestGetAnimeReview to test parsing MyAnimeList anime review page.
func TestGetAnimeReview(t *testing.T) {
	res, err := mal.GetAnimeReview(infoIdTest)
	if err != nil {
		t.Errorf("GetAnimeReview(%v) failed: %v", infoIdTest, err.Error())
	} else {
		t.Logf("GetAnimeReview(%v) success: parsed %v reviews", infoIdTest, len(res.Data))
	}
	time.Sleep(sleepTime)
}

// TestGetAnimeRecommendation to test parsing MyAnimeList anime recommendation page.
func TestGetAnimeRecommendation(t *testing.T) {
	res, err := mal.GetAnimeRecommendation(infoIdTest)
	if err != nil {
		t.Errorf("GetAnimeRecommendation(%v) failed: %v", infoIdTest, err.Error())
	} else {
		t.Logf("GetAnimeRecommendation(%v) success: parsed %v recommendation", infoIdTest, len(res.Data))
	}
	time.Sleep(sleepTime)
}

// TestGetAnimeStats to test parsing MyAnimeList anime stats page.
func TestGetAnimeStats(t *testing.T) {
	_, err := mal.GetAnimeStats(infoIdTest)
	if err != nil {
		t.Errorf("GetAnimeStats(%v) failed: %v", infoIdTest, err.Error())
	} else {
		t.Logf("GetAnimeStats(%v) success: parsed anime stats", infoIdTest)
	}
	time.Sleep(sleepTime)
}

// TestGetAnimeCharacter to test parsing MyAnimeList anime character list page.
func TestGetAnimeCharacter(t *testing.T) {
	res, err := mal.GetAnimeCharacter(infoIdTest)
	if err != nil {
		t.Errorf("GetAnimeCharacter(%v) failed: %v", infoIdTest, err.Error())
	} else {
		t.Logf("GetAnimeCharacter(%v) success: parsed %v characters", infoIdTest, len(res.Data))
	}
	time.Sleep(sleepTime)
}

// TestGetAnimeStaff to test parsing MyAnimeList anime staff list page.
func TestGetAnimeStaff(t *testing.T) {
	res, err := mal.GetAnimeStaff(infoIdTest)
	if err != nil {
		t.Errorf("GetAnimeStaff(%v) failed: %v", infoIdTest, err.Error())
	} else {
		t.Logf("GetAnimeStaff(%v) success: parsed %v characters", infoIdTest, len(res.Data))
	}
	time.Sleep(sleepTime)
}

// TestGetAnimePicture to test parsing MyAnimeList anime pictures page.
func TestGetAnimePicture(t *testing.T) {
	res, err := mal.GetAnimePicture(infoIdTest)
	if err != nil {
		t.Errorf("GetAnimePicture(%v) failed: %v", infoIdTest, err.Error())
	} else {
		t.Logf("GetAnimePicture(%v) success: parsed %v pictures", infoIdTest, len(res.Data))
	}
	time.Sleep(sleepTime)
}

// TestGetCharacter to test parsing MyAnimeList character main page.
func TestGetCharacter(t *testing.T) {
	res, err := mal.GetCharacter(charIdTest)
	if err != nil {
		t.Errorf("GetCharacter(%v) failed: %v", charIdTest, err.Error())
	} else {
		t.Logf("GetCharacter(%v) success: parsed character %v", charIdTest, res.Data.Name)
	}
	time.Sleep(sleepTime)
}

// TestGetCharacterPicture to test parsing MyAnimeList character pictures page.
func TestGetCharacterPicture(t *testing.T) {
	res, err := mal.GetCharacterPicture(charIdTest)
	if err != nil {
		t.Errorf("GetCharacterPicture(%v) failed: %v", charIdTest, err.Error())
	} else {
		t.Logf("GetCharacterPicture(%v) success: parsed %v pictures", charIdTest, len(res.Data))
	}
	time.Sleep(sleepTime)
}

// TestGetAnimeGenres to test parsing MyAnimeList anime genre list.
func TestGetAnimeGenres(t *testing.T) {
	res, err := mal.GetGenres("anime")
	if err != nil {
		t.Errorf("GetGenres(\"anime\") failed: %v", err.Error())
	} else {
		t.Logf("GetGenres(\"anime\") success: parsed %v anime genres", len(res.Data))
	}
	time.Sleep(sleepTime)
}

// TestGetMangaGenres to test parsing MyAnimeList manga genre list.
func TestGetMangaGenres(t *testing.T) {
	res, err := mal.GetGenres("manga")
	if err != nil {
		t.Errorf("GetGenres(\"manga\") failed: %v", err.Error())
	} else {
		t.Logf("GetGenres(\"manga\") success: parsed %v manga genres", len(res.Data))
	}
	time.Sleep(sleepTime)
}

// TestGetAnimeGenre to test parsing MyAnimeList genre's anime list.
func TestGetAnimeGenre(t *testing.T) {
	res, err := mal.GetAnimeWithGenre(animeGenreIdTest)
	if err != nil {
		t.Errorf("GetAnimeWithGenre(%v) failed: %v", animeGenreIdTest, err.Error())
	} else {
		t.Logf("GetAnimeWithGenre(%v) success: parsed anime genre with %v anime", animeGenreIdTest, len(res.Data))
	}
	time.Sleep(sleepTime)
}

// TestGetMangaGenre to test parsing MyAnimeList genre's manga list.
func TestGetMangaGenre(t *testing.T) {
	res, err := mal.GetMangaWithGenre(mangaGenreIdTest)
	if err != nil {
		t.Errorf("GetMangaWithGenre(%v) failed: %v", mangaGenreIdTest, err.Error())
	} else {
		t.Logf("GetMangaWithGenre(%v) success: parsed anime genre with %v anime", mangaGenreIdTest, len(res.Data))
	}
	time.Sleep(sleepTime)
}

// TestGetMagazines to test parsing MyAnimeList magazine list.
func TestGetMagazines(t *testing.T) {
	res, err := mal.GetMagazines()
	if err != nil {
		t.Errorf("GetMagazines() failed: %v", err.Error())
	} else {
		t.Logf("GetMagazines() success: parsed %v magazines", len(res.Data))
	}
	time.Sleep(sleepTime)
}

// TestGetMagazine to test parsing MyAnimeList magazine's manga list.
func TestGetMagazine(t *testing.T) {
	res, err := mal.GetMagazine(magazineIdTest)
	if err != nil {
		t.Errorf("GetMagazine(%v) failed: %v", magazineIdTest, err.Error())
	} else {
		t.Logf("GetMagazine(%v) success: parsed magazine with %v manga", magazineIdTest, len(res.Data))
	}
	time.Sleep(sleepTime)
}

// TestGetManga to test parsing MyAnimeList manga main page.
func TestGetManga(t *testing.T) {
	res, err := mal.GetManga(infoIdTest)
	if err != nil {
		t.Errorf("GetManga(%v) failed: %v", infoIdTest, err.Error())
	} else {
		t.Logf("GetManga(%v) success: parsed manga %v", infoIdTest, res.Data.Title)
	}
	time.Sleep(sleepTime)
}

// TestGetMangaReview to test parsing MyAnimeList manga review page.
func TestGetMangaReview(t *testing.T) {
	res, err := mal.GetMangaReview(infoIdTest)
	if err != nil {
		t.Errorf("GetMangaReview(%v) failed: %v", infoIdTest, err.Error())
	} else {
		t.Logf("GetMangaReview(%v) success: parsed %v reviews", infoIdTest, len(res.Data))
	}
	time.Sleep(sleepTime)
}

// TestGetMangaRecommendation to test parsing MyAnimeList manga recommendation page.
func TestGetMangaRecommendation(t *testing.T) {
	res, err := mal.GetMangaRecommendation(infoIdTest)
	if err != nil {
		t.Errorf("GetMangaRecommendation(%v) failed: %v", infoIdTest, err.Error())
	} else {
		t.Logf("GetMangaRecommendation(%v) success: parsed %v recommendation", infoIdTest, len(res.Data))
	}
	time.Sleep(sleepTime)
}

// TestGetMangaStats to test parsing MyAnimeList manga stats page.
func TestGetMangaStats(t *testing.T) {
	_, err := mal.GetMangaStats(infoIdTest)
	if err != nil {
		t.Errorf("GetMangaStats(%v) failed: %v", infoIdTest, err.Error())
	} else {
		t.Logf("GetMangaStats(%v) success: parsed manga stats", infoIdTest)
	}
	time.Sleep(sleepTime)
}

// TestGetMangaCharacter to test parsing MyAnimeList manga character list page.
func TestGetMangaCharacter(t *testing.T) {
	res, err := mal.GetMangaCharacter(infoIdTest)
	if err != nil {
		t.Errorf("GetMangaCharacter(%v) failed: %v", infoIdTest, err.Error())
	} else {
		t.Logf("GetMangaCharacter(%v) success: parsed %v characters", infoIdTest, len(res.Data))
	}
	time.Sleep(sleepTime)
}

// TestGetMangaPicture to test parsing MyAnimeList manga pictures page.
func TestGetMangaPicture(t *testing.T) {
	res, err := mal.GetMangaPicture(infoIdTest)
	if err != nil {
		t.Errorf("GetMangaPicture(%v) failed: %v", infoIdTest, err.Error())
	} else {
		t.Logf("GetMangaPicture(%v) success: parsed %v pictures", infoIdTest, len(res.Data))
	}
	time.Sleep(sleepTime)
}

// TestGetPeople to test parsing MyAnimeList people main page.
func TestGetPeople(t *testing.T) {
	res, err := mal.GetPeople(peopleIdTest)
	if err != nil {
		t.Errorf("GetPeople(%v) failed: %v", peopleIdTest, err.Error())
	} else {
		t.Logf("GetPeople(%v) success: parsed people %v", peopleIdTest, res.Data.Name)
	}
	time.Sleep(sleepTime)
}

// TestGetPeoplePicture to test parsing MyAnimeList people pictures page.
func TestGetPeoplePicture(t *testing.T) {
	res, err := mal.GetPeoplePicture(peopleIdTest)
	if err != nil {
		t.Errorf("GetPeoplePicture(%v) failed: %v", peopleIdTest, err.Error())
	} else {
		t.Logf("GetPeoplePicture(%v) success: parsed %v pictures", peopleIdTest, len(res.Data))
	}
	time.Sleep(sleepTime)
}

// TestGetProducers to test parsing MyAnimeList studio/producer/licensor list.
func TestGetProducers(t *testing.T) {
	res, err := mal.GetProducers()
	if err != nil {
		t.Errorf("GetProducers() failed: %v", err.Error())
	} else {
		t.Logf("GetProducers() success: parsed %v producers", len(res.Data))
	}
	time.Sleep(sleepTime)
}

// TestGetProducer to test parsing MyAnimeList producer's anime list.
func TestGetProducer(t *testing.T) {
	res, err := mal.GetProducer(studioIdTest)
	if err != nil {
		t.Errorf("GetProducer(%v) failed: %v", studioIdTest, err.Error())
	} else {
		t.Logf("GetProducer(%v) success: parsed producer with %v anime", studioIdTest, len(res.Data))
	}
	time.Sleep(sleepTime)
}

// TestGetRecommendations to test parsing MyAnimeList anime/manga recommendation page.
func TestGetRecommendations(t *testing.T) {
	res, err := mal.GetRecommendations("anime")
	if err != nil {
		t.Errorf("GetRecommendations(\"anime\") failed: %v", err.Error())
	} else {
		t.Logf("GetRecommendations(\"anime\") success: parsed %v recommendations", len(res.Data))
	}
	time.Sleep(sleepTime)
}

// TestGetRecommendation to test parsing MyAnimeList anime/manga recommendation page.
func TestGetRecommendation(t *testing.T) {
	res, err := mal.GetRecommendation("anime", animeRecommendationIdTest[0], animeRecommendationIdTest[1])
	if err != nil {
		t.Errorf("GetRecommendation(\"anime\", %v, %v) failed: %v", animeRecommendationIdTest[0], animeRecommendationIdTest[1], err.Error())
	} else {
		t.Logf("GetRecommendation(\"anime\", %v, %v) success: parsed %v recommendation", animeRecommendationIdTest[0], animeRecommendationIdTest[1], res.Data.Source.Liked.Title)
	}
	time.Sleep(sleepTime)
}

// TestGetReview to test parsing MyAnimeList anime & manga review page.
func TestGetReview(t *testing.T) {
	res, err := mal.GetReview(reviewIdTest)
	if err != nil {
		t.Errorf("GetReview(%v) failed: %v", reviewIdTest, err.Error())
	} else {
		t.Logf("GetReview(%v) success: parsed %v review", reviewIdTest, res.Data.Source.Title)
	}
	time.Sleep(sleepTime)
}

// TestGetReviews to test parsing MyAnimeList anime & manga review list page.
func TestGetReviews(t *testing.T) {
	res, err := mal.GetReviews("anime")
	if err != nil {
		t.Errorf("GetReviews(%v) failed: %v", "anime", err.Error())
	} else {
		t.Logf("GetReviews(%v) success: parsed %v reviews", "anime", len(res.Data))
	}
	time.Sleep(sleepTime)
}

// TestSearchAnime to test parsing MyAnimeList anime search result list.
func TestSearchAnime(t *testing.T) {
	res, err := mal.SearchAnime(searchTest)
	if err != nil {
		t.Errorf("SearchAnime(\"%v\") failed: %v", searchTest, err.Error())
	} else {
		t.Logf("SearchAnime(\"%v\") success: parsed %v results", searchTest, len(res.Data))
	}
	time.Sleep(sleepTime)
}

// TestAdvSearchAnime to test parsing MyAnimeList anime advanced search result list.
func TestAdvSearchAnime(t *testing.T) {
	res, err := mal.AdvSearchAnime(advSearchTest)
	if err != nil {
		t.Errorf("AdvSearchAnime() failed: %v", err.Error())
	} else {
		t.Logf("AdvSearchAnime() success: parsed %v results", len(res.Data))
	}
	time.Sleep(sleepTime)
}

// TestSearchManga to test parsing MyAnimeList manga search result list.
func TestSearchManga(t *testing.T) {
	res, err := mal.SearchManga(searchTest)
	if err != nil {
		t.Errorf("SearchManga(\"%v\") failed: %v", searchTest, err.Error())
	} else {
		t.Logf("SearchManga(\"%v\") success: parsed %v results", searchTest, len(res.Data))
	}
	time.Sleep(sleepTime)
}

// TestAdvSearchManga to test parsing MyAnimeList manga advanced search result list.
func TestAdvSearchManga(t *testing.T) {
	res, err := mal.AdvSearchManga(advSearchTest)
	if err != nil {
		t.Errorf("AdvSearchManga() failed: %v", err.Error())
	} else {
		t.Logf("AdvSearchManga() success: parsed %v results", len(res.Data))
	}
	time.Sleep(sleepTime)
}

// TestSearchCharacter to test parsing MyAnimeList character search result list.
func TestSearchCharacter(t *testing.T) {
	res, err := mal.SearchCharacter(searchTest)
	if err != nil {
		t.Errorf("SearchCharacter(\"%v\") failed: %v", searchTest, err.Error())
	} else {
		t.Logf("SearchCharacter(\"%v\") success: parsed %v results", searchTest, len(res.Data))
	}
	time.Sleep(sleepTime)
}

// TestSearchPeople to test parsing MyAnimeList people search result list.
func TestSearchPeople(t *testing.T) {
	res, err := mal.SearchPeople(searchTest)
	if err != nil {
		t.Errorf("SearchPeople(\"%v\") failed: %v", searchTest, err.Error())
	} else {
		t.Logf("SearchPeople(\"%v\") success: parsed %v results", searchTest, len(res.Data))
	}
	time.Sleep(sleepTime)
}

// TestSearchUser to test parsing MyAnimeList people search result list.
func TestSearchUser(t *testing.T) {
	res, err := mal.SearchUser(searchTest)
	if err != nil {
		t.Errorf("SearchUser(\"%v\") failed: %v", searchTest, err.Error())
	} else {
		t.Logf("SearchUser(\"%v\") success: parsed %v results", searchTest, len(res.Data))
	}
	time.Sleep(sleepTime)
}

// TestGetSeason to test parsing MyAnimeList seasonal anime list.
func TestGetSeason(t *testing.T) {
	res, err := mal.GetSeason()
	if err != nil {
		t.Errorf("GetSeason() failed: %v", err.Error())
	} else {
		t.Logf("GetSeason() success: parsed %v anime", len(res.Data))
	}
	time.Sleep(sleepTime)
}

// TestGetTopAnime to test parsing MyAnimeList top anime list.
func TestGetTopAnime(t *testing.T) {
	res, err := mal.GetTopAnime()
	if err != nil {
		t.Errorf("GetTopAnime() failed: %v", err.Error())
	} else {
		t.Logf("GetTopAnime() success: parsed %v anime", len(res.Data))
	}
	time.Sleep(sleepTime)
}

// TestGetTopManga to test parsing MyAnimeList top manga list.
func TestGetTopManga(t *testing.T) {
	res, err := mal.GetTopManga()
	if err != nil {
		t.Errorf("GetTopManga() failed: %v", err.Error())
	} else {
		t.Logf("GetTopManga() success: parsed %v manga", len(res.Data))
	}
	time.Sleep(sleepTime)
}

// TestGetTopCharacter to test parsing MyAnimeList top character list.
func TestGetTopCharacter(t *testing.T) {
	res, err := mal.GetTopCharacter()
	if err != nil {
		t.Errorf("GetTopCharacter() failed: %v", err.Error())
	} else {
		t.Logf("GetTopCharacter() success: parsed %v characters", len(res.Data))
	}
	time.Sleep(sleepTime)
}

// TestGetTopPeople to test parsing MyAnimeList top people list.
func TestGetTopPeople(t *testing.T) {
	res, err := mal.GetTopPeople()
	if err != nil {
		t.Errorf("GetTopPeople() failed: %v", err.Error())
	} else {
		t.Logf("GetTopPeople() success: parsed %v people", len(res.Data))
	}
	time.Sleep(sleepTime)
}

// TestGetUser to test parsing MyAnimeList user profile.
func TestGetUser(t *testing.T) {
	res, err := mal.GetUser(userTest)
	if err != nil {
		t.Errorf("GetUser(\"%v\") failed: %v", userTest, err.Error())
	} else {
		t.Logf("GetUser(\"%v\") success: parsed %v profile", userTest, res.Data.Username)
	}
	time.Sleep(sleepTime)
}

// TestGetUserFriend to test parsing MyAnimeList user's friend list.
func TestGetUserFriend(t *testing.T) {
	res, err := mal.GetUserFriend(userTest)
	if err != nil {
		t.Errorf("GetUserFriend(\"%v\") failed: %v", userTest, err.Error())
	} else {
		t.Logf("GetUserFriend(\"%v\") success: parsed %v friends", userTest, len(res.Data))
	}
	time.Sleep(sleepTime)
}

// TestGetUserHistory to test parsing MyAnimeList user's history list.
func TestGetUserHistory(t *testing.T) {
	res, err := mal.GetUserHistory(userTest)
	if err != nil {
		t.Errorf("GetUserHistory(\"%v\") failed: %v", userTest, err.Error())
	} else {
		t.Logf("GetUserHistory(\"%v\") success: parsed %v histories", userTest, len(res.Data))
	}
	time.Sleep(sleepTime)
}

// TestGetUserReview to test parsing MyAnimeList user's review list.
func TestGetUserReview(t *testing.T) {
	res, err := mal.GetUserReview(userTest2)
	if err != nil {
		t.Errorf("GetUserReview(\"%v\") failed: %v", userTest2, err.Error())
	} else {
		t.Logf("GetUserReview(\"%v\") success: parsed %v reviews", userTest2, len(res.Data))
	}
	time.Sleep(sleepTime)
}

// TestGetUserRecommendation to test parsing MyAnimeList user's recommendation list.
func TestGetUserRecommendation(t *testing.T) {
	res, err := mal.GetUserRecommendation(userTest2)
	if err != nil {
		t.Errorf("GetUserRecommendation(\"%v\") failed: %v", userTest2, err.Error())
	} else {
		t.Logf("GetUserRecommendation(\"%v\") success: parsed %v recommendations", userTest2, len(res.Data))
	}
	time.Sleep(sleepTime)
}
