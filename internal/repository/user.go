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

func (rep *user) GetAll(limit, offset int) (users []models.User, err error) {
	query := `SELECT * FROM users`

	rows, err := rep.db.Query(query)
	if err != nil {
		rep.logger.Errorf("error occured while getting all users, err: %s", err)
		return nil, err
	}

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
			continue
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
		last_activity = COALESCE(NULLIF($3, ''), last_activity), 
		role_id = COALESCE(NULLIF($4, ''), role_id), 
		status_id = COALESCE(NULLIF($5, ''), status_id)
		WHERE id = $1`

	result, err := tx.Exec(query,
		updateUser.ID,
		updateUser.LastActivity,
		updateUser.RoleID,
		updateUser.StatusID)
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
	query := `SELECT * FROM USERS 
    		WHERE ID = (SELECT user_id from chats_users 
			WHERE user_id = 1)`

	rows, err := rep.db.Query(query)
	if err != nil {
		rep.logger.Errorf("error occured while getting all users, err: %s", err)
		return nil, err
	}

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
			continue
		}
		users = append(users, user)
	}

	return users, nil
}

func (rep *user) AssignUserToChatAdmin(chatID, userID int) (err error) {
	tx, err := rep.db.Begin()
	if err != nil {
		return err
	}

	query := `UPDATE chats_users
		SET is_admin = TRUE
		WHERE chat_id = $1 AND user_id = $2`

	result, err := tx.Exec(query,
		chatID,
		userID)
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

func (rep *user) UnAssignUserFromChatAdmin(chatID, userID int) (err error) {
	tx, err := rep.db.Begin()
	if err != nil {
		return err
	}

	query := `UPDATE chats_users
		SET is_admin = FALSE
		WHERE chat_id = $1 AND user_id = $2`

	result, err := tx.Exec(query,
		chatID,
		userID)
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

func (rep *user) InviteUserToChat(chatID, userID int) (err error) {
	tx, err := rep.db.Begin()
	if err != nil {
		return err
	}

	query := `INSERT INTO chats_users(is_admin, chat_id, user_id) 
	values (FALSE, $1, $2) RETURNING id`

	if err = tx.QueryRow(query,
		chatID, userID).
		Err(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}

func (rep *user) RemoveUserFromChat(chatID, userID int) (err error) {
	tx, err := rep.db.Begin()
	if err != nil {
		return err
	}

	query := `DELETE FROM chats_users
	WHERE user_id = $1`

	if err = tx.QueryRow(query,
		chatID, userID).
		Err(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
