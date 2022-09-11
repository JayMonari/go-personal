package main

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"

	"example.xyz/bank/api"
	"example.xyz/bank/internal/db"
	"example.xyz/bank/internal/util"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
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

	driver, err := postgres.WithInstance(conn, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
	}
	m, err := migrate.NewWithDatabaseInstance("file://migration", cfg.DBDriver, driver)
	if err != nil {
		log.Fatal(err)
	}
	m.Up()
	defer m.Down()

	svr, err := api.NewServer(cfg, db.NewStore(conn))
	if err != nil {
		log.Fatal(err)
	}

	if err = svr.Start(cfg.ServerAddress); err != nil {
		log.Fatal(err)
	}
}
