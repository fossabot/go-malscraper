package constant

const (
	// MyAnimeListURL is MyAnimeList web base URL.
	MyAnimeListURL = "https://myanimelist.net"

	// SuccessMessage is a message for success response message.
	SuccessMessage = "success"
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
