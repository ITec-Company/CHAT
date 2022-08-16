CREATE TYPE user_status AS enum ('active',  'away', 'offline');

CREATE TYPE message_status AS enum ('readed', 'unreaded');

CREATE TYPE role AS enum ('student', 'teacher');

CREATE TABLE IF NOT EXISTS users(
  id SERIAL PRIMARY KEY , 
  profile_id INTEGER NOT NULL,  
  name TEXT NOT NULL,
  role role NOT NULL,
  last_activity TIMESTAMP NOT NULL, 
  status user_status NOT NULL 
);

CREATE TABLE IF NOT EXISTS personal_chats(
  id SERIAL PRIMARY KEY , 
  user_1 TEXT NOT NULL , 
  user_2 TEXT NOT NULL

);

CREATE TABLE IF NOT EXISTS group_chats(
  id SERIAL PRIMARY KEY , 
  name TEXT NOT NULL , 
  photo TEXT not null
);

CREATE TABLE IF NOT EXISTS chats_users(
  chats_id INTEGER REFERENCES group_chats(id) ON DELETE CASCADE NOT NULL, 
  user_id INTEGER REFERENCES users(id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE IF NOT EXISTS group_messages(
  id SERIAL PRIMARY KEY,
  chats_id INTEGER REFERENCES group_chats(id) ON DELETE CASCADE NOT NULL, 
  user_id INTEGER REFERENCES users(id) ON DELETE CASCADE NOT NULL,
  body TEXT ,
  created_at timestamp NOT NULL ,
  updated_at timestamp NOT NULL 
);

CREATE TABLE IF NOT EXISTS personal_messages(
  id SERIAL PRIMARY KEY,
  chats_id INTEGER REFERENCES personal_chats(id) ON DELETE CASCADE NOT NULL, 
  user_id INTEGER REFERENCES users(id) ON DELETE CASCADE NOT NULL,
  body TEXT ,
  created_at timestamp NOT NULL ,
  updated_at timestamp NOT NULL 
);

CREATE TABLE IF NOT EXISTS files(
  id SERIAL PRIMARY KEY , 
  group_messages_id INTEGER REFERENCES group_messages(id) ON DELETE CASCADE,
  personal_messages_id INTEGER REFERENCES personal_messages(id) ON DELETE CASCADE,

  data bytea NOT NULL
);

CREATE TABLE IF NOT EXISTS messages_read_by_users(
  group_messages_id INTEGER REFERENCES group_messages(id) ON DELETE CASCADE,
  personal_messages_id INTEGER REFERENCES personal_messages(id) ON DELETE CASCADE,
  user_id INTEGER REFERENCES users(id) ON DELETE CASCADE NOT NULL
);

