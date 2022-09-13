package wsHub

import (
	"errors"
	"fmt"

	"itec.chat/pkg/logging"
)

type Hub struct {
	id         int
	clients    map[*client]bool
	broadcast  chan []byte
	register   chan *client
	unregister chan *client
	logger     logging.Logger
}

var hubs = make(map[int]*Hub)

func NewHub(logger logging.Logger, id int) (*Hub, error) {
	if !isExist(id) {
		hub := &Hub{
			id:         id,
			clients:    make(map[*client]bool),
			broadcast:  make(chan []byte),
			register:   make(chan *client),
			unregister: make(chan *client),
			logger:     logger,
		}
		hubs[id] = hub

		return hub, nil
	}
	return nil, errors.New("error occured while creating hub. err: hub already exist")

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
			//deleting hub if if has no more clients
			if len(h.clients) == 0 {
				delete(hubs, h.id)
				close(h.broadcast)
				close(h.register)
				close(h.unregister)
				return
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

func GetHub(logger logging.Logger, id int) (*Hub, error) {
	hub, ok := hubs[id]
	if ok {
		return hub, nil
	} else {
		return nil, errors.New("error occured while getting hub. err: no hub with such id")
	}
}

func isExist(id int) bool {
	_, ok := hubs[id]
	return ok
}

func GetStringMaps() string {
	str := ""
	for key, value := range hubs {
		str = str + fmt.Sprintf("Key: %d and Hub clients: %s \n", key, value.getClientInfo())
	}

	return str
}

func (h *Hub) getClientInfo() string {
	str := ""
	for keys := range h.clients {
		str = str + fmt.Sprintf("Local adres of connections: %s \n", keys.conn.LocalAddr().String())

	}
	return str

}
