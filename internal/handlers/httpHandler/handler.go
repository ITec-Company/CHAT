package httpHandler

import (
	"github.com/gorilla/mux"
	"itec.chat/internal/service"
	"itec.chat/pkg/logging"
)

type Handler struct {
	logger      logging.Logger
	service     *service.Service
	userHandler *userHandler
}

func NewHandler(logger logging.Logger, service *service.Service) *Handler {
	return &Handler{
		logger:      logger,
		service:     service,
		userHandler: NewUserHandler(logger, service.UserService),
	}
}

func (h *Handler) InitRoutes() *mux.Router {
	router := mux.NewRouter()
	h.userHandler.Register(router)

	return router
}
