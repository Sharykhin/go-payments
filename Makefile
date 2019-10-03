.PHONY: serve stop stats migration migrate-up migrate-down migrate-status

serve:
	docker-compose up

stop:
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
