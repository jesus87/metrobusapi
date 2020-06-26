package main

import (
	"log"
	"os"
	"strconv"

	"github.com/jesus87/apidf/src/application/usecase"
	"github.com/jesus87/apidf/src/infrastructure/apiintegration/metrobus"
	"github.com/jesus87/apidf/src/infrastructure/orm/sqlxvendor"
	"github.com/jesus87/apidf/src/infrastructure/persistance"
)

func main() {

	log.Println("Running")
	url := os.Getenv("METROBUS_API_URL")
	pageSize, err := strconv.Atoi(os.Getenv("POSITIONS_API_PAGESIZE"))
	if err != nil {
		pageSize = 10
	}
	metrobusservice := metrobus.NewMetrobusService(url)

	orm := sqlxvendor.NewSqlxVendor("mysql", "root:P@ssw0rd@/metrobusdb?parseTime=true")
	repository := persistance.NewMetrobusRepository(orm)

	FetchAlcaldias := usecase.NewFetchAlcaldiasUseCase(metrobusservice, repository)
	FetchPositions := usecase.NewFetchPositionsUseCase(metrobusservice, repository, pageSize)
	err = FetchAlcaldias.Fetch()
	if err != nil {
		log.Println(err)
		panic(err)
	}

	err = FetchPositions.Fetch()
	if err != nil {
		log.Println(err)
		panic(err)
	}

}
