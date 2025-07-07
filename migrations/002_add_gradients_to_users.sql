alter table users add column gradient jsonb not null default '{}'::jsonb;

update users
set gradient = g.gradient
from (
  select distinct on (user_id) user_id, gradient
  from gradients
  order by user_id, created_at desc
) as g
where users.id = g.user_id;

---- create above / drop below ----

alter table users drop column gradient;
