-- name: CreateQuestion :one
INSERT INTO questions (id, name, url, solved, difficulty, updated_at, confidence)
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;

-- name: GetAllQuestions :many
SELECT * FROM questions;

-- name: GetRandomQuestionURL :one
SELECT url from questions ORDER BY RANDOM() LIMIT 1;

-- name: GetAllQuestionsByIds :many
SELECT * from questions WHERE id IN (sqlc.slice('ids'));

-- name: GetLatestQuestions :many
SELECT * FROM questions ORDER BY updated_at DESC LIMIT $1;

-- name: DeleteQuestionById :exec
DELETE FROM questions WHERE id = $1;

-- name: DeleteQuestionByURL :exec
DELETE FROM questions WHERE url = $1;
