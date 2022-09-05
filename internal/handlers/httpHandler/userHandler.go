package httpHandler

import (
	"net/http"

	"github.com/gorilla/mux"
	"itec.chat/internal/repository"
	"itec.chat/pkg/logging"
)

const (
	getUsers = "/users"
)

type userHandler struct {
	logger   logging.Logger
	userRepo repository.User
}

func NewUserHandler(logger logging.Logger, userRepo repository.User) *userHandler {
	return &userHandler{
		logger:   logger,
		userRepo: userRepo,
	}
}

func (uh *userHandler) Register(router *mux.Router) {
	router.HandleFunc(getUsers, uh.allUsers).Methods("GET")
}

func (uh *userHandler) allUsers(w http.ResponseWriter, r *http.Request) {
	uh.userRepo.GetAll(1, 1)
	w.Write([]byte("get all Users"))
	w.WriteHeader(http.StatusOK)
}
