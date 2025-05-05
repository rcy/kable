-- name: UserBySessionKey :one
select users.* from sessions join users on sessions.user_id = users.id where sessions.key = @key;

-- name: UserByID :one
select * from users where id = @id;

-- name: CreateParent :one
insert into users(email, username, is_parent) values(@email, @username, true) returning *;

-- name: ParentByID :one
select * from users where id = @id and is_parent = true;

-- name: UserByEmail :one
select * from users where email = @email;

-- name: UpdateAvatar :one
update users set avatar_url = @avatar_url where id = @id returning *;

-- name: RecentRoomMessages :many
select * from (
  select m.*, sender.avatar_url as sender_avatar_url
   from messages m
   join users sender on m.sender_id = sender.id
   where m.room_id = @room_id
   order by m.created_at desc
   limit 128
  ) t
order by created_at asc;

-- name: MessageByID :one
select * from messages where id = @id;

-- name: AdminRecentMessages :many
select
        m.*,
        sender.username as sender_username,
        sender.avatar_url as sender_avatar_url
 from messages m
 join users sender on m.sender_id = sender.id
 order by m.created_at desc
 limit 128;

-- name: AdminDeleteMessage :one
update messages
set body = '+++ deleted +++'
where id = @id
returning *;

-- name: UsersWithUnreadCounts :many
select users.*, count(*) unread_count
from deliveries
join users on sender_id = users.id
where recipient_id = @recipient_id and sent_at is null
group by users.username;


-- name: GetAttemptByID :one
select * from attempts where id = @id;

-- name: AttemptNextQuestion :one
select questions.* from questions
left join responses on responses.question_id = questions.id
where
  questions.id not in (select question_id from responses where responses.attempt_id = @attempt_id)
and
  questions.quiz_id = @quiz_id
order by random();

-- name: SetQuizPublished :one
update quizzes set published = @published where id = @id returning *;

-- name: QuestionCount :one
select count(*) from questions where quiz_id = @quiz_id;

-- name: ResponseCount :one
select count(*) from responses where attempt_id = @attempt_id;

-- name: UserByUsername :one
select * from users where username = @username;

-- name: CreateUser :one
insert into users(username) values(@username) returning *;

-- name: CreateKidParent :one
insert into kids_parents(kid_id, parent_id) values(@kid_id, @parent_id) returning *;

-- name: CreateFriend :one
insert into friends(a_id, b_id, b_role) values(@a_id, @b_id, @b_role) returning *;

-- name: AllUsers :many
select * from users order by created_at desc;

-- name: ParentsByKidID :many
select users.* from kids_parents join users on kids_parents.parent_id = users.id where kids_parents.kid_id = @kid_id;

-- name: KidsByParentID :many
select users.* from kids_parents join users on kids_parents.kid_id = users.id where kids_parents.parent_id = @parent_id order by kids_parents.created_at desc;

-- name: AttemptResponseIDs :many
select id from responses where attempt_id = @attempt_id;

-- name: CreateResponse :one
insert into responses(quiz_id, user_id, attempt_id, question_id, text) values(@quiz_id,@user_id,@attempt_id,@question_id,@text) returning *;

-- name: CreateAttempt :one
insert into attempts(quiz_id, user_id) values(@quiz_id,@user_id) returning *;

-- name: Delivery :one
select * from deliveries where id = @id;

-- name: UserGradient :one
select * from gradients where user_id = @user_id order by created_at desc limit 1;

-- name: Question :one
select * from questions where id = @id;

-- name: CreateQuestion :one
insert into questions(quiz_id, text, answer) values(@quiz_id,@text,@answer) returning *;

-- name: UpdateQuestion :one
update questions set text = @text, answer = @answer where id = @id returning *;

-- name: QuizQuestions :many
select * from questions where quiz_id = @quiz_id;

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

-- name: RoomByKey :one
select * from rooms where key = @key;

-- name: CreateRoom :one
insert into rooms(key) values(@key) returning *;

-- name: CreateRoomUser :one
insert into room_users(room_id, user_id) values(@room_id, @user_id) returning *;

-- name: GetConnection :one
select u.*,
       case
           when f1.a_id = @a_id then f1.b_role
           else ''
       end as role_out,
       case
           when f2.b_id = @a_id then f2.b_role
           else ''
       end as role_in
from users u
left join friends f1 on f1.b_id = u.id and f1.a_id = @a_id
left join friends f2 on f2.a_id = u.id and f2.b_id = @a_id
where
  u.id = @id;

-- name: GetCurrentAndPotentialParentConnections :many
select u.*,
       case
           when f1.a_id = @a_id then f1.b_role
           else ''
       end as role_out,
       case
           when f2.b_id = @a_id then f2.b_role
           else ''
       end as role_in
from users u
left join friends f1 on f1.b_id = u.id and f1.a_id = @a_id
left join friends f2 on f2.a_id = u.id and f2.b_id = @a_id
where
  u.id != @a_id
and
  is_parent = 1
order by role_in desc
limit 128;

-- name: GetFriends :many
select u.* from users u
join friends f1 on f1.b_id = u.id and f1.a_id = @a_id and f1.b_role = 'friend'
join friends f2 on f2.a_id = u.id and f2.b_id = @a_id and f2.b_role = 'friend';

-- name: GetConnections :many
select u.* from users u
join friends f1 on f1.b_id = u.id and f1.a_id = @a_id
join friends f2 on f2.a_id = u.id and f2.b_id = @a_id
where f1.b_role <> '' and f2.b_role <> '';

-- name: GetFriendsWithGradient :many
select u.*, g.gradient, max(g.created_at)
from users u
join friends f1 on f1.b_id = u.id and f1.a_id = @a_id
join friends f2 on f2.a_id = u.id and f2.b_id = @a_id
left outer join gradients g
on g.user_id = f1.b_id
where f1.b_role = 'friend'
group by u.id;

-- name: GetFamilyWithGradient :many
select u.*, g.gradient, max(g.created_at)
from users u
join friends f1 on f1.b_id = u.id and f1.a_id = @a_id
join friends f2 on f2.a_id = u.id and f2.b_id = @a_id
left outer join gradients g
on g.user_id = f1.b_id
where f1.b_role <> 'friend'
group by u.id;

-- name: GetConnectionsWithGradient :many
select u.*, g.gradient, max(g.created_at)
from users u
join friends f1 on f1.b_id = u.id and f1.a_id = @a_id
join friends f2 on f2.a_id = u.id and f2.b_id = @a_id
left outer join gradients g
on g.user_id = f1.b_id
group by u.id;

-- name: GetKids :many
select u.* from users u
join friends f1 on f1.b_id = u.id and f1.a_id = @a_id and f1.b_role = 'child'
join friends f2 on f2.a_id = u.id and f2.b_id = @a_id and f2.b_role = 'parent';

-- name: GetParents :many
select u.* from users u
join friends f1 on f1.b_id = u.id and f1.a_id = @a_id and f1.b_role = 'parent'
join friends f2 on f2.a_id = u.id and f2.b_id = @a_id and f2.b_role = 'child';

-- name: UserPostcardsSent :many
select p.*, r.username, r.avatar_url
from postcards p
join users r on p.recipient = r.id
where sender = @sender
order by p.created_at desc;

-- name: UserPostcardsReceived :many
select p.*, s.username, s.avatar_url
from postcards p
join users s on p.sender = s.id
where recipient = @recipient
order by p.created_at desc;

-- name: CreatePostcard :one
insert into postcards(sender, recipient, subject, body, state) values(@sender,@recipient,@subject,@body,@state) returning *;

-- name: CreateThread :one
insert into threads(user_id, thread_id, assistant_id) values(@user_id,@thread_id,@assistant_id) returning *;

-- name: AssistantThreads :many
select * from threads where assistant_id = @assistant_id and user_id = @user_id;

-- name: UserThreadByID :one
select * from threads where user_id = @user_id and thread_id = @thread_id;

-- name: CreateBot :one
insert into bots(owner_id, assistant_id, name, description) values(@owner_id,@assistant_id,@name,@description) returning *;

-- name: AllBots :many
select sqlc.embed(bots), sqlc.embed(users) from bots join users on bots.owner_id = users.id order by bots.created_at desc;

-- name: PublishedBots :many
select * from bots where published = 1;

-- name: UserVisibleBots :many
select * from bots where owner_id = @owner_id or published = 1;

-- name: Bot :one
select * from bots where id = @id;

-- name: UpdateBotDescription :one
update bots set description = @description, name = @name where id = @id and owner_id = @owner_id returning *;

-- name: PublishedNotes :many
select * from notes where published = 1;

-- name: UserNotes :many
select * from notes where owner_id = @owner_id order by created_at desc;

-- name: CreateNote :one
insert into notes(owner_id,body) values(@owner_id,@body) returning *;

-- name: UpdateNote :one
update notes set body = @body where id = @id and owner_id = @owner_id returning *;

-- name: DeleteNote :exec
delete from notes where id = @id and owner_id = @owner_id;

