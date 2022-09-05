package service

import (
	"itec.chat/internal/models"
	"itec.chat/internal/repository"
	"itec.chat/pkg/logging"
)

type UserService interface {
	GetByID(id int) (user *models.User, err error)
	GetAll(limit, offset int) (users []models.User, err error)
	Create(createUser *models.CreateUser) (id int, err error)
	Update(updateUser *models.UpdateUser) (err error)
	Delete(id int) (err error)

	GetUsersByChatID(id int) (users []models.User, err error)
	UpdateStatus(updateUserStatus *models.UpdateUserStatus) (err error)
}

type Service struct {
	UserService
}

func NewService(repo *repository.Repository, logger logging.Logger) *Service {
	return &Service{
		UserService: NewUserService(repo, logger),
	}
}
