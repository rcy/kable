-- name: CreateThread :one
insert into threads(user_id, bot_id) values(@user_id,@bot_id) returning *;

-- name: BotThreads :many
select * from threads where bot_id = @bot_id and user_id = @user_id;

-- name: UserThreadByID :one
select * from threads where user_id = @user_id and id = @id;
