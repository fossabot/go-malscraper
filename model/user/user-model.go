package user

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/helper"
	"github.com/rl404/go-malscraper/model"
)

// UserModel is an extended model from MainModel for user.
type UserModel struct {
	model.MainModel
	User string
	Data UserData
}

// InitUserModel to initiate fields in parent (MainModel) model.
func (u *UserModel) InitUserModel(user string) (UserData, int, string) {
	u.User = user
	u.Url = "/profile/" + u.User
	u.ParserArea = "#content"
	u.InitModel()

	if u.ResponseCode != 200 {
		return u.Data, u.ResponseCode, u.ErrorMessage
	}

	u.SetAllDetail()

	return u.Data, u.ResponseCode, u.ErrorMessage
}

// SetAllDetail to fill all user profile data.
func (u *UserModel) SetAllDetail() {
	u.SetUsername()
	u.SetImage()
	u.SetStatus()
	u.SetMoreStatus()
	u.SetSns()
	u.SetFriend()
	u.SetAbout()
	u.SetStatistic("anime")
	u.SetStatistic("manga")
	u.SetFavorite()
}

// SetUsername to set username.
func (u *UserModel) SetUsername() {
	u.Data.Username = u.User
}

// SetImage to set user image.
func (u *UserModel) SetImage() {
	image := u.Parser.Find(".container-left .user-profile")
	image = image.Find(".user-image img")
	imageSrc, _ := image.Attr("src")
	u.Data.Image = helper.ImageUrlCleaner(imageSrc)
}

// SetStatus to set user basic status info (last online, gender, etc).
func (u *UserModel) SetStatus() {
	statusList := make(map[string]string)
	statusArea := u.Parser.Find(".container-left .user-profile .user-status:nth-of-type(1)")
	statusArea.Find("li").Each(func(i int, eachStatus *goquery.Selection) {
		statusType := eachStatus.Find("span:nth-of-type(1)").Text()
		statusValue := eachStatus.Find("span:nth-of-type(2)").Text()
		statusList[statusType] = statusValue
	})

	u.Data.LastOnline = statusList["Last Online"]
	u.Data.Gender = statusList["Gender"]
	u.Data.Birthday = statusList["Birthday"]
	u.Data.Location = statusList["Location"]
	u.Data.JoinedDate = statusList["Joined"]
}

// SetMoreStatus to set more user basic status info (forum post, review, etc).
func (u *UserModel) SetMoreStatus() {
	statusArea := u.Parser.Find(".container-left .user-profile .user-status:nth-of-type(3)")

	u.Data.ForumPost = statusArea.Find("li:nth-of-type(1)").Find("span:nth-of-type(2)").Text()
	u.Data.Review = statusArea.Find("li:nth-of-type(2)").Find("span:nth-of-type(2)").Text()
	u.Data.Recommendation = statusArea.Find("li:nth-of-type(3)").Find("span:nth-of-type(2)").Text()
	u.Data.BlogPost = statusArea.Find("li:nth-of-type(4)").Find("span:nth-of-type(2)").Text()
	u.Data.Club = statusArea.Find("li:nth-of-type(5)").Find("span:nth-of-type(2)").Text()
}

// SetSns to set user sns.
func (u *UserModel) SetSns() {
	snsArea := u.Parser.Find(".container-left .user-profile .user-profile-sns")
	snsArea.Find("a").Each(func(i int, eachSns *goquery.Selection) {
		snsClass, _ := eachSns.Attr("class")
		if snsClass != "di-ib mb8" {
			snsHref, _ := eachSns.Attr("href")
			u.Data.Sns = append(u.Data.Sns, snsHref)
		}
	})
}

// SetFriend to set user friend info.
func (u *UserModel) SetFriend() {
	friendArea := u.Parser.Find(".container-left .user-profile .user-friends")
	friendCount := friendArea.Prev().Find("a").Text()

	r, _ := regexp.Compile(`\(\d+\)`)
	friendCount = r.FindString(friendCount)
	friendCount = strings.Replace(friendCount, "(", "", -1)
	friendCount = strings.Replace(friendCount, ")", "", -1)
	u.Data.Friend.Count = friendCount

	friendArea.Find("a").Each(func(i int, eachFriend *goquery.Selection) {
		var tempFriend FriendData

		friendImage, _ := eachFriend.Attr("data-bg")

		tempFriend.Name = eachFriend.Text()
		tempFriend.Image = helper.ImageUrlCleaner(friendImage)

		u.Data.Friend.Data = append(u.Data.Friend.Data, tempFriend)
	})
}

// SetAbout to set user about-me.
func (u *UserModel) SetAbout() {
	aboutArea := u.Parser.Find(".container-right table tr td div[class=word-break]")
	aboutContent, _ := aboutArea.Html()
	u.Data.About = strings.TrimSpace(aboutContent)
}

// SetStatistic to set user anime & manga progress statistic.
func (u *UserModel) SetStatistic(t string) {
	rightArea := u.Parser.Find(".container-right")
	statArea := rightArea.Find(".user-statistics")

	if t == "anime" {
		statArea = statArea.Find("div[class=\"user-statistics-stats mt16\"]:nth-of-type(1)")
		scoreArea := statArea.Find(".stat-score")

		u.Data.AnimeStat.Days = u.GetDaysScore(scoreArea, 1)
		u.Data.AnimeStat.MeanScore = u.GetDaysScore(scoreArea, 2)
		u.Data.AnimeStat.Status, _ = u.GetStatStatus(statArea, t)
		u.Data.AnimeStat.History = u.GetHistory(rightArea, t)
	} else {
		statArea = statArea.Find("div[class=\"user-statistics-stats mt16\"]:nth-of-type(2)")
		scoreArea := statArea.Find(".stat-score")

		u.Data.MangaStat.Days = u.GetDaysScore(scoreArea, 1)
		u.Data.MangaStat.MeanScore = u.GetDaysScore(scoreArea, 2)
		_, u.Data.MangaStat.Status = u.GetStatStatus(statArea, t)
		u.Data.MangaStat.History = u.GetHistory(rightArea, t)
	}
}

// GetDaysScore to get days & score for anime & manga statistic.
func (u *UserModel) GetDaysScore(scoreArea *goquery.Selection, nth int) string {
	area := scoreArea.Find("div:nth-of-type(" + strconv.Itoa(nth) + ")")
	tempArea := area.Find("span")
	return strings.Replace(area.Text(), tempArea.Text(), "", -1)
}

// GetStatStatus to get anime & manga progress.
func (u *UserModel) GetStatStatus(statArea *goquery.Selection, t string) (AnimeStatus, MangaStatus) {
	var animeStatus AnimeStatus
	var mangaStatus MangaStatus

	statStatus := make(map[string]string)

	aStatArea := statArea.Find("ul.stats-status")
	bStatArea := statArea.Find("ul.stats-data")

	statStatus["completed"] = u.GetStatStatusCount(aStatArea, 2, 1)
	statStatus["on_hold"] = u.GetStatStatusCount(aStatArea, 3, 1)
	statStatus["dropped"] = u.GetStatStatusCount(aStatArea, 4, 1)
	statStatus["total"] = u.GetStatStatusCount(bStatArea, 1, 2)

	if t == "anime" {
		animeStatus.Completed = statStatus["completed"]
		animeStatus.OnHold = statStatus["on_hold"]
		animeStatus.Dropped = statStatus["dropped"]
		animeStatus.Total = statStatus["total"]
		animeStatus.Watching = u.GetStatStatusCount(aStatArea, 1, 1)
		animeStatus.PlanToWatch = u.GetStatStatusCount(aStatArea, 5, 1)
		animeStatus.Rewatched = u.GetStatStatusCount(bStatArea, 2, 2)
		animeStatus.Episode = u.GetStatStatusCount(bStatArea, 3, 2)
	} else {
		mangaStatus.Completed = statStatus["completed"]
		mangaStatus.OnHold = statStatus["on_hold"]
		mangaStatus.Dropped = statStatus["dropped"]
		mangaStatus.Total = statStatus["total"]
		mangaStatus.Reading = u.GetStatStatusCount(aStatArea, 1, 1)
		mangaStatus.PlanToRead = u.GetStatStatusCount(aStatArea, 5, 1)
		mangaStatus.Reread = u.GetStatStatusCount(bStatArea, 2, 2)
		mangaStatus.Chapter = u.GetStatStatusCount(bStatArea, 3, 2)
		mangaStatus.Volume = u.GetStatStatusCount(bStatArea, 4, 2)
	}
	return animeStatus, mangaStatus
}

// GetStatStatusCount to get anime & manga progress count.
func (u *UserModel) GetStatStatusCount(aStatArea *goquery.Selection, liNo int, spanNo int) string {
	countStat := aStatArea.Find("li:nth-of-type(" + strconv.Itoa(liNo) + ") span:nth-of-type(" + strconv.Itoa(spanNo) + ")").Text()
	return strings.Replace(countStat, ",", "", -1)
}

// GetHistory to get anime & manga progress history.
func (u *UserModel) GetHistory(rightArea *goquery.Selection, t string) []History {
	var history []History
	historyArea := rightArea.Find("div.updates." + t)
	historyArea.Find(".statistics-updates").Each(func(i int, eachHistory *goquery.Selection) {
		var tempHistory History

		historyDataArea := eachHistory.Find(".data")
		historyProgress := u.GetHistoryProgress(historyDataArea)

		tempHistory.Image = u.GetHistoryImage(eachHistory)
		tempHistory.Id = u.GetHistoryId(historyDataArea)
		tempHistory.Title = u.GetHistoryTitle(historyDataArea)
		tempHistory.Date = u.GetHistoryDate(historyDataArea)
		tempHistory.Progress = historyProgress["progress"]
		tempHistory.Score = historyProgress["score"]
		tempHistory.Status = historyProgress["status"]

		history = append(history, tempHistory)
	})
	return history
}

// GetHistoryImage to get anime & manga history image.
func (u *UserModel) GetHistoryImage(eachHistory *goquery.Selection) string {
	image := eachHistory.Find("img")
	imageSrc, _ := image.Attr("src")
	return helper.ImageUrlCleaner(imageSrc)
}

// GetHistoryId to get anime & manga history id.
func (u *UserModel) GetHistoryId(historyDataArea *goquery.Selection) string {
	dataId := historyDataArea.Find("a")
	hrefId, _ := dataId.Attr("href")
	id := strings.Split(hrefId, "/")
	return id[4]
}

// GetHistoryTitle to get anime & manga history title.
func (u *UserModel) GetHistoryTitle(historyDataArea *goquery.Selection) string {
	return historyDataArea.Find("a").Text()
}

// GetHistoryDate to get anime & manga history date.
func (u *UserModel) GetHistoryDate(historyDataArea *goquery.Selection) string {
	r, _ := regexp.Compile(`\d*$`)
	historyDate := historyDataArea.Find("span").Text()
	historyDate = r.ReplaceAllString(historyDate, "")
	return strings.TrimSpace(historyDate)
}

// GetHistoryProgress to get anime & manga history progress.
func (u *UserModel) GetHistoryProgress(historyDataArea *goquery.Selection) map[string]string {
	tempHistory := make(map[string]string)

	progress := historyDataArea.Find(".graph-content").Next().Text()

	r, _ := regexp.Compile(`([\s])+`)
	progress = r.ReplaceAllString(progress, " ")
	progress = strings.TrimSpace(progress)
	progressSplit := strings.Split(progress, "·")
	progressSplit2 := strings.Split(progressSplit[0], " ")

	if len(progressSplit2) > 3 {
		tempHistory["status"] = strings.ToLower(progressSplit[0])
		tempHistory["progress"] = "-"
	} else {
		tempHistory["status"] = strings.ToLower(progressSplit2[0])
		tempHistory["progress"] = progressSplit2[1]
	}

	tempHistory["score"] = strings.TrimSpace(strings.Replace(progressSplit[1], "Scored", "", -1))

	return tempHistory
}

// SetFavorite to set user favorite.
func (u *UserModel) SetFavorite() {
	favoriteArea := u.Parser.Find(".container-right .user-favorites-outer")

	u.Data.Favorite.Anime, _, _ = u.GetFavList(favoriteArea, "anime")
	u.Data.Favorite.Manga, _, _ = u.GetFavList(favoriteArea, "manga")
	_, u.Data.Favorite.Character, _ = u.GetFavList(favoriteArea, "characters")
	_, _, u.Data.Favorite.People = u.GetFavList(favoriteArea, "people")
}

// GetFavList to get anime, manga, character, and people favorite list.
func (u *UserModel) GetFavList(favoriteArea *goquery.Selection, t string) ([]FavAnimeManga, []FavCharacter, []FavPeople) {
	var favAnimeManga []FavAnimeManga
	var favCharacter []FavCharacter
	var favPeople []FavPeople

	favoriteArea = favoriteArea.Find("ul.favorites-list." + t)

	if favoriteArea.Text() != "" {
		favoriteArea.Find("li").Each(func(i int, eachFavorite *goquery.Selection) {
			if t == "anime" || t == "manga" {
				favAnimeManga = append(favAnimeManga, FavAnimeManga{
					Id:    u.GetFavId(eachFavorite),
					Image: u.GetFavImage(eachFavorite),
					Title: u.GetFavTitle(eachFavorite),
					Type:  u.GetFavTypeYear(eachFavorite, 0),
					Year:  u.GetFavTypeYear(eachFavorite, 1),
				})
			} else {
				if t == "characters" {
					favCharacter = append(favCharacter, FavCharacter{
						Id:         u.GetFavId(eachFavorite),
						Image:      u.GetFavImage(eachFavorite),
						Name:       u.GetFavTitle2(eachFavorite),
						MediaId:    u.GetFavMedia(eachFavorite, 2),
						MediaTitle: u.GetFavMediaTitle(eachFavorite),
						MediaType:  u.GetFavMedia(eachFavorite, 1),
					})
				} else {
					favPeople = append(favPeople, FavPeople{
						Id:    u.GetFavId(eachFavorite),
						Image: u.GetFavImage(eachFavorite),
						Name:  u.GetFavTitle(eachFavorite),
					})
				}
			}
		})
	}

	return favAnimeManga, favCharacter, favPeople
}

// GetFavImage to get favorite image.
func (u *UserModel) GetFavImage(eachFavorite *goquery.Selection) string {
	imageArea := eachFavorite.Find("a")
	imageStyle, _ := imageArea.Attr("style")

	r, _ := regexp.Compile(`\'([^\'])*`)
	imageUrl := r.FindString(imageStyle)

	return helper.ImageUrlCleaner(imageUrl[1:])
}

// GetFavId to get favorite id.
func (u *UserModel) GetFavId(eachFavorite *goquery.Selection) string {
	idArea := eachFavorite.Find("a")
	hrefId, _ := idArea.Attr("href")
	id := strings.Split(hrefId, "/")
	return id[4]
}

// GetFavTitle to get favorite title.
func (u *UserModel) GetFavTitle(eachFavorite *goquery.Selection) string {
	return eachFavorite.Find(".data a").Text()
}

// GetFavTypeYear to get favorite type and year.
func (u *UserModel) GetFavTypeYear(eachFavorite *goquery.Selection, i int) string {
	tempType := eachFavorite.Find("span").Text()
	favType := strings.Split(tempType, "·")
	return strings.TrimSpace(favType[i])
}

// GetFavTitle2 to get favorite character name.
func (u *UserModel) GetFavTitle2(eachFavorite *goquery.Selection) string {
	tempTitle := eachFavorite.Find(".data a").Text()
	favTitle := strings.Split(tempTitle, "\n")
	return strings.TrimSpace(favTitle[0])
}

// GetFavMedia to get favorite anime & manga id and type.
func (u *UserModel) GetFavMedia(eachFavorite *goquery.Selection, i int) string {
	tempMedia := eachFavorite.Find(".data .fn-grey2 a")
	mediaHref, _ := tempMedia.Attr("href")
	favMedia := strings.Split(mediaHref, "/")
	return favMedia[i]
}

// GetFavMediaTitle to get favorite anime & manga title.
func (u *UserModel) GetFavMediaTitle(eachFavorite *goquery.Selection) string {
	mediaTitle := eachFavorite.Find(".data .fn-grey2 a").Text()
	return strings.TrimSpace(mediaTitle)
}
