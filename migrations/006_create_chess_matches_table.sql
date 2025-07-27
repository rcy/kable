create table chess_matches (
    id bigserial primary key,
    created_at timestamp with time zone default now() not null,
    updated_at timestamp with time zone default now() not null,
    white_user_id bigint references users not null,
    black_user_id bigint references users not null,
    pgn text not null
    --turn text not null default 'white',
    --outcome text not null default '*'
);

---- create above / drop below ----

drop table chess_matches;
