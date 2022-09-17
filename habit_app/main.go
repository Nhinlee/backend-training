package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"habits.com/habit/api"
	db "habits.com/habit/db/sqlc"
)

const (
	dbDriver      = "postgres"
	dbSource      = "postgresql://root:nhin123456@localhost:5432/habits?sslmode=disable"
	serverAddress = "0.0.0.0:8080"
)

func main() {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(serverAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
