package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"task/cmd"
	"task/db"
)

func main() {
	cache, err := os.UserCacheDir()
	if err != nil {
		log.Fatal(err)
	}
	dbPath := filepath.Join(cache, "tasks.db")
	must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())
}

func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
