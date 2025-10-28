package websocket

import (
	"encoding/json"
	"log"
	"sync"
)

// Hub maintains active WebSocket connections and broadcasts messages
type Hub struct {
	// Registered clients (boardID -> list of clients)
	boards map[uint]map[*Client]bool

	// Register requests from clients
	register chan *Client

	// Unregister requests from clients
	unregister chan *Client

	// Broadcast messages to all clients in a board
	broadcast chan *Message

	// Mutex to protect concurrent access
	mu sync.RWMutex
}

// Message represents a WebSocket message
type Message struct {
	Type    string      `json:"type"` // "card_moved", "card_created", "comment_added", etc.
	BoardID uint        `json:"board_id"`
	Data    interface{} `json:"data"`
}

// NewHub creates a new Hub
func NewHub() *Hub {
	return &Hub{
		boards:     make(map[uint]map[*Client]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		broadcast:  make(chan *Message, 256),
	}
}

// Run starts the hub
func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.mu.Lock()
			if h.boards[client.BoardID] == nil {
				h.boards[client.BoardID] = make(map[*Client]bool)
			}
			h.boards[client.BoardID][client] = true
			h.mu.Unlock()
			log.Printf("✅ Client registered for board %d. Total clients: %d", client.BoardID, len(h.boards[client.BoardID]))

		case client := <-h.unregister:
			h.mu.Lock()
			if clients, ok := h.boards[client.BoardID]; ok {
				if _, ok := clients[client]; ok {
					delete(clients, client)
					close(client.Send)
					if len(clients) == 0 {
						delete(h.boards, client.BoardID)
					}
				}
			}
			h.mu.Unlock()
			log.Printf("❌ Client unregistered from board %d", client.BoardID)

		case message := <-h.broadcast:
			h.mu.RLock()
			clients := h.boards[message.BoardID]
			h.mu.RUnlock()

			messageJSON, err := json.Marshal(message)
			if err != nil {
				log.Printf("Error marshaling message: %v", err)
				continue
			}

			for client := range clients {
				select {
				case client.Send <- messageJSON:
				default:
					close(client.Send)
					h.mu.Lock()
					delete(h.boards[message.BoardID], client)
					h.mu.Unlock()
				}
			}
		}
	}
}

// Register registers a client to the hub
func (h *Hub) Register(client *Client) {
	h.register <- client
}

// Unregister unregisters a client from the hub
func (h *Hub) Unregister(client *Client) {
	h.unregister <- client
}

// BroadcastToBoard sends a message to all clients in a board
func (h *Hub) BroadcastToBoard(boardID uint, messageType string, data interface{}) {
	message := &Message{
		Type:    messageType,
		BoardID: boardID,
		Data:    data,
	}
	h.broadcast <- message
}
