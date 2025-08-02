alter table room_users add constraint unique_room_user unique (room_id, user_id);
---- create above / drop below ----
alter table room_users drop constraint unique_room_user;
