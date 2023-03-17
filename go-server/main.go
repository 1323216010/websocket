package main

import (
	"encoding/json"
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
	var wsList []*websocket.Conn
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ws, _ := upgrader.Upgrade(w, r, nil)
		wsList = append(wsList, ws)

		messageType, bt, _ := wsList[0].ReadMessage()
		_ = wsList[0].WriteMessage(messageType, bt)
	})

	http.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
		str := "消息"
		bt, _ := json.Marshal(str)
		_ = wsList[0].WriteMessage(1, bt)
	})

	http.ListenAndServe(":2021", nil)
}
