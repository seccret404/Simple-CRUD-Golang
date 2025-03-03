-- name: CreateUser :execresult
INSERT INTO users(username, password) VALUES(?, ?);

-- name: GetUserByUsername :one
SELECT id, username, password FROM users WHERE username= ?;