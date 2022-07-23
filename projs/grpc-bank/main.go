package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"

	"example.xyz/bank/api"
	"example.xyz/bank/internal/db"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:password@localhost:5432/simple_bank?sslmode=disable"
	address  = "0.0.0.0:9001"
)

func main() {

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal(err)
	}
	server := api.NewServer(db.NewStore(conn))
	if err := server.Start(address); err != nil {
		log.Fatal(err)
	}
}
