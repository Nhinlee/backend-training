package sqlc_test

import (
	"database/sql"
	"log"
	"os"
	"testing"
	"v1/config"
	db "v1/db/sqlc"

	_ "github.com/lib/pq"
)

var testQueries *db.Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	// load configs
	config, err := config.LoadConfig("../../..")
	if err != nil {
		return
	}

	testDB, err = sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	testQueries = db.New(testDB)

	os.Exit(m.Run())
}
