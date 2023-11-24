-- name: UserBySessionKey :one
select users.* from sessions join users on sessions.user_id = users.id where sessions.key = ?;

-- name: UserByID :one
select * from users where id = ?;

-- name: CreateParent :one
insert into users(email, username, is_parent) values(?, ?, true) returning *;

-- name: UserByEmail :one
select * from users where email = ?;

-- name: RecentMessages :many
select * from (
  select m.*, sender.avatar_url as sender_avatar_url
   from messages m
   join users sender on m.sender_id = sender.id
   where m.room_id = ?
   order by m.created_at desc
   limit 128
  ) t
order by created_at asc;

-- name: UsersWithUnreadCounts :many
select users.*, count(*) unread_count
from deliveries
join users on sender_id = users.id
where recipient_id = ? and sent_at is null
group by users.username;


-- name: GetAttemptByID :one
select * from attempts where id = ?;

-- name: AttemptNextQuestion :one
select questions.* from questions
left join responses on responses.question_id = questions.id
where
  questions.id not in (select question_id from responses where responses.attempt_id = ?)
and
  questions.quiz_id = ?
order by random();

-- name: SetQuizPublished :one
update quizzes set published = ? where id = ? returning *;

-- name: QuestionCount :one
select count(*) from questions where quiz_id = ?;

-- name: ResponseCount :one
select count(*) from responses where attempt_id = ?;

-- name: UserByUsername :one
select * from users where username = ?;

-- name: CreateUser :one
insert into users(username) values(?) returning *;

-- name: CreateKidParent :one
insert into kids_parents(kid_id, parent_id) values(?, ?) returning *;

-- name: CreateFriend :one
insert into friends(a_id, b_id, b_role) values(?, ?, ?) returning *;

-- name: AllUsers :many
select * from users order by created_at desc;

-- name: ParentsByKidID :many
select users.* from kids_parents join users on kids_parents.parent_id = users.id where kids_parents.kid_id = ?;

-- name: KidsByParentID :many
select users.* from kids_parents join users on kids_parents.kid_id = users.id where kids_parents.parent_id = ? order by kids_parents.created_at desc;

-- name: AttemptResponseIDs :many
select id from responses where attempt_id = ?;

-- name: CreateResponse :one
insert into responses(quiz_id, user_id, attempt_id, question_id, text) values(?,?,?,?,?) returning *;

-- name: CreateAttempt :one
insert into attempts(quiz_id, user_id) values(?,?) returning *;

-- name: Delivery :one
select * from deliveries where id = ?;

-- name: UserGradient :one
select * from gradients where user_id = ? order by created_at desc limit 1;

-- name: Question :one
select * from questions where id = ?;

-- name: CreateQuestion :one
insert into questions(quiz_id, text, answer) values(?,?,?) returning *;

-- name: UpdateQuestion :one
update questions set text = ?, answer = ? where id = ? returning *;

-- name: QuizQuestions :many
select * from questions where quiz_id = ?;

-- name: AllQuizzes :many
select * from quizzes order by created_at desc;

-- name: PublishedQuizzes :many
select * from quizzes where published = true order by created_at desc;

-- name: Quiz :one
select * from quizzes where id = ?;

-- name: UpdateQuiz :one
update quizzes set name = ?, description = ? where id = ? returning *;

-- name: CreateQuiz :one
insert into quizzes(name,description) values(?,?) returning *;

-- name: Responses :many
select
   responses.*,
   questions.answer question_answer,
   questions.text question_text,
   questions.answer = responses.text is_correct
from responses
 join questions on responses.question_id = questions.id
 where attempt_id = ?
order by responses.created_at;

-- name: RoomByKey :one
select * from rooms where key = ?;

-- name: CreateRoom :one
insert into rooms(key) values(?) returning *;

-- name: CreateRoomUser :one
insert into room_users(room_id, user_id) values(?, ?) returning *;
