ALTER TABLE IF EXISTS users
RENAME COLUMN hashed_password to password;

ALTER TABLE users
DROP COLUMN IF EXISTS created_at;

ALTER TABLE users
DROP COLUMN IF EXISTS password_changed_at;