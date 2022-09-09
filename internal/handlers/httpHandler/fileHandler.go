package httpHandler

import (
	"net/http"

	"github.com/gorilla/mux"
	"itec.chat/internal/repository"
	"itec.chat/pkg/logging"
)

const (
	getFiles = "/file"
)

type fileHandler struct {
	logger   logging.Logger
	fileRepo repository.File
}

func newFileHandler(logger logging.Logger, fileRepo repository.File) *fileHandler {
	return &fileHandler{
		logger:   logger,
		fileRepo: fileRepo,
	}
}

func (fh *fileHandler) register(router *mux.Router) {
	router.HandleFunc(getFiles, fh.allFiles).Methods("GET")
}

func (fh *fileHandler) allFiles(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get all Files"))
	w.WriteHeader(http.StatusOK)
}
