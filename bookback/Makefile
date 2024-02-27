.PHONY: create-migration

bin:
	go build -o ./bin/bookback ./cmd/bookback/main.go

run:
	DSN="host=77.91.126.204 port=5429 user=postgres password=secret dbname=book_dev sslmode=disable" go run ./cmd/bookback/main.go

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
	goose -dir migrations create $(mn) sql

migrate:
	goose -dir migrations postgres "postgresql://postgres:secret@77.91.126.204:5429/book_dev?sslmode=disable" up

gen_mock:
	mockgen \
 	-source=pkg/db/db.go \
 	-destination=internal/mocks/db.go \
 	-package=mocks
