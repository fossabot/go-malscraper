package constant

const (
	RedisGetAnime               = "anime"                // GetAnime(1) -> anime:1
	RedisGetAnimeVideo          = "anime-video"          // GetAnimeVideo(1, 2) -> anime-video:1,2
	RedisGetAnimeEpisode        = "anime-episode"        // GetAnimeEpisode(1, 2) -> anime-episode:1,2
	RedisGetAnimeReview         = "anime-review"         // GetAnimeReview(1, 2) -> anime-review:1,2
	RedisGetAnimeRecommendation = "anime-recommendation" // GetAnimeRecommendation(1) -> anime-recommendation:1
	RedisGetAnimeStats          = "anime-stats"          // GetAnimeStats(1, 2) -> anime-stats:1,2
	RedisGetAnimeCharacter      = "anime-character"      // GetAnimeCharacter(1) -> anime-character:1
	RedisGetAnimeStaff          = "anime-staff"          // GetAnimeStaff(1) -> anime-staff:1
	RedisGetAnimePicture        = "anime-picture"        // GetAnimePicture(1) -> anime-picture:1

	RedisGetManga               = "manga"                // GetManga(1) -> manga:1
	RedisGetMangaReview         = "manga-review"         // GetMangaReview(1, 2) -> manga-review:1,2
	RedisGetMangaRecommendation = "manga-recommendation" // GetMangaRecommendation(1) -> manga-recommendation:1
	RedisGetMangaStats          = "manga-stats"          // GetMangaStats(1, 2) -> manga-stats:1,2
	RedisGetMangaCharacter      = "manga-character"      // GetMangaCharacter(1) -> manga-character:1
	RedisGetMangaPicture        = "manga-picture"        // GetMangaPicture(1) -> manga-picture:1

	RedisGetCharacter        = "character"         // GetCharacter(1) -> character:1
	RedisGetCharacterPicture = "character-picture" // GetCharacterPicture(1) -> character-picture:1

	RedisGetPeople        = "people"         // GetPeople(1) -> people:1
	RedisGetPeoplePicture = "people-picture" // GetPeoplePicture(1) -> people-picture:1

	RedisGetProducers = "producers" // GetProducers() -> producers
	RedisGetProducer  = "producer"  // GetProducer(1, 2) -> producer:1,2

	RedisGetMagazines = "magazines" // GetMagazines() -> magazines
	RedisGetMagazine  = "magazine"  // GetMagazine(1, 2) -> magazine:1,2

	RedisGetGenres         = "genres"           // GetGenres(anime) -> genres:anime
	RedisGetAnimeWithGenre = "anime-with-genre" // GetAnimeWithGenre(1, 2) -> anime-with-genre:1,2
	RedisGetMangaWithGenre = "manga-with-genre" // GetMangaWithGenre(1, 2) -> manga-with-genre:1,2

	RedisGetReviews = "reviews" // GetReviews(anime, 2) -> reviews:anime,2
	RedisGetReview  = "review"  // GetReview(1) -> review:1

	RedisGetRecommendations = "recommendations" // GetRecommendations(anime, 2) -> recommendations:anime,2
	RedisGetRecommendation  = "recommendation"  // GetRecommendation(anime, 1, 2) -> recommendation:anime,1,2

	RedisGetUser               = "user"                // GetUser(rl404) -> user:rl404
	RedisGetUserFriend         = "user-friend"         // GetUserFriend(rl404, 2) -> user-friend:rl404,2
	RedisGetUserHistory        = "user-history"        // GetUserHistory(rl404, anime) -> user-history:rl404,anime
	RedisGetUserReview         = "user-review"         // GetUserReview(rl404, 2) -> user-review:rl404,2
	RedisGetUserRecommendation = "user-recommendation" // GetUserRecommendation(rl404, 2) -> user-recommendation:rl404,2

	RedisSearchCharacter = "search-character" // SearchCharacter(query, 2) -> search-character:query,2
	RedisSearchPeople    = "search-people"    // SearchPeople(query, 2) -> search-people:query,2
	RedisSearchUser      = "search-user"      // SearchUser(query, 2) -> search-user:query,2

	RedisGetSeason = "season" // GetSeason(2019, winter) -> season:2019,winter

	RedisGetTopAnime     = "top-anime"     // GetTopAnime(1, 2) -> top-anime:1,2
	RedisGetTopManga     = "top-manga"     // GetTopManga(1, 2) -> top-manga:1,2
	RedisGetTopCharacter = "top-character" // GetTopCharacter(1) -> top-character:1
	RedisGetTopPeople    = "top-people"    // GetTopPeople(1) -> top-people:1
)
