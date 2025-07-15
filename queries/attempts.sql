-- name: GetAttemptByID :one
select * from attempts where id = @id;

-- name: CreateAttempt :one
insert into attempts(quiz_id, user_id) values(@quiz_id,@user_id) returning *;

