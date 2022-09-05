package service

import (
	"itec.chat/internal/models"
	"itec.chat/internal/repository"
	"itec.chat/pkg/logging"
)

type userService struct {
	logger logging.Logger
	repo   *repository.Repository
}

func NewUserService(repo *repository.Repository, logger logging.Logger) UserService {
	return &userService{
		logger: logger,
		repo:   repo,
	}
}

func (u *userService) GetByID(id int) (user *models.User, err error) {
	return nil, nil
}
func (u *userService) GetAll(limit, offset int) (users []models.User, err error) {
	return nil, nil
}
func (u *userService) Create(createUser *models.CreateUser) (id int, err error) {
	return 0, nil
}
func (u *userService) Update(updateUser *models.UpdateUser) (err error) {
	return nil
}
func (u *userService) Delete(id int) (err error) {
	return nil
}

func (u *userService) GetUsersByChatID(id int) (users []models.User, err error) {
	return nil, nil
}
func (u *userService) UpdateStatus(updateUserStatus *models.UpdateUserStatus) (err error) {
	return nil
}
