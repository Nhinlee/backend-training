package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
	"habits.com/habit/utils"
)

var testQueries *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	// Load configs
	config, err := utils.LoadConfig("../..")
	if err != nil {
		return
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}
