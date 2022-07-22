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

.PHONY: postgres createdb dropdb migrateup migratedown
