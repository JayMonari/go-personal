package handler

import (
	"net/http"
	"strconv"
	"time"
)

var GlobalBAD = map[int]time.Time{}

// AddRaceCondition accesses a global map which is written to by many connections.
func AddRaceCondition(w http.ResponseWriter, r *http.Request) {
	a, _ := strconv.Atoi(r.PostFormValue("numberOne"))
	b, _ := strconv.Atoi(r.PostFormValue("numberTwo"))
	GlobalBAD[a+b] = time.Now()
	if _, err := w.Write([]byte(strconv.Itoa(a + b))); err != nil {
		panic(err)
	}
}

// Add takes two numbers in the POST form data and returns their sum.
func Add(w http.ResponseWriter, r *http.Request) {
	a, _ := strconv.Atoi(r.PostFormValue("numberOne"))
	b, _ := strconv.Atoi(r.PostFormValue("numberTwo"))
	if _, err := w.Write([]byte(strconv.Itoa(a + b))); err != nil {
		panic(err)
	}
}
