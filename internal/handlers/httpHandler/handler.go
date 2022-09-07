package httpHandler

import (
	"github.com/gorilla/mux"
	"itec.chat/internal/handlers"
	"itec.chat/internal/wsHub"
	"itec.chat/pkg/logging"
)

type Handler struct {
	logger logging.Logger
	hub    *wsHub.Hub
	//repository      *repository.Repository
	//userHandler     *userHandler
	//fileHandler     *fileHandler
	//chatHandler     *chatHandler
	//messageHandler  *messageHandler
	//statusHandler   *statusHandler
	//userRoleHandler *userRoleHandler
}

func NewHandler(logger logging.Logger , hub *wsHub.Hub/*, repository *repository.Repository*/) *Handler {
	return &Handler{
		logger: logger,
		hub: hub,
		//repository:      repository,
		//userHandler:     NewUserHandler(logger, repository.User),
		//fileHandler:     NewFileHandler(logger, repository.File),
		//	chatHandler:     NewChatHandler(logger, repository.Chat),
		//	messageHandler:  NewMessageHandler(logger, repository.Message),
		//	statusHandler:   NewStatusHandler(logger, repository.Status),
		//	userRoleHandler: NewUserRoleHandler(logger, repository.UserRole),
	}
}

func (h *Handler) InitRoutes() *mux.Router {
	router := mux.NewRouter()
	handlers.Register(router, h.hub) //ws
	/*	h.userHandler.Register(router)
		h.chatHandler.Register(router)
		h.fileHandler.Register(router)
		h.messageHandler.Register(router)
		h.statusHandler.Register(router)
		h.statusHandler.Register(router)*/
	return router
}
