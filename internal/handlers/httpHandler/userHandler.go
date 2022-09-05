package httpHandler

import (
	"net/http"

	"github.com/gorilla/mux"
	"itec.chat/pkg/logging"
)

const (
	getUsers = "/users"
)

type userHandler struct {
	logger logging.Logger
}

func NewUserHandler(logger logging.Logger) *userHandler {
	return &userHandler{
		logger: logger,
	}
}

func (uh *userHandler) Register(router *mux.Router) {
	router.HandleFunc(getUsers, uh.allUsers).Methods("GET")
}

func (uh *userHandler) allUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get all Users"))
	w.WriteHeader(http.StatusOK)
}
