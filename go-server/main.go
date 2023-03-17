package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"

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
	i := 0
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ws, _ := upgrader.Upgrade(w, r, nil)
		wsList = append(wsList, ws)

		messageType, bt, _ := wsList[0].ReadMessage()
		_ = wsList[i].WriteMessage(messageType, bt)
		i++
	})

	http.HandleFunc("/send", func(w http.ResponseWriter, r *http.Request) {
		str := "消息"
		bt, _ := json.Marshal(str)
		_ = wsList[0].WriteMessage(1, bt)
	})

	go http.ListenAndServe(":2021", nil)

	r := gin.Default()

	r.GET("/send", func(c *gin.Context) {
		str := c.Query("msg")
		bt, _ := json.Marshal(str)
		id, _ := strconv.Atoi(c.Query("id"))
		_ = wsList[id].WriteMessage(1, bt)
	})

	r.Run(":8082")
}
