package handlers

import (
	"github.com/gorilla/mux"
	"itec.chat/internal/handlers/http"
	"itec.chat/pkg/logging"
)

type Router struct {
	Logger logging.Logger
	Router mux.Router
	//repository repository.Repository

	UserHandler http.UserHandler
}

func NewRouter(logger logging.Logger /*, repository *repository.Repository*/) *Router {
	return &Router{
		Logger: logger,
		Router: *mux.NewRouter(),
		//	repository:  *repository,
		UserHandler: *http.NewUserHandler(logger /*, repository.User*/),
	}
}

func (r *Router) InitRoutes() {
	r.UserHandler.Register(&r.Router)
}
