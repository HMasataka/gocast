package gocast

import "io"

type Client interface {
	io.WriteCloser
	Error(err error)
}

type hub struct {
	// Registered clients.
	clients map[Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan Client

	// Unregister requests from clients.
	unregister chan Client
}

func NewHub() *hub {
	return &hub{
		broadcast:  make(chan []byte),
		register:   make(chan Client),
		unregister: make(chan Client),
		clients:    make(map[Client]bool),
	}
}

func (h *hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)

				if err := client.Close(); err != nil {
					client.Error(err)
				}
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				if _, err := client.Write(message); err != nil {
					client.Error(err)
				}
			}
		}
	}
}

func (h *hub) Broadcast(message []byte) {
	h.broadcast <- message
}

func (h *hub) Register(client Client) {
	h.register <- client
}

func (h *hub) Unregister(client Client) {
	h.unregister <- client
}
