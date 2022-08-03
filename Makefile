POSTGRESQL_URL='postgres://postgres:537j04222@localhost:5432/postgres?sslmode=disable'

up:
	migrate -database ${POSTGRESQL_URL} -path db/migrations up
down:
	migrate -database ${POSTGRESQL_URL} -path db/migrations down