package usecase

import (
	"log"

	"github.com/jesus87/apidf/src/domain/entity"
	"github.com/jesus87/apidf/src/infrastructure/apiintegration/metrobus"
	"github.com/jesus87/apidf/src/infrastructure/persistance"
)

type FetchAlcaldiasUseCase struct {
	_metrobusService    *metrobus.MetrobusService
	_metrobusRepository *persistance.MetrobusRepository
}

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

func (u *FetchAlcaldiasUseCase) GetAlcaldias() ([]entity.Alcaldia, error) {
	return u._metrobusRepository.GetAlcaldias()
}

func NewFetchAlcaldiasUseCase(
	metrobusService *metrobus.MetrobusService,
	metrobusRepository *persistance.MetrobusRepository,
) *FetchAlcaldiasUseCase {
	return &FetchAlcaldiasUseCase{
		_metrobusService:    metrobusService,
		_metrobusRepository: metrobusRepository,
	}
}
