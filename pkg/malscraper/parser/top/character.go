package top

import (
	"strconv"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/pkg/malscraper/config"
	"github.com/rl404/go-malscraper/pkg/malscraper/constant"
	model "github.com/rl404/go-malscraper/pkg/malscraper/model/top"
	"github.com/rl404/go-malscraper/pkg/malscraper/parser"
	"github.com/rl404/go-malscraper/pkg/malscraper/utils"
)

// CharacterParser is parser for MyAnimeList top character list.
// Example: https://myanimelist.net/character.php
type CharacterParser struct {
	parser.BaseParser
	Page int
	Data []model.Character
}

// InitCharacterParser to initiate all fields and data of CharacterParser.
func InitCharacterParser(config config.Config, page ...int) (character CharacterParser, err error) {
	character.Page = 0
	character.Config = config

	if len(page) > 0 {
		character.Page = 50 * (page[0] - 1)
	}

	// Checking to redis if using redis in config.
	// Redis key's pattern is `top-character:{page}`.
	redisKey := constant.RedisGetTopCharacter + ":" + strconv.Itoa(character.Page)
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
	err = character.InitParser("/character.php?limit="+strconv.Itoa(character.Page), "#content")
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

// setAllDetail to set all top character detail.
func (cp *CharacterParser) setAllDetail() {
	var characters []model.Character
	area := cp.Parser.Find(".characters-favorites-ranking-table")
	area.Find("tr.ranking-list").Each(func(i int, eachChar *goquery.Selection) {
		nameArea := eachChar.Find(".people")

		characters = append(characters, model.Character{
			Rank:         cp.getRank(eachChar),
			ID:           cp.getID(nameArea),
			Name:         cp.getName(nameArea),
			JapaneseName: cp.getJapaneseName(nameArea),
			Image:        cp.getImage(nameArea),
			Animeography: cp.getRole(eachChar, ".animeography"),
			Mangaography: cp.getRole(eachChar, ".mangaography"),
			Favorite:     cp.getFavorite(eachChar),
		})
	})

	cp.Data = characters
}

// getRank to get character rank.
func (cp *CharacterParser) getRank(eachChar *goquery.Selection) int {
	rank := eachChar.Find("td").First().Find("span").Text()
	return utils.StrToNum(rank)
}

// getID to get character id.
func (cp *CharacterParser) getID(nameArea *goquery.Selection) int {
	id, _ := nameArea.Find("a").First().Attr("href")
	id = utils.GetValueFromSplit(id, "/", 4)
	return utils.StrToNum(id)
}

// getName to get character name.
func (cp *CharacterParser) getName(nameArea *goquery.Selection) string {
	return nameArea.Find(".information a").Text()
}

// getJapaneseName to get character japanese name.
func (cp *CharacterParser) getJapaneseName(nameArea *goquery.Selection) string {
	japName := nameArea.Find(".information span").Text()

	if japName != "" {
		japName = japName[1 : len(japName)-1]
	}

	return japName
}

// getImage to get character image.
func (cp *CharacterParser) getImage(nameArea *goquery.Selection) string {
	image, _ := nameArea.Find("img").First().Attr("data-src")
	return utils.URLCleaner(image, "image", cp.Config.CleanImageURL)
}

// getRole to get character role in anime + manga.
func (cp *CharacterParser) getRole(eachChar *goquery.Selection, areaClass string) []model.Ography {
	var roleList []model.Ography
	area := eachChar.Find(areaClass)
	area.Find(".title").Each(func(i int, eachRole *goquery.Selection) {
		linkA := eachRole.Find("a").First()
		link, _ := linkA.Attr("href")
		id := utils.GetValueFromSplit(link, "/", 4)
		roleList = append(roleList, model.Ography{
			ID:    utils.StrToNum(id),
			Title: linkA.Text(),
		})
	})

	return roleList
}

// getFavorite to get character favorite number.
func (cp *CharacterParser) getFavorite(eachChar *goquery.Selection) int {
	fav := eachChar.Find(".favorites").Text()
	return utils.StrToNum(fav)
}
