package constant

const (
	// MyAnimeListURL is MyAnimeList web base URL.
	MyAnimeListURL = "https://myanimelist.net"

	// SuccessCode is response code for success response.
	SuccessCode = 200

	// SuccessMessage is a message for success response message.
	SuccessMessage = "success"

	// BadRequestCode is reponse code when the request param is invalid.
	BadRequestCode = 400

	// InternalErrorCode is response code if there is something wrong
	// when processing the request.
	InternalErrorCode = 500
)

var (
	// MainType is a valid general types.
	MainType = []string{"anime", "manga"}

	// AnimeSeasons is anime season list.
	AnimeSeasons = []string{"winter", "spring", "summer", "fall"}

	// TopAnimeTypes is type list of top anime list.
	TopAnimeTypes = []string{"", "airing", "upcoming", "tv", "movie", "ova", "special", "bypopularity", "favorite"}

	// TopMangaTypes is type list of top manga list.
	TopMangaTypes = []string{"", "manga", "novels", "oneshots", "doujin", "manhwa", "manhua", "bypopularity", "favorite"}

	// ReviewTypes is type list of review list.
	ReviewTypes = []string{"anime", "manga", "bestvoted"}
)
