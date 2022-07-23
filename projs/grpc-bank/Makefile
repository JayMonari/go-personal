postgres:
	docker run --name bank_psql -p 5432:5432 -e POSTGRES_PASSWORD=password -d postgres:12-alpine

createdb:
	docker exec -it bank_psql createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -it bank_psql dropdb -U postgres simple_bank

migrateup:
	migrate -path db/migration -database "postgres://postgres:password@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgres://postgres:password@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	mkdir -p /tmp/gocover
	go test -v -timeout 10s -coverprofile=/tmp/gocover/codecover ./...

coverage:
	go tool cover -html=/tmp/gocover/codecover

server:
	go run main.go

mock:
	mockgen -package mockdb -destination internal/db/mock/store.go --build_flags=--mod=mod example.xyz/bank/internal/db Store

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test coverage server mock
