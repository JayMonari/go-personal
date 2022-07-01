package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter().StrictSlash(true)
	r.HandleFunc("/cars", returnAllCars).Methods("GET")
	r.HandleFunc("/cars/", createCar).Methods("POST")
	r.HandleFunc("/cars/{id}", readCar).Methods("GET")
	r.HandleFunc("/cars/{id}", updateCar).Methods("PUT")
	r.HandleFunc("/cars/{id}", deleteCar).Methods("DELETE")
	r.HandleFunc("/cars/make/{make}", byBrand).Methods("GET")
	log.Fatal(http.ListenAndServe(":3000", r))
}

type Vehicle struct {
	ID    int
	Make  string
	Model string
	Price int
}

var vehicles = []Vehicle{
	{
		ID:    1,
		Make:  "Toyota",
		Model: "Corolla",
		Price: 10000,
	},
	{
		ID:    2,
		Make:  "Toyota",
		Model: "Camry",
		Price: 20000,
	},
	{
		ID:    3,
		Make:  "Ford",
		Model: "Tacoma",
		Price: 30000,
	},
	{
		ID:    4,
		Make:  "Honda",
		Model: "Civic",
		Price: 40000,
	},
}

func returnAllCars(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(vehicles); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
}

func createCar(w http.ResponseWriter, r *http.Request) {
	var car Vehicle
	json.NewDecoder(r.Body).Decode(&car)
	r.Body.Close()
	vehicles = append(vehicles, car)

	if err := json.NewEncoder(w).Encode(vehicles); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Add("Content-Type", "application/json")
}

func readCar(w http.ResponseWriter, r *http.Request) {
	carID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	for _, c := range vehicles {
		if c.ID != carID {
			continue
		}
		if err := json.NewEncoder(w).Encode(c); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	w.Header().Add("Content-Type", "application/json")
}

func updateCar(w http.ResponseWriter, r *http.Request) {
	var car Vehicle
	json.NewDecoder(r.Body).Decode(&car)
	r.Body.Close()
	carID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	for i, c := range vehicles {
		if c.ID != carID {
			continue
		}
		vehicles[i] = car
		if err := json.NewEncoder(w).Encode(vehicles); err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	w.Header().Add("Content-Type", "application/json")
}

func deleteCar(w http.ResponseWriter, r *http.Request) {
	carID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	for i, c := range vehicles {
		if c.ID != carID {
			continue
		}
		vehicles = append(vehicles[:i], vehicles[i+1:]...)
		if err := json.NewEncoder(w).Encode(vehicles); err != nil {
			log.Println(err)
		}
		break
	}
	w.Header().Add("Content-Type", "application/json")
}

func byBrand(w http.ResponseWriter, r *http.Request) {
	carM := mux.Vars(r)["make"]
	var cars []Vehicle
	for _, c := range vehicles {
		if c.Make == carM {
			cars = append(cars, c)
		}
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cars)
}
