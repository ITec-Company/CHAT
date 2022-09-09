package wsHub

import (
	"itec.chat/pkg/logging"
)

type Hub struct {
	clients    map[*client]bool
	broadcast  chan []byte
	register   chan *client
	unregister chan *client
	logger     logging.Logger
}

func NewHub(logger logging.Logger) *Hub {
	return &Hub{
		clients:    make(map[*client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *client),
		unregister: make(chan *client),
		logger:     logger,
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			h.logger.Info("new user: ", client.conn.LocalAddr().String())

			h.clients[client] = true

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				h.logger.Info(" user unregister: ", client.conn.LocalAddr().String())
				delete(h.clients, client)
				close(client.send)
			}

		case message := <-h.broadcast:

			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}

		}
	}
}
