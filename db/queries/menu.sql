-- name: CreateMenu :execresult
INSERT INTO menus (name_product, price, description_product, image_product, stock_product) 
VALUES (?, ?, ?, ?, ?);

-- name: GetMenuByID :one
SELECT * FROM menus WHERE id = ?;

-- name: ListMenus :many
SELECT * FROM menus ORDER BY created_at DESC;

-- name: UpdateMenu :exec
UPDATE menus 
SET name_product = ?, price = ?, description_product = ?, image_product = ?, stock_product = ?
WHERE id = ?;

-- name: DeleteMenu :exec
DELETE FROM menus WHERE id = ?;
