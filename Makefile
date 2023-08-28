DB_URL="postgresql://root:secret@localhost:5432/simple_shop?sslmode=disable"

postgres:
	docker run --name simple_shop_postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15.4-alpine3.18

createdb:
	docker exec -it simple_shop_postgres createdb --username=root --owner=root simple_shop

dropdb:
	docker exec -it simple_shop_postgres dropdb simple_shop

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt paths=source_relative \
    proto/*.proto

server:
	go run main.go

.PHONY: postgres createdb dropdb new_migration migrateup migrateup1 migratedown migratedown1 proto