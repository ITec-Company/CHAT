package httpHandler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gorilla/mux"
	"itec.chat/internal/models"
	"itec.chat/internal/repository"
	"itec.chat/pkg/logging"
)

const (
	getUsers         = "/user/all"
	creatUser        = "/user/create"
	updateUser       = "/user/update"
	getUserByID      = "/user/{id:[0-9]+}"
	deleteUser       = "/user/delete/{id:[0-9]+}"
	getUsersByChatID = "/user/chat/{id:[0-9]+}"
	updateUserStatus = "/user/update/status"
)

type userHandler struct {
	logger   logging.Logger
	userRepo repository.User
}

func newUserHandler(logger logging.Logger /*, userRepo repository.User*/) *userHandler {
	return &userHandler{
		logger: logger,
		//userRepo: userRepo,
	}
}

func (uh *userHandler) register(router *mux.Router) {
	router.HandleFunc(getUsers, uh.allUsers).Methods(http.MethodGet)
	router.HandleFunc(getUserByID, uh.getUserByID).Methods(http.MethodGet)
	router.HandleFunc(getUsersByChatID, uh.getUsersByChatID).Methods(http.MethodGet)
	router.HandleFunc(creatUser, uh.creatUser).Methods(http.MethodPost)
	router.HandleFunc(updateUser, uh.updateUser).Methods(http.MethodPost)
	router.HandleFunc(updateUserStatus, uh.updateUserStatus).Methods(http.MethodPost)
	router.HandleFunc(deleteUser, uh.deleteUser).Methods(http.MethodDelete)

}

func (uh *userHandler) allUsers(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode("get all Users")
}

func (uh *userHandler) getUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	json.NewEncoder(w).Encode(fmt.Sprintf("user with id: %s", vars["id"]))

}

func (uh *userHandler) creatUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	req := &models.CreateUser{}

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		uh.logger.Errorf("Error occurred while parsing json. err: ", err)
		json.NewEncoder(w).Encode(fmt.Sprintf("err: %s", err))
		return
	}

	_, err = govalidator.ValidateStruct(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		uh.logger.Errorf("Error occurred while validating. err: ", err)
		json.NewEncoder(w).Encode(fmt.Sprintf("err: %s", err))
		return
	}

	req.LastActivity = time.Now()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(req)
}

func (uh *userHandler) deleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	json.NewEncoder(w).Encode(fmt.Sprintf("deleted user with id: %s", vars["id"]))

}

func (uh *userHandler) updateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	req := &models.UpdateUser{}

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		uh.logger.Errorf("Error occurred while parsing json. err: ", err)
		json.NewEncoder(w).Encode(fmt.Sprintf("err: %s", err))
		return
	}
	_, err = govalidator.ValidateStruct(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		uh.logger.Errorf("Error occurred while validating. err: ", err)
		json.NewEncoder(w).Encode(fmt.Sprintf("err: %s", err))
		return
	}
	req.LastActivity = time.Now()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(req)
}

func (uh *userHandler) updateUserStatus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	req := &models.UpdateUserStatus{}

	err := json.NewDecoder(r.Body).Decode(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		uh.logger.Errorf("Error occurred while parsing json. err: ", err)
		json.NewEncoder(w).Encode(fmt.Sprintf("err: %s", err))
		return
	}
	_, err = govalidator.ValidateStruct(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		uh.logger.Errorf("Error occurred while validating. err: ", err)
		json.NewEncoder(w).Encode(fmt.Sprintf("err: %s", err))
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(req)
}

func (uh *userHandler) getUsersByChatID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	json.NewEncoder(w).Encode(fmt.Sprintf("user with chat id: %s", vars["id"]))

}
