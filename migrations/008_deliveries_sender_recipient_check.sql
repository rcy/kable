alter table deliveries add check(sender_id <> recipient_id);
---- create above / drop below ----
alter table deliveries drop constraint deliveries_check;

