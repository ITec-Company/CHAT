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
	user, err = u.repo.User.GetByID(id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (u *userService) GetAll(limit, offset int) (users []models.User, err error) {
	users, err = u.repo.User.GetAll(limit, offset)
	if err != nil {
		return nil, err
	}
	return users, nil
}
func (u *userService) Create(createUser *models.CreateUser) (id int, err error) {
	id, err = u.repo.User.Create(createUser)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (u *userService) Update(updateUser *models.UpdateUser) (err error) {
	err = u.repo.User.Update(updateUser)
	if err != nil {
		return err
	}
	return nil
}
func (u *userService) Delete(id int) (err error) {
	err = u.repo.User.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (u *userService) GetUsersByChatID(id int) (users []models.User, err error) {
	users, err = u.repo.User.GetUsersByChatID(id)
	if err != nil {
		return nil, err
	}
	return users, nil
}
func (u *userService) UpdateStatus(updateUserStatus *models.UpdateUserStatus) (err error) {
	err = u.repo.User.UpdateStatus(updateUserStatus)
	if err != nil {
		return err
	}
	return nil
}
