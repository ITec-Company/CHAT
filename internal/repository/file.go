package repository

import (
	"database/sql"
	"itec.chat/internal/domain"
	"itec.chat/pkg/logging"
)

type file struct {
	db     *sql.DB
	logger *logging.Logger
}

func NewFileRepository(db *sql.DB, logger *logging.Logger) (fileRepository File) {
	return &file{
		db:     db,
		logger: logger,
	}
}

func (rep *file) GetByID(id int) (file *domain.File, err error) {
	return
}

func (rep *file) GetAll(limit, offset int) (files []File, err error) {
	return
}

func (rep *file) Create(createFile *domain.CreateFile) (id int, err error) {
	return
}

func (rep *file) Update(updateFile *domain.UpdateFile) (err error) {
	return
}

func (rep *file) Delete(id int) (err error) {
	return
}
