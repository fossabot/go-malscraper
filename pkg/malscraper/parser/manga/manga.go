package manga

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/pkg/malscraper/model/common"
	model "github.com/rl404/go-malscraper/pkg/malscraper/model/manga"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser"
	"github.com/rl404/go-malscraper/pkg/malscraper/utils"
)

// MangaParser is parser for MyAnimeList manga information.
// Example: https://myanimelist.net/manga/1
type MangaParser struct {
	parser.BaseParser
	ID   int
	Data model.Manga
}

// InitMangaParser to initiate all fields and data of MangaParser.
func InitMangaParser(id int) (mangaParser MangaParser, err error) {
	mangaParser.ID = id

	err = mangaParser.InitParser("/manga/"+strconv.Itoa(mangaParser.ID), "#content")
	if err != nil {
		return mangaParser, err
	}

	mangaParser.setDetails()
	return mangaParser, nil
}

// setDetails to set manga details information.
func (mp *MangaParser) setDetails() {
	mp.setID()
	mp.setCover()
	mp.setTitle()
	mp.setAltTitle()
	mp.setSynopsis()
	mp.setScore()
	mp.setVoter()
	mp.setRank()
	mp.setPopularity()
	mp.setMember()
	mp.setFavorite()
	mp.setOtherInfo()
	mp.setRelated()
	mp.setCharacter()
	mp.setReview()
	mp.setRecommendation()
}

// setID to set manga id.
func (mp *MangaParser) setID() {
	mp.Data.ID = mp.ID
}

// setCover to set manga cover image.
func (mp *MangaParser) setCover() {
	image, _ := mp.Parser.Find("img.ac").Attr("data-src")
	mp.Data.Cover = utils.ImageURLCleaner(image)
}

// setTitle to set manga title.
func (mp *MangaParser) setTitle() {
	mp.Data.Title = mp.Parser.Parent().Find("h1.h1 span").Text()
}

// setAltTitle to set manga alternative titles.
func (mp *MangaParser) setAltTitle() {
	area := mp.Parser.Find(".js-scrollfix-bottom")

	mp.Data.AlternativeTitles.English = mp.getAltTitle(area, "English")
	mp.Data.AlternativeTitles.Synonym = mp.getAltTitle(area, "Synonyms")
	mp.Data.AlternativeTitles.Japanese = mp.getAltTitle(area, "Japanese")
}

// getAltTitle to get manga alternative titles (english, synonym, japanese).
func (mp *MangaParser) getAltTitle(area *goquery.Selection, t string) string {
	altTitle, _ := area.Html()

	r := regexp.MustCompile(`(` + t + `:</span>)([^<]*)`)
	altTitle = r.FindString(altTitle)
	altTitle = strings.Replace(altTitle, t+":</span>", "", -1)

	return strings.TrimSpace(altTitle)
}

// setSynopsis to set manga synopsis.
func (mp *MangaParser) setSynopsis() {
	synopsisArea := mp.Parser.Find("span[itemprop=description]")

	r := regexp.MustCompile(`\n[^\S\n]*`)
	synopsis := r.ReplaceAllString(synopsisArea.Text(), "\n")

	mp.Data.Synopsis = strings.TrimSpace(synopsis)
}

// setScore to set manga score.
func (mp *MangaParser) setScore() {
	scoreArea := mp.Parser.Find("div[class=\"fl-l score\"]")
	score := strings.TrimSpace(scoreArea.Text())

	if score != "N/A" {
		mp.Data.Score = utils.StrToFloat(score)
	} else {
		mp.Data.Score = 0.0
	}
}

// setVoter to set number who vote the score.
func (mp *MangaParser) setVoter() {
	voter, _ := mp.Parser.Find("div[class=\"fl-l score\"]").Attr("data-user")

	replacer := strings.NewReplacer("users", "", "user", "", ",", "")
	voter = replacer.Replace(voter)

	mp.Data.Voter = utils.StrToNum(voter)
}

// setRank to set manga rank.
func (mp *MangaParser) setRank() {
	rank := mp.Parser.Find("span[class=\"numbers ranked\"] strong").Text()
	rank = strings.Replace(rank, "#", "", -1)

	if rank == "N/A" {
		rank = ""
	}

	mp.Data.Rank = utils.StrToNum(rank)
}

// setPopularity to set manga popularity rank.
func (mp *MangaParser) setPopularity() {
	popularity := mp.Parser.Find("span[class=\"numbers popularity\"] strong").Text()
	popularity = strings.Replace(popularity, "#", "", -1)
	mp.Data.Popularity = utils.StrToNum(popularity)
}

// setMember to set manga number of member.
func (mp *MangaParser) setMember() {
	member := mp.Parser.Find("span[class=\"numbers members\"] strong").Text()
	mp.Data.Member = utils.StrToNum(member)
}

// setFavorite to set manga number of favorite.
func (mp *MangaParser) setFavorite() {
	favoriteArea := mp.Parser.Find("div[data-id=info2]").Next().Next().Next()
	favoriteArea.Find("span").Remove()
	mp.Data.Favorite = utils.StrToNum(favoriteArea.Text())
}

// setOtherInfo to set manga other details.
func (mp *MangaParser) setOtherInfo() {
	mp.Parser.Find(".js-scrollfix-bottom").Find("h2").Each(func(i int, area *goquery.Selection) {
		if area.Text() == "Information" {
			area = area.Next()
			for {
				infoType := area.Find("span").First().Text()
				infoType = strings.ToLower(infoType)
				infoType = strings.Replace(infoType, ":", "", -1)

				if infoType == "type" {
					mp.Data.Type = mp.getCleanInfo(area)
				}

				if infoType == "volumes" {
					mp.Data.Volume = utils.StrToNum(mp.getCleanInfo(area))
				}

				if infoType == "chapters" {
					mp.Data.Chapter = utils.StrToNum(mp.getCleanInfo(area))
				}

				if infoType == "status" {
					mp.Data.Status = mp.getCleanInfo(area)
				}

				if infoType == "published" {
					infoValue := mp.getCleanInfo(area)
					mp.Data.StartDate.Start, mp.Data.StartDate.End = mp.getAiringInfo(infoValue)
				}

				if infoType == "serialization" {
					infoValue := mp.getCleanInfo(area)
					mp.Data.Serializations = mp.getIDNameInfo(area, infoType, infoValue)
				}

				if infoType == "authors" {
					infoValue := mp.getCleanInfo(area)
					mp.Data.Authors = mp.getIDNameInfo(area, infoType, infoValue)
				}

				if infoType == "genres" {
					infoValue := mp.getCleanInfo(area)
					mp.Data.Genres = mp.getIDNameInfo(area, infoType, infoValue)
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

// getCleanInfo to get manga clean details.
func (mp *MangaParser) getCleanInfo(area *goquery.Selection) string {
	area.Find("span:nth-of-type(1)").Remove()

	replacer := strings.NewReplacer(", add some", "", "?", "", "Not yet aired", "", "Unknown", "")

	infoValue := area.Text()
	infoValue = strings.TrimSpace(infoValue)
	infoValue = replacer.Replace(infoValue)

	return infoValue
}

// getAiringInfo to get manga airing date.
func (mp *MangaParser) getAiringInfo(infoValue string) (string, string) {
	if infoValue != "Not available" {
		r := regexp.MustCompile(`\s+`)
		infoValue = r.ReplaceAllString(infoValue, " ")
		splitDate := strings.Split(infoValue, " to ")
		if len(splitDate) > 1 {
			return splitDate[0], splitDate[1]
		}
		return splitDate[0], ""
	}
	return "", ""
}

// getIDNameInfo to get manga producer, licensor, studio, and genre.
func (mp *MangaParser) getIDNameInfo(infoArea *goquery.Selection, infoType string, infoValue string) []common.IDName {
	var IDNameList []common.IDName
	if infoValue != "None found" {
		infoArea.Find("a").Each(func(i int, name *goquery.Selection) {
			link, _ := name.Attr("href")
			splitLink := strings.Split(link, "/")

			infoID := utils.StrToNum(splitLink[3])
			if infoType == "authors" {
				infoID = utils.StrToNum(splitLink[2])
			}

			IDNameList = append(IDNameList, common.IDName{
				ID:   infoID,
				Name: name.Text(),
			})
		})
	}
	return IDNameList
}

// setRelated to set related manga or manga.
func (mp *MangaParser) setRelated() {
	result := make(map[string][]model.Related)
	relatedArea := mp.Parser.Find(".anime_detail_related_anime")

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

	mp.Data.Related = result
}

// setCharacter to set manga related characters.
func (mp *MangaParser) setCharacter() {
	var characters []model.Character
	area := mp.Parser.Find("div[class^=detail-characters-list]:nth-of-type(1)")

	var areaList []*goquery.Selection
	areaList = append(areaList, area.Find("div[class*=fl-l]"))
	areaList = append(areaList, area.Find("div[class*=fl-r]"))

	for _, characterSide := range areaList {
		characterSide.Find("table[width=\"100%\"]").Each(func(i int, character *goquery.Selection) {
			charArea := character.Find("tr td:nth-of-type(2)")

			characters = append(characters, model.Character{
				ID:    mp.getCharID(charArea),
				Name:  mp.getCharName(charArea),
				Role:  mp.getCharRole(charArea),
				Image: mp.getCharImage(character),
			})
		})
	}

	mp.Data.Characters = characters
}

// getCharID to get manga character id.
func (mp *MangaParser) getCharID(charArea *goquery.Selection) int {
	found, _ := charArea.Html()

	if found == "" {
		return 0
	}

	charLink, _ := charArea.Find("a").Attr("href")
	id := utils.GetValueFromSplit(charLink, "/", 4)
	return utils.StrToNum(id)
}

// getCharName to get manga character, va or staff name.
func (mp *MangaParser) getCharName(charArea *goquery.Selection) string {
	r := regexp.MustCompile(`\s+`)
	charName := charArea.Find("a").Text()
	charName = r.ReplaceAllString(charName, " ")
	return strings.TrimSpace(charName)
}

// getCharRole to get manga character, va or staff role.
func (mp *MangaParser) getCharRole(charArea *goquery.Selection) string {
	charRole := charArea.Find("small").Text()
	return strings.TrimSpace(charRole)
}

// getCharImage to get manga character or staff image.
func (mp *MangaParser) getCharImage(eachCharacter *goquery.Selection) string {
	charImage, _ := eachCharacter.Find("tr td img").Attr("data-src")
	return utils.ImageURLCleaner(charImage)
}

// setReview to set manga review.
func (mp *MangaParser) setReview() {
	var reviews []model.Review
	area := mp.Parser.Find(".js-scrollfix-bottom-rel table tr:nth-of-type(2)")
	area.Find(".borderDark[style^=padding]").Each(func(i int, review *goquery.Selection) {
		topArea := review.Find(".spaceit:nth-of-type(1)")
		bottomArea := topArea.Next()
		veryBottomArea := bottomArea.Next()

		reviews = append(reviews, model.Review{
			ID:       mp.getReviewID(veryBottomArea),
			Username: mp.getReviewUsername(topArea),
			Image:    mp.getReviewImage(topArea),
			Helpful:  mp.getReviewHelpful(topArea),
			Date:     mp.getReviewDate(topArea),
			Chapter:  mp.getReviewProgress(topArea),
			Score:    mp.getReviewScore(bottomArea),
			Review:   mp.getReviewContent(bottomArea),
		})
	})

	mp.Data.Reviews = reviews
}

// getReviewID to get manga review id.
func (mp *MangaParser) getReviewID(veryBottomArea *goquery.Selection) int {
	idLink, _ := veryBottomArea.Find("a").Attr("href")
	id := utils.GetValueFromSplit(idLink, "?id=", 1)
	return utils.StrToNum(id)
}

// getReviewUsername to get manga review username.
func (mp *MangaParser) getReviewUsername(topArea *goquery.Selection) string {
	reviewUsername := topArea.Find("table td:nth-of-type(2)").Find("a").Text()
	return strings.Replace(reviewUsername, "All reviews", "", -1)
}

// getReviewImage to get manga review user image.
func (mp *MangaParser) getReviewImage(topArea *goquery.Selection) string {
	reviewImage, _ := topArea.Find("table td img").Attr("src")
	return utils.ImageURLCleaner(reviewImage)
}

// getReviewHelpful to get manga review helpful number.
func (mp *MangaParser) getReviewHelpful(topArea *goquery.Selection) int {
	reviewHelpful := topArea.Find("table td:nth-of-type(2) strong span[id^=rhelp]").Text()
	return utils.StrToNum(reviewHelpful)
}

// getReviewDate to get manga review date.
func (mp *MangaParser) getReviewDate(topArea *goquery.Selection) common.DateTime {
	reviewDate := topArea.Find("div:nth-of-type(1)").Find("div:nth-of-type(1)")
	dateDate := reviewDate.Text()
	dateTime, _ := reviewDate.Attr("title")
	return common.DateTime{
		Date: dateDate,
		Time: dateTime,
	}
}

// getReviewProgress to get manga review episode.
func (mp *MangaParser) getReviewProgress(topArea *goquery.Selection) string {
	reviewProgress := topArea.Find("div:nth-of-type(1)").Find("div:nth-of-type(2)").Text()
	replacer := strings.NewReplacer("episodes seen", "", "chapters read", "")
	return strings.TrimSpace(replacer.Replace(reviewProgress))
}

// getReviewScore to get manga review score.
func (mp *MangaParser) getReviewScore(bottomArea *goquery.Selection) map[string]int {
	reviewScore := make(map[string]int)
	area := bottomArea.Find("table")
	area.Find("tr").Each(func(i int, score *goquery.Selection) {
		scoreType := strings.ToLower(score.Find("td:nth-of-type(1)").Text())
		reviewScore[scoreType] = utils.StrToNum(score.Find("td:nth-of-type(2)").Text())
	})
	return reviewScore
}

// getReviewContent to get manga review content.
func (mp *MangaParser) getReviewContent(bottomArea *goquery.Selection) string {
	bottomArea.Find("div[id^=score]").Remove()
	bottomArea.Find("div[id^=revhelp_output]").Remove()
	bottomArea.Find("a[id^=reviewToggle]").Remove()

	r := regexp.MustCompile(`[^\S\r\n]+`)
	reviewContent := r.ReplaceAllString(bottomArea.Text(), " ")

	r = regexp.MustCompile(`(\n\n \n)`)
	reviewContent = r.ReplaceAllString(reviewContent, "")

	return strings.TrimSpace(reviewContent)
}

// setRecommendation to set manga recommendation.
func (mp *MangaParser) setRecommendation() {
	var recommendations []model.MangaRecommendation
	area := mp.Parser.Find("#manga_recommendation")
	area.Find("li.btn-anime").Each(func(i int, recommendation *goquery.Selection) {
		recommendations = append(recommendations, model.MangaRecommendation{
			ID:    mp.getRecomID(recommendation),
			Title: mp.getRecomTitle(recommendation),
			Image: mp.getRecomImage(recommendation),
			Count: mp.getRecomCount(recommendation),
		})
	})

	mp.Data.Recommendations = recommendations
}

// getRecomID to get manga recommendation id.
func (mp *MangaParser) getRecomID(recommendation *goquery.Selection) int {
	recomLink, _ := recommendation.Find("a").Attr("href")
	splitLink := strings.Split(recomLink, "/")

	if recommendation.Find(".users").Text() == "AutoRec" {
		return utils.StrToNum(splitLink[4])
	}

	splitLink = strings.Split(splitLink[5], "-")

	if splitLink[0] == strconv.Itoa(mp.ID) {
		return utils.StrToNum(splitLink[1])
	}
	return utils.StrToNum(splitLink[0])
}

// getRecomTitle to get manga recommendation title.
func (mp *MangaParser) getRecomTitle(recommendation *goquery.Selection) string {
	return recommendation.Find("span:nth-of-type(1)").Text()
}

// getRecomImage to get manga recommendation image.
func (mp *MangaParser) getRecomImage(recommendation *goquery.Selection) string {
	recomImage, _ := recommendation.Find("img").Attr("data-src")
	return utils.ImageURLCleaner(recomImage)
}

// getRecomCount to get manga user recommendation count.
func (mp *MangaParser) getRecomCount(recommendation *goquery.Selection) int {
	recomCount := recommendation.Find(".users").Text()

	if recomCount == "AutoRec" {
		return 0
	}
	replacer := strings.NewReplacer("Users", "", "User", "")
	return utils.StrToNum(replacer.Replace(recomCount))
}
