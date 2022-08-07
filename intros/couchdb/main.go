package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"text/template"

	_ "github.com/go-kivik/couchdb/v3"
	"github.com/go-kivik/kivik/v3"
)

var tmpl = template.Must(template.ParseFiles("customers.html"))

type customer struct {
	ID  string `json:"id"`
	Rev string `json:"rev"`

	Name    string `json:"name"`
	Email   string `json:"email"`
	Address struct {
		Street string `json:"street"`
		City   string `json:"city"`
		State  string `json:"state"`
	}
}

func main() {
	client, err := kivik.New("couch", "http://admin:password@localhost:5984/")
	if err != nil {
		log.Fatal(err)
	}
	db := client.DB(context.TODO(), "customers")
	if db.Err() != nil {
		log.Fatal(err)
	}

	// ServerMux
	s := http.NewServeMux()
	s.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rows, err := db.Query(context.TODO(), "_design/all_customers", "_view/all")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "There was a problem fetching all the rows: %v", err)
			return
		}
		var cust []customer
		for rows.Next() {
			var c customer
			if err := rows.ScanValue(&c); err != nil {
				panic(err)
			}
			c.ID = rows.Key()
			cust = append(cust, c)
		}
		tmpl.Execute(w, cust)
	}))

	s.Handle("/customer/add", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		c := customer{
			Name:  r.Form["name"][0],
			Email: r.Form["email"][0],
		}
		if _, _, err := db.CreateDoc(context.TODO(), c); err != nil {
			fmt.Fprintf(w, "Could not create customer: %v\n", err)
			return
		}
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}))

	s.Handle("/customer/delete", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id := r.URL.Query().Get("id")
		rev := r.URL.Query().Get("rev")
		if _, err := db.Delete(context.TODO(), id, rev); err != nil {
			fmt.Printf("could not delete customer ID: %q, rev: %q: %v\n", id, rev, err)
		}
		http.Redirect(w, r, "/", http.StatusMovedPermanently)
	}))
	log.Fatalln(http.ListenAndServe(":9001", s))
}
