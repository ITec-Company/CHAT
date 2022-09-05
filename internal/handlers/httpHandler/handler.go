package httpHandler

import (
	"github.com/gorilla/mux"
	"itec.chat/pkg/logging"
)

type Handler struct {
	logger      logging.Logger
	userHandler *userHandler
}

func NewHandler(logger logging.Logger) *Handler {
	return &Handler{
		logger:      logger,
		userHandler: NewUserHandler(logger),
	}
}

func (h *Handler) InitRoutes() *mux.Router {
	router := mux.NewRouter()
	h.userHandler.Register(router)

	return router
}
