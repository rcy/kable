alter table users add column avatar jsonb not null default '{}'::jsonb;

---- create above / drop below ----

alter table users drop column avatar;

