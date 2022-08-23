package repository

import (
	"database/sql"

	"itec.chat/internal/models"
	"itec.chat/pkg/logging"
)

type user struct {
	db     *sql.DB
	logger *logging.Logg
}

func NewUserRepository(db *sql.DB, logger *logging.Logg) (userRepository User) {
	return &user{
		db:     db,
		logger: logger,
	}
}

func (rep *user) GetByID(id int) (user *models.User, err error) {
	query := `SELECT *
  FROM users
  WHERE id = $1`

	if err = rep.db.QueryRow(query, id).Scan(
		&user.ID,
		&user.ProfileID,
		&user.Name,
		&user.LastActivity,
		&user.Role.ID,
		&user.Status.ID,
	); err != nil {
		rep.logger.Errorf("error occured while getting user by id, err: %s", err)
		return nil, err
	}

	return user, nil
}

func (rep *user) GetAll(limit, offset int) (users []User, err error) {

	return
}

func (rep *user) Create(createUser *models.CreateUser) (id int, err error) {
	tx, err := rep.db.Begin()
	if err != nil {
		return 0, err
	}

	query := `INSERT INTO users(profile_id, name, last_activity, role_id, status_id) 
	values ($1, $2, $3, $4, $5) RETURNING id`

	if err = tx.QueryRow(query,
		createUser.ProfileID, createUser.Name, createUser.LastActivity, createUser.RoleID, createUser.StatusID).Scan(
		&id,
	); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, nil
}

func (rep *user) Update(updateUser *models.UpdateUser) (err error) {
	/*		tx, err := rep.db.Begin()
	if err != nil {
			return err
		}

		query := `UPDATE users
		SET`*/
	return
}

func (rep *user) Delete(id int) (err error) {
	tx, err := rep.db.Begin()
	if err != nil {
		return err
	}

	query := `DELETE FROM users 
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
