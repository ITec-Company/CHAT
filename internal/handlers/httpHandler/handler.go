package httpHandler

import (
	"github.com/gorilla/mux"
	"itec.chat/internal/repository"
	"itec.chat/pkg/logging"
)

type Handler struct {
	logger      logging.Logger
	repository  *repository.Repository
	userHandler *userHandler
}

func NewHandler(logger logging.Logger, repository *repository.Repository) *Handler {
	return &Handler{
		logger:      logger,
		repository:  repository,
		userHandler: NewUserHandler(logger, repository.User),
	}
}

func (h *Handler) InitRoutes() *mux.Router {
	router := mux.NewRouter()
	h.userHandler.Register(router)

	return router
}
