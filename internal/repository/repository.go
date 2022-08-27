package repository

import (
	"database/sql"
	"errors"
	"itec.chat/internal/models"
	"itec.chat/pkg/logging"
)

var (
	ErrNoRowsAffected = errors.New("now rows affected")
)

type Chat interface {
	GetByID(id int) (chat *models.ChatResponse, err error)
	GetByUserID(id int) (chats []models.ChatByUser, err error)
	Create(createChat *models.CreateChat) (id int, err error)
	Update(updateChat *models.UpdateChat) (err error)
	Delete(id int) (err error)
	AddUserToChat(userID, chatID int) (err error)
	RemoveUserFromChat(userID, chatID int) (err error)
	PromoteUserToAdmin(userID, chatID int) (err error)
	LowerAdminToUser(userID, chatID int) (err error)
}

type File interface {
	GetByID(id int) (file *models.File, err error)
	GetByChatID(id int) (files []models.FileResponse, err error)
	Create(createFile *models.CreateFile) (id int, err error)
	Update(updateFile *models.UpdateFile) (err error)
	Delete(id int) (err error)
}

type Message interface {
	GetByID(id int) (message *models.Message, err error)
	Create(createMessage *models.CreateMessage) (id int, err error)
	Update(updateMessage *models.UpdateMessage) (err error)
	Delete(id int) (err error)
	ReceiveMessage(messageID, userID int) (message *models.ReceiveMessage, err error)
}

type Status interface {
	GetByID(id int) (status *models.Status, err error)
	GetAll() (statuses []models.Status, err error)
	Create(createStatus *models.CreateStatus) (id int, err error)
	Update(updateStatus *models.UpdateStatus) (err error)
	Delete(id int) (err error)
}

type UserRole interface {
	GetByID(id int) (role *models.UserRole, err error)
	GetAll() (roles []models.UserRole, err error)
	Create(createRole *models.CreateUserRole) (id int, err error)
	Update(updateRole *models.UpdateUserRole) (err error)
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
