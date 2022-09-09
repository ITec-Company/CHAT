package httpHandler

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"itec.chat/internal/wsHub"
	"itec.chat/pkg/logging"
)

type websocketHandler struct {
	logger logging.Logger
	//hub    *wsHub.Hub
}

func newWebsocketHandler(logger logging.Logger /*, hub *wsHub.Hub*/) *websocketHandler {
	return &websocketHandler{
		logger: logger,
		//hub:    hub,
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func (wh *websocketHandler) register(router *mux.Router) {
	router.HandleFunc("/ws/{id:[0-9]+}", wh.handleWebsocket)
	router.HandleFunc("/map", wh.getHubsMap)

}

func (wh *websocketHandler) handleWebsocket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		wh.logger.Errorf("err : ", err)
	}

	hub := wsHub.GetHub(wh.logger, id)
	go hub.Run()

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		wh.logger.Errorf("error occurred while upgrading ws conn", err)
		return
	}

	client := wsHub.NewClient(hub, ws, wh.logger)

	go client.WritePump()
	go client.ReadPump()

}

func (wh *websocketHandler) getHubsMap(w http.ResponseWriter, r *http.Request) {
	str := wsHub.GetStringMaps()

	w.Write([]byte(str))
}
