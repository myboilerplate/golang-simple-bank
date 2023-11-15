package db

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:postgres@localhost:5432/postgres?sslmode=disable"
)

var testQuery *Queries
var testDB *sql.DB

func TestMain(m *testing.M) {
	var err error

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	testQuery = New(conn)
	testDB = conn

	os.Exit(m.Run())
}
