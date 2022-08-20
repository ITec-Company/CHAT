package repository

import (
	"database/sql"
	"itec.chat/internal/models"
	"itec.chat/pkg/logging"
)

type status struct {
	db     *sql.DB
	logger *logging.Logger
}

func NewStatusRepository(db *sql.DB, logger *logging.Logger) (statusRepository Status) {
	return &status{
		db:     db,
		logger: logger,
	}
}

func (rep *status) GetByID(id int) (status *models.Status, err error) {
	return
}

func (rep *status) GetAll(limit, offset int) (statuses []Status, err error) {
	return
}

func (rep *status) Create(createStatus *models.CreateStatus) (id int, err error) {
	return
}

func (rep *status) Update(updateStatus *models.UpdateStatus) (err error) {
	return
}

func (rep *status) Delete(id int) (err error) {
	return
}
