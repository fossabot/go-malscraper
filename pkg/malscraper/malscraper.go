package malscraper

import (
	"github.com/rl404/go-malscraper/pkg/malscraper/parser/anime"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser/character"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser/genre"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser/magazine"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser/manga"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser/people"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser/producer"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser/recommendation"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser/review"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser/search"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser/season"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser/top"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser/user"

	"github.com/rl404/go-malscraper/pkg/malscraper/config"
	searchModel "github.com/rl404/go-malscraper/pkg/malscraper/model/search"
)

// MalService for all go-malscraper service.
type MalService struct {
	Config config.Config
}

// New creates new MalService object with user set config.
func New(malConfig config.Config) *MalService {
	malConfig.Init()
	return &MalService{
		Config: malConfig,
	}
}

// Default creates new MalService object with default config.
func Default() *MalService {
	return &MalService{
		Config: config.DefaultConfig,
	}
}

// GetAnime to get anime information.
func (m *MalService) GetAnime(id int) (anime.AnimeParser, error) {
	return anime.InitAnimeParser(m.Config, id)
}

// GetAnimeVideo to get anime's video list.
func (m *MalService) GetAnimeVideo(id int, page ...int) (anime.VideoParser, error) {
	return anime.InitVideoParser(m.Config, id, page...)
}

// GetAnimeEpisode to get anime's episode list.
func (m *MalService) GetAnimeEpisode(id int, page ...int) (anime.EpisodeParser, error) {
	return anime.InitEpisodeParser(m.Config, id, page...)
}

// GetAnimeReview to get anime's review list.
func (m *MalService) GetAnimeReview(id int, page ...int) (anime.ReviewParser, error) {
	return anime.InitReviewParser(m.Config, id, page...)
}

// GetAnimeRecommendation to get anime's recommendation list.
func (m *MalService) GetAnimeRecommendation(id int) (anime.RecommendationParser, error) {
	return anime.InitRecommendationParser(m.Config, id)
}

// GetAnimeStats to get anime's stats information.
func (m *MalService) GetAnimeStats(id int, page ...int) (anime.StatsParser, error) {
	return anime.InitStatsParser(m.Config, id, page...)
}

// GetAnimeCharacter to get anime's character list.
func (m *MalService) GetAnimeCharacter(id int) (anime.CharacterParser, error) {
	return anime.InitCharacterParser(m.Config, id)
}

// GetAnimeStaff to get anime's staff list.
func (m *MalService) GetAnimeStaff(id int) (anime.StaffParser, error) {
	return anime.InitStaffParser(m.Config, id)
}

// GetAnimePicture to get anime's picture list.
func (m *MalService) GetAnimePicture(id int) (anime.PictureParser, error) {
	return anime.InitPictureParser(m.Config, id)
}

// GetManga to get manga information.
func (m *MalService) GetManga(id int) (manga.MangaParser, error) {
	return manga.InitMangaParser(m.Config, id)
}

// GetMangaReview to get manga's review list.
func (m *MalService) GetMangaReview(id int, page ...int) (manga.ReviewParser, error) {
	return manga.InitReviewParser(m.Config, id, page...)
}

// GetMangaRecommendation to get manga's recommendation list.
func (m *MalService) GetMangaRecommendation(id int) (manga.RecommendationParser, error) {
	return manga.InitRecommendationParser(m.Config, id)
}

// GetMangaStats to get manga's stats information.
func (m *MalService) GetMangaStats(id int, page ...int) (manga.StatsParser, error) {
	return manga.InitStatsParser(m.Config, id, page...)
}

// GetMangaCharacter to get manga's character list.
func (m *MalService) GetMangaCharacter(id int) (manga.CharacterParser, error) {
	return manga.InitCharacterParser(m.Config, id)
}

// GetMangaPicture to get manga's picture list.
func (m *MalService) GetMangaPicture(id int) (manga.PictureParser, error) {
	return manga.InitPictureParser(m.Config, id)
}

// GetCharacter to get character information.
func (m *MalService) GetCharacter(id int) (character.CharacterParser, error) {
	return character.InitCharacterParser(m.Config, id)
}

// GetCharacterPicture to get character's picture list.
func (m *MalService) GetCharacterPicture(id int) (character.PictureParser, error) {
	return character.InitPictureParser(m.Config, id)
}

// GetPeople to get people information.
func (m *MalService) GetPeople(id int) (people.PeopleParser, error) {
	return people.InitPeopleParser(m.Config, id)
}

// GetPeoplePicture to get people's pictures list.
func (m *MalService) GetPeoplePicture(id int) (people.PeoplePictureParser, error) {
	return people.InitPeoplePictureParser(m.Config, id)
}

// GetProducers to get all producers, studios, and licensors.
func (m *MalService) GetProducers() (producer.ProducersParser, error) {
	return producer.InitProducersParser(m.Config)
}

// GetProducer to get producer's anime list.
func (m *MalService) GetProducer(id int, page ...int) (producer.ProducerParser, error) {
	return producer.InitProducerParser(m.Config, id, page...)
}

// GetMagazines to get all magazines, and serializations.
func (m *MalService) GetMagazines() (magazine.MagazinesParser, error) {
	return magazine.InitMagazinesParser(m.Config)
}

// GetMagazine to get magazine's manga list.
func (m *MalService) GetMagazine(id int, page ...int) (magazine.MagazineParser, error) {
	return magazine.InitMagazineParser(m.Config, id, page...)
}

// GetGenres to get all anime & manga genres.
func (m *MalService) GetGenres(gType string) (genre.GenresParser, error) {
	return genre.InitGenresParser(m.Config, gType)
}

// GetAnimeWithGenre to get anime list having specific genre.
func (m *MalService) GetAnimeWithGenre(id int, page ...int) (genre.AnimeParser, error) {
	return genre.InitAnimeParser(m.Config, id, page...)
}

// GetMangaWithGenre to get manga list having specific genre.
func (m *MalService) GetMangaWithGenre(id int, page ...int) (genre.MangaParser, error) {
	return genre.InitMangaParser(m.Config, id, page...)
}

// GetReviews to get anime/manga review list.
func (m *MalService) GetReviews(params ...interface{}) (review.ReviewsParser, error) {
	return review.InitReviewsParser(m.Config, params...)
}

// GetReview to get anime/manga review.
func (m *MalService) GetReview(id int) (review.ReviewParser, error) {
	return review.InitReviewParser(m.Config, id)
}

// GetRecommendations to get anime & manga recommendation list.
func (m *MalService) GetRecommendations(rType string, page ...int) (recommendation.RecommendationsParser, error) {
	return recommendation.InitRecommendationsParser(m.Config, rType, page...)
}

// GetRecommendation to get anime & manga's recommendation.
func (m *MalService) GetRecommendation(rType string, id1 int, id2 int) (recommendation.RecommendationParser, error) {
	return recommendation.InitRecommendationParser(m.Config, rType, id1, id2)
}

// GetUser to get user profile information.
func (m *MalService) GetUser(username string) (user.UserParser, error) {
	return user.InitUserParser(m.Config, username)
}

// GetUserFriend to get user friend list.
func (m *MalService) GetUserFriend(username string, page ...int) (user.UserFriendParser, error) {
	return user.InitUserFriendParser(m.Config, username, page...)
}

// GetUserHistory to get user history list.
func (m *MalService) GetUserHistory(username string, historyType ...string) (user.UserHistoryParser, error) {
	return user.InitUserHistoryParser(m.Config, username, historyType...)
}

// GetUserReview to get user review list.
func (m *MalService) GetUserReview(username string, page ...int) (user.ReviewParser, error) {
	return user.InitReviewParser(m.Config, username, page...)
}

// GetUserRecommendation to get user recommendation list.
func (m *MalService) GetUserRecommendation(username string, page ...int) (user.RecommendationParser, error) {
	return user.InitRecommendationParser(m.Config, username, page...)
}

// SearchAnime to simple search anime.
func (m *MalService) SearchAnime(query string, page ...int) (search.AnimeParser, error) {
	return search.InitAnimeParser(query, page...)
}

// AdvSearchAnime to advanced search anime.
func (m *MalService) AdvSearchAnime(queryObj searchModel.Query) (search.AnimeParser, error) {
	return search.InitAdvAnimeParser(queryObj)
}

// SearchManga to simple search manga.
func (m *MalService) SearchManga(query string, page ...int) (search.MangaParser, error) {
	return search.InitMangaParser(query, page...)
}

// AdvSearchManga to advanced search manga.
func (m *MalService) AdvSearchManga(queryObj searchModel.Query) (search.MangaParser, error) {
	return search.InitAdvMangaParser(queryObj)
}

// SearchCharacter to search character.
func (m *MalService) SearchCharacter(query string, page ...int) (search.CharacterParser, error) {
	return search.InitCharacterParser(m.Config, query, page...)
}

// SearchPeople to search character.
func (m *MalService) SearchPeople(query string, page ...int) (search.PeopleParser, error) {
	return search.InitPeopleParser(m.Config, query, page...)
}

// SearchUser to search user.
func (m *MalService) SearchUser(query string, page ...int) (search.UserParser, error) {
	return search.InitUserParser(m.Config, query, page...)
}

// GetSeason to get seasonal anime list.
func (m *MalService) GetSeason(params ...interface{}) (season.SeasonParser, error) {
	return season.InitSeasonParser(m.Config, params...)
}

// GetTopAnime to get top anime list.
func (m *MalService) GetTopAnime(params ...int) (top.AnimeParser, error) {
	return top.InitAnimeParser(m.Config, params...)
}

// GetTopManga to get top manga list.
func (m *MalService) GetTopManga(params ...int) (top.MangaParser, error) {
	return top.InitMangaParser(m.Config, params...)
}

// GetTopCharacter to get top character list.
func (m *MalService) GetTopCharacter(page ...int) (top.CharacterParser, error) {
	return top.InitCharacterParser(m.Config, page...)
}

// GetTopPeople to get top people list.
func (m *MalService) GetTopPeople(page ...int) (top.PeopleParser, error) {
	return top.InitPeopleParser(m.Config, page...)
}
