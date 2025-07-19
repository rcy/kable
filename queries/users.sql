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

-- name: UpdateUserAvatar :one
update users set avatar = @avatar where id = @id returning *;

-- name: UsersWithUnreadCounts :many
select users.*, count(*) unread_count
from deliveries
join users on sender_id = users.id
where recipient_id = @recipient_id and sent_at is null
group by users.id;

-- name: UserByUsername :one
select * from users where username = @username;

-- name: CreateUser :one
insert into users(username) values(@username) returning *;

-- name: AllUsers :many
select * from users order by created_at desc;

-- name: ParentsByKidID :many
select users.* from kids_parents join users on kids_parents.parent_id = users.id where kids_parents.kid_id = @kid_id;

-- name: KidsByParentID :many
select users.* from kids_parents join users on kids_parents.kid_id = users.id where kids_parents.parent_id = @parent_id order by kids_parents.created_at desc;

-- name: GetConnection :one
select sqlc.embed(u),
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
select sqlc.embed(u),
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
  is_parent = true
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

-- name: GetFamily :many
select u.*
from users u
join friends f1 on f1.b_id = u.id and f1.a_id = @a_id
join friends f2 on f2.a_id = u.id and f2.b_id = @a_id
where f1.b_role <> 'friend';

-- name: GetKids :many
select u.* from users u
join friends f1 on f1.b_id = u.id and f1.a_id = @a_id and f1.b_role = 'child'
join friends f2 on f2.a_id = u.id and f2.b_id = @a_id and f2.b_role = 'parent';

-- name: GetParents :many
select u.* from users u
join friends f1 on f1.b_id = u.id and f1.a_id = @a_id and f1.b_role = 'parent'
join friends f2 on f2.a_id = u.id and f2.b_id = @a_id and f2.b_role = 'child';

-- name: UpdateUserGradient :one
update users set gradient = $1 where id = @user_id returning *;
