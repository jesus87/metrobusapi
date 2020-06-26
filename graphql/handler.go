package main

import (
	"encoding/json"
	"os"
	"strconv"

	"github.com/jesus87/metrobusapi/src/application/usecase"
	"github.com/jesus87/metrobusapi/src/infrastructure/apiintegration/metrobus"
	"github.com/jesus87/metrobusapi/src/infrastructure/orm/sqlxvendor"
	"github.com/jesus87/metrobusapi/src/infrastructure/persistance"
	"github.com/valyala/fasthttp"

	"context"
	"net/http"
)

type GraphqlHandler struct {
}

//RequestOptions type for request option for getting query, vars and operations
type RequestOptions struct {
	Query         string                 `json:"query" url:"query" schema:"query"`
	Variables     map[string]interface{} `json:"variables" url:"variables" schema:"variables"`
	OperationName string                 `json:"operationName" url:"operationName" schema:"operationName"`
}

//getFromArgs get values from args
func getFromArgs(values *fasthttp.Args) *RequestOptions {
	query := values.Peek("query")
	if query != nil {
		// get variables map
		variables := make(map[string]interface{}, values.Len())
		variablesStr := values.Peek("variables")
		if variablesStr != nil {
			err := json.Unmarshal(variablesStr, &variables)
			if err != nil {
				return nil
			}
		}

		return &RequestOptions{
			Query:         string(query),
			Variables:     variables,
			OperationName: string(values.Peek("operationName")),
		}
	}

	return nil
}

// ContextHandler provides an entrypoint into executing graphQL queries with a
// user-provided context.
func (c *GraphqlHandler) ContextHandler(ctx context.Context, ctxreq *fasthttp.RequestCtx) {
	// get query
	reqOpt := getFromArgs(ctxreq.URI().QueryArgs())

	if reqOpt == nil {
		ctxreq.Response.SetStatusCode(http.StatusBadRequest)
		return
	}

	var buff []byte
	var err error
	switch reqOpt.OperationName {
	case "vehicle":
		buff, err = getAvailableVehicles()
		break
	case "positions":
		buff, err = getAlcaldias()
		break
	case "alcaldias":
		buff, err = getPositionHistory(reqOpt.Variables["vehicleid"].(int))
		break
	}

	if err != nil {
		ctxreq.Response.SetStatusCode(http.StatusInternalServerError)
		b, _ := json.MarshalIndent(err, "", "\t")
		ctxreq.Response.AppendBody(b)
		return
	}

	ctxreq.Response.SetStatusCode(http.StatusOK)
	ctxreq.Response.AppendBody(buff)
}

//getAvailableVehicles list vehicles
func getAvailableVehicles() ([]byte, error) {
	pageSize, err := strconv.Atoi(os.Getenv("POSITIONS_API_PAGESIZE"))
	if err != nil {
		pageSize = 10
	}
	metrobusservice := metrobus.NewMetrobusService(os.Getenv("METROBUS_API_URL"))

	orm := sqlxvendor.NewSqlxVendor(os.Getenv("DB_CONTROLLER"), os.Getenv("CONNECTION_STRING"))
	repository := persistance.NewMetrobusRepository(orm)
	useCase := usecase.NewFetchPositionsUseCase(metrobusservice, repository, pageSize)
	vehicles, err := useCase.GetVehicles()
	if err != nil {
		return nil, err
	}
	buff, err := json.MarshalIndent(vehicles, "", "\t")
	if err != nil {
		return nil, err
	}

	return buff, nil
}

//getPositionHistory get historical positions
func getPositionHistory(vehicleID int) ([]byte, error) {
	pageSize, err := strconv.Atoi(os.Getenv("POSITIONS_API_PAGESIZE"))
	if err != nil {
		pageSize = 10
	}
	metrobusservice := metrobus.NewMetrobusService(os.Getenv("METROBUS_API_URL"))

	orm := sqlxvendor.NewSqlxVendor(os.Getenv("DB_CONTROLLER"), os.Getenv("CONNECTION_STRING"))
	repository := persistance.NewMetrobusRepository(orm)

	useCase := usecase.NewFetchPositionsUseCase(metrobusservice, repository, pageSize)
	vehicles, err := useCase.GetVehiclePositions(vehicleID)
	if err != nil {
		return nil, err
	}
	buff, err := json.MarshalIndent(vehicles, "", "\t")
	if err != nil {
		return nil, err
	}

	return buff, nil
}

//getAlcaldias list alcaldias
func getAlcaldias() ([]byte, error) {
	metrobusservice := metrobus.NewMetrobusService(os.Getenv("METROBUS_API_URL"))

	orm := sqlxvendor.NewSqlxVendor(os.Getenv("DB_CONTROLLER"), os.Getenv("CONNECTION_STRING"))
	repository := persistance.NewMetrobusRepository(orm)

	useCase := usecase.NewFetchAlcaldiasUseCase(metrobusservice, repository)
	alcaldias, err := useCase.GetAlcaldias()
	if err != nil {
		return nil, err
	}
	buff, err := json.MarshalIndent(alcaldias, "", "\t")
	if err != nil {
		return nil, err
	}

	return buff, nil
}
//Handle handle incoming request
func (c *GraphqlHandler) Handle(ctx *fasthttp.RequestCtx) {
	c.ContextHandler(context.Background(), ctx)
}
//NewGraphqlHandler instance
func NewGraphqlHandler() *GraphqlHandler {
	return &GraphqlHandler{}
}
