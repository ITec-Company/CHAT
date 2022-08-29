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
	query := `SELECT profile_id, name, last_activity, role_id, status_id
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

func (rep *user) GetAll(limit, offset int) (users []models.User, err error) {

	query := `SELECT profile_id, name, last_activity, role_id, status_id
	FROM users
	LIMIT $1 OFFSET $2`

	rows, err := rep.db.Query(query, limit, offset)
	if err != nil {
		rep.logger.Errorf("error occured while getting all users, err: %s", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := models.User{}
		if err = rows.Scan(
			&user.ID,
			&user.ProfileID,
			&user.Name,
			&user.LastActivity,
			&user.Role.ID,
			&user.Status.ID,
		); err != nil {
			rep.logger.Errorf("error occured while getting all users, err: %s", err)
			return
		}
		users = append(users, user)
	}

	return users, nil
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
	tx, err := rep.db.Begin()
	if err != nil {
		return err
	}

	query := `UPDATE users
		SET name = COALESCE(NULLIF($2, ''), name), 
		last_activity = COALESCE($3, last_activity), 
		role_id = COALESCE(NULLIF($4, 0), role_id), 
		status_id = COALESCE(NULLIF($5, 0), status_id)
		WHERE id = $1`

	result, err := tx.Exec(query,
		updateUser.ID,
		updateUser.LastActivity,
		updateUser.RoleID)
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

func (rep *user) GetUsersByChatID(id int) (users []models.User, err error) {
	query := `SELECT profile_id, name, last_activity, role_id, status_id
			FROM USERS 
    		WHERE ID IN (SELECT user_id from chats_users 
			WHERE chat_id = $1)`

	rows, err := rep.db.Query(query, id)
	if err != nil {
		rep.logger.Errorf("error occured while getting all users, err: %s", err)
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		user := models.User{}
		if err = rows.Scan(
			&user.ID,
			&user.ProfileID,
			&user.Name,
			&user.LastActivity,
			&user.Role.ID,
			&user.Status.ID,
		); err != nil {
			rep.logger.Errorf("error occured while getting all users, err: %s", err)
			return
		}
		users = append(users, user)
	}

	return users, nil
}

func (rep *user) UpdateStatus(updateUserStatus *models.UpdateUserStatus) (err error) {

	tx, err := rep.db.Begin()
	if err != nil {
		return err
	}

	query := `UPDATE users
		SET status_id = $2
		WHERE id = $1`

	result, err := tx.Exec(query,
		updateUserStatus.ID,
		updateUserStatus.StatusID)
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
