CREATE TABLE IF NOT EXISTS url (
  id       INTEGER PRIMARY KEY,
  original text    NOT NULL,
  shorten  text    NOT NULL UNIQUE
);
