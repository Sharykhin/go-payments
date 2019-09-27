.PHONY: serve stop migration migrate migrate-down migrate-status

serve:
	docker-compose up

stop:
	docker-compose down

migration:
	docker-compose run migration goose -dir /migrations create ${NAME} sql

migrate:
	docker-compose run migration goose -dir /migrations postgres "host=postgres user=root password=root dbname=payments sslmode=disable port=5432" up

migrate-down:
	docker-compose run migration goose -dir /migrations postgres "host=postgres user=root password=root dbname=payments sslmode=disable port=5432" down

migrate-status:
	docker-compose run migration goose -dir /migrations postgres "host=postgres user=root password=root dbname=payments sslmode=disable port=5432" status
