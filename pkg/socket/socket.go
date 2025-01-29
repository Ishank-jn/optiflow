package socket

import (
    "log"
    "net/http"
    "github.com/gorilla/websocket"
	"github.com/gorilla/mux"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(r *http.Request) bool {
        return true // Allow all connections (customize for production)
    },
}

// WebSocket connection handler
func HandleConnections(w http.ResponseWriter, r *http.Request) {
    // Upgrade initial GET request to a WebSocket
    ws, err := upgrader.Upgrade(w, r, nil)
    if err != nil {
        log.Printf("WebSocket upgrade failed: %v", err)
        return
    }
    defer ws.Close()

    // Register new client
    clients[ws] = true
    log.Println("New WebSocket client connected")

    for {
        // Read message from client
        var msg Message
        err := ws.ReadJSON(&msg)
        if err != nil {
            log.Printf("WebSocket read error: %v", err)
            delete(clients, ws)
            break
        }

        // Broadcast message to all clients
        broadcast <- msg
    }
}

// Broadcast messages to all connected clients
func HandleMessages() {
    for {
        msg := <-broadcast
        for client := range clients {
            err := client.WriteJSON(msg)
            if err != nil {
                log.Printf("WebSocket write error: %v", err)
                client.Close()
                delete(clients, client)
            }
        }
    }
}

// Message struct for WebSocket communication
type Message struct {
    Type    string `json:"type"`
    Content string `json:"content"`
}

var (
    clients   = make(map[*websocket.Conn]bool) // Connected clients
    broadcast = make(chan Message)             // Broadcast channel
)

// Initialize WebSocket handlers
func InitWebSocket(router *mux.Router) {
    go HandleMessages()
    router.HandleFunc("/ws", HandleConnections)
}