package repository

import (
	"database/sql"
	"itec.chat/internal/domain"
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

func (rep *message) GetByID(id int) (message *domain.Message, err error) {
	return
}

func (rep *message) GetAll(limit, offset int) (messages []Message, err error) {
	return
}

func (rep *message) Create(createMessage *domain.CreateMessage) (id int, err error) {
	return
}

func (rep *message) Update(updateMessage *domain.UpdateMessage) (err error) {
	return
}

func (rep *message) Delete(id int) (err error) {
	return
}
