CREATE TYPE user_status AS enum ('active',  'away', 'offline');

CREATE TYPE message_status AS enum ('readed', 'unreaded');

CREATE TABLE IF NOt EXISTS roles(
  id SERIAL PRIMARY KEY, 
  user_role TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS users(
  id SERIAL PRIMARY KEY , 
  profile_id INTEGER NOT NULL,  
  name TEXT NOT NULL,
  role role NOT NULL,
  last_activity TIMESTAMP NOT NULL,
  role_id INTEGER REFERENCES roles(id) ON DELETE CASCADE NOT NULL, 
  status user_status NOT NULL 
);

CREATE TABLE IF NOT EXISTS group_chats(
  id SERIAL PRIMARY KEY , 
  name TEXT NOT NULL , 
  photo TEXT not null
);

CREATE TABLE IF NOT EXISTS chats_users(
  chat_id INTEGER REFERENCES group_chats(id) ON DELETE CASCADE NOT NULL, 
  user_id INTEGER REFERENCES users(id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE IF NOT EXISTS group_messages(
  id SERIAL PRIMARY KEY,
  chat_id INTEGER REFERENCES group_chats(id) ON DELETE CASCADE NOT NULL, 
  user_id INTEGER REFERENCES users(id) ON DELETE CASCADE NOT NULL,
  body TEXT ,
  created_at timestamp NOT NULL ,
  updated_at timestamp NOT NULL 
);

CREATE TABLE IF NOT EXISTS private_messages(
  id SERIAL PRIMARY KEY,
  user_id INTEGER REFERENCES users(id) ON DELETE CASCADE NOT NULL,
  body TEXT,  
  created_at timestamp NOT NULL,
  updated_at timestamp NOT NULL,
  status message_status NOT NULL
);

CREATE TABLE IF NOT EXISTS files(
  id SERIAL PRIMARY KEY , 
  group_messages_id INTEGER REFERENCES group_messages(id) ON DELETE CASCADE,
  data bytea NOT NULL
);

CREATE TABLE IF NOT EXISTS messages_read_by_users(
  group_messages_id INTEGER REFERENCES group_messages(id) ON DELETE CASCADE,
  user_id INTEGER REFERENCES users(id) ON DELETE CASCADE NOT NULL
);

