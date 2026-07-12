-- name: CreateBotMessage :one
insert into bot_messages(thread_id, role, content) values(@thread_id,@role,@content) returning *;

-- name: ThreadMessages :many
select * from bot_messages where thread_id = @thread_id order by created_at asc;
