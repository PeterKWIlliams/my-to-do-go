-- name: CreateNote :exec
INSERT INTO tasks (id, created_at, updated_at, note, user_id)
VALUES (?, ?, ?, ?, ?);
--

-- name: GetNote :one
SELECT * FROM tasks WHERE id = ?;
--

-- name: GetNotesForUser :many
SELECT * FROM tasks WHERE user_id = ?;
--
