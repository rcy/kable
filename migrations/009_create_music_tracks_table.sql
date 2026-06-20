create table music_tracks (
    id bigserial primary key,
    created_at timestamp with time zone default now() not null,
    user_id bigint not null references users(id),
    url text not null,
    title text,
    uploader text,
    filename text not null,
    status text not null default 'done',
    error text
);

---- create above / drop below ----

drop table music_tracks;
