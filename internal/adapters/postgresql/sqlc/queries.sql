-- name: ListAll :many
SELECT * FROM subscribtions;

-- name: FindSubscribtionByID :one
SELECT * FROM subscribtions WHERE id = $1;

-- name: CreateSubscribtion :one
INSERT INTO subscribtions (service_name, price_in_ruble, id_user, start_date, end_date) 
VALUES ($1, $2, $3, $4, $5);

-- name: UpdateSubscribtionByID :one
UPDATE subscribtions
SET 
service_name = $2,
price_in_rubles = $3,
id_user = $4,
start_date = $5,
end_date = $6
WHERE id = $1;

-- name: DeleteSubscribtionByID :one
DELETE FROM subscribtions WHERE id = $1;