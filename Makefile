.PHONY: serve stop migration migrate migrate-down migrate-status

serve:
	docker-compose up

stop:
	docker-compose down

migration:
	goose -dir migrations create ${NAME} sql

migrate:
	goose -dir migrations postgres "host=localhost user=root password=root dbname=payments sslmode=disable port=54320" up

migrate-down:
	goose -dir migrations postgres "host=localhost user=root password=root dbname=payments sslmode=disable port=54320" down

migrate-status:
	goose -dir migrations postgres "host=localhost user=root password=root dbname=payments sslmode=disable port=54320" status
