package usecase

import (
	"testing"

	"github.com/jesus87/apidf/src/infrastructure/apiintegration/metrobus"
	"github.com/jesus87/apidf/src/infrastructure/orm/sqlxvendor"
	"github.com/jesus87/apidf/src/infrastructure/persistance"
)

func TestFetchAlcaldiasSuccess(t *testing.T) {
	metrobusservice := metrobus.NewMetrobusService("https://datos.cdmx.gob.mx/api/records/1.0/search/?")

	orm := sqlxvendor.NewSqlxVendor("mysql", "root:P@ssw0rd@/metrobusdb?parseTime=true")
	repository := persistance.NewMetrobusRepository(orm)

	usecase := NewFetchAlcaldiasUseCase(metrobusservice, repository)
	err := usecase.Fetch()
	if err != nil {
		t.Fail()
	}

}
