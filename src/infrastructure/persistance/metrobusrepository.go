package persistance

import (
	"fmt"
	"strings"

	"github.com/jesus87/metrobusapi/src/domain/entity"
	"github.com/jesus87/metrobusapi/src/infrastructure/orm"
)

type MetrobusRepository struct {
	_orm orm.OrmManager
}

func (r *MetrobusRepository) GetPositions(vehicleID int) ([]entity.VehiclePosition, error) {
	result := []entity.VehiclePosition{}

	restrictions := []string{}
	parameters := map[string]interface{}{}
	query := "SELECT * FROM positions"
	restrictions = append(restrictions, "vehicleId = :id")
	parameters["id"] = vehicleID

	if len(restrictions) > 0 {
		query = query + " WHERE " + strings.Join(restrictions, " AND ")
	}

	err := r._orm.Select(&result, query, parameters)

	return result, err
}

func (r *MetrobusRepository) ExistPosition(position *entity.VehiclePosition) (bool, error) {
	result := []entity.VehiclePosition{}

	restrictions := []string{}
	parameters := map[string]interface{}{}
	query := "SELECT id FROM positions"
	restrictions = append(restrictions, "id = :id")
	parameters["id"] = position.Id

	if len(restrictions) > 0 {
		query = query + " WHERE " + strings.Join(restrictions, " AND ")
	}

	err := r._orm.Select(&result, query, parameters)
	if err != nil {
		return false, err
	}

	if len(result) > 0 {
		return true, nil
	}
	return false, nil
}

func (r *MetrobusRepository) GetVehicles() ([]entity.Vehicle, error) {
	result := []entity.Vehicle{}

	parameters := map[string]interface{}{}
	query := "SELECT distinct vehicleId,label FROM positions"

	err := r._orm.Select(&result, query, parameters)

	return result, err
}

func (r *MetrobusRepository) SavePosition(position *entity.VehiclePosition) error {
	mapper := r.positionMapper(position)

	fields := []string{"id",
		"vehicleId",
		"startDate",
		"lastUpdate",
		"longitude",
		"parentTripId",
		"positionSpeed",
		"latitude",
		"routeId",
		"label",
		"positionOdometer",
		"tripId",
		"vehicleStatus",
		"alcaldia",
	}

	query := fmt.Sprintf("INSERT INTO positions (%v) values (:%v)", strings.Join(fields, ","), strings.Join(fields, ",:"))

	if err := r._orm.Save(query, mapper); err != nil {
		return err
	}

	return nil
}

func (r *MetrobusRepository) positionMapper(position *entity.VehiclePosition) map[string]interface{} {
	return map[string]interface{}{
		"id":               position.Id,
		"vehicleId":        position.PositionAttributes.VehicleId,
		"startDate":        position.PositionAttributes.StartDate,
		"lastUpdate":       position.PositionAttributes.LastUpdate,
		"longitude":        position.PositionAttributes.Longitude,
		"parentTripId":     position.PositionAttributes.ParentTripId,
		"positionSpeed":    position.PositionAttributes.PositionSpeed,
		"latitude":         position.PositionAttributes.Latitude,
		"routeId":          position.PositionAttributes.RouteId,
		"label":            position.PositionAttributes.Label,
		"positionOdometer": position.PositionAttributes.PostitionOdometer,
		"tripId":           position.PositionAttributes.TripId,
		"vehicleStatus":    position.PositionAttributes.VehicleStatus,
		"alcaldia":         position.PositionAttributes.Alcaldia,
	}
}

func (r *MetrobusRepository) GetAlcaldiaByPosition(longitude float64, latitude float64) (*entity.Alcaldia, error) {
	result := []entity.Alcaldia{}

	restrictions := []string{}
	parameters := map[string]interface{}{}
	query := "SELECT id, name FROM alcaldia"
	restrictions = append(restrictions, "ST_Contains(poligono, point(:latitude, :longitude)) = 1")
	parameters["latitude"] = longitude
	parameters["longitude"] = latitude

	query = query + " WHERE " + strings.Join(restrictions, " AND ")

	err := r._orm.Select(&result, query, parameters)
	if err != nil {
		return nil, err
	}

	if len(result) > 0 {
		return &result[0], nil
	}
	return nil, nil
}

func (r *MetrobusRepository) ExistAlcaldia(alcaldia *entity.Alcaldia) (bool, error) {
	result := []entity.Alcaldia{}

	restrictions := []string{}
	parameters := map[string]interface{}{}
	query := "SELECT id, name FROM alcaldia"
	restrictions = append(restrictions, "id = :id")
	parameters["id"] = alcaldia.Id

	if len(restrictions) > 0 {
		query = query + " WHERE " + strings.Join(restrictions, " AND ")
	}

	err := r._orm.Select(&result, query, parameters)
	if err != nil {
		return false, err
	}

	if len(result) > 0 {
		return true, nil
	}
	return false, nil
}

func (r *MetrobusRepository) GetAlcaldias() ([]entity.Alcaldia, error) {
	result := []entity.Alcaldia{}

	parameters := map[string]interface{}{}
	query := "SELECT * FROM alcaldia"

	err := r._orm.Select(&result, query, parameters)

	return result, err
}

func (r *MetrobusRepository) SaveAlcaldia(alcaldia *entity.Alcaldia) error {
	mapper := r.alcaldiaMapper(alcaldia)

	if err := r._orm.Save("INSERT INTO alcaldia (id,name, poligono) values (:id,:name,ST_PolygonFromText(:poligono))", mapper); err != nil {
		return err
	}

	return nil
}

func (r *MetrobusRepository) alcaldiaMapper(alcaldia *entity.Alcaldia) map[string]interface{} {
	return map[string]interface{}{
		"id":       alcaldia.Id,
		"name":     alcaldia.Attributes.Name,
		"poligono": alcaldia.GetPolygon(),
	}
}

func NewMetrobusRepository(orm orm.OrmManager) *MetrobusRepository {
	return &MetrobusRepository{
		_orm: orm,
	}
}
