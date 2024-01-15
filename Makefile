postgres:
	docker run --name postgres16_1 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:16.1-alpine

createdb:
	docker exec -it postgres16_1 createdb --username=root --owner=root simple-bank

dropdb:
	docker exec -it postgres16_1 dropdb simple-bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple-bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple-bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

## a label for a set of commands that should be executed when the make createdb command is run
.PHONY: postgres createdb dropdb migrateup migratedown