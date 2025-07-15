-- name: UserPostcardsSent :many
select p.*, r.username, r.avatar
from postcards p
join users r on p.recipient = r.id
where sender = @sender
order by p.created_at desc;

-- name: UserPostcardsReceived :many
select p.*, s.username, s.avatar
from postcards p
join users s on p.sender = s.id
where recipient = @recipient
order by p.created_at desc;

-- name: CreatePostcard :one
insert into postcards(sender, recipient, subject, body, state) values(@sender,@recipient,@subject,@body,@state) returning *;

