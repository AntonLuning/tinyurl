CREATE TABLE IF NOT EXISTS url (
  id         INTEGER PRIMARY KEY,
  original   text    NOT NULL,
  shorten_id text    NOT NULL UNIQUE
);
