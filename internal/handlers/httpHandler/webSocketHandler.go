package httpHandler

import (
	"fmt"
	"net/http"
	"runtime"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"itec.chat/internal/wsHub"
	"itec.chat/pkg/logging"
)

const (
	connect = "/ws/{id:[0-9]+}"
	newHub  = "/hub/new/{id:[0-9]+}"
)

type websocketHandler struct {
	logger logging.Logger
}

func newWebsocketHandler(logger logging.Logger) *websocketHandler {

	return &websocketHandler{
		logger: logger,
	}
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func (wh *websocketHandler) register(router *mux.Router) {
	router.HandleFunc(connect, wh.handleWebsocket)
	router.HandleFunc(newHub, wh.NewHub(wh.handleWebsocket)).Methods(http.MethodPost)

	router.HandleFunc("/map", wh.getHubsMap)
	router.HandleFunc("/g", wh.numGoroutine)

}

func (wh *websocketHandler) handleWebsocket(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		wh.logger.Errorf("err : ", err)
	}

	hub, err := wsHub.GetHub(wh.logger, id)
	if err != nil {
		wh.logger.Errorf("err :", err)
		w.WriteHeader(http.StatusBadRequest)
		str := fmt.Sprintf("Hub doest exist. id:%d ", id)
		w.Write([]byte(str))
		return
	}

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

func (wh *websocketHandler) numGoroutine(w http.ResponseWriter, r *http.Request) {
	str := strconv.Itoa(runtime.NumGoroutine())

	w.Write([]byte(str))
}

func (wh *websocketHandler) NewHub(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			wh.logger.Errorf("err: ", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		hub, err := wsHub.NewHub(wh.logger, id)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte("Hub alreary xist"))
			fmt.Println("3")

		} else {
			go hub.Run()
			w.WriteHeader(http.StatusCreated)
			w.Write([]byte("Hub is created"))
			fmt.Println("4")

		}

	}
}
