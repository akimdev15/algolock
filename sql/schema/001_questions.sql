-- +goose Up

CREATE TABLE questions (
    id TEXT PRIMARY KEY,
    name TEXT NOT NULL,
    url TEXT NOT NULL,
    solved TEXT NOT NULL,
    difficulty TEXT NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    confidence TEXT NOT NULL
);

-- +goose Down
DROP TABLE questions;