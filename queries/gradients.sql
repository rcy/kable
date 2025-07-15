-- name: InsertGradient :one
insert into gradients(user_id, gradient) values(@user_id, @gradient) returning *;

