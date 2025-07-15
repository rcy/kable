-- name: CreateThread :one
insert into threads(user_id, thread_id, assistant_id) values(@user_id,@thread_id,@assistant_id) returning *;

-- name: AssistantThreads :many
select * from threads where assistant_id = @assistant_id and user_id = @user_id;

-- name: UserThreadByID :one
select * from threads where user_id = @user_id and thread_id = @thread_id;

