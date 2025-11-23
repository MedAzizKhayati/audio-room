package handler

import (
	"encoding/base64"
	"encoding/json"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// WSMessage represents the standard JSON message format
type WSMessage struct {
	Type    string `json:"type"`              // "join", "audio", "users"
	Sender  string `json:"sender,omitempty"`  // Name of the sender
	Payload string `json:"payload,omitempty"` // Base64 encoded audio or other data
	Count   int    `json:"count,omitempty"`   // User count
	Users  []string `json:"users,omitempty"`   // List of usernames
}

type AudioHandler struct {
	// Registered clients map: Conn -> Name
	clients map[*websocket.Conn]string
	mu      sync.RWMutex

	// Inbound messages
	broadcast chan WSMessage

	// Register/Unregister
	register   chan *websocket.Conn
	unregister chan *websocket.Conn
}

func NewAudioHandler() *AudioHandler {
	h := &AudioHandler{
		broadcast:  make(chan WSMessage),
		register:   make(chan *websocket.Conn),
		unregister: make(chan *websocket.Conn),
		clients:    make(map[*websocket.Conn]string),
	}
	go h.run()
	return h
}

func (h *AudioHandler) run() {
	for {
		select {
		case conn := <-h.register:
			h.mu.Lock()
			h.clients[conn] = "" // Name not set yet
			h.mu.Unlock()
			h.broadcastUserList()

		case conn := <-h.unregister:
			h.mu.Lock()
			if _, ok := h.clients[conn]; ok {
				delete(h.clients, conn)
				conn.Close()
			}
			h.mu.Unlock()
			h.broadcastUserList()

		case message := <-h.broadcast:
			h.mu.RLock()
			for conn := range h.clients {
				// Don't send audio back to sender (echo cancellation)
				// But do send "users" updates to everyone
				if message.Type == "audio" && h.clients[conn] == message.Sender {
					continue
				}

				err := conn.WriteJSON(message)
				if err != nil {
					println("Error writing to client:", err.Error())
					conn.Close()
					// We can't safely delete from map while iterating with RLock
					// Ideally, queue for deletion or let the read loop handle it
				}
			}
			h.mu.RUnlock()
		}
	}
}

func (h *AudioHandler) broadcastUserList() {
	h.mu.RLock()
	count := len(h.clients)
	h.mu.RUnlock()

	msg := WSMessage{
		Type:  "users",
		Count: count,
		Users: func() []string {
			h.mu.RLock()
			defer h.mu.RUnlock()
			usernames := make([]string, 0, len(h.clients))
			for _, name := range h.clients {
				if name != "" {
					usernames = append(usernames, name)
				}
			}
			return usernames
		}(),
	}

	// We need to send this to everyone
	// In a real app, we'd put this in the broadcast channel,
	// but to avoid deadlock if channel is full, we launch a goroutine
	go func() { h.broadcast <- msg }()
}

func (h *AudioHandler) RegisterRoutes(router *gin.Engine) {
	createPath := GetControllerPathCreator("/audio")
	router.GET(createPath("/room"), WrapWSHandler(h.handleAudioRoom))
}

func (h *AudioHandler) handleAudioRoom(c *gin.Context, conn *websocket.Conn) error {
	println("AudioHandler: Client connected")
	h.register <- conn

	defer func() {
		println("AudioHandler: Client disconnected")
		h.unregister <- conn
	}()

	for {
		_, rawMsg, err := conn.ReadMessage()
		if err != nil {
			println("Error reading from client:", err.Error())
			return err
		}

		// Try to parse as JSON first (for "join")
		var msg WSMessage
		if err := json.Unmarshal(rawMsg, &msg); err == nil {
			if msg.Type == "join" {
				println("User joined:", msg.Sender)
				h.mu.Lock()
				h.clients[conn] = msg.Sender
				h.mu.Unlock()
				h.broadcastUserList()
				continue
			}
			// If it's audio sent as JSON
			if msg.Type == "audio" {
				h.mu.RLock()
				msg.Sender = h.clients[conn] // Ensure sender name is correct
				h.mu.RUnlock()
				h.broadcast <- msg
				continue
			}
		}

		// If we are here, it might be raw binary audio (if client sends blobs)
		// We wrap it in our JSON format
		h.mu.RLock()
		senderName := h.clients[conn]
		h.mu.RUnlock()

		if senderName == "" {
			println("Ignoring audio from unnamed user")
			continue // Ignore audio from unnamed users
		}

		encoded := base64.StdEncoding.EncodeToString(rawMsg)
		h.broadcast <- WSMessage{
			Type:    "audio",
			Sender:  senderName,
			Payload: encoded,
		}
	}
}
