package repository

import (
	"database/sql"
	"itec.chat/internal/domain"
	"itec.chat/pkg/logging"
)

type chat struct {
	db     *sql.DB
	logger *logging.Logg
}

func NewChatRepository(db *sql.DB, logger *logging.Logg) (chatRepository Chat) {
	return &chat{
		db:     db,
		logger: logger,
	}
}

func (rep *chat) GetByID(id int) (chat *domain.Chat, err error) {
	query := `SELECT id, name, photo_url, created_at, updated_at 
			FROM chat 
			WHERE id = $1 
			GROUP BY id, name, photo_url, created_at, updated_at`

	if err = rep.db.QueryRow(query, id).
		Scan(
			&chat.ID,
			&chat.Name,
			&chat.PhotoURL,
			&chat.CreatedAt,
			&chat.UpdatedAt,
		); err != nil {
		rep.logger.Errorf("error occured while getting chat by id, err: %s", err)
		return nil, err
	}

	return chat, nil
}

func (rep *chat) GetAll(limit, offset int) (chats []Chat, err error) {
	return
}

func (rep *chat) Create(createChat *domain.CreateChat) (id int, err error) {
	tx, err := rep.db.Begin()
	if err != nil {
		return 0, err
	}

	query := `INSERT INTO chat(name, photo_url) values ($1, $2) RETURNING id`

	if err = tx.QueryRow(query,
		createChat.Name, createChat.PhotoURL).
		Scan(
			&id,
		); err != nil {
		tx.Rollback()
		return 0, err
	}

	return id, nil
}

func (rep *chat) Update(updateChat *domain.UpdateChat) (err error) {
	tx, err := rep.db.Begin()
	if err != nil {
		return err
	}

	query := `UPDATE chat 
			SET name = $2,
				photo_url = $3 
			WHERE id = $1`

	result, err := tx.Exec(query, updateChat.Name, updateChat.PhotoURL)
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
		tx.Rollback()
		return ErrNoRowsAffected
	}

	return tx.Commit()
}

func (rep *chat) Delete(id int) (err error) {
	tx, err := rep.db.Begin()
	if err != nil {
		return err
	}

	query := `DELETE FROM  chat 
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
		tx.Rollback()
		return ErrNoRowsAffected
	}

	return tx.Commit()
}
