package anime

import (
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/pkg/malscraper/config"
	"github.com/rl404/go-malscraper/pkg/malscraper/constant"
	model "github.com/rl404/go-malscraper/pkg/malscraper/model/anime"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser"
	"github.com/rl404/go-malscraper/pkg/malscraper/utils"
)

// CharacterParser is parser for MyAnimeList anime character list.
// Example: https://myanimelist.net/anime/1/Cowboy_Bebop/characters
type CharacterParser struct {
	parser.BaseParser
	ID   int
	Data []model.Character
}

// InitCharacterParser to initiate all fields and data of CharacterParser.
func InitCharacterParser(config config.Config, id int) (character CharacterParser, err error) {
	character.ID = id
	character.Config = config

	// Checking to redis if using redis in config.
	// Redis key's pattern is `anime-character:{id}`.
	redisKey := constant.RedisGetAnimeCharacter + ":" + strconv.Itoa(character.ID)
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
	err = character.InitParser("/anime/"+strconv.Itoa(character.ID)+"/a/characters", ".js-scrollfix-bottom-rel")
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
func (csp *CharacterParser) setAllDetail() {
	var characters []model.Character

	csp.Parser.Find("article").Remove()

	charArea := csp.Parser.Find(".js-scrollfix-bottom-rel h2").First().Next()

	for goquery.NodeName(charArea) == "table" {

		charNameArea := charArea.Find("td:nth-of-type(2)")
		vaArea := charArea.Find("td:nth-of-type(3)")

		characters = append(characters, model.Character{
			ID:          csp.getCharID(charNameArea),
			Image:       csp.getCharImage(charArea),
			Name:        csp.getCharName(charNameArea),
			Role:        csp.getCharRole(charNameArea),
			VoiceActors: csp.getVa(vaArea),
		})

		charArea = charArea.Next()
	}

	csp.Data = characters
}

// getCharID to get character id.
func (csp *CharacterParser) getCharID(charNameArea *goquery.Selection) int {
	id, _ := charNameArea.Find("a").First().Attr("href")
	id = utils.GetValueFromSplit(id, "/", 4)
	return utils.StrToNum(id)
}

// getCharImage to get character image.
func (csp *CharacterParser) getCharImage(charArea *goquery.Selection) string {
	image, _ := charArea.Find("td .picSurround img").Attr("data-src")
	return utils.URLCleaner(image, "image", csp.Config.CleanImageURL)
}

// getCharName to get character name.
func (csp *CharacterParser) getCharName(charNameArea *goquery.Selection) string {
	return charNameArea.Find("a").First().Text()
}

// getCharRole to get character role.
func (csp *CharacterParser) getCharRole(charNameArea *goquery.Selection) string {
	return charNameArea.Find("small").First().Text()
}

// getVa to get character's va list.
func (csp *CharacterParser) getVa(vaArea *goquery.Selection) []model.VoiceActor {
	var vaList []model.VoiceActor

	vaArea = vaArea.Find("table")
	vaArea.Find("tr").Each(func(i int, eachVa *goquery.Selection) {

		vaNameArea := eachVa.Find("td").First()

		vaList = append(vaList, model.VoiceActor{
			ID:    csp.getCharID(vaNameArea),
			Image: csp.getCharImage(eachVa),
			Name:  csp.getCharName(vaNameArea),
			Role:  csp.getCharRole(vaNameArea),
		})
	})

	return vaList
}
