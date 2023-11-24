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

test:
	go test -v -cover -short ./...