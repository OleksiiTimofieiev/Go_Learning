.PHONY: build

build:
	@ go build -v -o "apiServer" ./cmd/*.go

.PHONY: test
test:
	@ go test -v -race -timeout 30s "../src/REST_Api/apiserver/"
	@ go test -v -race -timeout 30s "../src/REST_Api/store/sqlstore"
	@ go test -v -race -timeout 30s "../src/REST_Api/store/teststore"
	@ go test -v -race -timeout 30s "../src/REST_Api/model/"


# sudo service postgresql start
# migrate create -ext sql -dir ./src/REST_Api/migrations create_user
# migrate -path ./src/REST_Api/migrations/ -database "postgres://postgres:postgres@localhost:5432/restapi_dev?sslmode=disable" up
.DEFAULT_GOAL := build