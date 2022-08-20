package repository

import (
	"database/sql"
	"itec.chat/internal/models"
	"itec.chat/pkg/logging"
)

type userRole struct {
	db     *sql.DB
	logger *logging.Logger
}

func NewUserRoleRepository(db *sql.DB, logger *logging.Logger) (userRoleRepository UserRole) {
	return &userRole{
		db:     db,
		logger: logger,
	}
}

func (rep *userRole) GetByID(id int) (role *models.UserRole, err error) {
	return
}

func (rep *userRole) GetAll(limit, offset int) (roles []UserRole, err error) {
	return
}

func (rep *userRole) Create(createRole *models.CreateUserRole) (id int, err error) {
	return
}

func (rep *userRole) Update(updateRole *models.UpdateUserRole) (err error) {
	return
}

func (rep *userRole) Delete(id int) (err error) {
	return
}
