package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var todoList []string

func webSocketConnection(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	defer conn.Close()
	if err != nil {
		log.Println(err)
		return
	}
}

func runServer() {
	fmt.Println("Starting server on port :8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(fmt.Sprintf("there was an error running the server: \"%s\"", err))
	}
}
