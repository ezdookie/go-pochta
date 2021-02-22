MIGRATIONS_DIR = ./migrations
DB_CONNECTION = postgres://postgres:postgres@127.0.0.1:5432/pochta?sslmode=disable

create-migration:
	migrate create -ext sql -dir ${MIGRATIONS_DIR} -seq ${name}

migrate:
	migrate -path ${MIGRATIONS_DIR} -database "${DB_CONNECTION}" up

down-migrate:
	migrate -path ${MIGRATIONS_DIR} -database "${DB_CONNECTION}" down