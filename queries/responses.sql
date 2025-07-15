-- name: ResponseCount :one
select count(*) from responses where attempt_id = @attempt_id;

-- name: AttemptResponseIDs :many
select id from responses where attempt_id = @attempt_id;

-- name: CreateResponse :one
insert into responses(quiz_id, user_id, attempt_id, question_id, text) values(@quiz_id,@user_id,@attempt_id,@question_id,@text) returning *;

-- name: Responses :many
select
   responses.*,
   questions.answer question_answer,
   questions.text question_text,
   questions.answer = responses.text is_correct
from responses
 join questions on responses.question_id = questions.id
 where attempt_id = @attempt_id
order by responses.created_at;
