package general

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/rl404/go-malscraper/helper"
	"github.com/rl404/go-malscraper/model"
)

// CharacterModel is an extended model from MainModel for character.
type CharacterModel struct {
	model.MainModel
	Id   int
	Data CharacterData
}

// InitCharacterModel to initiate fields in parent (MainModel) model.
func (i *CharacterModel) InitCharacterModel(id int) (CharacterData, int, string) {
	i.Id = id
	i.InitModel("/character/"+strconv.Itoa(i.Id), "#contentWrapper")

	i.Parser = i.GetParser(i.Url)

	if i.ResponseCode != 200 {
		return i.Data, i.ResponseCode, i.ErrorMessage
	}

	i.SetAllDetail()

	return i.Data, i.ResponseCode, i.ErrorMessage
}

// SetAllDetail to fill all character data.
func (i *CharacterModel) SetAllDetail() {
	i.SetId()
	i.SetImage()
	i.SetNickname()
	i.SetNameKanji()
	i.SetName()
	i.SetFavorite()
	i.SetAbout()
	i.SetMedia("anime")
	i.SetMedia("manga")
	i.SetVa()
}

// SetId to set character id.
func (i *CharacterModel) SetId() {
	i.Data.Id = strconv.Itoa(i.Id)
}

// SetImage to set character image.
func (i *CharacterModel) SetImage() {
	image := i.Parser.Find("#content table tr")
	image = i.Parser.Find("td div a img")

	i.Data.Image, _ = image.Attr("src")
}

// SetNickname to set character nickname.
func (i *CharacterModel) SetNickname() {
	nickname := i.Parser.Find("h1")

	r, _ := regexp.Compile(`\s+`)
	nickRegex := r.ReplaceAllString(nickname.Text(), " ")
	nickRegex = strings.TrimSpace(nickRegex)

	r, _ = regexp.Compile(`\"([^"])*`)
	nickRegex = r.FindString(nickRegex)

	if nickRegex != "" {
		i.Data.Nickname = nickRegex[1:]
	} else {
		i.Data.Nickname = ""
	}
}

// SetName to set character name.
func (i *CharacterModel) SetName() {
	areaName := i.Parser.Find("#content table tr td").Next()
	areaName = areaName.Find("div.normal_header[style^=height]")
	areaName.Find("span").Remove()

	i.Data.Name = strings.TrimSpace(areaName.Text())
}

// SetNameKanji to set character kanji name.
func (i *CharacterModel) SetNameKanji() {
	areaName := i.Parser.Find("#content table tr td").Next()
	areaName = areaName.Find("div.normal_header small")

	r, _ := regexp.Compile(`(\(|\))`)
	nameKanji := r.ReplaceAllString(areaName.Text(), "")

	i.Data.NameKanji = nameKanji
}

// SetFavorite to set number of user who favorite the character.
func (i *CharacterModel) SetFavorite() {
	favorite := i.Parser.Find("#content table tr td").Text()

	r, _ := regexp.Compile(`(Member Favorites: ).+`)
	regexFav := r.FindString(favorite)
	regexFav = strings.TrimSpace(regexFav)

	splitFav := strings.Split(regexFav, ": ")
	splitFav[1] = strings.Replace(splitFav[1], ",", "", -1)

	i.Data.Favorite = splitFav[1]
}

// SetAbout to set character about.
func (i *CharacterModel) SetAbout() {
	aboutArea := i.Parser.Find("#content table tr td").Next()
	aboutHtml, _ := aboutArea.Html()

	r, _ := regexp.Compile(`(?ms)(<div class="normal_header" style="height: 15px;">).*(<div class="normal_header">)`)
	regexAbout := r.FindString(aboutHtml)

	aboutGoQuery, _ := goquery.NewDocumentFromReader(strings.NewReader(regexAbout))

	aboutGoQuery.Find("div").Remove()

	cleanAbout := aboutGoQuery.Text()

	if cleanAbout != "No biography written." {
		i.Data.About = cleanAbout
	}
}

// SetMedia to set animeography and mangaography of the character.
func (i *CharacterModel) SetMedia(t string) {
	var mediaList []IdTitleImageRole

	mediaArea := i.Parser.Find("#content table tr td:nth-of-type(1)")

	if t == "anime" {
		mediaArea = mediaArea.Find("table:nth-of-type(1)")
	} else {
		mediaArea = mediaArea.Find("table:nth-of-type(2)")
	}

	mediaArea = mediaArea.Find("tr").Each(func(j int, eachMedia *goquery.Selection) {
		mediaImage := eachMedia.Find("td:nth-of-type(1)")
		eachArea := eachMedia.Find("td:nth-of-type(2)")

		mediaList = append(mediaList, IdTitleImageRole{
			Id:    i.GetVaId(eachArea),
			Title: i.GetVaName(eachArea),
			Image: i.GetVaImage(mediaImage),
			Role:  i.GetVaRole(eachArea),
		})
	})

	if t == "anime" {
		i.Data.Animeography = mediaList
	} else {
		i.Data.Mangaography = mediaList
	}
}

// SetVa to set VA of the character.
func (i *CharacterModel) SetVa() {
	var vaList []Staff

	vaArea := i.Parser.Find("#content table tr td").Next()
	vaArea.Find("table[width=\"100%\"]").Each(func(j int, eachVa *goquery.Selection) {
		vaNameArea := eachVa.Find("td:nth-of-type(2)")

		vaList = append(vaList, Staff{
			Id:    i.GetVaId(vaNameArea),
			Name:  i.GetVaName(vaNameArea),
			Role:  i.GetVaRole(vaNameArea),
			Image: i.GetVaImage(eachVa),
		})
	})

	i.Data.Va = vaList
}

// GetVaId to get animeography, mangaography and staff id.
func (i *CharacterModel) GetVaId(vaArea *goquery.Selection) string {
	id, _ := vaArea.Find("a").Attr("href")
	parsedId := strings.Split(id, "/")

	return parsedId[4]
}

// GetVaName to get animeography, mangaography and staff name.
func (i *CharacterModel) GetVaName(vaArea *goquery.Selection) string {
	name, _ := vaArea.Find("a:nth-of-type(1)").Html()
	return name
}

// GetVaImage to get animeography, mangaography and staff image.
func (i *CharacterModel) GetVaImage(vaArea *goquery.Selection) string {
	image, _ := vaArea.Find("img").Attr("src")
	return helper.ImageUrlCleaner(image)
}

// GetVaRole to get animeography, mangaography and staff role.
func (i *CharacterModel) GetVaRole(vaArea *goquery.Selection) string {
	return vaArea.Find("div small").Text()
}
