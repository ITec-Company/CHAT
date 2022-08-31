package http

import (
	"net/http"

	"github.com/gorilla/mux"
	"itec.chat/internal/repository"
	"itec.chat/pkg/logging"
)

const (
	getUsers = "/users"
)

type UserHandler struct {
	logger     logging.Logger
	repository *repository.User
}

func NewUserHandler(logger logging.Logger /*, user repository.User*/) *UserHandler {
	return &UserHandler{
		logger: logger,
		//repository: user,
	}
}

func (uh *UserHandler) Register(router *mux.Router) {
	router.HandleFunc(getUsers, uh.allUsers).Methods("GET")
}

func (uh *UserHandler) allUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get all Users"))
	w.WriteHeader(http.StatusOK)
}
