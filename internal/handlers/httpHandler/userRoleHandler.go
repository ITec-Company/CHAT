package httpHandler

import (
	"net/http"

	"github.com/gorilla/mux"
	"itec.chat/internal/repository"
	"itec.chat/pkg/logging"
)

const (
	getRoles = "/roles"
)

type userRoleHandler struct {
	logger    logging.Logger
	rolesRepo repository.UserRole
}

func newUserRoleHandler(logger logging.Logger, rolesRepo repository.UserRole) *userRoleHandler {
	return &userRoleHandler{
		logger:    logger,
		rolesRepo: rolesRepo,
	}
}

func (urh *userRoleHandler) register(router *mux.Router) {
	router.HandleFunc(getRoles, urh.allRoles).Methods("GET")
}

func (urh *userRoleHandler) allRoles(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get all Roles"))
	w.WriteHeader(http.StatusOK)
}
