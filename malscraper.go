package malscraper

import (
	"github.com/rl404/go-malscraper/model/general"
	"github.com/rl404/go-malscraper/model/user"
)

// MalScraper is the main function to call all method in go-malscraper.
func MalScraper() *MalService {
	m := &MalService{}
	return m
}

// MalService for all go-malscraper service.
type MalService struct {
	User    *UserService
	General *GeneralService
	// Additional	*AdditionalService
	// Lists		*ListsService
	// Search		*SearchService
	// Seasonal		*SeasonalService
	// Top			*TopService
}

// UserService for all user-related method.
type UserService struct{}

// GetUser to get user profile information.
func (u *UserService) GetUser(username string) (user.UserData, int, string) {
	var userModel user.UserModel
	return userModel.InitUserModel(username)
}

// GeneralService for all general method.
type GeneralService struct{}

// GetInfo to get anime & manga information.
func (g *GeneralService) GetInfo(t string, id int) (general.InfoData, int, string) {
	var InfoModel general.InfoModel
	return InfoModel.InitInfoModel(t, id)
}

// GetCharacter to get character information.
func (g *GeneralService) GetCharacter(id int) (general.CharacterData, int, string) {
	var CharacterModel general.CharacterModel
	return CharacterModel.InitCharacterModel(id)
}

// GetPeople to get people information.
func (g *GeneralService) GetPeople(id int) (general.PeopleData, int, string) {
	var PeopleModel general.PeopleModel
	return PeopleModel.InitPeopleModel(id)
}

// GetProducer to get studio/producer information.
func (g *GeneralService) GetProducer(id int, page int) ([]general.ProducerData, int, string) {
	var ProducerModel general.ProducerModel
	return ProducerModel.InitProducerModel("anime", "producer", id, page)
}

// GetMagazine to get magazine information.
func (g *GeneralService) GetMagazine(id int, page int) ([]general.ProducerData, int, string) {
	var ProducerModel general.ProducerModel
	return ProducerModel.InitProducerModel("manga", "producer", id, page)
}

// GetGenre to get magazine information.
func (g *GeneralService) GetGenre(t string, id int, page int) ([]general.ProducerData, int, string) {
	var ProducerModel general.ProducerModel
	return ProducerModel.InitProducerModel(t, "genre", id, page)
}
