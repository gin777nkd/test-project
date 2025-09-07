-- name: CreateUser :execlastid
INSERT INTO users (name, email) VALUES (?, ?);

-- name: GetUser :one
SELECT id, name, email FROM users WHERE id = ?;

-- name: ListUsers :many
SELECT id, name, email FROM users;
