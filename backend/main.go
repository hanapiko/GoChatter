package main

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}
var clients = make(map[*websocket.Conn]bool) // Track connected clients
var mutex = sync.Mutex{}

func broadcast(msg []byte) {
	mutex.Lock()
	defer mutex.Unlock()
	for client := range clients {
		client.WriteMessage(websocket.TextMessage, msg)
	}
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer conn.Close()

	// Add client to map
	mutex.Lock()
	clients[conn] = true
	mutex.Unlock()

	// Remove client when they disconnect
	defer func() {
		mutex.Lock()
		delete(clients, conn)
		mutex.Unlock()
	}()

	for {
		_, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		broadcast(msg) // Send to all clients
	}

}

func main() {
	http.HandleFunc("/ws", handleWebSocket)
	http.ListenAndServe(":8080", nil)
}
