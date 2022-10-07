ALTER TABLE IF EXISTS users
RENAME COLUMN password to hashed_password;

ALTER TABLE users
ADD COLUMN created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP;

ALTER TABLE users
ADD COLUMN password_changed_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP;