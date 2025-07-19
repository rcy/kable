-- name: SetQuizPublished :one
update quizzes set published = @published where not is_deleted and id = @id returning *;

-- name: AllQuizzes :many
select * from quizzes where not is_deleted order by created_at desc;

-- name: PublishedQuizzes :many
select * from quizzes where not is_deleted and published order by created_at desc;

-- name: Quiz :one
select * from quizzes where not is_deleted and id = @id;

-- name: UpdateQuiz :one
update quizzes set name = @name, description = @description where id = @id returning *;

-- name: CreateQuiz :one
insert into quizzes(name,description,user_id) values(@name,@description,@user_id) returning *;

-- name: PublishedUserQuizzes :many
select * from quizzes where not is_deleted and published and user_id = @user_id order by created_at desc;

-- name: AllUserQuizzes :many
select * from quizzes where not is_deleted and user_id = @user_id order by published desc;

-- name: DeleteQuiz :exec
update quizzes set is_deleted = true where id = @quiz_id and user_id = @user_id;
