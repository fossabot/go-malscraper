package character

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/pkg/malscraper/config"
	"github.com/rl404/go-malscraper/pkg/malscraper/constant"
	model "github.com/rl404/go-malscraper/pkg/malscraper/model/character"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser"
	"github.com/rl404/go-malscraper/pkg/malscraper/utils"
)

// CharacterParser is parser for MyAnimeList character information.
// Example: https://myanimelist.net/character/1/Spike_Spiegel
type CharacterParser struct {
	parser.BaseParser
	ID   int
	Data model.Character
}

// InitCharacterParser to initiate all fields and data of CharacterParser.
func InitCharacterParser(config config.Config, id int) (character CharacterParser, err error) {
	character.ID = id
	character.Config = config

	// Checking to redis if using redis in config.
	// Redis key's pattern is `character:{id}`.
	redisKey := constant.RedisGetCharacter + ":" + strconv.Itoa(character.ID)
	if config.RedisClient != nil {
		found, err := utils.UnmarshalFromRedis(config.RedisClient, redisKey, &character.Data)
		if err != nil {
			character.SetResponse(500, err.Error())
			return character, err
		}

		if found {
			character.SetResponse(200, constant.SuccessMessage)
			return character, nil
		}
	}

	// Get MyAnimeList HTML source page and initiate the parser.
	err = character.InitParser("/character/"+strconv.Itoa(character.ID), "#contentWrapper")
	if err != nil {
		return character, err
	}

	// Fill in data field.
	character.setAllDetail()

	// Save data field to redis if using redis in config.
	if config.RedisClient != nil {
		go utils.SaveToRedis(config.RedisClient, redisKey, character.Data, config.CacheTime)
	}

	return character, nil
}

// setAllDetail to set all character detail information.
func (cp *CharacterParser) setAllDetail() {
	cp.setID()
	cp.setImage()
	cp.setNickname()
	cp.setName()
	cp.setKanjiName()
	cp.setFavorite()
	cp.setAbout()
	cp.setMedia("anime")
	cp.setMedia("manga")
	cp.setVa()
}

// setID to set character's id.
func (cp *CharacterParser) setID() {
	cp.Data.ID = cp.ID
}

// setImage to set character's image.
func (cp *CharacterParser) setImage() {
	image := cp.Parser.Find("#content table tr td div a img")
	cp.Data.Image, _ = image.Attr("data-src")
}

// setNickname to set character's nickname.
func (cp *CharacterParser) setNickname() {
	nickname := cp.Parser.Find("h1")

	r := regexp.MustCompile(`\s+`)
	nickRegex := r.ReplaceAllString(nickname.Text(), " ")
	nickRegex = strings.TrimSpace(nickRegex)

	r = regexp.MustCompile(`\"([^"])*`)
	nickRegex = r.FindString(nickRegex)

	if nickRegex != "" {
		cp.Data.Nickname = nickRegex[1:]
	}
}

// setName to set character's name.
func (cp *CharacterParser) setName() {
	area := cp.Parser.Find("#content table tr td").Next()
	area = area.Find("div.normal_header[style^=height]")
	area.Find("span").Remove()

	cp.Data.Name = strings.TrimSpace(area.Text())
}

// setKanjiName to set character's kanji name.
func (cp *CharacterParser) setKanjiName() {
	area := cp.Parser.Find("#content table tr td").Next()
	area = area.Find("div.normal_header small")

	r := regexp.MustCompile(`(\(|\))`)
	kanjiName := r.ReplaceAllString(area.Text(), "")

	cp.Data.KanjiName = kanjiName
}

// setFavorite to set character's number of favorite.
func (cp *CharacterParser) setFavorite() {
	favorite := cp.Parser.Find("#content table tr td").Text()

	r := regexp.MustCompile(`(Member Favorites: ).+`)
	regexFav := r.FindString(favorite)
	regexFav = strings.TrimSpace(regexFav)

	fav := utils.GetValueFromSplit(regexFav, ": ", 1)

	cp.Data.Favorite = utils.StrToNum(fav)
}

// setAbout to set character's about.
func (cp *CharacterParser) setAbout() {
	aboutHtml, _ := cp.Parser.Find("#content table tr td").Next().Html()

	r := regexp.MustCompile(`(?ms)(<div class="normal_header" style="height: 15px;">).*(<div class="normal_header">)`)
	regexAbout := r.FindString(aboutHtml)

	aboutGoQuery, _ := goquery.NewDocumentFromReader(strings.NewReader(regexAbout))
	cleanAbout := aboutGoQuery.Text()

	if cleanAbout != "No biography written." {
		cp.Data.About = strings.TrimSpace(cleanAbout)
	}
}

// setMedia to set character's animeography and mangaography.
func (cp *CharacterParser) setMedia(t string) {
	var medias []model.Ography

	area := cp.Parser.Find("#content table tr td:nth-of-type(1)")

	if t == "anime" {
		area = area.Find("table:nth-of-type(1)")
	} else {
		area = area.Find("table:nth-of-type(2)")
	}

	area.Find("tr").Each(func(i int, media *goquery.Selection) {
		mediaImage := media.Find("td:nth-of-type(1)")
		eachArea := media.Find("td:nth-of-type(2)")

		medias = append(medias, model.Ography{
			ID:    cp.getVaID(eachArea),
			Title: cp.getVaName(eachArea),
			Image: cp.getVaImage(mediaImage),
			Role:  cp.getVaRole(eachArea),
		})
	})

	if t == "anime" {
		cp.Data.Animeography = medias
	} else {
		cp.Data.Mangaography = medias
	}
}

// setVa to set character's VA.
func (cp *CharacterParser) setVa() {
	var vas []model.VoiceActor

	vaArea := cp.Parser.Find("#content table tr td").Next()
	vaArea.Find("table[width=\"100%\"]").Each(func(i int, va *goquery.Selection) {
		vaNameArea := va.Find("td:nth-of-type(2)")

		vas = append(vas, model.VoiceActor{
			ID:    cp.getVaID(vaNameArea),
			Name:  cp.getVaName(vaNameArea),
			Role:  cp.getVaRole(vaNameArea),
			Image: cp.getVaImage(va),
		})
	})

	cp.Data.VoiceActors = vas
}

// getVaID to get character's animeography, mangaography and va id.
func (cp *CharacterParser) getVaID(vaArea *goquery.Selection) int {
	id, _ := vaArea.Find("a").Attr("href")
	id = utils.GetValueFromSplit(id, "/", 4)
	return utils.StrToNum(id)
}

// getVaName to get character's animeography, mangaography and va name.
func (cp *CharacterParser) getVaName(vaArea *goquery.Selection) string {
	name, _ := vaArea.Find("a:nth-of-type(1)").Html()
	return name
}

// getVaImage to get character's animeography, mangaography and va image.
func (cp *CharacterParser) getVaImage(vaArea *goquery.Selection) string {
	image, _ := vaArea.Find("img").Attr("data-src")
	return utils.URLCleaner(image, "image", cp.Config.CleanImageURL)
}

// getVaRole to get character's animeography, mangaography and va role.
func (cp *CharacterParser) getVaRole(vaArea *goquery.Selection) string {
	return vaArea.Find("div small").Text()
}
