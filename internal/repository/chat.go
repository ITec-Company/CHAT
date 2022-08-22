package repository

import (
	"database/sql"
	"itec.chat/internal/models"
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

func (rep *chat) GetByID(id int) (chat *models.Chat, err error) {
	query := `SELECT id, name, photo_url, created_at, updated_at 
			FROM chats 
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

func (rep *chat) GetByUserID(id int) (chats []models.ChatByUser, err error) {
	query := `WITH chat_ids AS (
			SELECT chat_id, is_admin
			FROM chats_users 
			WHERE user_id = $1
		)
    	SELECT id, name, photo_url
			FROM chats 
			WHERE id = any (chat_ids)
			GROUP BY id, name, photo_url, is_admin`

	rows, err := rep.db.Query(query, id)
	if err != nil {
		rep.logger.Errorf("error occurred while getting chats by userID. err: %s", err)
		return nil, err
	}

	for rows.Next() {
		chat := models.ChatByUser{}
		err := rows.Scan(
			&chat.ID,
			&chat.Name,
			&chat.PhotoURL,
			&chat.IsAdmin,
		)

		if err != nil {
			rep.logger.Errorf("error occurred while getting chats by userID. err: %s", err)
			continue
		}

		chats = append(chats, chat)
	}

	return chats, nil
}

func (rep *chat) Create(createChat *models.CreateChat) (id int, err error) {
	tx, err := rep.db.Begin()
	if err != nil {
		return 0, err
	}

	query := `INSERT INTO chats(name, photo_url) values ($1, $2) RETURNING id`

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

func (rep *chat) Update(updateChat *models.UpdateChat) (err error) {
	tx, err := rep.db.Begin()
	if err != nil {
		return err
	}

	query := `UPDATE chats 
			SET name = COALESCE(NULLIF($2, ''), name),
				photo_url = COALESCE(NULLIF($3, ''), photo_url) 
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
		tx.Commit()
		return ErrNoRowsAffected
	}

	return tx.Commit()
}

func (rep *chat) Delete(id int) (err error) {
	tx, err := rep.db.Begin()
	if err != nil {
		return err
	}

	query := `DELETE FROM  chats 
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
