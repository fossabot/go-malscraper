package user

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/pkg/malscraper/config"
	"github.com/rl404/go-malscraper/pkg/malscraper/constant"
	model "github.com/rl404/go-malscraper/pkg/malscraper/model/user"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser"
	"github.com/rl404/go-malscraper/pkg/malscraper/utils"
)

// UserParser is parser for MyAnimeList user profile.
// Example: https://myanimelist.net/profile/rl404
type UserParser struct {
	parser.BaseParser
	Username string
	Data     model.User
}

// InitUserParser to initiate all fields and data of UserParser.
func InitUserParser(config config.Config, username string) (user UserParser, err error) {
	user.Username = username
	user.Config = config

	// Checking to redis if using redis in config.
	// Redis key's pattern is `user:{name}`.
	redisKey := constant.RedisGetUser + ":" + user.Username
	if config.RedisClient != nil {
		found, err := utils.UnmarshalFromRedis(config.RedisClient, redisKey, &user.Data)
		if err != nil {
			user.SetResponse(500, err.Error())
			return user, err
		}

		if found {
			user.SetResponse(200, constant.SuccessMessage)
			return user, nil
		}
	}

	// Get MyAnimeList HTML source page and initiate the parser.
	err = user.InitParser("/profile/"+user.Username, "#content")
	if err != nil {
		return user, err
	}

	// Fill in data field.
	user.setAllDetail()

	// Save data field to redis if using redis in config.
	if config.RedisClient != nil {
		go utils.SaveToRedis(config.RedisClient, redisKey, user.Data, config.CacheTime)
	}

	return user, nil
}

// setAllDetail to set all user detail information.
func (user *UserParser) setAllDetail() {
	user.setUsername()
	user.setImage()
	user.setInfo()
	user.setMoreInfo()
	user.setSns()
	user.setFriend()
	user.setAbout()
	user.setStat("anime")
	user.setStat("manga")
	user.setFavorite()
}

// setUsername to set user's username.
func (user *UserParser) setUsername() {
	user.Data.Username = user.Username
}

// setImage to set user's image.
func (user *UserParser) setImage() {
	imageSrc, _ := user.Parser.Find(".container-left .user-profile .user-image img").Attr("data-src")
	user.Data.Image = utils.URLCleaner(imageSrc, "image", user.Config.CleanImageURL)
}

// setInfo to set user's basic info (last online, gender, etc).
func (user *UserParser) setInfo() {
	infoList := make(map[string]string)
	infoArea := user.Parser.Find(".container-left .user-profile .user-status:nth-of-type(1)")
	infoArea.Find("li").Each(func(i int, info *goquery.Selection) {
		infoType := info.Find("span:nth-of-type(1)").Text()
		infoValue := info.Find("span:nth-of-type(2)").Text()
		infoList[infoType] = infoValue
	})

	user.Data.LastOnline = infoList["Last Online"]
	user.Data.Gender = infoList["Gender"]
	user.Data.Birthday = infoList["Birthday"]
	user.Data.Location = infoList["Location"]
	user.Data.JoinedDate = infoList["Joined"]
}

// setMoreInfo to set user's more info (forum post, review, etc).
func (user *UserParser) setMoreInfo() {
	infoArea := user.Parser.Find(".container-left .user-profile .user-status:nth-of-type(3)")

	user.Data.ForumPost = utils.StrToNum(infoArea.Find("li:nth-of-type(1)").Find("span:nth-of-type(2)").Text())
	user.Data.Review = utils.StrToNum(infoArea.Find("li:nth-of-type(2)").Find("span:nth-of-type(2)").Text())
	user.Data.Recommendation = utils.StrToNum(infoArea.Find("li:nth-of-type(3)").Find("span:nth-of-type(2)").Text())
	user.Data.BlogPost = utils.StrToNum(infoArea.Find("li:nth-of-type(4)").Find("span:nth-of-type(2)").Text())
	user.Data.Club = utils.StrToNum(infoArea.Find("li:nth-of-type(5)").Find("span:nth-of-type(2)").Text())
}

// setSns to set user's sns list.
func (user *UserParser) setSns() {
	snsArea := user.Parser.Find(".container-left .user-profile .user-profile-sns")
	snsArea.Find("a").Each(func(i int, sns *goquery.Selection) {
		snsClass, _ := sns.Attr("class")
		if snsClass != "di-ib mb8" {
			snsHref, _ := sns.Attr("href")
			user.Data.Sns = append(user.Data.Sns, snsHref)
		}
	})
}

// setFriend to set user's simple friends info.
func (user *UserParser) setFriend() {
	friendArea := user.Parser.Find(".container-left .user-profile .user-friends")
	friendCount := friendArea.Prev().Find("a").Text()

	r := regexp.MustCompile(`\(\d+\)`)
	friendCount = r.FindString(friendCount)

	replacer := strings.NewReplacer("(", "", ")", "")
	friendCount = replacer.Replace(friendCount)
	user.Data.Friend.Count = utils.StrToNum(friendCount)

	var friends []model.SimpleFriend
	friendArea.Find("a").Each(func(i int, friend *goquery.Selection) {
		friendImage, _ := friend.Attr("data-bg")

		friends = append(friends, model.SimpleFriend{
			Name:  friend.Text(),
			Image: utils.URLCleaner(friendImage, "image", user.Config.CleanImageURL),
		})
	})

	user.Data.Friend.Friends = friends
}

// setAbout to set user's about-me.
func (user *UserParser) setAbout() {
	aboutArea := user.Parser.Find(".container-right table tr td div[class=word-break]")
	aboutContent, _ := aboutArea.Html()
	user.Data.About = strings.TrimSpace(aboutContent)
}

// setStat to set user's anime & manga stats.
func (user *UserParser) setStat(t string) {
	rightArea := user.Parser.Find(".container-right")
	statArea := rightArea.Find(".user-statistics")

	if t == "anime" {
		statArea = statArea.Find("div[class=\"user-statistics-stats mt16\"]:nth-of-type(1)")
		scoreArea := statArea.Find(".stat-score")

		user.Data.AnimeStats.Days = user.getDaysScore(scoreArea, 1)
		user.Data.AnimeStats.MeanScore = user.getDaysScore(scoreArea, 2)
		user.Data.AnimeStats.Status, _ = user.getStatStatus(statArea, t)
		user.Data.AnimeStats.History = user.getHistory(rightArea, t)
	} else {
		statArea = statArea.Find("div[class=\"user-statistics-stats mt16\"]:nth-of-type(2)")
		scoreArea := statArea.Find(".stat-score")

		user.Data.MangaStats.Days = user.getDaysScore(scoreArea, 1)
		user.Data.MangaStats.MeanScore = user.getDaysScore(scoreArea, 2)
		_, user.Data.MangaStats.Status = user.getStatStatus(statArea, t)
		user.Data.MangaStats.History = user.getHistory(rightArea, t)
	}
}

// getDaysScore to get user's anime & manga total days and mean score.
func (user *UserParser) getDaysScore(scoreArea *goquery.Selection, nth int) float64 {
	area := scoreArea.Find("div:nth-of-type(" + strconv.Itoa(nth) + ")")
	area.Find("span").Remove()
	return utils.StrToFloat(area.Text())
}

// getStatStatus to get user's anime & manga stats count.
func (user *UserParser) getStatStatus(statArea *goquery.Selection, t string) (model.AnimeStatsCount, model.MangaStatsCount) {
	var animeStatus model.AnimeStatsCount
	var mangaStatus model.MangaStatsCount

	leftStatArea := statArea.Find("ul.stats-status")
	rightStatArea := statArea.Find("ul.stats-data")

	if t == "anime" {
		animeStatus.Completed = user.getStatStatusCount(leftStatArea, 2, 1)
		animeStatus.OnHold = user.getStatStatusCount(leftStatArea, 3, 1)
		animeStatus.Dropped = user.getStatStatusCount(leftStatArea, 4, 1)
		animeStatus.Total = user.getStatStatusCount(rightStatArea, 1, 2)
		animeStatus.Watching = user.getStatStatusCount(leftStatArea, 1, 1)
		animeStatus.PlanToWatch = user.getStatStatusCount(leftStatArea, 5, 1)
		animeStatus.Rewatched = user.getStatStatusCount(rightStatArea, 2, 2)
		animeStatus.Episode = user.getStatStatusCount(rightStatArea, 3, 2)
	} else {
		mangaStatus.Completed = user.getStatStatusCount(leftStatArea, 2, 1)
		mangaStatus.OnHold = user.getStatStatusCount(leftStatArea, 3, 1)
		mangaStatus.Dropped = user.getStatStatusCount(leftStatArea, 4, 1)
		mangaStatus.Total = user.getStatStatusCount(rightStatArea, 1, 2)
		mangaStatus.Reading = user.getStatStatusCount(leftStatArea, 1, 1)
		mangaStatus.PlanToRead = user.getStatStatusCount(leftStatArea, 5, 1)
		mangaStatus.Reread = user.getStatStatusCount(rightStatArea, 2, 2)
		mangaStatus.Chapter = user.getStatStatusCount(rightStatArea, 3, 2)
		mangaStatus.Volume = user.getStatStatusCount(rightStatArea, 4, 2)
	}
	return animeStatus, mangaStatus
}

// getStatStatusCount to get user's each anime & manga status count.
func (user *UserParser) getStatStatusCount(statArea *goquery.Selection, liNo int, spanNo int) int {
	countStat := statArea.Find("li:nth-of-type(" + strconv.Itoa(liNo) + ") span:nth-of-type(" + strconv.Itoa(spanNo) + ")").Text()
	return utils.StrToNum(countStat)
}

// getHistory to get user's anime & manga progress history.
func (user *UserParser) getHistory(historyArea *goquery.Selection, t string) []model.History {
	var histories []model.History
	historyArea = historyArea.Find("div.updates." + t)
	historyArea.Find(".statistics-updates").Each(func(i int, history *goquery.Selection) {
		historyDataArea := history.Find(".data")
		historyProgress := user.getHistoryProgress(historyDataArea)

		histories = append(histories, model.History{
			ID:       user.getHistoryID(historyDataArea),
			Title:    user.getHistoryTitle(historyDataArea),
			Image:    user.getHistoryImage(history),
			Date:     user.getHistoryDate(historyDataArea),
			Progress: historyProgress["progress"].(string),
			Score:    historyProgress["score"].(int),
			Status:   historyProgress["status"].(string),
		})
	})
	return histories
}

// getHistoryID to get user's anime & manga history id.
func (user *UserParser) getHistoryID(historyDataArea *goquery.Selection) int {
	href, _ := historyDataArea.Find("a").Attr("href")
	id := utils.GetValueFromSplit(href, "/", 4)
	return utils.StrToNum(id)
}

// getHistoryTitle to get user's anime & manga history title.
func (user *UserParser) getHistoryTitle(historyDataArea *goquery.Selection) string {
	return historyDataArea.Find("a").Text()
}

// getHistoryImage to get user's anime & manga history image.
func (user *UserParser) getHistoryImage(history *goquery.Selection) string {
	imageSrc, _ := history.Find("img").Attr("data-src")
	return utils.URLCleaner(imageSrc, "image", user.Config.CleanImageURL)
}

// getHistoryDate to get user's anime & manga history date.
func (user *UserParser) getHistoryDate(historyDataArea *goquery.Selection) string {
	r := regexp.MustCompile(`\d*$`)
	historyDate := historyDataArea.Find("span").Text()
	historyDate = r.ReplaceAllString(historyDate, "")
	return strings.TrimSpace(historyDate)
}

// getHistoryProgress to get user's anime & manga history progress.
func (user *UserParser) getHistoryProgress(historyDataArea *goquery.Selection) map[string]interface{} {
	tempHistory := make(map[string]interface{})

	progress := historyDataArea.Find(".graph-content").Next().Text()

	r := regexp.MustCompile(`([\s])+`)
	progress = r.ReplaceAllString(progress, " ")
	progress = strings.TrimSpace(progress)
	progressSplit := strings.Split(progress, "·")

	progressSplit2 := strings.Split(progressSplit[0], " ")

	if len(progressSplit2) > 3 {
		tempHistory["status"] = strings.TrimSpace(strings.ToLower(progressSplit[0]))
		tempHistory["progress"] = "-"
	} else {
		tempHistory["status"] = strings.TrimSpace(strings.ToLower(progressSplit2[0]))
		tempHistory["progress"] = progressSplit2[1]
	}

	scoreStr := strings.TrimSpace(strings.Replace(progressSplit[1], "Scored", "", -1))
	if scoreStr == "-" {
		scoreStr = "0"
	}

	tempHistory["score"], _ = strconv.Atoi(scoreStr)

	return tempHistory
}

// setFavorite to set user's favorite anime, manga, character and people.
func (user *UserParser) setFavorite() {
	favoriteArea := user.Parser.Find(".container-right .user-favorites-outer")

	user.Data.Favorite.Anime, _, _ = user.getFavList(favoriteArea, "anime")
	user.Data.Favorite.Manga, _, _ = user.getFavList(favoriteArea, "manga")
	_, user.Data.Favorite.Character, _ = user.getFavList(favoriteArea, "characters")
	_, _, user.Data.Favorite.People = user.getFavList(favoriteArea, "people")
}

// getFavList to get user's favorite anime, manga, character, and people.
func (user *UserParser) getFavList(favoriteArea *goquery.Selection, t string) ([]model.FavAnimeManga, []model.FavCharacter, []model.FavPeople) {
	var favAnimeManga []model.FavAnimeManga
	var favCharacter []model.FavCharacter
	var favPeople []model.FavPeople

	favoriteArea = favoriteArea.Find("ul.favorites-list." + t)

	if favoriteArea.Text() != "" {
		favoriteArea.Find("li").Each(func(i int, favorite *goquery.Selection) {
			if t == "anime" || t == "manga" {
				favAnimeManga = append(favAnimeManga, model.FavAnimeManga{
					ID:    user.getFavID(favorite),
					Title: user.getFavTitle(favorite),
					Image: user.getFavImage(favorite),
					Type:  user.getFavTypeYear(favorite, 0).(string),
					Year:  user.getFavTypeYear(favorite, 1).(int),
				})
			} else {
				if t == "characters" {
					favCharacter = append(favCharacter, model.FavCharacter{
						ID:          user.getFavID(favorite),
						Name:        user.getFavTitle2(favorite),
						Image:       user.getFavImage(favorite),
						SourceID:    user.getFavMedia(favorite, 2).(int),
						SourceTitle: user.getFavMediaTitle(favorite),
						SourceType:  user.getFavMedia(favorite, 1).(string),
					})
				} else {
					favPeople = append(favPeople, model.FavPeople{
						ID:    user.getFavID(favorite),
						Name:  user.getFavTitle(favorite),
						Image: user.getFavImage(favorite),
					})
				}
			}
		})
	}

	return favAnimeManga, favCharacter, favPeople
}

// getFavID to get user's favorite anime, manga, character and people id.
func (user *UserParser) getFavID(favorite *goquery.Selection) int {
	href, _ := favorite.Find("a").Attr("href")
	id := utils.GetValueFromSplit(href, "/", 4)
	return utils.StrToNum(id)
}

// getFavTitle to get user's favorite anime, manga, character and people name/title.
func (user *UserParser) getFavTitle(favorite *goquery.Selection) string {
	return favorite.Find(".data a").Text()
}

// getFavImage to get user's favorite anime, manga, character and people image.
func (user *UserParser) getFavImage(favorite *goquery.Selection) string {
	image, _ := favorite.Find("img").First().Attr("data-src")
	return utils.URLCleaner(image, "image", user.Config.CleanImageURL)
}

// getFavTypeYear to get user's favorite anime, and manga type and year.
func (user *UserParser) getFavTypeYear(favorite *goquery.Selection, i int) interface{} {
	favType := favorite.Find("span").Text()
	favType = utils.GetValueFromSplit(favType, "·", i)
	typeYear := strings.TrimSpace(favType)
	if i == 1 {
		return utils.StrToNum(typeYear)
	}
	return typeYear
}

// getFavTitle2 to get user's favorite character name.
func (user *UserParser) getFavTitle2(favorite *goquery.Selection) string {
	favTitle := favorite.Find(".data a").Text()
	favTitle = utils.GetValueFromSplit(favTitle, "\n", 0)
	return strings.TrimSpace(favTitle)
}

// getFavMedia to get user's favorite caharacter's anime & manga id and type.
func (user *UserParser) getFavMedia(favorite *goquery.Selection, i int) interface{} {
	mediaHref, _ := favorite.Find(".data .fn-grey2 a").Attr("href")
	favMedia := utils.GetValueFromSplit(mediaHref, "/", i)
	if i == 2 {
		return utils.StrToNum(favMedia)
	}
	return favMedia
}

// getFavMediaTitle to get user's favorite character's anime & manga title.
func (user *UserParser) getFavMediaTitle(favorite *goquery.Selection) string {
	mediaTitle := favorite.Find(".data .fn-grey2 a").Text()
	return strings.TrimSpace(mediaTitle)
}
