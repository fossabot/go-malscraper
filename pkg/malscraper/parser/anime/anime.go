package anime

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/pkg/malscraper/config"
	"github.com/rl404/go-malscraper/pkg/malscraper/constant"
	model "github.com/rl404/go-malscraper/pkg/malscraper/model/anime"
	"github.com/rl404/go-malscraper/pkg/malscraper/model/common"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser"
	"github.com/rl404/go-malscraper/pkg/malscraper/utils"
)

// AnimeParser is parser for MyAnimeList anime information.
// Example: https://myanimelist.net/anime/1
type AnimeParser struct {
	parser.BaseParser
	ID   int
	Data model.Anime
}

// InitAnimeParser to initiate all fields and data of AnimeParser.
func InitAnimeParser(config config.Config, id int) (animeParser AnimeParser, err error) {
	animeParser.ID = id
	animeParser.Config = config

	// Checking to redis if using redis in config.
	// Redis key's pattern is `anime:{id}`.
	redisKey := constant.RedisGetAnime + ":" + strconv.Itoa(animeParser.ID)
	if config.RedisClient != nil {
		found, err := utils.UnmarshalFromRedis(config.RedisClient, redisKey, &animeParser.Data)
		if err != nil {
			animeParser.SetResponse(500, err.Error())
			return animeParser, err
		}

		if found {
			animeParser.SetResponse(200, constant.SuccessMessage)
			return animeParser, nil
		}
	}

	// Get MyAnimeList HTML source page and initiate the parser.
	err = animeParser.InitParser("/anime/"+strconv.Itoa(animeParser.ID), "#content")
	if err != nil {
		return animeParser, err
	}

	// Fill in data field.
	animeParser.setDetails()

	// Save data field to redis if using redis in config.
	if config.RedisClient != nil {
		go utils.SaveToRedis(config.RedisClient, redisKey, animeParser.Data, config.CacheTime)
	}

	return animeParser, err
}

// setDetails to set anime details information.
func (ap *AnimeParser) setDetails() {
	ap.setID()
	ap.setCover()
	ap.setTitle()
	ap.setAltTitle()
	ap.setVideo()
	ap.setSynopsis()
	ap.setScore()
	ap.setVoter()
	ap.setRank()
	ap.setPopularity()
	ap.setMember()
	ap.setFavorite()
	ap.setOtherInfo()
	ap.setRelated()
	ap.setCharacter()
	ap.setStaff()
	ap.setSong()
	ap.setReview()
	ap.setRecommendation()
}

// setID to set anime id.
func (ap *AnimeParser) setID() {
	ap.Data.ID = ap.ID
}

// setCover to set anime cover image.
func (ap *AnimeParser) setCover() {
	image, _ := ap.Parser.Find("img.ac").Attr("data-src")
	ap.Data.Cover = utils.URLCleaner(image, "image", ap.Config.CleanImageURL)
}

// setTitle to set anime title.
func (ap *AnimeParser) setTitle() {
	ap.Data.Title = ap.Parser.Parent().Find("h1.h1 span").Text()
}

// setAltTitle to set anime alternative titles.
func (ap *AnimeParser) setAltTitle() {
	area := ap.Parser.Find(".js-scrollfix-bottom")

	ap.Data.AlternativeTitles.English = ap.getAltTitle(area, "English")
	ap.Data.AlternativeTitles.Synonym = ap.getAltTitle(area, "Synonyms")
	ap.Data.AlternativeTitles.Japanese = ap.getAltTitle(area, "Japanese")
}

// getAltTitle to get anime alternative titles (english, synonym, japanese).
func (ap *AnimeParser) getAltTitle(area *goquery.Selection, t string) string {
	altTitle, _ := area.Html()

	r := regexp.MustCompile(`(` + t + `:</span>)([^<]*)`)
	altTitle = r.FindString(altTitle)
	altTitle = strings.Replace(altTitle, t+":</span>", "", -1)

	return strings.TrimSpace(altTitle)
}

// setVideo to set anime promotion video.
func (ap *AnimeParser) setVideo() {
	video, _ := ap.Parser.Find(".video-promotion a").Attr("href")
	ap.Data.Video = utils.URLCleaner(video, "video", ap.Config.CleanVideoURL)
}

// setSynopsis to set anime synopsis.
func (ap *AnimeParser) setSynopsis() {
	synopsisArea := ap.Parser.Find("span[itemprop=description]")

	r := regexp.MustCompile(`\n[^\S\n]*`)
	synopsis := r.ReplaceAllString(synopsisArea.Text(), "\n")

	ap.Data.Synopsis = strings.TrimSpace(synopsis)
}

// setScore to set anime score.
func (ap *AnimeParser) setScore() {
	scoreArea := ap.Parser.Find("div[class=\"fl-l score\"]")
	score := strings.TrimSpace(scoreArea.Text())

	if score != "N/A" {
		ap.Data.Score = utils.StrToFloat(score)
	} else {
		ap.Data.Score = 0.0
	}
}

// setVoter to set number who vote the score.
func (ap *AnimeParser) setVoter() {
	voter, _ := ap.Parser.Find("div[class=\"fl-l score\"]").Attr("data-user")

	replacer := strings.NewReplacer("users", "", "user", "", ",", "")
	voter = replacer.Replace(voter)

	ap.Data.Voter = utils.StrToNum(voter)
}

// setRank to set anime rank.
func (ap *AnimeParser) setRank() {
	rank := ap.Parser.Find("span[class=\"numbers ranked\"] strong").Text()
	rank = strings.Replace(rank, "#", "", -1)

	if rank == "N/A" {
		rank = ""
	}

	ap.Data.Rank = utils.StrToNum(rank)
}

// setPopularity to set anime popularity rank.
func (ap *AnimeParser) setPopularity() {
	popularity := ap.Parser.Find("span[class=\"numbers popularity\"] strong").Text()
	popularity = strings.Replace(popularity, "#", "", -1)
	ap.Data.Popularity = utils.StrToNum(popularity)
}

// setMember to set anime number of member.
func (ap *AnimeParser) setMember() {
	member := ap.Parser.Find("span[class=\"numbers members\"] strong").Text()
	ap.Data.Member = utils.StrToNum(member)
}

// setFavorite to set anime number of favorite.
func (ap *AnimeParser) setFavorite() {
	favoriteArea := ap.Parser.Find("div[data-id=info2]").Next().Next().Next()
	favoriteArea.Find("span").Remove()
	ap.Data.Favorite = utils.StrToNum(favoriteArea.Text())
}

// setOtherInfo to set anime other details.
func (ap *AnimeParser) setOtherInfo() {
	ap.Parser.Find(".js-scrollfix-bottom").Find("h2").Each(func(i int, area *goquery.Selection) {
		if area.Text() == "Information" {
			area = area.Next()
			for {
				infoType := area.Find("span").First().Text()
				infoType = strings.ToLower(infoType)
				infoType = strings.Replace(infoType, ":", "", -1)

				if infoType == "type" {
					ap.Data.Type = ap.getCleanInfo(area)
				}

				if infoType == "episodes" {
					ap.Data.Episode = utils.StrToNum(ap.getCleanInfo(area))
				}

				if infoType == "status" {
					ap.Data.Status = ap.getCleanInfo(area)
				}

				if infoType == "premiered" {
					ap.Data.Premiered = ap.getCleanInfo(area)
				}

				if infoType == "broadcast" {
					ap.Data.Broadcast = ap.getCleanInfo(area)
				}

				if infoType == "source" {
					ap.Data.Source = ap.getCleanInfo(area)
				}

				if infoType == "duration" {
					ap.Data.Duration = ap.getCleanInfo(area)
				}

				if infoType == "rating" {
					ap.Data.Rating = ap.getCleanInfo(area)
				}

				if infoType == "aired" {
					infoValue := ap.getCleanInfo(area)
					ap.Data.StartDate.Start, ap.Data.StartDate.End = ap.getAiringInfo(infoValue)
				}

				if infoType == "producers" {
					infoValue := ap.getCleanInfo(area)
					ap.Data.Producers = ap.getIDNameInfo(area, infoType, infoValue)
				}

				if infoType == "licensors" {
					infoValue := ap.getCleanInfo(area)
					ap.Data.Licensors = ap.getIDNameInfo(area, infoType, infoValue)
				}

				if infoType == "studios" {
					infoValue := ap.getCleanInfo(area)
					ap.Data.Studios = ap.getIDNameInfo(area, infoType, infoValue)
				}

				if infoType == "genres" {
					infoValue := ap.getCleanInfo(area)
					ap.Data.Genres = ap.getIDNameInfo(area, infoType, infoValue)
				}

				area = area.Next()
				if goquery.NodeName(area) == "h2" || goquery.NodeName(area) == "br" {
					break
				}
			}
			return
		}
	})
}

// getCleanInfo to get anime clean details.
func (ap *AnimeParser) getCleanInfo(area *goquery.Selection) string {
	area.Find("span:nth-of-type(1)").Remove()

	replacer := strings.NewReplacer(", add some", "", "?", "", "Not yet aired", "", "Unknown", "")

	infoValue := area.Text()
	infoValue = strings.TrimSpace(infoValue)
	infoValue = replacer.Replace(infoValue)

	return infoValue
}

// getAiringInfo to get anime airing date.
func (ap *AnimeParser) getAiringInfo(infoValue string) (string, string) {
	if infoValue != "Not available" {
		splitDate := strings.Split(infoValue, " to ")
		if len(splitDate) > 1 {
			return splitDate[0], splitDate[1]
		}
		return splitDate[0], ""
	}
	return "", ""
}

// getIDNameInfo to get anime producer, licensor, studio, and genre.
func (ap *AnimeParser) getIDNameInfo(infoArea *goquery.Selection, infoType string, infoValue string) []common.IDName {
	var IDNameList []common.IDName
	if infoValue != "None found" {
		infoArea.Find("a").Each(func(i int, name *goquery.Selection) {
			link, _ := name.Attr("href")
			link = utils.GetValueFromSplit(link, "/", 3)
			IDNameList = append(IDNameList, common.IDName{
				ID:   utils.StrToNum(link),
				Name: name.Text(),
			})
		})
	}
	return IDNameList
}

// setRelated to set related anime or manga.
func (ap *AnimeParser) setRelated() {
	result := make(map[string][]model.Related)
	relatedArea := ap.Parser.Find(".anime_detail_related_anime")

	relatedArea.Find("tr").Each(func(i int, related *goquery.Selection) {
		var relatedList []model.Related

		relatedType := related.Find("td:nth-of-type(1)").Text()
		relatedType = strings.Replace(relatedType, ":", "", -1)
		relatedType = strings.Replace(relatedType, " ", "_", -1)
		relatedType = strings.TrimSpace(strings.ToLower(relatedType))

		relatedData := related.Find("td:nth-of-type(2)")
		relatedData.Find("a").Each(func(i int, data *goquery.Selection) {
			relatedLink, _ := data.Attr("href")
			splitLink := strings.Split(relatedLink, "/")

			relatedList = append(relatedList, model.Related{
				ID:    utils.StrToNum(splitLink[2]),
				Title: data.Text(),
				Type:  splitLink[1],
			})
		})
		result[relatedType] = relatedList
	})

	ap.Data.Related = result
}

// setCharacter to set anime related characters.
func (ap *AnimeParser) setCharacter() {
	var characters []model.AnimeCharacter
	area := ap.Parser.Find("div[class^=detail-characters-list]:nth-of-type(1)")

	var areaList []*goquery.Selection
	areaList = append(areaList, area.Find("div[class*=fl-l]"))
	areaList = append(areaList, area.Find("div[class*=fl-r]"))

	for _, characterSide := range areaList {
		characterSide.Find("table[width=\"100%\"]").Each(func(i int, character *goquery.Selection) {
			charArea := character.Find("tr td:nth-of-type(2)")
			vaArea := character.Find("td:nth-of-type(3) table td")

			characters = append(characters, model.AnimeCharacter{
				ID:      ap.getCharVaStaffID(charArea),
				Name:    ap.getCharVaStaffName(charArea),
				Role:    ap.getCharVaStaffRole(charArea),
				Image:   ap.getCharStaffImage(character),
				VaId:    ap.getCharVaStaffID(vaArea),
				VaName:  ap.getCharVaStaffName(vaArea),
				VaRole:  ap.getCharVaStaffRole(vaArea),
				VaImage: ap.getVaImage(character),
			})
		})
	}

	ap.Data.Characters = characters
}

// getCharVaStaffID to get anime character, va or staff id.
func (ap *AnimeParser) getCharVaStaffID(charArea *goquery.Selection) int {
	found, _ := charArea.Html()

	if found == "" {
		return 0
	}

	charLink, _ := charArea.Find("a").Attr("href")
	id := utils.GetValueFromSplit(charLink, "/", 4)
	return utils.StrToNum(id)
}

// getCharVaStaffName to get anime character, va or staff name.
func (ap *AnimeParser) getCharVaStaffName(charArea *goquery.Selection) string {
	r := regexp.MustCompile(`\s+`)
	charName := charArea.Find("a").Text()
	charName = r.ReplaceAllString(charName, " ")
	return strings.TrimSpace(charName)
}

// getCharVaStaffRole to get anime character, va or staff role.
func (ap *AnimeParser) getCharVaStaffRole(charArea *goquery.Selection) string {
	charRole := charArea.Find("small").Text()
	return strings.TrimSpace(charRole)
}

// getCharStaffImage to get anime character or staff image.
func (ap *AnimeParser) getCharStaffImage(eachCharacter *goquery.Selection) string {
	charImage, _ := eachCharacter.Find("tr td img").Attr("data-src")
	return utils.URLCleaner(charImage, "image", ap.Config.CleanImageURL)
}

// getVaImage to get anime va image.
func (ap *AnimeParser) getVaImage(eachCharacter *goquery.Selection) string {
	vaImage, _ := eachCharacter.Find("td:nth-of-type(3) table td:nth-of-type(2) img").Attr("data-src")
	return utils.URLCleaner(vaImage, "image", ap.Config.CleanImageURL)
}

// setStaff to set anime staff involved.
func (ap *AnimeParser) setStaff() {
	var staffList []model.Staff
	area := ap.Parser.Find("a[name=staff]").Next().Next()

	var areaList []*goquery.Selection
	areaList = append(areaList, area.Find("div[class*=fl-l]"))
	areaList = append(areaList, area.Find("div[class*=fl-r]"))

	for _, staffSide := range areaList {
		staffSide.Find("table[width=\"100%\"]").Each(func(i int, staff *goquery.Selection) {
			stArea := staff.Find("tr td:nth-of-type(2)")

			staffList = append(staffList, model.Staff{
				ID:    ap.getCharVaStaffID(stArea),
				Name:  ap.getCharVaStaffName(stArea),
				Role:  ap.getCharVaStaffRole(stArea),
				Image: ap.getCharStaffImage(staff),
			})
		})
	}

	ap.Data.Staff = staffList
}

// setSong to set anime opening and ending song.
func (ap *AnimeParser) setSong() {
	ap.Data.Song.Opening = ap.getCleanSong("div[class*=\"theme-songs opnening\"]")
	ap.Data.Song.Closing = ap.getCleanSong("div[class*=\"theme-songs ending\"]")
}

// getCleanSong to get clean song from div.
func (ap *AnimeParser) getCleanSong(div string) []string {
	var songs []string
	area := ap.Parser.Find(div)
	area.Find("span.theme-song").Each(func(i int, eachSong *goquery.Selection) {
		r := regexp.MustCompile(`#\d*:\s`)
		song := eachSong.Text()
		song = r.ReplaceAllString(song, " ")
		song = strings.TrimSpace(song)
		songs = append(songs, song)
	})
	return songs
}

// setReview to set anime review.
func (ap *AnimeParser) setReview() {
	var reviews []model.Review
	area := ap.Parser.Find(".js-scrollfix-bottom-rel table tr:nth-of-type(2)")
	area.Find(".borderDark[style^=padding]").Each(func(i int, review *goquery.Selection) {
		topArea := review.Find(".spaceit:nth-of-type(1)")
		bottomArea := topArea.Next()
		veryBottomArea := bottomArea.Next()

		reviews = append(reviews, model.Review{
			ID:       ap.getReviewID(veryBottomArea),
			Username: ap.getReviewUsername(topArea),
			Image:    ap.getReviewImage(topArea),
			Helpful:  ap.getReviewHelpful(topArea),
			Date:     ap.getReviewDate(topArea),
			Episode:  ap.getReviewProgress(topArea),
			Score:    ap.getReviewScore(bottomArea),
			Review:   ap.getReviewContent(bottomArea),
		})
	})

	ap.Data.Reviews = reviews
}

// getReviewID to get anime review id.
func (ap *AnimeParser) getReviewID(veryBottomArea *goquery.Selection) int {
	idLink, _ := veryBottomArea.Find("a").Attr("href")
	id := utils.GetValueFromSplit(idLink, "?id=", 1)
	return utils.StrToNum(id)
}

// getReviewUsername to get anime review username.
func (ap *AnimeParser) getReviewUsername(topArea *goquery.Selection) string {
	reviewUsername := topArea.Find("table td:nth-of-type(2)").Find("a").Text()
	return strings.Replace(reviewUsername, "All reviews", "", -1)
}

// getReviewImage to get anime review user image.
func (ap *AnimeParser) getReviewImage(topArea *goquery.Selection) string {
	reviewImage, _ := topArea.Find("table td img").Attr("src")
	return utils.URLCleaner(reviewImage, "image", ap.Config.CleanImageURL)
}

// getReviewHelpful to get anime review helpful number.
func (ap *AnimeParser) getReviewHelpful(topArea *goquery.Selection) int {
	reviewHelpful := topArea.Find("table td:nth-of-type(2) strong span[id^=rhelp]").Text()
	return utils.StrToNum(reviewHelpful)
}

// getReviewDate to get anime review date.
func (ap *AnimeParser) getReviewDate(topArea *goquery.Selection) common.DateTime {
	reviewDate := topArea.Find("div:nth-of-type(1)").Find("div:nth-of-type(1)")
	dateDate := reviewDate.Text()
	dateTime, _ := reviewDate.Attr("title")
	return common.DateTime{
		Date: dateDate,
		Time: dateTime,
	}
}

// getReviewProgress to get anime review episode.
func (ap *AnimeParser) getReviewProgress(topArea *goquery.Selection) string {
	reviewProgress := topArea.Find("div:nth-of-type(1)").Find("div:nth-of-type(2)").Text()
	replacer := strings.NewReplacer("episodes seen", "", "chapters read", "")
	return strings.TrimSpace(replacer.Replace(reviewProgress))
}

// getReviewScore to get anime review score.
func (ap *AnimeParser) getReviewScore(bottomArea *goquery.Selection) map[string]int {
	reviewScore := make(map[string]int)
	area := bottomArea.Find("table")
	area.Find("tr").Each(func(i int, score *goquery.Selection) {
		scoreType := strings.ToLower(score.Find("td:nth-of-type(1)").Text())
		reviewScore[scoreType] = utils.StrToNum(score.Find("td:nth-of-type(2)").Text())
	})
	return reviewScore
}

// getReviewContent to get anime review content.
func (ap *AnimeParser) getReviewContent(bottomArea *goquery.Selection) string {
	bottomArea.Find("div[id^=score]").Remove()
	bottomArea.Find("div[id^=revhelp_output]").Remove()
	bottomArea.Find("a[id^=reviewToggle]").Remove()

	r := regexp.MustCompile(`[^\S\r\n]+`)
	reviewContent := r.ReplaceAllString(bottomArea.Text(), " ")

	r = regexp.MustCompile(`(\n\n \n)`)
	reviewContent = r.ReplaceAllString(reviewContent, "")

	return strings.TrimSpace(reviewContent)
}

// setRecommendation to set anime recommendation.
func (ap *AnimeParser) setRecommendation() {
	var recommendations []model.AnimeRecommendation
	area := ap.Parser.Find("#anime_recommendation")
	area.Find("li.btn-anime").Each(func(i int, recommendation *goquery.Selection) {
		recommendations = append(recommendations, model.AnimeRecommendation{
			ID:    ap.getRecomID(recommendation),
			Title: ap.getRecomTitle(recommendation),
			Image: ap.getRecomImage(recommendation),
			Count: ap.getRecomCount(recommendation),
		})
	})

	ap.Data.Recommendations = recommendations
}

// getRecomID to get anime recommendation id.
func (ap *AnimeParser) getRecomID(recommendation *goquery.Selection) int {
	recomLink, _ := recommendation.Find("a").Attr("href")
	splitLink := strings.Split(recomLink, "/")

	if recommendation.Find(".users").Text() == "AutoRec" {
		return utils.StrToNum(splitLink[4])
	}

	splitLink = strings.Split(splitLink[5], "-")

	if splitLink[0] == strconv.Itoa(ap.ID) {
		return utils.StrToNum(splitLink[1])
	}
	return utils.StrToNum(splitLink[0])
}

// getRecomTitle to get anime recommendation title.
func (ap *AnimeParser) getRecomTitle(recommendation *goquery.Selection) string {
	return recommendation.Find("span:nth-of-type(1)").Text()
}

// getRecomImage to get anime recommendation image.
func (ap *AnimeParser) getRecomImage(recommendation *goquery.Selection) string {
	recomImage, _ := recommendation.Find("img").Attr("data-src")
	return utils.URLCleaner(recomImage, "image", ap.Config.CleanImageURL)
}

// getRecomCount to get anime user recommendation count.
func (ap *AnimeParser) getRecomCount(recommendation *goquery.Selection) int {
	recomCount := recommendation.Find(".users").Text()

	if recomCount == "AutoRec" {
		return 0
	}
	replacer := strings.NewReplacer("Users", "", "User", "")
	return utils.StrToNum(replacer.Replace(recomCount))
}
