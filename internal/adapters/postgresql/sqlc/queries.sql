-- name: ListAll :many
SELECT * FROM subscribtions;

-- name: FindProductByID :one
SELECT * FROM subscribtions WHERE id = $1;