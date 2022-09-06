package httpHandler

import (
	"net/http"

	"github.com/gorilla/mux"
	"itec.chat/internal/repository"
	"itec.chat/pkg/logging"
)

const (
	getMessages = "/messages"
)

type messageHandler struct {
	logger   logging.Logger
	messRepo repository.Message
}

func NewMessageHandler(logger logging.Logger, messRepo repository.Message) *messageHandler {
	return &messageHandler{
		logger:   logger,
		messRepo: messRepo,
	}
}

func (fh *messageHandler) Register(router *mux.Router) {
	router.HandleFunc(getMessages, fh.allMessages).Methods("GET")
}

func (fh *messageHandler) allMessages(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get all mess"))
	w.WriteHeader(http.StatusOK)
}
