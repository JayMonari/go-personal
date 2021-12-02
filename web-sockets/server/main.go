package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func socketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Fatal("Error during connection upgrade:", err)
	}
	defer conn.Close()

	for {
		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			log.Fatal("Error during message reading:", err)
		}
		log.Println("Received:", string(msg))
		err = conn.WriteMessage(msgType, msg)
		if err != nil {
			log.Fatal("Error during message writing:", err)
		}
	}
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Index Page")
}

func main() {
	http.HandleFunc("/socket", socketHandler)
	http.HandleFunc("/", home)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
