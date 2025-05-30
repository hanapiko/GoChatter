package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

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
	timestamp := time.Now().Format("15:04:05")
	formattedMsg := fmt.Sprintf("[%s] %s", timestamp, msg)

	for client := range clients {
		client.WriteMessage(websocket.TextMessage, []byte(formattedMsg))
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

func homePage(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`
		<h1>Welcome to GoChatter!</h1>
		<p>Connect via WebSocket at <code>ws://localhost:8080/ws</code></p>
	`))
}

func main() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/ws", handleWebSocket)
	http.ListenAndServe(":8080", nil)
}
