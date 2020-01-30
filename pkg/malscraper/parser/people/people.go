package people

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/grokify/html-strip-tags-go"
	"github.com/rl404/go-malscraper/pkg/malscraper/config"
	"github.com/rl404/go-malscraper/pkg/malscraper/constant"
	model "github.com/rl404/go-malscraper/pkg/malscraper/model/people"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser"
	"github.com/rl404/go-malscraper/pkg/malscraper/utils"
)

// PeopleParser is parser for MyAnimeList people information.
// Example: https://myanimelist.net/people/1
type PeopleParser struct {
	parser.BaseParser
	ID   int
	Data model.People
}

// InitPeopleParser to initiate all fields and data of PeopleParser.
func InitPeopleParser(config config.Config, id int) (people PeopleParser, err error) {
	people.ID = id
	people.Config = config

	// Checking to redis if using redis in config.
	// Redis key's pattern is `people:{id}`.
	redisKey := constant.RedisGetPeople + ":" + strconv.Itoa(people.ID)
	if config.RedisClient != nil {
		found, err := utils.UnmarshalFromRedis(config.RedisClient, redisKey, &people.Data)
		if err != nil {
			people.SetResponse(500, err.Error())
			return people, err
		}

		if found {
			people.SetResponse(200, constant.SuccessMessage)
			return people, nil
		}
	}

	// Get MyAnimeList HTML source page and initiate the parser.
	err = people.InitParser("/people/"+strconv.Itoa(people.ID), "#contentWrapper")
	if err != nil {
		return people, err
	}

	// Fill in data field.
	people.setAllDetail()

	// Save data field to redis if using redis in config.
	if config.RedisClient != nil {
		go utils.SaveToRedis(config.RedisClient, redisKey, people.Data, config.CacheTime)
	}

	return people, nil
}

// setAllDetail to set all people detail information.
func (pp *PeopleParser) setAllDetail() {
	pp.setID()
	pp.setName()
	pp.setImage()
	pp.cleanBiodata()
	pp.setBiodata()
	pp.setMore()
	pp.setVa()
	pp.setStaffManga(false)
	pp.setStaffManga(true)
}

// setID to set people's id.
func (pp *PeopleParser) setID() {
	pp.Data.ID = pp.ID
}

// setName to set people's name.
func (pp *PeopleParser) setName() {
	pp.Data.Name = pp.Parser.Find("h1").Text()
}

// setImage to set people's name.
func (pp *PeopleParser) setImage() {
	pp.Data.Image, _ = pp.Parser.Find("#content table tr td img").Attr("data-src")
}

// cleanBiodata to clean people's biodata area.
func (pp *PeopleParser) cleanBiodata() {
	area := pp.Parser.Find("#content table tr td")
	area.Find("div").EachWithBreak(func(i int, trash *goquery.Selection) bool {
		if i < 5 {
			trash.Remove()
			return true
		}
		return false
	})
}

// setBiodata to set people's detail biodata.
func (pp *PeopleParser) setBiodata() {
	pp.Data.GivenName, _ = pp.getBiodata("Given name")
	pp.Data.FamilyName, _ = pp.getBiodata("Family name")
	_, pp.Data.AlternativeNames = pp.getBiodata("Alternate names")
	pp.Data.Birthday, _ = pp.getBiodata("Birthday")
	pp.Data.Website = pp.getBioWeb()
	pp.Data.Favorite = pp.getBioFavorite()
}

// getBiodata to get people's detail biodata.
func (pp *PeopleParser) getBiodata(t string) (string, []string) {
	area, _ := pp.Parser.Find("#content table tr td").Html()

	r := regexp.MustCompile(`(` + t + `:<\/span>)[^<]*`)
	bioRegex := r.FindString(area)
	bioRegex = strip.StripTags(bioRegex)

	if bioRegex != "" {
		splitBio := strings.Split(bioRegex, ": ")

		splitBio[1] = strings.TrimSpace(splitBio[1])

		r = regexp.MustCompile(`\s+`)
		splitBio[1] = r.ReplaceAllString(splitBio[1], " ")

		if t == "Alternate names" {
			splitName := strings.Split(splitBio[1], ", ")
			return "", splitName
		}

		if t == "Member Favorites" {
			splitBio[1] = strings.Replace(splitBio[1], ",", "", -1)
		}

		return splitBio[1], nil
	}

	return "", nil
}

// getBioWeb to get people's website.
func (pp *PeopleParser) getBioWeb() string {
	area, _ := pp.Parser.Find("#content table tr td").Html()

	r := regexp.MustCompile(`(Website:<\/span> <a)[^<]*`)
	bioRegex := r.FindString(area)

	r = regexp.MustCompile(`".+"`)
	bioRegex = r.FindString(bioRegex)

	if bioRegex != "\"http://\"" {
		return strings.Replace(bioRegex, "\"", "", -1)
	}

	return ""
}

// getBioFavorite to get people's member favorite.
func (pp *PeopleParser) getBioFavorite() int {
	area, _ := pp.Parser.Find("#content table tr td").Html()

	r := regexp.MustCompile(`(Member Favorites:<\/span>)[^<]*`)
	bioRegex := r.FindString(area)
	bioRegex = strip.StripTags(bioRegex)

	if bioRegex != "" {
		splitBio := utils.GetValueFromSplit(bioRegex, ": ", 1)
		return utils.StrToNum(splitBio)
	}

	return 0
}

// setMore to set people's more detail biodata.
func (pp *PeopleParser) setMore() {
	pp.Data.More = pp.Parser.Find("#content table tr td div[class^=people-informantion-more]").Text()
}

// setVa to set people's anime voice actor list.
func (pp *PeopleParser) setVa() {
	var actors []model.VoiceActor

	area := pp.Parser.Find("#content table tr td").Next()
	area = area.Find(".normal_header").First().Next()

	if goquery.NodeName(area) == "table" {
		area.Find("tr").Each(func(i int, va *goquery.Selection) {
			animeImageArea := va.Find("td:nth-of-type(1)")
			animeArea := va.Find("td:nth-of-type(2)")
			charImageArea := va.Find("td:nth-of-type(4)")
			charArea := va.Find("td:nth-of-type(3)")

			actors = append(actors, model.VoiceActor{
				Anime: model.Anime{
					ID:    pp.getAnimeID(animeArea),
					Title: pp.getAnimeTitle(animeArea),
					Image: pp.getAnimeImage(animeImageArea),
				},
				Character: model.Character{
					ID:    pp.getAnimeID(charArea),
					Name:  pp.getAnimeTitle(charArea),
					Role:  pp.getAnimeRole(charArea),
					Image: pp.getAnimeImage(charImageArea),
				},
			})
		})

		pp.Data.VoiceActors = actors
	}
}

// getAnimeID to get people's va's anime id.
func (pp *PeopleParser) getAnimeID(animeArea *goquery.Selection) int {
	animeID, _ := animeArea.Find("a").Attr("href")
	id := utils.GetValueFromSplit(animeID, "/", 4)
	return utils.StrToNum(id)
}

// getAnimeImage to get people's va's anime image.
func (pp *PeopleParser) getAnimeImage(animeImageArea *goquery.Selection) string {
	animeImage, _ := animeImageArea.Find("img").Attr("data-src")
	return utils.URLCleaner(animeImage, "image", pp.Config.CleanImageURL)
}

// getAnimeTitle to get people's va's anime title.
func (pp *PeopleParser) getAnimeTitle(animeArea *goquery.Selection) string {
	return animeArea.Find("a").First().Text()
}

// getAnimeRole to get people's va's anime role.
func (pp *PeopleParser) getAnimeRole(animeArea *goquery.Selection) string {
	return strings.TrimSpace(animeArea.Find("div").Text())
}

// setStaff to set people's anime staff and published manga list.
func (pp *PeopleParser) setStaffManga(isManga bool) {
	var staffList []model.Staff

	area := pp.Parser.Find("#content table tr td").Next()

	area.Find(".normal_header").First().Remove()

	if isManga {
		area = area.Find(".normal_header").First().Next()
	} else {
		area = area.Find(".normal_header").First().Next()
	}

	if goquery.NodeName(area) == "table" {
		area.Find("tr").Each(func(i int, staff *goquery.Selection) {
			animeImageArea := staff.Find("td:nth-of-type(1)")
			stArea := staff.Find("td:nth-of-type(2)")

			staffList = append(staffList, model.Staff{
				ID:    pp.getAnimeID(stArea),
				Title: pp.getAnimeTitle(stArea),
				Image: pp.getAnimeImage(animeImageArea),
				Role:  pp.getStaffRole(stArea),
			})

		})

		if isManga {
			pp.Data.PublishedManga = staffList
		} else {
			pp.Data.Staff = staffList
		}
	}
}

// getStaffRole to get people's anime title of the voice actor play.
func (pp *PeopleParser) getStaffRole(stArea *goquery.Selection) string {
	return strings.TrimSpace(stArea.Find("small").Text())
}
