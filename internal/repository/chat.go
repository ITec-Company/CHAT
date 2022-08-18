package repository

import (
	"database/sql"
	"itec.chat/internal/domain"
	"itec.chat/pkg/logging"
)

type chat struct {
	db     *sql.DB
	logger *logging.Logger
}

func NewChatRepository(db *sql.DB, logger *logging.Logger) (chatRepository Chat) {
	return &chat{
		db:     db,
		logger: logger,
	}
}

func (rep *chat) GetByID(id int) (chat *domain.Chat, err error) {
	return
}

func (rep *chat) GetAll(limit, offset int) (chats []Chat, err error) {
	return
}

func (rep *chat) Create(createChat *domain.CreateChat) (id int, err error) {
	return
}

func (rep *chat) Update(updateChat *domain.UpdateChat) (err error) {
	return
}

func (rep *chat) Delete(id int) (err error) {
	return
}
