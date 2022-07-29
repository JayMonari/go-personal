postgres:
	docker run --name bank_psql --network bank-network -p 5432:5432 -e POSTGRES_PASSWORD=password -d postgres:12-alpine

createdb:
	docker exec -it bank_psql createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -it bank_psql dropdb -U postgres simple_bank

migrateup:
	migrate -path db/migration -database "postgres://postgres:password@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://postgres:password@localhost:5432/simple_bank?sslmode=disable" -verbose down

migrateup1:
	migrate -path db/migration -database "postgres://postgres:password@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown1:
	migrate -path db/migration -database "postgres://postgres:password@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	mkdir -p /tmp/gocover
	go test -v -timeout 10s -coverprofile=/tmp/cover.out ./...

coverage:
	go tool cover -html=/tmp/cover.out

server:
	go run main.go

gen:
	go generate ./...

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test coverage server gen migrateup1 migratedown1
