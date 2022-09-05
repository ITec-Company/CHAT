package httpHandler

import (
	"net/http"

	"github.com/gorilla/mux"
	"itec.chat/internal/service"
	"itec.chat/pkg/logging"
)

const (
	getUsers = "/users"
)

type userHandler struct {
	logger  logging.Logger
	service service.UserService
}

func NewUserHandler(logger logging.Logger, service service.UserService) *userHandler {
	return &userHandler{
		logger:  logger,
		service: service,
	}
}

func (uh *userHandler) Register(router *mux.Router) {
	router.HandleFunc(getUsers, uh.allUsers).Methods("GET")
}

func (uh *userHandler) allUsers(w http.ResponseWriter, r *http.Request) {
	uh.service.GetByID(2)
	w.Write([]byte("get all Users"))
	w.WriteHeader(http.StatusOK)
}
