postgres:
	docker run --name postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=nhin123456 -d postgres:14.5-alpine

createdb:
	docker exec -it postgres createdb --username=root --owner=root habits

dropdb:
	docker exec -it postgres dropdb habits

migrate-up:
	migrate -path db/migration -database "postgresql://root:nhin123456@localhost:5432/hobbies?sslmode=disable" -verbose up

migrate-down:
	migrate -path db/migration -database "postgresql://root:nhin123456@localhost:5432/hobbies?sslmode=disable" -verbose down

sqlc:
	sqlc generate
.PHONY: postgres createdb dropdb migrate-up migrate-down sqlc