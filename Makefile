.PHONY: build up stop stats migration migrate-up migrate-down migrate-status

build:
	docker-compose build

up:
	docker-compose up

down:
	docker-compose down

stats:
	docker stats $$(docker ps --filter network=go_payments --format="{{.Names}}")

migration:
	docker-compose run migration goose -dir /migrations create ${NAME} sql

migrate-up:
	docker-compose run migration goose -dir /migrations postgres "host=postgres user=root password=root dbname=payments sslmode=disable port=5432" up

migrate-down:
	docker-compose run migration goose -dir /migrations postgres "host=postgres user=root password=root dbname=payments sslmode=disable port=5432" down

migrate-status:
	docker-compose run migration goose -dir /migrations postgres "host=postgres user=root password=root dbname=payments sslmode=disable port=5432" status
