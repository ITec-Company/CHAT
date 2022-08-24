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

type User interface {
	GetByID(id int) (user *models.User, err error)
	GetAll(limit, offset int) (users []models.User, err error)
	Create(createUser *models.CreateUser) (id int, err error)
	Update(updateUser *models.UpdateUser) (err error)
	Delete(id int) (err error)

	GetUsersByChatID(id int) (users []models.User, err error)
	AssignUserToChatAdmin(chatID, userID int) (err error)
	UnAssignUserFromChatAdmin(chatID, userID int) (err error)
	InviteUserToChat(chatID, userID int) (err error)
	RemoveUserFromChat(chatID, userID int) (err error)
}

type Chat interface {
	GetByID(id int) (chat *models.Chat, err error)
	GetByUserID(id int) (chats []models.ChatByUser, err error)
	Create(createChat *models.CreateChat) (id int, err error)
	Update(updateChat *models.UpdateChat) (err error)
	Delete(id int) (err error)
}

type File interface {
	GetByID(id int) (file *models.File, err error)
	GetAll(limit, offset int) (files []File, err error)
	Create(createFile *models.CreateFile) (id int, err error)
	Update(updateFile *models.UpdateFile) (err error)
	Delete(id int) (err error)
}

type Message interface {
	GetByID(id int) (message *models.Message, err error)
	GetAll(limit, offset int) (messages []Message, err error)
	Create(createMessage *models.CreateMessage) (id int, err error)
	Update(updateMessage *models.UpdateMessage) (err error)
	Delete(id int) (err error)
}

type Status interface {
	GetByID(id int) (status *models.Status, err error)
	GetAll(limit, offset int) (statuses []Status, err error)
	Create(createStatus *models.CreateStatus) (id int, err error)
	Update(updateStatus *models.UpdateStatus) (err error)
	Delete(id int) (err error)
}

type UserRole interface {
	GetByID(id int) (role *models.UserRole, err error)
	GetAll(limit, offset int) (roles []UserRole, err error)
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
	User
}

func New(db *sql.DB, logger logging.Logger) (repository *repository) {
	return repository
}
