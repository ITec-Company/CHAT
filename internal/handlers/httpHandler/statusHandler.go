package httpHandler

import (
	"net/http"

	"github.com/gorilla/mux"
	"itec.chat/internal/repository"
	"itec.chat/pkg/logging"
)

const (
	getStatuses = "/statuses"
)

type statusHandler struct {
	logger   logging.Logger
	statusRepo repository.Status
}

func NewStatusHandler(logger logging.Logger, statusRepo repository.Status) *statusHandler {
	return &statusHandler{
		logger:   logger,
		statusRepo: statusRepo,
	}
}

func (uh *statusHandler) Register(router *mux.Router) {
	router.HandleFunc(getStatuses, uh.allStatuses).Methods("GET")
}

func (uh *statusHandler) allStatuses(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get all statuses"))
	w.WriteHeader(http.StatusOK)
}
