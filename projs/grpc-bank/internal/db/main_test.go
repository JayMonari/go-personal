package db_test

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"

	"example.xyz/bank/internal/db"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:password@localhost:5432/simple_bank?sslmode=disable"
)

var testQueries *db.Queries

func TestMain(m *testing.M) {
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal(err)
	}
	testQueries = db.New(conn)

	os.Exit(m.Run())
}
