  
-- name: FindAll :many
SELECT p_name FROM item;

-- name: CreateItem :exec
INSERT INTO item (p_name, price) VALUES ($1,$2);