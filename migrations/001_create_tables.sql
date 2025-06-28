CREATE TABLE users (
    id bigserial primary key,
    created_at timestamp with time zone default now() not null,
    username text not null,
    email text,
    avatar_url text DEFAULT 'https://www.gravatar.com/avatar/?d=mp'::text not null,
    is_parent boolean DEFAULT false not null,
    bio text DEFAULT ''::text not null,
    become_user_id bigint references users(id),
    admin boolean DEFAULT false not null
);

CREATE TABLE quizzes (
    id bigserial primary key,
    created_at timestamp with time zone default now() not null,
    name text not null,
    description text not null,
    published boolean DEFAULT false not null
);

CREATE TABLE attempts (
    id bigserial primary key,
    created_at timestamp with time zone default now(),
    quiz_id bigint not null references quizzes(id),
    user_id bigint not null references users(id)
);

CREATE TABLE bots (
    id bigserial primary key,
    created_at timestamp with time zone default now() not null,
    owner_id bigint not null references users(id),
    assistant_id text not null,
    name text not null,
    description text not null,
    published boolean DEFAULT false not null
);

CREATE TABLE codes (
    id bigserial primary key,
    created_at timestamp with time zone default now() not null,
    code text not null,
    nonce text not null,
    email text not null
);

CREATE TABLE rooms (
    id bigserial primary key,
    created_at timestamp with time zone default now() not null,
    key text not null
);


CREATE TABLE messages (
    id bigserial primary key,
    created_at timestamp with time zone default now() not null,
    sender_id bigint not null references users(id),
    room_id bigint not null references rooms(id),
    body text not null
);

CREATE TABLE deliveries (
    id bigserial primary key,
    created_at timestamp with time zone default now() not null,
    message_id bigint not null references messages(id),
    room_id bigint not null references rooms(id),
    recipient_id bigint not null references users(id),
    sender_id bigint not null references users(id),
    sent_at timestamp with time zone
);

CREATE TABLE friends (
    id bigserial primary key,
    created_at timestamp with time zone default now() not null,
    a_id bigint not null references users(id),
    b_id bigint not null references users(id),
    b_role text not null
);

CREATE TABLE gradients (
    id bigserial primary key,
    created_at timestamp with time zone default now() not null,
    user_id bigint not null references users(id),
    gradient jsonb not null
);

CREATE TABLE images (
    id bigserial primary key,
    created_at timestamp with time zone default now() not null,
    url text not null,
    user_id bigint not null references users(id)
);


CREATE TABLE kids_codes (
    id bigserial primary key,
    created_at timestamp with time zone default now() not null,
    code text not null,
    nonce text not null,
    user_id bigint not null references users(id)
);

CREATE TABLE kids_parents (
    id bigserial primary key,
    created_at timestamp with time zone default now() not null,
    kid_id bigint not null references users(id),
    parent_id bigint not null references users(id)
);


CREATE TABLE notes (
    id bigserial primary key,
    created_at timestamp with time zone default now() not null,
    owner_id bigint not null references users(id),
    body text not null,
    published boolean DEFAULT false not null
);

CREATE TABLE postcards (
    id bigserial primary key,
    created_at timestamp with time zone default now() not null,
    sender bigint not null references users(id),
    recipient bigint not null references users(id),
    subject text not null,
    body text not null,
    state text DEFAULT 'draft'::text not null
);

CREATE TABLE questions (
    id bigserial primary key,
    created_at timestamp with time zone default now() not null,
    quiz_id bigint not null references quizzes(id),
    text text not null,
    answer text not null
);

CREATE TABLE responses (
    id bigserial primary key,
    created_at timestamp with time zone default now() not null,
    quiz_id bigint not null references quizzes(id),
    user_id bigint not null references users(id),
    attempt_id bigint not null references attempts(id),
    question_id bigint not null references questions(id),
    text text not null
);


CREATE TABLE room_users (
    id bigserial primary key,
    created_at timestamp with time zone default now() not null,
    room_id bigint not null references rooms(id),
    user_id bigint not null references users(id)
);



CREATE TABLE sessions (
    id bigserial primary key,
    user_id bigint not null references users(id),
    key text not null
--    created_at timestamp with time zone default now() not null
);


CREATE TABLE threads (
    id bigserial primary key,
    created_at timestamp with time zone default now() not null,
    thread_id text not null,
    assistant_id text not null,
    user_id bigint not null references users(id)
);

---- create above / drop below ----

drop TABLE threads;
drop TABLE sessions;
drop TABLE room_users;
drop TABLE responses;
drop TABLE questions;
drop TABLE postcards;
drop TABLE notes;
drop TABLE kids_parents;
drop TABLE kids_codes;
drop TABLE images;
drop TABLE gradients;
drop TABLE friends;
drop TABLE deliveries;
drop TABLE messages;
drop TABLE rooms;
drop TABLE codes;
drop TABLE bots;
drop TABLE attempts;
drop TABLE quizzes;
drop TABLE users;
