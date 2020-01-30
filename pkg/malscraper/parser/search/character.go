package search

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/pkg/malscraper/config"
	"github.com/rl404/go-malscraper/pkg/malscraper/constant"
	"github.com/rl404/go-malscraper/pkg/malscraper/model/common"
	model "github.com/rl404/go-malscraper/pkg/malscraper/model/search"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser"
	"github.com/rl404/go-malscraper/pkg/malscraper/utils"
)

// CharacterParser is parser for MyAnimeList character search result list.
// Example: https://myanimelist.net/character.php?q=luffy
type CharacterParser struct {
	parser.BaseParser
	Query string
	Page  int
	Data  []model.Character
}

// InitCharacterParser to initiate all fields and data of CharacterParser.
func InitCharacterParser(config config.Config, query string, page ...int) (character CharacterParser, err error) {
	character.Query = query
	character.Page = 0
	character.Config = config

	if len(page) > 0 {
		character.Page = 50 * (page[0] - 1)
	}

	if len(character.Query) < 3 {
		character.ResponseCode = 400
		return character, common.Err3LettersSearch
	}

	// Checking to redis if using redis in config.
	// Redis key's pattern is `search-character:{query},{page}`.
	redisKey := constant.RedisSearchCharacter + ":" + character.Query + "," + strconv.Itoa(character.Page)
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
	err = character.InitParser("/character.php?q="+character.Query+"&show="+strconv.Itoa(character.Page), "#content")
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
	var characters []model.Character

	area := cp.Parser.Find("table")
	area.Find("tr").EachWithBreak(func(i int, eachSearch *goquery.Selection) bool {
		if i == 0 {
			return true
		}

		nameArea := eachSearch.Find("td:nth-of-type(2)")

		characters = append(characters, model.Character{
			Image:    cp.getImage(eachSearch),
			ID:       cp.getID(nameArea),
			Name:     cp.getName(nameArea),
			Nickname: cp.getNickname(nameArea),
			Anime:    cp.getRole("anime", eachSearch),
			Manga:    cp.getRole("manga", eachSearch),
		})

		return true
	})

	cp.Data = characters
}

// getImage to get character image.
func (cp *CharacterParser) getImage(eachSearch *goquery.Selection) string {
	image, _ := eachSearch.Find("td div.picSurround a img").Attr("data-src")
	return utils.URLCleaner(image, "image", cp.Config.CleanImageURL)
}

// getID to get character id.
func (cp *CharacterParser) getID(nameArea *goquery.Selection) int {
	id, _ := nameArea.Find("a").First().Attr("href")
	id = utils.GetValueFromSplit(id, "/", 4)
	return utils.StrToNum(id)
}

// getName to get character name.
func (cp *CharacterParser) getName(nameArea *goquery.Selection) string {
	return nameArea.Find("a").First().Text()
}

// getNickname to get character nickname.
func (cp *CharacterParser) getNickname(nameArea *goquery.Selection) string {
	nick := nameArea.Find("small").First().Text()
	if nick != "" {
		return nick[1 : len(nick)-1]
	}
	return nick
}

// getRole to get character anime/manga role.
func (cp *CharacterParser) getRole(t string, eachSearch *goquery.Selection) []model.Role {
	var roles []model.Role
	area := eachSearch.Find("td:nth-of-type(3) small")
	area.Find("a").Each(func(i int, eachRole *goquery.Selection) {
		id, _ := eachRole.Attr("href")
		splitID := strings.Split(id, "/")
		roleType := splitID[1]

		if t == roleType && splitID[2] != "" {
			roles = append(roles, model.Role{
				ID:    utils.StrToNum(splitID[2]),
				Title: eachRole.Text(),
			})
		}
	})
	return roles
}
