-- name: CreateRoomUser :one
insert into room_users(room_id, user_id) values(@room_id, @user_id) returning *;

