package usecase

import (
	"log"

	"github.com/jesus87/metrobusapi/src/domain/entity"
	"github.com/jesus87/metrobusapi/src/infrastructure/apiintegration/metrobus"
	"github.com/jesus87/metrobusapi/src/infrastructure/persistance"
)

//FetchAlcaldiasUseCase tye for fetching data
type FetchAlcaldiasUseCase struct {
	_metrobusService    *metrobus.MetrobusService
	_metrobusRepository *persistance.MetrobusRepository
}

//Fetch method for fetching alcaldias information
func (u *FetchAlcaldiasUseCase) Fetch() error {

	alcaldias, err := u._metrobusService.FetchAlcaldias()
	if err != nil {
		return err
	}

	for _, alcaldia := range alcaldias {
		exist, err := u._metrobusRepository.ExistAlcaldia(alcaldia)
		if err != nil {
			log.Println(err)
			continue
		}

		if exist {
			continue
		}

		err = u._metrobusRepository.SaveAlcaldia(alcaldia)
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}

//GetAlcaldias method or getting alcaldias from the repository
func (u *FetchAlcaldiasUseCase) GetAlcaldias() ([]entity.Alcaldia, error) {
	return u._metrobusRepository.GetAlcaldias()
}

//NewFetchAlcaldiasUseCase instance for type fetch
func NewFetchAlcaldiasUseCase(
	metrobusService *metrobus.MetrobusService,
	metrobusRepository *persistance.MetrobusRepository,
) *FetchAlcaldiasUseCase {
	return &FetchAlcaldiasUseCase{
		_metrobusService:    metrobusService,
		_metrobusRepository: metrobusRepository,
	}
}
