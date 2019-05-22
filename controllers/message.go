package controllers

import (
	"fmt"
	"github.com/gorilla/websocket"
	"html/template"
	"log"
	"net/http"
)

var clients = make(map[*websocket.Conn]bool) // connected clients
var broadcast = make(chan Message)           // broadcast channel

type Message struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Message  string `json:"message"`
}

var upgrader = websocket.Upgrader{}

func GetMessage(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("views/message.html")
	_ = t.Execute(w, nil)
}

// 监听WebSocket
func Handle(w http.ResponseWriter, r *http.Request) {

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Connect fail")
	}
	defer conn.Close()
	clients[conn] = true

	for {
		var msg Message

		err := conn.ReadJSON(&msg)
		if err != nil {
			log.Printf("line 44 error: %v", err)
			delete(clients,conn)
			conn.Close()
			return
		}
		fmt.Println(msg)

		//broadcast <- msg
		handleMessages(msg)
	}
}

// 发送消息
func handleMessages(msg Message) {
	fmt.Println(msg)
	for client := range clients {
		err := client.WriteJSON(msg)
		if err != nil {
			log.Printf("line 61 error: %v", err)
		}
	}
	fmt.Println(clients)
}
