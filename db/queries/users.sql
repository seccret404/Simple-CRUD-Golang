-- name: CreateUser :execresult
INSERT INTO users(username, password) VALUES(?, ?);

-- name: GetUserByID :one
SELECT id, username, password FROM users WHERE id= ?;