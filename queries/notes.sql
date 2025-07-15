-- name: PublishedNotes :many
select * from notes where published = 1;

-- name: UserNotes :many
select * from notes where owner_id = @owner_id order by created_at desc;

-- name: CreateNote :one
insert into notes(owner_id,body) values(@owner_id,@body) returning *;

-- name: UpdateNote :one
update notes set body = @body where id = @id and owner_id = @owner_id returning *;

-- name: DeleteNote :exec
delete from notes where id = @id and owner_id = @owner_id;

