package httpHandler

import (
	"net/http"

	"github.com/gorilla/mux"
	"itec.chat/internal/repository"
	"itec.chat/pkg/logging"
)

const (
	getChats = "/chats"
)

type chatHandler struct {
	logger   logging.Logger
	chatRepo repository.Chat
}

func NewChatHandler(logger logging.Logger, chatRepo repository.Chat) *chatHandler {
	return &chatHandler{
		logger:   logger,
		chatRepo: chatRepo,
	}
}

func (ch *chatHandler) Register(router *mux.Router) {
	router.HandleFunc(getChats, ch.allChats)
}

func (uh *chatHandler) allChats(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get all shats"))
	w.WriteHeader(http.StatusOK)
}
