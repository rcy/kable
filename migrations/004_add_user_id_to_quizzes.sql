alter table quizzes add column user_id bigint references users;
update quizzes set user_id = 1;
alter table quizzes alter column user_id set not null;

---- create above / drop below ----
alter table quizzes drop column user_id;
