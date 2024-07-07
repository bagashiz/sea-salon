-- +goose Up
CREATE TABLE IF NOT EXISTS sessions (
	token TEXT PRIMARY KEY,
	data BYTEA NOT NULL,
	expiry TIMESTAMPTZ NOT NULL
);
CREATE INDEX sessions_expiry_key ON sessions (expiry);

-- +goose Down
DROP TABLE IF EXISTS sessions;
