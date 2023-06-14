postgres:
	docker run --name master-class -p 5432:5432 -e POSTGRES_PASSWORD=secret -d postgres:alpine
createdb:
	docker exec -it master-class createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -it master-class dropdb -U postgres simple_bank

migratup:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratdown:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: postgres createdb dropdb migratup migratdown sqlc test