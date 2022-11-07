postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=nhin123456 -d postgres:14.5-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root habits

dropdb:
	docker exec -it postgres dropdb habits

migrate-up:
	migrate -path db/migration -database "postgresql://root:nhin123456@localhost:5432/habits?sslmode=disable" -verbose up

migrate-up1:
	migrate -path db/migration -database "postgresql://root:nhin123456@localhost:5432/habits?sslmode=disable" -verbose up 1

migrate-down:
	migrate -path db/migration -database "postgresql://root:nhin123456@localhost:5432/habits?sslmode=disable" -verbose down

migrate-down1:
	migrate -path db/migration -database "postgresql://root:nhin123456@localhost:5432/habits?sslmode=disable" -verbose down 1

test:
	go test -v -cover ./...

sqlc:
	sqlc generate

gen-proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto

evans:
	evans --host localhost --port 9090 -r repl


server:
	nodemon --exec go run main.go --signal SIGTERM

.PHONY: postgres createdb dropdb migrate-up migrate-down migrate-up1 migrate-down1 sqlc test server gen-proto evans