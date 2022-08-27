package repository

import (
	"database/sql"
	"itec.chat/internal/models"
	"itec.chat/pkg/logging"
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

	query := `INSERT INTO messages(chat_id, created_by) values ($1, $2, $3) RETURNING id`

	if err = tx.QueryRow(query,
		createMessage.ChatID, createMessage.UserID, createMessage.Body).
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
			SET body = $2
			WHERE id = $1`

	result, err := tx.Exec(query, updateMessage.Body)
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
			SET is_deleted = true
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
