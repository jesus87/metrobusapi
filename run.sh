export VERSION="1.0"
export APINAME="graphql"
export SCHEDULER_NAME="scheduler"
export SCHEDULER_TIME="10s"
export SCHEDULER_ATTEMPS=5
export NAMESPACE="jesus87"
export GO111MODULE=on
export GRAPHQL_PORT=8199
export METROBUS_API_URL="https://datos.cdmx.gob.mx/api/records/1.0/search/?"
export POSITIONS_API_PAGESIZE=10
export DB_CONTROLLER="mysql"
export CONNECTION_STRING="root:P@ssw0rd@tcp(metrobusdb-mysql:3306)/metrobusdb?parseTime=true"

echo Step 1/12 : Shuting down enviroment
docker-compose down

echo Step 2/12 : Building Scheduler 
cd scheduler
GOOS=linux CGO_ENABLED=0 go build -ldflags "-s -w" -o build/$SCHEDULER_NAME .

echo Step 3/12 : Attempting to stop container $SCHEDULER_NAME
docker stop $SCHEDULER_NAME

echo Step 4/12 : Remove previous container if exist $SCHEDULER_NAME
docker rm $SCHEDULER_NAME

echo Step 5/12 : Remove previous docker image if exist $SCHEDULER_NAME
docker rmi $NAMESPACE/$SCHEDULER_NAME:$VERSION

echo Step 6/12 : Building docker image
docker build -t $NAMESPACE/$SCHEDULER_NAME:$VERSION .
cd ..

echo Step 7/12 : Building Graphql API 
cd graphql
GOOS=linux CGO_ENABLED=0 go build -ldflags "-s -w" -o build/$APINAME .

echo Step 8/12 : Attempting to stop container $APINAME
docker stop $APINAME

echo Step 9/12 : Remove previous container if exist $APINAME
docker rm $APINAME

echo Step 10/12 : Remove previous docker image if exist $APINAME
docker rmi $NAMESPACE/$APINAME:$VERSION

echo Step 11/12 : Building docker image
docker build -t $NAMESPACE/$APINAME:$VERSION .
cd ..

echo Step 12/12 :  Creating new enviroment
docker-compose up -d