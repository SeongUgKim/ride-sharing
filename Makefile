postgres:
	docker run --name postgres12 -p 5433:5433 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root uber

dropdb:
	docker exec -it postgres12 dropdb uber

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/uber?sslmode=disable" -verbose up

migratedown:
	 migrate -path db/migration -database "postgresql://root:secret@localhost:5433/uber?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover -short ./...
.PHONY: postgres createdb dropdb migrateup, migratedown sqlc test