package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type Name struct {
	NConst    string `json:"nconst"`
	Name      string `json:"name"`
	BirthYear string `json:"birthYear"`
	DeathYear string `json:"deathYear"`
}

type Error struct {
	Message string `json:"message"`
}

func main() {
	db, err := NewPostgreSQL()
	if err != nil {
		log.Fatalf("Could not initialize Database connection %v", err)
	}
	defer db.Close()

	mc, err := NewMemcached()
	if err != nil {
		log.Fatalf("Could not initialize Memcached client %v", err)
	}

	router := mux.NewRouter()
	renderJSON := func(w http.ResponseWriter, val any, statusCode int) {
		w.WriteHeader(statusCode)
		if err := json.NewEncoder(w).Encode(val); err != nil {
			log.Printf("Could not encode value, %v", err)
		}
	}
	router.HandleFunc("/names/{id}", func(w http.ResponseWriter, r *http.Request) {
		id := mux.Vars(r)["id"]

		if val, err := mc.GetName(id); err == nil {
			renderJSON(w, &val, http.StatusOK)
			return
		}

		name, err := db.FindByNConst(id)
		if err != nil {
			renderJSON(w, &Error{Message: err.Error()}, http.StatusInternalServerError)
			return
		}
		if err := mc.SetName(name); err != nil {
			renderJSON(w, &Error{Message: err.Error()}, http.StatusInternalServerError)
			return
		}
		renderJSON(w, &name, http.StatusOK)
	})

	fmt.Println("Starting server :8080")
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      router,
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
	}
	log.Fatal(srv.ListenAndServe())
}
