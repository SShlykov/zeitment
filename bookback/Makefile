include .env
export

run:
	go run ./cmd/bookback/main.go

.PHONY: create-migration

bin:
	go build -o ./bin/bookback ./cmd/bookback/main.go

docs:
	swag init -g ./cmd/bookback/main.go

dock:
	docker compose -f ./deployments/docker-compose.yml up --build

lint:
	golangci-lint run

migration:
ifeq ($(mn),)
	$(error mn is not set)
endif
	mkdir -p migrations
	goose -dir database/migrations create $(mn) sql

migrate:
	goose -dir migrations postgres ... up

gen_mock:
	mockgen \
 	-source=internal/domain/services/chapter_service.go \
 	-destination=tests/mocks/domain/services/chapter_service_mock.go \
 	-package=mocks
