package httpHandler

import (
	"github.com/gorilla/mux"
	"itec.chat/internal/wsHub"
	"itec.chat/pkg/logging"
)

type Handler struct {
	logger           logging.Logger
	websocketHandler *websocketHandler
	//repository      *repository.Repository
	//userHandler     *userHandler
	//fileHandler     *fileHandler
	//chatHandler     *chatHandler
	//messageHandler  *messageHandler
	//statusHandler   *statusHandler
	//userRoleHandler *userRoleHandler
}

func NewHandler(logger logging.Logger, hub *wsHub.Hub /*, repository *repository.Repository*/) *Handler {
	return &Handler{
		logger:           logger,
		websocketHandler: newWebsocketHandler(logger, hub),
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
	h.websocketHandler.register(router) //ws

	/*	h.userHandler.register(router)
		h.chatHandler.register(router)
		h.fileHandler.register(router)
		h.messageHandler.register(router)
		h.statusHandler.register(router)
		h.statusHandler.register(router)*/
	return router
}
