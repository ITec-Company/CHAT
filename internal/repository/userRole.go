package repository

import (
	"database/sql"
	"itec.chat/internal/models"
	"itec.chat/pkg/logging"
)

type userRole struct {
	db     *sql.DB
	logger *logging.Logg
}

func NewUserRoleRepository(db *sql.DB, logger *logging.Logg) (userRoleRepository UserRole) {
	return &userRole{
		db:     db,
		logger: logger,
	}
}

func (rep *userRole) GetByID(id int) (userRole *models.UserRole, err error) {
	query := `SELECT id, name
			FROM roles 
			WHERE id = $1 
			GROUP BY id, name`

	if err = rep.db.QueryRow(query, id).
		Scan(
			&userRole.ID,
			&userRole.Name,
		); err != nil {
		rep.logger.Errorf("error occured while getting userRole by id, err: %s", err)
		return nil, err
	}

	return userRole, nil
}

func (rep *userRole) GetAll(limit, offset int) (userRoles []UserRole, err error) {
	return
}

func (rep *userRole) Create(createUserRole *models.CreateUserRole) (id int, err error) {
	tx, err := rep.db.Begin()
	if err != nil {
		return 0, err
	}

	query := `INSERT INTO roles(name) values ($1) RETURNING id`

	if err = tx.QueryRow(query,
		createUserRole.Name).
		Scan(
			&id,
		); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, nil
}

func (rep *userRole) Update(updateUserRole *models.UpdateUserRole) (err error) {
	tx, err := rep.db.Begin()
	if err != nil {
		return err
	}

	query := `UPDATE roles 
			SET nane = $2
			WHERE id = $1`

	result, err := tx.Exec(query, updateUserRole.Name)
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

func (rep *userRole) Delete(id int) (err error) {
	tx, err := rep.db.Begin()
	if err != nil {
		return err
	}

	query := `DELETE FROM  roles 
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
