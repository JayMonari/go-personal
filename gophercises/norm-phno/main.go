package main

import (
	"database/sql"
	"fmt"
	phonedb "norm/db"
	"strings"
	"unicode"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "user"
	password = "password"
	dbname   = "phone"
)

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s sslmode=disable",
		host, port, user, password)
	must(phonedb.Reset("postgres", dsn, dbname))

	dsn = fmt.Sprintf("%s dbname=%s", dsn, dbname)
	must(phonedb.Migrate("postgres", dsn))
	db, err := phonedb.Open("postgres", dsn)
	must(err)
	defer db.Close()

	must(db.Seed())

	phones, err := db.AllPhones()
	must(err)
	for _, p := range phones {
		fmt.Printf("Working on... %+v\n", p)
		number := normalize(p.Number)
		if number != p.Number {
			existing, err := db.FindPhone(number)
			must(err)
			if existing != nil {
				must(db.DeletePhone(p.ID))
			} else {
				p.Number = number
				must(db.UpdatePhone(&p))
			}
		} else {
			fmt.Println("No need to format")
		}
	}
}

func getPhone(db *sql.DB, id int) (string, error) {
	var phNo string
	must(db.QueryRow("SELECT value FROM phone_numbers WHERE id=$1", id).Scan(&phNo))
	return phNo, nil
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

func normalize(phoneNo string) string {
	b := strings.Builder{}
	for _, r := range phoneNo {
		if unicode.IsDigit(r) {
			b.WriteRune(r)
		}
	}
	return b.String()
}
