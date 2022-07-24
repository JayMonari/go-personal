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
	svr, err:= api.NewServer(cfg, db.NewStore(conn))
	if err != nil {
		log.Fatal(err)
	}
	if  err = svr.Start(cfg.ServerAddress); err != nil {
		log.Fatal(err)
	}
}
