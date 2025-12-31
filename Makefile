postgres: 
	docker run --name postgres-db \
	-e POSTGRES_USER=postgres \
	-e POSTGRES_PASSWORD=postgres \
	-e POSTGRES_DB=postgres \
	-p 5432:5432 \
	-d postgres:17

createdb: 
	docker exec -it postgres-db createdb --username=postgres --owner=postgres simple_bank

dropdb: 
	docker exec -it postgres-db dropdb simple_bank

migrateup: 
	migrate -path ./db/migration -database "postgresql://postgres:postgres@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown: 
	migrate -path ./db/migration -database "postgresql://postgres:postgres@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc: 
	sqlc generate

test: 
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migrateup sqlc test
