alter table users add column avatar jsonb not null default '{}'::jsonb;
alter table users rename column avatar_url to avatar_url_deprecated;
---- create above / drop below ----

alter table users rename column avatar_url_deprecated to avatar_url;
alter table users drop column avatar;
