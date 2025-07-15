-- name: CreateKidParent :one
insert into kids_parents(kid_id, parent_id) values(@kid_id, @parent_id) returning *;
