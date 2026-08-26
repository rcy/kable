alter table bots drop column assistant_id;

delete from threads;
alter table threads drop column thread_id;
alter table threads drop column assistant_id;
alter table threads add column bot_id bigint not null references bots(id);

create table bot_messages (
    id bigserial primary key,
    created_at timestamp with time zone default now() not null,
    thread_id bigint not null references threads(id),
    role text not null,
    content text not null
);
