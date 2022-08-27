package repository

import (
	"database/sql"
	"github.com/lib/pq"
	"itec.chat/internal/models"
	"itec.chat/pkg/logging"
	"strconv"
	"strings"
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

func (rep *chat) GetByID(id int) (chat *models.ChatResponse, err error) {
	query := `WITH a AS (select ARRAY (
										  (
											  SELECT admin_id
											  FROM chats_admins
											  WHERE chat_id = $1
										  )
									  ) admins_array,
									ARRAY (
										   SELECT user_id
										   FROM chats_users cu
										   WHERE chat_id = $1
									   ) users_array)
				select c.name, c.photo_url, array_agg(DISTINCT (u.id, u.name)) as admins, array_agg(DISTINCT (ua.id, ua.name)) as users
				from chats as c, a
				left join users u on u.id = any (a.admins_array)
				left join users ua on ua.id = any (a.users_array)
				where c.id = $1
				group by c.name, c.photo_url`

	var users []string
	var admins []string
	if err = rep.db.QueryRow(query, id).
		Scan(
			&chat.ID,
			&chat.Name,
			&chat.PhotoURL,
			&chat.CreatedAt,
			&chat.UpdatedAt,
			pq.Array(&admins),
			pq.Array(&users),
		); err != nil {
		rep.logger.Errorf("error occured while getting chat by id, err: %s", err)
		return nil, err
	}

	for _, t := range admins {
		t = strings.Replace(t, "(", "", -1)
		t = strings.Replace(t, ")", "", -1)
		t = strings.Replace(t, "\"", "", -1)
		data := strings.Split(t, ",")
		var admin models.User
		admin.ID, _ = strconv.Atoi(data[0])
		admin.Name = data[1]
		chat.Admins = append(chat.Admins, admin)
	}

	for _, t := range users {
		t = strings.Replace(t, "(", "", -1)
		t = strings.Replace(t, ")", "", -1)
		t = strings.Replace(t, "\"", "", -1)
		data := strings.Split(t, ",")
		var user models.User
		user.ID, _ = strconv.Atoi(data[0])
		user.Name = data[1]
		chat.Users = append(chat.Users, user)
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
			pq.Array(&chat.AdminsIDs),
			pq.Array(&chat.UsersIDs),
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

func (rep *chat) AddUserToChat(userID, chatID int) (err error) {
	query := `INSERT INTO chats_users (user_id, chat_id) values ($1, $2)`

	result, err := rep.db.Exec(query,
		userID,
		chatID)

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		return ErrNoRowsAffected
	}

	return nil
}

func (rep *chat) RemoveUserFromChat(userID, chatID int) (err error) {
	query := `DELETE FROM chats_users WHERE user_id = $1 AND chat_id = $2`

	result, err := rep.db.Exec(query,
		userID,
		chatID)

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		return ErrNoRowsAffected
	}

	return nil
}

func (rep *chat) PromoteUserToAdmin(userID, chatID int) (err error) {
	query := `INSERT INTO chats_admins (admin_id, chat_id) values ($1, $2)`

	result, err := rep.db.Exec(query,
		userID,
		chatID)

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		return ErrNoRowsAffected
	}

	return nil
}

func (rep *chat) LowerAdminToUser(userID, chatID int) (err error) {
	query := `DELETE FROM chats_admins WHERE admin_id = $1 AND chat_id = $2`

	result, err := rep.db.Exec(query,
		userID,
		chatID)

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected < 1 {
		return ErrNoRowsAffected
	}

	return nil
}
