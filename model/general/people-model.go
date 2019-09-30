package general

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/grokify/html-strip-tags-go"
	"github.com/rl404/go-malscraper/helper"
	"github.com/rl404/go-malscraper/model"
)

// PeopleModel is an extended model from MainModel for people.
type PeopleModel struct {
	model.MainModel
	Id   int
	Data PeopleData
}

// InitPeopleModel to initiate fields in parent (MainModel) model.
func (i *PeopleModel) InitPeopleModel(id int) (PeopleData, int, string) {
	i.Id = id
	i.InitModel("/people/"+strconv.Itoa(i.Id), "#contentWrapper")

	i.Parser = i.GetParser(i.Url)

	if i.ResponseCode != 200 {
		return i.Data, i.ResponseCode, i.ErrorMessage
	}

	i.SetAllDetail()

	return i.Data, i.ResponseCode, i.ErrorMessage
}

// SetAllDetail to fill all people data.
func (i *PeopleModel) SetAllDetail() {
	i.SetId()
	i.SetName()
	i.SetImage()
	i.CleanBiodata()
	i.SetBiodata()
	i.SetMore()
	i.SetVa()
	i.SetStaff(false)
	i.SetStaff(true)
}

// SetId to set people id.
func (i *PeopleModel) SetId() {
	i.Data.Id = strconv.Itoa(i.Id)
}

// SetName to set people name.
func (i *PeopleModel) SetName() {
	i.Data.Name = i.Parser.Find("h1").Text()
}

// SetImage to set people name.
func (i *PeopleModel) SetImage() {
	i.Data.Image, _ = i.Parser.Find("#content table tr td img").Attr("src")
}

// CleanBiodata to clean people biodata area.
func (i *PeopleModel) CleanBiodata() {
	bioArea := i.Parser.Find("#content table tr td")
	bioArea.Find("div").EachWithBreak(func(j int, eachTrash *goquery.Selection) bool {
		if j < 5 {
			eachTrash.Remove()
			return true
		} else {
			return false
		}
	})
}

// SetBiodata to set people detail biodata.
func (i *PeopleModel) SetBiodata() {
	i.Data.GivenName, _ = i.GetBiodata("Given name")
	i.Data.FamilyName, _ = i.GetBiodata("Family name")
	_, i.Data.AlternativeName = i.GetBiodata("Alternate names")
	i.Data.Birthday, _ = i.GetBiodata("Birthday")
	i.Data.Website = i.GetBioWeb()
	i.Data.Favorite, _ = i.GetBiodata("Member Favorites")
}

// GetBiodata to get people each detail biodata.
func (i *PeopleModel) GetBiodata(t string) (string, []string) {
	bioArea, _ := i.Parser.Find("#content table tr td").Html()

	r, _ := regexp.Compile(`(` + t + `:<\/span>)[^<]*`)
	bioRegex := r.FindString(bioArea)
	bioRegex = strip.StripTags(bioRegex)

	if bioRegex != "" {
		splitBio := strings.Split(bioRegex, ": ")

		splitBio[1] = strings.TrimSpace(splitBio[1])

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

// GetBioWeb to get people website.
func (i *PeopleModel) GetBioWeb() string {
	bioArea, _ := i.Parser.Find("#content table tr td").Html()

	r, _ := regexp.Compile(`(Website:<\/span> <a)[^<]*`)
	bioRegex := r.FindString(bioArea)

	r, _ = regexp.Compile(`".+"`)
	bioRegex = r.FindString(bioRegex)

	if bioRegex != "http://" {
		return strings.Replace(bioRegex, "\"", "", -1)
	}

	return ""
}

// SetMore to set people more detail biodata.
func (i *PeopleModel) SetMore() {
	i.Data.More = i.Parser.Find("#content table tr td div[class^=people-informantion-more]").Text()
}

// SetVa to set people voice actor list
func (i *PeopleModel) SetVa() {
	var animeCharacter []AnimeCharacter

	vaArea := i.Parser.Find("#content table tr td").Next()
	vaArea = vaArea.Find(".normal_header").First().Next()

	if goquery.NodeName(vaArea) == "table" {
		vaArea.Find("tr").Each(func(j int, eachVa *goquery.Selection) {
			animeImageArea := eachVa.Find("td:nth-of-type(1)")
			animeArea := eachVa.Find("td:nth-of-type(2)")
			charImageArea := eachVa.Find("td:nth-of-type(4)")
			charArea := eachVa.Find("td:nth-of-type(3)")

			animeCharacter = append(animeCharacter, AnimeCharacter{
				Anime: IdImageTitle{
					Id:    i.GetAnimeId(animeArea),
					Image: i.GetAnimeImage(animeImageArea),
					Title: i.GetAnimeTitle(animeArea),
				},
				Character: Staff{
					Id:    i.GetAnimeId(charArea),
					Name:  i.GetAnimeTitle(charArea),
					Role:  i.GetAnimeRole(charArea),
					Image: i.GetAnimeImage(charImageArea),
				},
			})
		})

		i.Data.Va = animeCharacter
	}
}

// GetAnimeId to get anime id of the voice actor play.
func (i *PeopleModel) GetAnimeId(animeArea *goquery.Selection) string {
	animeId, _ := animeArea.Find("a").Attr("href")
	splitId := strings.Split(animeId, "/")

	return splitId[4]
}

// GetAnimeImage to get anime image of the voice actor play.
func (i *PeopleModel) GetAnimeImage(animeImageArea *goquery.Selection) string {
	animeImage, _ := animeImageArea.Find("img").Attr("data-src")
	return helper.ImageUrlCleaner(animeImage)
}

// GetAnimeTitle to get anime title of the voice actor play.
func (i *PeopleModel) GetAnimeTitle(animeArea *goquery.Selection) string {
	return animeArea.Find("a").First().Text()
}

// GetAnimeRole to get anime role of the voice actor play.
func (i *PeopleModel) GetAnimeRole(animeArea *goquery.Selection) string {
	return strings.TrimSpace(animeArea.Find("div").Text())
}

// SetStaff to set people staff list.
func (i *PeopleModel) SetStaff(isManga bool) {
	var staffList []IdTitleImageRole

	staffArea := i.Parser.Find("#content table tr td").Next()

	staffArea.Find(".normal_header").First().Remove()

	if isManga {
		staffArea = staffArea.Find(".normal_header").First().Next()
	} else {
		staffArea = staffArea.Find(".normal_header").First().Next()
	}

	if goquery.NodeName(staffArea) == "table" {
		staffArea.Find("tr").Each(func(j int, eachStaff *goquery.Selection) {
			animeImageArea := eachStaff.Find("td:nth-of-type(1)")
			stArea := eachStaff.Find("td:nth-of-type(2)")

			staffList = append(staffList, IdTitleImageRole{
				Id:    i.GetAnimeId(stArea),
				Title: i.GetAnimeTitle(stArea),
				Image: i.GetAnimeImage(animeImageArea),
				Role:  i.GetStaffRole(stArea),
			})

		})

		if isManga {
			i.Data.PublishedManga = staffList
		} else {
			i.Data.Staff = staffList
		}
	}
}

// GetStaffRole to get anime title of the voice actor play.
func (i *PeopleModel) GetStaffRole(stArea *goquery.Selection) string {
	return strings.TrimSpace(stArea.Find("small").Text())
}
