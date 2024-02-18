.PHONY: create-migration

bin:
	go build -o ./bin/bookback ./cmd/bookback/main.go

run:
	DSN="host=10.102.37.150 user=admin password=adX12 dbname=book_dev sslmode=disable" go run ./cmd/bookback/main.go

docs:
	swag init -g ./cmd/bookback/main.go

dock:
	docker compose -f ./deployments/docker-compose.yml up --build

migration:
ifeq ($(mn),)
	$(error mn is not set)
endif
	mkdir -p migrations
	goose -dir migrations create $(mn) sql

migrate:
	goose -dir migrations postgres "postgresql://admin:adX12@10.102.37.150:5432/book_dev?sslmode=disable" up