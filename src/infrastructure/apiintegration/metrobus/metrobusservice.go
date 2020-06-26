package metrobus

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/jesus87/metrobusapi/src/domain/entity"
)

//MetrobusService type for manipulate metrobusservice
type MetrobusService struct {
	_url string
}

//FetchPositions fecht positions from  df api
func (m *MetrobusService) FetchPositions(pageSize int) ([]*entity.VehiclePosition, error) {

	const queryString = "dataset=prueba_fetchdata_metrobus&q=&rows="
	_url := fmt.Sprintf("%v%v%v", m._url, queryString, pageSize)

	body, err := m.fetchData(_url)
	if err != nil {
		return nil, err
	}

	var positions entity.Positions
	err = json.Unmarshal(body, &positions)
	if err != nil {
		return nil, err
	}

	return positions.VehiclePositions, nil
}

//FetchAlcaldias fecth alcaldias from df api
func (m *MetrobusService) FetchAlcaldias() ([]*entity.Alcaldia, error) {
	const queryString = "dataset=alcaldias&q=&facet=nomgeo&facet=cve_mun&facet=municipio"
	_url := fmt.Sprintf("%v%v", m._url, queryString)

	body, err := m.fetchData(_url)
	if err != nil {
		return nil, err
	}

	var catalog entity.AlcaldiaCatalog
	err = json.Unmarshal(body, &catalog)
	if err != nil {
		return nil, err
	}

	return catalog.Alcaldias, nil
}

// fetchData method for fetching retriving data from api
func (m *MetrobusService) fetchData(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func NewMetrobusService(url string) *MetrobusService {
	return &MetrobusService{
		_url: url,
	}
}
