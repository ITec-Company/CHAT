package repository

import (
	"database/sql"
	"itec.chat/internal/models"
	"itec.chat/pkg/logging"
	"time"
)

type message struct {
	db     *sql.DB
	logger *logging.Logg
}

func NewMessageRepository(db *sql.DB, logger *logging.Logg) (messageRepository Message) {
	return &message{
		db:     db,
		logger: logger,
	}
}

func (rep *message) GetByID(id int) (message *models.Message, err error) {
	query := `SELECT id, chat_id, created_by, body, is_deleted, created_at, updated_at
			FROM messages 
			WHERE id = $1`

	if err = rep.db.QueryRow(query, id).
		Scan(
			&message.ID,
			&message.Chat.ID,
			&message.User.ID,
			&message.Body,
			&message.IsDeleted,
			&message.CreatedAt,
			&message.UpdatedAt,
		); err != nil {
		rep.logger.Errorf("error occured while getting message by id, err: %s", err)
		return nil, err
	}

	return message, nil
}

func (rep *message) Create(createMessage *models.CreateMessage) (id int, err error) {
	tx, err := rep.db.Begin()
	if err != nil {
		return 0, err
	}

	query := `WITH insert_message AS (
									INSERT INTO messages(chat_id, created_by, body)
										VALUES ($1, $2, $3) RETURNING id
										), users AS (
									SELECT user_id
										FROM chats_users
									   WHERE chat_id = $1 AND user_id != $2)
				INSERT INTO messages_unread_by_users(message_id, user_id)
				SELECT insert_message.id, users.user_id FROM insert_message, users`

	if err = tx.QueryRow(query,
		createMessage.ChatID,
		createMessage.UserID,
		createMessage.Body).
		Scan(
			&id,
		); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, nil
}

func (rep *message) Update(updateMessage *models.UpdateMessage) (err error) {
	tx, err := rep.db.Begin()
	if err != nil {
		return err
	}

	query := `UPDATE messages 
			SET updated_at = $2, 
			    body = $3
			WHERE id = $1`

	result, err := tx.Exec(query,
		time.Now(),
		updateMessage.Body)
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

func (rep *message) Delete(id int) (err error) {
	tx, err := rep.db.Begin()
	if err != nil {
		return err
	}

	query := `UPDATE messages 
			SET is_deleted = true,
			    updated_at = $2
			WHERE id = $1`

	result, err := tx.Exec(query,
		id,
		time.Now())
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

func (rep *message) ReceiveMessage(userID, messageID int) (message *models.ReceiveMessage, err error) {
	query := `WITH delete AS (
						DELETE FROM messages_unread_by_users 
						    WHERE user_id = $1 AND message_id = $2)
				SELECT messages.id, created_by, u.name, body, created_at, updated_at
				FROM messages
				JOIN users u on u.id = messages.created_by
				WHERE messages.id = $2`

	var updatedAt time.Time
	if err = rep.db.QueryRow(query,
		userID,
		messageID).
		Scan(
			&message.ID,
			&message.User.ID,
			&message.User.Name,
			&message.Body,
			&message.CreatedAt,
			&updatedAt,
		); err != nil {
		rep.logger.Errorf("error occured while getting message by id, err: %s", err)
		return nil, err
	}

	if !updatedAt.IsZero() {
		message.IsUpdated = true
	}

	return message, nil
}
