package repository

import (
	"database/sql"
	"itec.chat/internal/models"
	"itec.chat/pkg/logging"
)

type message struct {
	db     *sql.DB
	logger *logging.Logger
}

func NewMessageRepository(db *sql.DB, logger *logging.Logger) (messageRepository Message) {
	return &message{
		db:     db,
		logger: logger,
	}
}

func (rep *message) GetByID(id int) (message *models.Message, err error) {
	return
}

func (rep *message) GetAll(limit, offset int) (messages []Message, err error) {
	return
}

func (rep *message) Create(createMessage *models.CreateMessage) (id int, err error) {
	return
}

func (rep *message) Update(updateMessage *models.UpdateMessage) (err error) {
	return
}

func (rep *message) Delete(id int) (err error) {
	return
}
