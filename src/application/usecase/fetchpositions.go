package usecase

import (
	"log"

	"github.com/jesus87/metrobusapi/src/domain/entity"
	"github.com/jesus87/metrobusapi/src/infrastructure/apiintegration/metrobus"
	"github.com/jesus87/metrobusapi/src/infrastructure/persistance"
)

//FetchPositionsUseCase type for fetching positions
type FetchPositionsUseCase struct {
	_metrobusService    *metrobus.MetrobusService
	_metrobusRepository *persistance.MetrobusRepository
	_pageSize           int
}

//Fetch method for fetching positions
func (u *FetchPositionsUseCase) Fetch() error {

	positions, err := u._metrobusService.FetchPositions(u._pageSize)
	if err != nil {
		return err
	}

	for _, position := range positions {

		exist, err := u._metrobusRepository.ExistPosition(position)
		if err != nil {
			log.Println(err)
			continue
		}

		if exist {
			continue
		}

		alcaldia, err := u._metrobusRepository.GetAlcaldiaByPosition(position.PositionAttributes.Latitude, position.PositionAttributes.Longitude)
		if err != nil {
			log.Println(err)
		}

		position.Alcaldia = ""
		if alcaldia != nil {
			position.Alcaldia = alcaldia.Attributes.Name
		}

		err = u._metrobusRepository.SavePosition(position)
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}

//GetVehicles get vehicles from repoository
func (u *FetchPositionsUseCase) GetVehicles() ([]entity.Vehicle, error) {
	return u._metrobusRepository.GetVehicles()
}

//GetVehiclePositions get positions for vehicles from repository
func (u *FetchPositionsUseCase) GetVehiclePositions(vehicleID int) ([]entity.VehiclePosition, error) {
	return u._metrobusRepository.GetPositions(vehicleID)
}

//NewFetchPositionsUseCase instance for type fetchpositions
func NewFetchPositionsUseCase(
	metrobusService *metrobus.MetrobusService,
	metrobusRepository *persistance.MetrobusRepository,
	pageSize int,
) *FetchPositionsUseCase {
	return &FetchPositionsUseCase{
		_metrobusService:    metrobusService,
		_metrobusRepository: metrobusRepository,
		_pageSize:           pageSize,
	}
}
