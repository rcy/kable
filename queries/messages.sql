-- name: RecentRoomMessages :many
select sqlc.embed(m), sqlc.embed(sender)
from messages m
join users sender on m.sender_id = sender.id
where m.room_id = @room_id
order by m.created_at desc
limit 128;

-- name: MessageByID :one
select * from messages where id = @id;

-- name: AdminRecentMessages :many
select sqlc.embed(m), sqlc.embed(sender)
 from messages m
 join users sender on m.sender_id = sender.id
 order by m.created_at desc
 limit 128;

-- name: AdminDeleteMessage :one
update messages
set body = '+++ deleted +++'
where id = @id
returning *;
