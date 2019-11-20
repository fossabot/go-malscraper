package general

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/helper"
	"github.com/rl404/go-malscraper/model"
)

// InfoModel is an extended model from MainModel for user.
type InfoModel struct {
	model.MainModel
	Type string
	Id   int
	Data InfoData
}

// InitInfoModel to initiate fields in parent (MainModel) model.
func (i *InfoModel) InitInfoModel(t string, id int) (InfoData, int, string) {
	i.Type = t
	i.Id = id
	i.InitModel("/"+i.Type+"/"+strconv.Itoa(i.Id), "#content")

	i.Parser = i.GetParser(i.Url)

	if i.ResponseCode != 200 {
		return i.Data, i.ResponseCode, i.ErrorMessage
	}

	i.SetAllDetail()

	return i.Data, i.ResponseCode, i.ErrorMessage
}

// SetAllDetail to fill all anime & manga data.
func (i *InfoModel) SetAllDetail() {
	i.SetId()
	i.SetCover()
	i.SetTitle()
	i.SetTitle2()
	i.SetVideo()
	i.SetSynopsis()
	i.SetScore()
	i.SetVoter()
	i.SetRank()
	i.SetPopularity()
	i.SetMembers()
	i.SetFavorite()
	i.SetOtherInfo()
	i.SetRelated()
	i.SetCharacter()
	i.SetStaff()
	i.SetSong()
	i.SetReview()
	i.SetRecommendation()
}

// SetId to set anime or manga id.
func (i *InfoModel) SetId() {
	i.Data.Id = i.Id
}

// SetCover to set anime or manga cover image.
func (i *InfoModel) SetCover() {
	imageSrc, _ := i.Parser.Find("img.ac").Attr("src")
	i.Data.Cover = helper.ImageUrlCleaner(imageSrc)
}

// SetTitle to set anime or manga title.
func (i *InfoModel) SetTitle() {
	title := i.Parser.Find("h1.h1 span")
	i.Data.Title = title.Text()
}

// SetTitle2 to set anime or manga alternative titles.
func (i *InfoModel) SetTitle2() {
	infoArea := i.Parser.Find(".js-scrollfix-bottom")

	i.Data.Title2.English = i.GetTitle2(infoArea, "English")
	i.Data.Title2.Synonym = i.GetTitle2(infoArea, "Synonyms")
	i.Data.Title2.Japanese = i.GetTitle2(infoArea, "Japanese")
}

// GetTitle2 to get each anime or manga alternative titles.
func (i *InfoModel) GetTitle2(infoArea *goquery.Selection, t string) string {
	title2, _ := infoArea.Html()

	r, _ := regexp.Compile(`(` + t + `:</span>)([^<]*)`)
	title3 := r.FindString(title2)
	title3 = strings.Replace(title3, t+":</span>", "", -1)

	return strings.TrimSpace(title3)
}

// SetVideo to set anime or manga video url.
func (i *InfoModel) SetVideo() {
	videoUrl, _ := i.Parser.Find(".video-promotion a").Attr("href")
	i.Data.Video = helper.VideoUrlCleaner(videoUrl)
}

// SetSynopsis to set anime or manga synopsis.
func (i *InfoModel) SetSynopsis() {
	synopsisArea := i.Parser.Find("span[itemprop=description]")

	r, _ := regexp.Compile(`\n[^\S\n]*`)
	synopsis := r.ReplaceAllString(synopsisArea.Text(), "\n")

	i.Data.Synopsis = strings.TrimSpace(synopsis)
}

// SetScore to set anime or manga score.
func (i *InfoModel) SetScore() {
	scoreArea := i.Parser.Find("div[class=\"fl-l score\"]")
	score := strings.TrimSpace(scoreArea.Text())

	if score != "N/A" {
		i.Data.Score, _ = strconv.ParseFloat(score, 64)
	} else {
		i.Data.Score = 0
	}
}

// SetVoter to set number of user who vote the score.
func (i *InfoModel) SetVoter() {
	voter, _ := i.Parser.Find("div[class=\"fl-l score\"]").Attr("data-user")

	voter = strings.Replace(voter, "users", "", -1)
	voter = strings.Replace(voter, "user", "", -1)
	voter = strings.Replace(voter, ",", "", -1)

	i.Data.Voter, _ = strconv.Atoi(strings.TrimSpace(voter))
}

// SetRank to set rank of the anime or manga.
func (i *InfoModel) SetRank() {
	rank := i.Parser.Find("span[class=\"numbers ranked\"] strong").Text()

	if rank == "N/A" {
		rank = ""
	}

	i.Data.Rank, _ = strconv.Atoi(strings.Replace(rank, "#", "", -1))
}

// SetPopularity to set popularity rank of the anime or manga.
func (i *InfoModel) SetPopularity() {
	popularity := i.Parser.Find("span[class=\"numbers popularity\"] strong").Text()
	i.Data.Popularity, _ = strconv.Atoi(strings.Replace(popularity, "#", "", -1))
}

// SetMembers to set number of member of the anime or manga.
func (i *InfoModel) SetMembers() {
	member := i.Parser.Find("span[class=\"numbers members\"] strong").Text()
	i.Data.Members, _ = strconv.Atoi(strings.Replace(member, ",", "", -1))
}

// SetFavorite to set number of user who favorite the anime or manga.
func (i *InfoModel) SetFavorite() {
	favoriteArea := i.Parser.Find("div[data-id=info2]").Next().Next().Next()
	favoriteTitle := favoriteArea.Find("span").Text()
	favorite := favoriteArea.Text()
	favorite = strings.Replace(favorite, favoriteTitle, "", -1)
	favorite = strings.Replace(favorite, ",", "", -1)

	i.Data.Favorite, _ = strconv.Atoi(strings.TrimSpace(favorite))
}

// SetOtherInfo to set other detail anime or manga info.
func (i *InfoModel) SetOtherInfo() {
	i.Parser.Find(".js-scrollfix-bottom").Find("h2").Each(func(j int, infoArea *goquery.Selection) {
		if infoArea.Text() == "Information" {
			infoArea = infoArea.Next()
			for true {
				infoType := infoArea.Find("span").Text()
				infoType = strings.ToLower(infoType)
				infoType = strings.Replace(infoType, ":", "", -1)

				if infoType == "type" {
					i.Data.Type = i.GetCleanInfo(infoArea)
				}

				if infoType == "episodes" {
					i.Data.Episodes, _ = strconv.Atoi(i.GetCleanInfo(infoArea))
				}

				if infoType == "volumes" {
					i.Data.Volumes, _ = strconv.Atoi(i.GetCleanInfo(infoArea))
				}

				if infoType == "chapters" {
					i.Data.Chapters, _ = strconv.Atoi(i.GetCleanInfo(infoArea))
				}

				if infoType == "status" {
					i.Data.Status = i.GetCleanInfo(infoArea)
				}

				if infoType == "premiered" {
					i.Data.Premiered = i.GetCleanInfo(infoArea)
				}

				if infoType == "broadcast" {
					i.Data.Broadcast = i.GetCleanInfo(infoArea)
				}

				if infoType == "source" {
					i.Data.Source = i.GetCleanInfo(infoArea)
				}

				if infoType == "serialization" {
					i.Data.Serialization = i.GetCleanInfo(infoArea)
				}

				if infoType == "duration" {
					i.Data.Duration = i.GetCleanInfo(infoArea)
				}

				if infoType == "rating" {
					i.Data.Rating = i.GetCleanInfo(infoArea)
				}

				if infoType == "published" {
					infoValue := i.GetCleanInfo(infoArea)
					i.Data.Published.Start, i.Data.Published.End = i.GetAiringInfo(infoValue)
				}

				if infoType == "aired" {
					infoValue := i.GetCleanInfo(infoArea)
					i.Data.Aired.Start, i.Data.Aired.End = i.GetAiringInfo(infoValue)
				}

				if infoType == "producers" {
					infoValue := i.GetCleanInfo(infoArea)
					i.Data.Producers = i.GetIdNameInfo(infoArea, infoType, infoValue)
				}

				if infoType == "licensors" {
					infoValue := i.GetCleanInfo(infoArea)
					i.Data.Licensors = i.GetIdNameInfo(infoArea, infoType, infoValue)
				}

				if infoType == "studios" {
					infoValue := i.GetCleanInfo(infoArea)
					i.Data.Studios = i.GetIdNameInfo(infoArea, infoType, infoValue)
				}

				if infoType == "genres" {
					infoValue := i.GetCleanInfo(infoArea)
					i.Data.Genres = i.GetIdNameInfo(infoArea, infoType, infoValue)
				}

				if infoType == "authors" {
					infoValue := i.GetCleanInfo(infoArea)
					i.Data.Authors = i.GetIdNameInfo(infoArea, infoType, infoValue)
				}

				infoArea = infoArea.Next()
				if goquery.NodeName(infoArea) == "h2" || goquery.NodeName(infoArea) == "br" {
					break
				}
			}
			return
		}
	})
}

// GetCleanInfo to get clean detail anime or manga info.
func (i *InfoModel) GetCleanInfo(infoArea *goquery.Selection) string {
	infoArea.Find("span:nth-of-type(1)").Remove()

	infoValue := infoArea.Text()
	infoValue = strings.TrimSpace(infoValue)

	replacer := strings.NewReplacer(", add some", "", "?", "", "Not yet aired", "", "Unknown", "")

	infoValue = replacer.Replace(infoValue)

	return infoValue
}

// GetAiringInfo to get airing and published date of anime or manga.
func (i *InfoModel) GetAiringInfo(infoValue string) (string, string) {
	if infoValue != "Not available" {
		splitDate := strings.Split(infoValue, " to ")
		if len(splitDate) > 1 {
			return splitDate[0], splitDate[1]
		} else {
			return splitDate[0], ""
		}
	}
	return "", ""
}

// GetIdNameInfo to get producer, licensor, studio, genre, author of anime or manga.
func (i *InfoModel) GetIdNameInfo(infoArea *goquery.Selection, infoType string, infoValue string) []IdName {
	var IdNameList []IdName
	if infoValue != "None found" {
		infoArea.Find("a").Each(func(j int, eachName *goquery.Selection) {
			link, _ := eachName.Attr("href")
			splitLink := strings.Split(link, "/")

			infoId, _ := strconv.Atoi(splitLink[3])
			if infoType == "authors" {
				infoId, _ = strconv.Atoi(splitLink[2])
			}

			IdNameList = append(IdNameList, IdName{
				Id:   infoId,
				Name: eachName.Text(),
			})
		})
	}

	return IdNameList
}

// SetRelated to set related anime or manga.
func (i *InfoModel) SetRelated() {
	result := make(map[string][]IdTitleType)
	relatedArea := i.Parser.Find(".anime_detail_related_anime")

	relatedArea.Find("tr").Each(func(i int, eachRelated *goquery.Selection) {
		var relatedList []IdTitleType

		relatedType := eachRelated.Find("td:nth-of-type(1)").Text()
		relatedType = strings.Replace(relatedType, ":", "", -1)
		relatedType = strings.TrimSpace(strings.ToLower(relatedType))

		relatedData := eachRelated.Find("td:nth-of-type(2)")
		relatedData.Find("a").Each(func(i int, eachData *goquery.Selection) {
			relatedLink, _ := eachData.Attr("href")
			splitLink := strings.Split(relatedLink, "/")
			idInt, _ := strconv.Atoi(splitLink[2])

			relatedList = append(relatedList, IdTitleType{
				Id:    idInt,
				Title: eachData.Text(),
				Type:  splitLink[1],
			})
		})
		result[relatedType] = relatedList
	})

	i.Data.Related = result
}

// SetCharacter to set character in anime or manga
func (i *InfoModel) SetCharacter() {
	var characterList []Character
	characterArea := i.Parser.Find("div[class^=detail-characters-list]:nth-of-type(1)")

	var areaList []*goquery.Selection
	areaList = append(areaList, characterArea.Find("div[class*=fl-l]"))
	areaList = append(areaList, characterArea.Find("div[class*=fl-r]"))

	for _, characterSide := range areaList {
		characterSide.Find("table[width=\"100%\"]").Each(func(j int, eachCharacter *goquery.Selection) {
			charArea := eachCharacter.Find("tr td:nth-of-type(2)")
			vaArea := eachCharacter.Find("td:nth-of-type(3) table td")

			characterList = append(characterList, Character{
				Id:      i.GetCharId(charArea),
				Name:    i.GetCharName(charArea),
				Role:    i.GetCharRole(charArea),
				Image:   i.GetCharImage(eachCharacter),
				VaId:    i.GetCharId(vaArea),
				VaName:  i.GetCharName(vaArea),
				VaRole:  i.GetCharRole(vaArea),
				VaImage: i.GetVaImage(eachCharacter),
			})
		})
	}

	i.Data.Character = characterList
}

// GetCharId to get character or staff id from link.
func (i *InfoModel) GetCharId(charArea *goquery.Selection) int {
	found, _ := charArea.Html()

	if found == "" {
		return 0
	}

	charLink, _ := charArea.Find("a").Attr("href")
	splitLink := strings.Split(charLink, "/")
	idInt, _ := strconv.Atoi(splitLink[4])
	return idInt
}

// GetCharName to get character or staff name from link.
func (i *InfoModel) GetCharName(charArea *goquery.Selection) string {
	r, _ := regexp.Compile(`\s+`)
	charName := charArea.Find("a").Text()
	charName = r.ReplaceAllString(charName, " ")
	return strings.TrimSpace(charName)
}

// GetCharRole to get character role in anime or manga.
func (i *InfoModel) GetCharRole(charArea *goquery.Selection) string {
	charRole := charArea.Find("small").Text()
	return strings.TrimSpace(charRole)
}

// GetCharImage to get character image.
func (i *InfoModel) GetCharImage(eachCharacter *goquery.Selection) string {
	charImage, _ := eachCharacter.Find("tr td img").Attr("data-src")
	return helper.ImageUrlCleaner(charImage)
}

// GetVaImage to get va image.
func (i *InfoModel) GetVaImage(eachCharacter *goquery.Selection) string {
	vaImage, _ := eachCharacter.Find("td:nth-of-type(3) table td:nth-of-type(2) img").Attr("data-src")
	return helper.ImageUrlCleaner(vaImage)
}

// SetStaff to set staff involved in anime or manga.
func (i *InfoModel) SetStaff() {
	var staffList []Staff
	staffArea := i.Parser.Find("a[name=staff]").Next().Next()

	var areaList []*goquery.Selection
	areaList = append(areaList, staffArea.Find("div[class*=fl-l]"))
	areaList = append(areaList, staffArea.Find("div[class*=fl-r]"))

	for _, staffSide := range areaList {
		staffSide.Find("table[width=\"100%\"]").Each(func(j int, eachStaff *goquery.Selection) {
			stArea := eachStaff.Find("tr td:nth-of-type(2)")

			staffList = append(staffList, Staff{
				Id:    i.GetCharId(stArea),
				Name:  i.GetCharName(stArea),
				Role:  i.GetCharRole(stArea),
				Image: i.GetCharImage(eachStaff),
			})
		})
	}

	i.Data.Staff = staffList
}

// SetSong to set anime opening and ending song.
func (i *InfoModel) SetSong() {
	i.Data.Song.Opening = i.GetCleanSong("div[class*=\"theme-songs opnening\"]")
	i.Data.Song.Closing = i.GetCleanSong("div[class*=\"theme-songs ending\"]")
}

// GetCleanSong to get clean song from div.
func (i *InfoModel) GetCleanSong(div string) []string {
	var songList []string
	songArea := i.Parser.Find(div)
	songArea.Find("span.theme-song").Each(func(j int, eachSong *goquery.Selection) {
		r, _ := regexp.Compile(`#\d*:\s`)
		song := eachSong.Text()
		song = r.ReplaceAllString(song, " ")
		song = strings.TrimSpace(song)
		songList = append(songList, song)
	})
	return songList
}

// SetReview to set anime or manga review.
func (i *InfoModel) SetReview() {
	var reviewList []Review
	reviewArea := i.Parser.Find(".js-scrollfix-bottom-rel table tr:nth-of-type(2)")
	reviewArea.Find(".borderDark[style^=padding]").Each(func(j int, eachReview *goquery.Selection) {
		topArea := eachReview.Find(".spaceit:nth-of-type(1)")
		bottomArea := topArea.Next()
		veryBottomArea := bottomArea.Next()

		reviewList = append(reviewList, Review{
			Id:       i.GetReviewId(veryBottomArea),
			Username: i.GetReviewUsername(topArea),
			Image:    i.GetReviewImage(topArea),
			Helpful:  i.GetReviewHelpful(topArea),
			Date:     i.GetReviewDate(topArea),
			Episode:  i.GetReviewProgress(topArea, "anime"),
			Chapter:  i.GetReviewProgress(topArea, "manga"),
			Score:    i.GetReviewScore(bottomArea),
			Review:   i.GetReviewContent(bottomArea),
		})
	})

	i.Data.Review = reviewList
}

// GetReviewId to get anime or manga review id.
func (i *InfoModel) GetReviewId(veryBottomArea *goquery.Selection) int {
	idLink, _ := veryBottomArea.Find("a").Attr("href")
	splitLink := strings.Split(idLink, "?id=")
	idInt, _ := strconv.Atoi(splitLink[1])
	return idInt
}

// GetReviewUsername to get anime or manga review username.
func (i *InfoModel) GetReviewUsername(topArea *goquery.Selection) string {
	reviewUsername := topArea.Find("table td:nth-of-type(2)").Find("a").Text()
	return strings.Replace(reviewUsername, "All reviews", "", -1)
}

// GetReviewImage to get review user image.
func (i *InfoModel) GetReviewImage(topArea *goquery.Selection) string {
	reviewImage, _ := topArea.Find("table td img").Attr("src")
	return helper.ImageUrlCleaner(reviewImage)
}

// GetReviewHelpful to get review helpful number.
func (i *InfoModel) GetReviewHelpful(topArea *goquery.Selection) int {
	reviewHelpful := topArea.Find("table td:nth-of-type(2) strong span[id^=rhelp]").Text()
	helpfulInt, _ := strconv.Atoi(strings.TrimSpace(reviewHelpful))
	return helpfulInt
}

// GetReviewDate to get review date.
func (i *InfoModel) GetReviewDate(topArea *goquery.Selection) ReviewDate {
	reviewDate := topArea.Find("div:nth-of-type(1)").Find("div:nth-of-type(1)")
	dateDate := reviewDate.Text()
	dateTime, _ := reviewDate.Attr("title")
	return ReviewDate{
		Date: dateDate,
		Time: dateTime,
	}
}

// GetReviewProgress to get anime review episode.
func (i *InfoModel) GetReviewProgress(topArea *goquery.Selection, t string) string {
	if i.Type == t {
		reviewProgress := topArea.Find("div:nth-of-type(1)").Find("div:nth-of-type(2)").Text()
		reviewProgress = strings.Replace(reviewProgress, "episodes seen", "", -1)
		reviewProgress = strings.Replace(reviewProgress, "chapters read", "", -1)
		return strings.TrimSpace(reviewProgress)
	}
	return ""
}

// GetReviewScore to get review score.
func (i *InfoModel) GetReviewScore(bottomArea *goquery.Selection) map[string]int {
	reviewScore := make(map[string]int)
	scoreArea := bottomArea.Find("table")
	scoreArea.Find("tr").Each(func(j int, eachScore *goquery.Selection) {
		scoreType := strings.ToLower(eachScore.Find("td:nth-of-type(1)").Text())
		scoreValue, _ := strconv.Atoi(eachScore.Find("td:nth-of-type(2)").Text())
		reviewScore[scoreType] = scoreValue
	})
	return reviewScore
}

// GetReviewContent to get review content.
func (i *InfoModel) GetReviewContent(bottomArea *goquery.Selection) string {
	bottomArea.Find("div[id^=score]").Remove()
	bottomArea.Find("div[id^=revhelp_output]").Remove()
	bottomArea.Find("a[id^=reviewToggle]").Remove()

	r, _ := regexp.Compile(`[^\S\r\n]+`)
	reviewContent := r.ReplaceAllString(bottomArea.Text(), " ")

	r, _ = regexp.Compile(`(\n\n \n)`)
	reviewContent = r.ReplaceAllString(reviewContent, "")

	return strings.TrimSpace(reviewContent)
}

// SetRecommendation to set anime or manga recommendation.
func (i *InfoModel) SetRecommendation() {
	var recommendationList []Recommendation
	recommendationArea := i.Parser.Find("#" + i.Type + "_recommendation")
	recommendationArea.Find("li.btn-anime").Each(func(j int, eachRecommendation *goquery.Selection) {
		recommendationList = append(recommendationList, Recommendation{
			Id:    i.GetRecomId(eachRecommendation),
			Title: i.GetRecomTitle(eachRecommendation),
			Image: i.GetRecomImage(eachRecommendation),
			User:  i.GetRecomUser(eachRecommendation),
		})
	})

	i.Data.Recommendation = recommendationList
}

// GetRecomId to get recommendation id.
func (i *InfoModel) GetRecomId(eachRecommendation *goquery.Selection) int {
	recomLink, _ := eachRecommendation.Find("a").Attr("href")
	splitLink := strings.Split(recomLink, "/")

	if eachRecommendation.Find(".users").Text() == "AutoRec" {
		idInt, _ := strconv.Atoi(splitLink[4])
		return idInt
	}

	splitLink = strings.Split(splitLink[5], "-")

	if splitLink[0] == strconv.Itoa(i.Id) {
		idInt, _ := strconv.Atoi(splitLink[1])
		return idInt
	} else {
		idInt, _ := strconv.Atoi(splitLink[0])
		return idInt
	}
}

// GetRecomTitle to get recommendation anime or manga title.
func (i *InfoModel) GetRecomTitle(eachRecommendation *goquery.Selection) string {
	return eachRecommendation.Find("span:nth-of-type(1)").Text()
}

// GetRecomImage to get recommendation anime or manga image.
func (i *InfoModel) GetRecomImage(eachRecommendation *goquery.Selection) string {
	recomImage, _ := eachRecommendation.Find("img").Attr("data-src")
	return helper.ImageUrlCleaner(recomImage)
}

// GetRecomUser to get number of user who recommend the anime or manga.
func (i *InfoModel) GetRecomUser(eachRecommendation *goquery.Selection) int {
	recomUser := eachRecommendation.Find(".users").Text()

	if recomUser == "AutoRec" {
		return 0
	}

	recomUser = strings.Replace(recomUser, "Users", "", -1)
	recomUser = strings.Replace(recomUser, "User", "", -1)
	recomInt, _ := strconv.Atoi(strings.TrimSpace(recomUser))
	return recomInt
}
