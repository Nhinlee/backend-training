package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
	"habits.com/habit/api"
	db "habits.com/habit/db/sqlc"
	"habits.com/habit/utils"
)

func main() {
	// Load configs
	config, err := utils.LoadConfig(".")
	if err != nil {
		return
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)

	runHTTPServer(config, store)
}

func runHTTPServer(config utils.Config, store db.Store) {
	server, err := api.NewServer(&config, store)
	if err != nil {
		log.Fatal("cannot create server: ", err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}
