-- name: SetQuizPublished :one
update quizzes set published = @published where id = @id returning *;

-- name: QuestionCount :one
select count(*) from questions where quiz_id = @quiz_id;

-- name: AllQuizzes :many
select * from quizzes order by created_at desc;

-- name: PublishedQuizzes :many
select * from quizzes where published = true order by created_at desc;

-- name: Quiz :one
select * from quizzes where id = @id;

-- name: UpdateQuiz :one
update quizzes set name = @name, description = @description where id = @id returning *;

-- name: CreateQuiz :one
insert into quizzes(name,description) values(@name,@description) returning *;

