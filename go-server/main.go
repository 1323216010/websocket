package main

import (
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		conn, _ := upgrader.Upgrade(w, r, nil)

		msgType, msg, err := conn.ReadMessage()
		if err != nil {
			return
		}

		if err = conn.WriteMessage(msgType, msg); err != nil {
			return
		}
	})

	http.ListenAndServe(":2021", nil)
}
