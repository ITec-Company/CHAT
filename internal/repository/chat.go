package repository

import (
	"database/sql"
	"itec.chat/internal/models"
	"itec.chat/pkg/logging"
	"time"
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
	query := `SELECT c.id, c.name, c.photo_url, c.created_at, c.updated_at, a.admins_ids, u.users_ids
				FROM chats c, LATERAL (
					SELECT ARRAY (
								   SELECT admin_id
								   FROM chats_admins ca
								   WHERE ca.chat_id = $1
							   ) as admins_ids
					) a, LATERAL (
					SELECT ARRAY (
								   SELECT user_id
								   FROM chats_users ca
								   WHERE ca.chat_id = $1
							   ) as users_ids
					) u
				WHERE id = $1
				GROUP BY c.id, c.name, c.photo_url, c.created_at, c.updated_at, a.admins_ids, u.users_ids;`

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
	query := `SELECT c.chat_id, chat. name, chat.photo_url, users_array, admins_array
				FROM chats_users c, LATERAL (
					SELECT ARRAY (
						SELECT user_id
						FROM chats_users cu
						WHERE chat_id = c.chat_id)
					) users_array, LATERAL (
						SELECT ARRAY (
							SELECT admin_id
							FROM chats_admins ca
							WHERE chat_id = c.chat_id)
					) admins_array, LATERAL (
						SELECT name, photo_url
						FROM chats
						WHERE id = c.chat_id
					) chat
				WHERE user_id = $1`

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
			SET updated_at = $2,
			    name = COALESCE(NULLIF($3, ''), name),
				photo_url = COALESCE(NULLIF($4, ''), photo_url)
			WHERE id = $1`

	result, err := tx.Exec(query,
		updateChat.ID,
		time.Now(),
		updateChat.Name,
		updateChat.PhotoURL)
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

	query := `UPDATE chats 
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
