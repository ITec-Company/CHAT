package handlers

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"itec.chat/internal/chatsession"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func Register(router *mux.Router) {
	router.HandleFunc("/ws/{id:[0-9]+}", WebsocketHandler)
	router.HandleFunc("/info", InfoWSHandler)

}

func WebsocketHandler(w http.ResponseWriter, r *http.Request) {
	chatId := mux.Vars(r)

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print(err)
		return
	}
	id, _ := strconv.Atoi(chatId["id"])
	chatSession := chatsession.NewChatSession(ws.RemoteAddr().String(), id, ws)

	chatSession.Start()
}

func InfoWSHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(chatsession.MapInfo()))
}
