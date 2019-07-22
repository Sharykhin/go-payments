.PHONY: dev

dev:
	docker-compose -f docker-compose.dev.yml up

api:
	API_ADDR=:8000 DATABASE_HOST=localhost DATABASE_PORT=54320 DATABASE_USER=root DATABASE_PASSWORD=root DATABASE_NAME=payments go run cmd/api/main.go

web:
	WEB_ADDR=:8081 API_ADDR=:8000 go run cmd/web/main.go

migration:
	goose -dir migrations create ${NAME} sql

migrate:
	goose -dir migrations postgres "host=localhost user=root password=root dbname=payments sslmode=disable port=54320" up
