package handlers

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"itec.chat/internal/wsHub"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func Register(router *mux.Router, hub *wsHub.Hub) {
	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		WebsocketHandler(w, r, hub)
	})

}

func WebsocketHandler(w http.ResponseWriter, r *http.Request, hub *wsHub.Hub) {

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("error occurred while upgrading ws conn", err)
		return
	}

	client := wsHub.NewClient(hub, ws)

	go client.WritePump()
	go client.ReadPump()

}
