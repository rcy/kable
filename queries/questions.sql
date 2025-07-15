-- name: AttemptNextQuestion :one
select questions.* from questions
left join responses on responses.question_id = questions.id
where
  questions.id not in (select question_id from responses where responses.attempt_id = @attempt_id)
and
  questions.quiz_id = @quiz_id
order by random();

-- name: Question :one
select * from questions where id = @id;

-- name: CreateQuestion :one
insert into questions(quiz_id, text, answer) values(@quiz_id,@text,@answer) returning *;

-- name: UpdateQuestion :one
update questions set text = @text, answer = @answer where id = @id returning *;

-- name: QuizQuestions :many
select * from questions where quiz_id = @quiz_id;

