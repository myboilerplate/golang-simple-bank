postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -d postgres

migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose down

migratedrop:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable" -verbose drop -f

sqlc:
	sqlc generate

cleanup:
	go mod tidy

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/myboilerplate/golang-simple-bank/db/sqlc Store

test:
	go test -v -cover -short ./...

.PHONY: postgres migrateup migratedown migratedrop sqlc test server mock