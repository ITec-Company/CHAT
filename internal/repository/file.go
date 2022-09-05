package repository

import (
	"database/sql"
	"itec.chat/internal/models"
	"itec.chat/pkg/logging"
)

type file struct {
	db     *sql.DB
	logger *logging.Logg
}

func NewFileRepository(db *sql.DB, logger *logging.Logg) (fileRepository File) {
	return &file{
		db:     db,
		logger: logger,
	}
}

func (rep *file) GetByID(id int) (file *models.File, err error) {
	query := `SELECT id, message_id, data_url
			FROM files 
			WHERE id = $1`

	if err = rep.db.QueryRow(query, id).
		Scan(
			&file.ID,
			&file.Message.ID,
			&file.URL,
		); err != nil {
		rep.logger.Errorf("error occured while getting file by id, err: %s", err)
		return nil, err
	}

	return file, nil
}

func (rep *file) GetByChatID(id int) (files []models.FileResponse, err error) {
	query := `SELECT  f.id, m.id, f.data_url
				FROM messages m
				RIGHT JOIN files f ON m.id = f.message_id
				WHERE chat_id = $1 AND is_deleted = FALSE`

	rows, err := rep.db.Query(query, id)
	if err != nil {
		rep.logger.Errorf("error occurred while getting files by userID. err: %s", err)
		return nil, err
	}

	for rows.Next() {
		file := models.FileResponse{}
		err := rows.Scan(
			&file.ID,
			&file.MessageID,
			&file.URL,
		)

		if err != nil {
			rep.logger.Errorf("error occurred while getting files by userID. err: %s", err)
			continue
		}

		files = append(files, file)
	}

	return files, nil
}

func (rep *file) Create(createFile *models.CreateFile) (id int, err error) {
	tx, err := rep.db.Begin()
	if err != nil {
		return 0, err
	}

	query := `INSERT INTO files(message_id, data_url) values ($1, $2) RETURNING id`

	if err = tx.QueryRow(query,
		createFile.MessageID, createFile.URL).
		Scan(
			&id,
		); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, nil
}

func (rep *file) Update(updateFile *models.UpdateFile) (err error) {
	tx, err := rep.db.Begin()
	if err != nil {
		return err
	}

	query := `UPDATE files 
			SET data_url = $2
			WHERE id = $1`

	result, err := tx.Exec(query, updateFile.URL)
	if err != nil {
		tx.Rollback()
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}

	if rowsAffected < 1 {
		tx.Commit()
		return ErrNoRowsAffected
	}

	return tx.Commit()
}

func (rep *file) Delete(id int) (err error) {
	tx, err := rep.db.Begin()
	if err != nil {
		return err
	}

	query := `DELETE FROM  files 
			WHERE id = $1`

	result, err := tx.Exec(query, id)
	if err != nil {
		tx.Rollback()
		return err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		tx.Rollback()
		return err
	}

	if rowsAffected < 1 {
		tx.Commit()
		return ErrNoRowsAffected
	}

	return tx.Commit()
}
