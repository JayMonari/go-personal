package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"example.xyz/bank/api"
	"example.xyz/bank/internal/db"
	"example.xyz/bank/internal/util"
)

func main() {
	cfg, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal(err)
	}
	conn, err := sql.Open(cfg.DBDriver, cfg.DBSource)
	if err != nil {
		log.Fatal(err)
	}
	if err := api.NewServer(db.NewStore(conn)).Start(cfg.ServerAddress); err != nil {
		log.Fatal(err)
	}
}
