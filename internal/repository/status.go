package repository

import (
	"database/sql"
	"itec.chat/internal/models"
	"itec.chat/pkg/logging"
)

type status struct {
	db     *sql.DB
	logger *logging.Logg
}

func NewStatusRepository(db *sql.DB, logger *logging.Logg) (statusRepository Status) {
	return &status{
		db:     db,
		logger: logger,
	}
}

func (rep *status) GetByID(id int) (status *models.Status, err error) {
	query := `SELECT id, name
			FROM statuses 
			WHERE id = $1 
			GROUP BY id, name`

	if err = rep.db.QueryRow(query, id).
		Scan(
			&status.ID,
			&status.Name,
		); err != nil {
		rep.logger.Errorf("error occured while getting status by id, err: %s", err)
		return nil, err
	}

	return status, nil
}

func (rep *status) GetAll(limit, offset int) (statuss []Status, err error) {
	return
}

func (rep *status) Create(createStatus *models.CreateStatus) (id int, err error) {
	tx, err := rep.db.Begin()
	if err != nil {
		return 0, err
	}

	query := `INSERT INTO statuses(name) values ($1) RETURNING id`

	if err = tx.QueryRow(query,
		createStatus.Name).
		Scan(
			&id,
		); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, nil
}

func (rep *status) Update(updateStatus *models.UpdateStatus) (err error) {
	tx, err := rep.db.Begin()
	if err != nil {
		return err
	}

	query := `UPDATE statuss 
			SET nane = $2
			WHERE id = $1`

	result, err := tx.Exec(query, updateStatus.Name)
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

func (rep *status) Delete(id int) (err error) {
	tx, err := rep.db.Begin()
	if err != nil {
		return err
	}

	query := `DELETE FROM  statuss 
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
