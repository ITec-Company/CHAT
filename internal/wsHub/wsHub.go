package wsHub

import (
	"log"
)

type Hub struct {
	clients    map[*client]bool
	broadcast  chan []byte
	register   chan *client
	unregister chan *client
}

func NewHub() *Hub {
	return &Hub{
		clients:    make(map[*client]bool),
		broadcast:  make(chan []byte),
		register:   make(chan *client),
		unregister: make(chan *client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.register:
			log.Print("new user: ", client.conn.LocalAddr().String())

			h.clients[client] = true

		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				log.Print(" user unregister: ", client.conn.LocalAddr().String())

				delete(h.clients, client)
				close(client.send)
			}

		case message := <-h.broadcast:
			log.Print("Hub GETTING BROADCAST MESSGAE ", message)

			for client := range h.clients {
				log.Print("range clients ", message)
				select {
				case client.send <- message:
					log.Print("HUB client.send <- message: ", message)

				default:
					log.Print("HUB default hyb", message)

					close(client.send)
					delete(h.clients, client)
				}
			}

		}
	}
}
