# METROBUS SERVICE

## PREREQUISITES

* GO Compiler
* Same repository path using $GOPATH

## GLOBAL VARIABLES

Following Global enviroment variables should be defined in order to execute the service

|Enviroment Variable| Description|
|----|----|
|NAMESPACE| Containers name space in case you want to push them|
|VERSION| Containers version - all containers share same version|
|METROBUS_API_URL| CDMX Api Url |

## SCHEDULE SERVICE

Polling metrobus information from source:
* Alcaldias Catalog
* Vehicles position history

### Requirements

|Enviroment Variable| Description|
|----|----|
|SCHEDULER_NAME| Container and build file names|
|SCHEDULER_TIME| Execution interval|
|SCHEDULER_ATTEMPTS| Max execution attempts in case of failure|
|METROBUS_API_URL| CDMX Api Url |
|POSITIONS_API_PAGESIZE|Page Size for vehicles positions history|

## GRAPHQL API

GraphQL Api for querying metrobus information as:
* Alcaldia by geo point
* Available vehicles
* Current position
* Position history by vehicle

### Requirements

|Enviroment Variable| Description|
|----|----|
|APINAME| Container and build file names|
|GRAPHQL_PORT| Access port for Graphql Api|
|METROBUS_API_URL| CDMX Api Url |

Previous requirements are defined on run.sh script.

## RUN ENVIROMENT STACK.

In order to run Metrobus Service you may need to execute run.sh script, it will handle al the environment using docker compose, cleaning previous versions and re building the code creating new images.
