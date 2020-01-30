package manga

import (
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/pkg/malscraper/config"
	"github.com/rl404/go-malscraper/pkg/malscraper/constant"
	model "github.com/rl404/go-malscraper/pkg/malscraper/model/manga"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser"
	"github.com/rl404/go-malscraper/pkg/malscraper/utils"
)

// CharacterParser is parser for MyAnimeList manga character list.
// Example: https://myanimelist.net/manga/1/Monster/characters
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
	// Redis key's pattern is `manga-character:{id}`.
	redisKey := constant.RedisGetMangaCharacter + ":" + strconv.Itoa(character.ID)
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
	err = character.InitParser("/manga/"+strconv.Itoa(character.ID)+"/a/characters", ".js-scrollfix-bottom-rel")
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

// setAllDetail to set all character and staff detail information.
func (csp *CharacterParser) setAllDetail() {
	var characters []model.Character

	csp.Parser.Find("article").Remove()

	charArea := csp.Parser.Find(".js-scrollfix-bottom-rel h2").First().Next()

	for goquery.NodeName(charArea) == "table" {

		charNameArea := charArea.Find("td:nth-of-type(2)")

		characters = append(characters, model.Character{
			ID:    csp.getCharID(charNameArea),
			Image: csp.getCharImage(charArea),
			Name:  csp.getCharName(charNameArea),
			Role:  csp.getCharRole(charNameArea),
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
