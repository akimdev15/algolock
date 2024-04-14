-- name: CreateQuestion :one
INSERT INTO questions (id, name, url, solved, updated_at)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetAllQuestions :many
SELECT * FROM questions;

-- name: GetRandomQuestionURL :one
SELECT url from questions ORDER BY RANDOM() LIMIT 1;