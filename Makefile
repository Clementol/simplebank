postgres:
	docker run --name master-class --network bank-network -p 5432:5432 -e POSTGRES_PASSWORD=secret -d postgres:alpine
createdb:
	docker exec -it master-class createdb --username=postgres --owner=postgres simple_bank

dropdb:
	docker exec -it master-class dropdb -U postgres simple_bank

migratup:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratup1:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratdown:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratdown1:
	migrate -path db/migration -database "postgresql://postgres:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 1


sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -build_flags=--mod=mod -package mockdb -destination db/mock/store.go github.com/Clementol/simplebank/db/sqlc Store

.PHONY: postgres createdb dropdb migratup migratdown migratup1 migratdown1 sqlc test server
