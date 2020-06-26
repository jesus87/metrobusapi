package metrobus

import "testing"

func TestFectchAlcaldiasSuccess(t *testing.T) {
	metrobusService := NewMetrobusService("https://datos.cdmx.gob.mx/api/records/1.0/search/?")

	_, err := metrobusService.FetchAlcaldias()
	if err != nil {
		t.Fail()
	}

}

func TestFectchPositionsSuccess(t *testing.T) {
	metrobusService := NewMetrobusService("https://datos.cdmx.gob.mx/api/records/1.0/search/?")

	_, err := metrobusService.FetchPositions(10)
	if err != nil {
		t.Fail()
	}

}
