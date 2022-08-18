package repository

import (
	"database/sql"
	"errors"
	"itec.chat/internal/domain"
	"itec.chat/pkg/logging"
)

var (
	ErrNoRowsAffected = errors.New("now rows affected")
)

type Chat interface {
	GetByID(id int) (chat *domain.Chat, err error)
	GetAll(limit, offset int) (chats []Chat, err error)
	Create(createChat *domain.CreateChat) (id int, err error)
	Update(updateChat *domain.UpdateChat) (err error)
	Delete(id int) (err error)
}

type File interface {
	GetByID(id int) (file *domain.File, err error)
	GetAll(limit, offset int) (files []File, err error)
	Create(createFile *domain.CreateFile) (id int, err error)
	Update(updateFile *domain.UpdateFile) (err error)
	Delete(id int) (err error)
}

type Message interface {
	GetByID(id int) (message *domain.Message, err error)
	GetAll(limit, offset int) (messages []Message, err error)
	Create(createMessage *domain.CreateMessage) (id int, err error)
	Update(updateMessage *domain.UpdateMessage) (err error)
	Delete(id int) (err error)
}

type Status interface {
	GetByID(id int) (status *domain.Status, err error)
	GetAll(limit, offset int) (statuses []Status, err error)
	Create(createStatus *domain.CreateStatus) (id int, err error)
	Update(updateStatus *domain.UpdateStatus) (err error)
	Delete(id int) (err error)
}

type UserRole interface {
	GetByID(id int) (role *domain.UserRole, err error)
	GetAll(limit, offset int) (roles []UserRole, err error)
	Create(createRole *domain.CreateUserRole) (id int, err error)
	Update(updateRole *domain.UpdateUserRole) (err error)
	Delete(id int) (err error)
}

type repository struct {
	Chat
	File
	Message
	Status
	UserRole
}

func New(db *sql.DB, logger logging.Logger) (repository *repository) {
	return repository
}
