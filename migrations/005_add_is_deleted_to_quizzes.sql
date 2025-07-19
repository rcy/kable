alter table quizzes add column is_deleted bool default false not null;
---- create above / drop below ----
alter table quizzes drop column is_deleted;
