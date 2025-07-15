-- name: RoomByKey :one
select * from rooms where key = @key;

-- name: CreateRoom :one
insert into rooms(key) values(@key) returning *;
