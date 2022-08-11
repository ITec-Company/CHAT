CREATE TYPE user_status AS enum ('active',  'away', 'offline');

CREATE TYPE message_status AS enum ('readed', 'unreaded');


CREATE TABLE IF NOT EXISTS  users(
  id INTEGER PRIMARY KEY , 
  name TEXT NOT NULL,
  status user_status NOT NULL 
);

CREATE TABLE IF NOT EXISTS  chats(
  id INTEGER PRIMARY KEY , 
  name TEXT NOT NULL
);

CREATE TABLE IF NOT EXISTS  chats_users(
  chats_id INTEGER REFERENCES chats(id) ON DELETE CASCADE NOT NULL, 
  user_id INTEGER REFERENCES users(id) ON DELETE CASCADE NOT NULL
);

CREATE TABLE IF NOT EXISTS  files(
  id INTEGER PRIMARY KEY , 
  data bytea NOT NULL
);

CREATE TABLE IF NOT EXISTS  massages(
  id INTEGER PRIMARY KEY,
  chats_id INTEGER REFERENCES chats(id) ON DELETE CASCADE NOT NULL, 
  user_id INTEGER REFERENCES users(id) ON DELETE CASCADE NOT NULL,
  body TEXT ,
  file_id INTEGER REFERENCES files(id) ON DELETE CASCADE NOT NULL,
  created_at timestamp NOT NULL ,
  updated_at timestamp NOT NULL 
);

CREATE TABLE IF NOT EXISTS  messages_users(
  messages_id INTEGER REFERENCES massages(id) ON DELETE CASCADE NOT NULL, 
  user_id INTEGER REFERENCES users(id) ON DELETE CASCADE NOT NULL, 
  status message_status NOT NULL 
);

