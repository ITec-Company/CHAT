CREATE TABLE IF NOT EXISTS statuses (
                                        id SERIAL PRIMARY KEY,
                                        name VARCHAR(200) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS roles (
                                        id SERIAL PRIMARY KEY,
                                        name VARCHAR(200) NOT NULL UNIQUE
);

CREATE TABLE IF NOT EXISTS users (
                                        id SERIAL PRIMARY KEY,
                                        profile_id INTEGER NOT NULL UNIQUE,
                                        name VARCHAR(200) NOT NULL,
                                        last_activity TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                        role_id INTEGER DEFAULT 0 NOT NULL ,
                                        status_id INTEGER DEFAULT 0 NOT NULL ,
                                        FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE SET DEFAULT ,
                                        FOREIGN KEY (status_id) REFERENCES statuses(id) ON DELETE SET DEFAULT
);

CREATE TABLE IF NOT EXISTS chats (
                                        id SERIAL PRIMARY KEY,
                                        name VARCHAR(200) NOT NULL,
                                        photo_url VARCHAR(200) NOT NULL DEFAULT '',
                                        is_deleted bool NOT NULL DEFAULT false,
                                        created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                        updated_at timestamp
);

CREATE TABLE IF NOT EXISTS chats_users (
                                        id SERIAL PRIMARY KEY,
                                        is_admin BOOL DEFAULT FALSE,
                                        chat_id INTEGER NOT NULL,
                                        user_id INTEGER NOT NULL,
                                        FOREIGN KEY (chat_id) REFERENCES chats(id) ON DELETE CASCADE,
                                        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS messages (
                                        id SERIAL PRIMARY KEY,
                                        chat_id INTEGER NOT NULL,
                                        created_by INTEGER NOT NULL,
                                        body TEXT,
                                        is_deleted bool NOT NULL DEFAULT false,
                                        created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
                                        updated_at timestamp,
                                        FOREIGN KEY (chat_id) REFERENCES chats(id) ON DELETE CASCADE,
                                        FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS files (
                                        id SERIAL PRIMARY KEY,
                                        message_id INTEGER NOT NULL ,
                                        data_url VARCHAR(250) NOT NULL UNIQUE,
                                        FOREIGN KEY (message_id) REFERENCES messages(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS messages_unread_by_users (
                                        message_id INTEGER NOT NULL,
                                        user_id INTEGER NOT NULL,
                                        PRIMARY KEY(message_id, user_id),
                                        FOREIGN KEY (message_id) REFERENCES messages(id) ON DELETE CASCADE,
                                        FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);