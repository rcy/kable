-- name: CreateFriend :one
insert into friends(a_id, b_id, b_role) values(@a_id, @b_id, @b_role) returning *;
