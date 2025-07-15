-- name: CreateBot :one
insert into bots(owner_id, assistant_id, name, description) values(@owner_id,@assistant_id,@name,@description) returning *;

-- name: AllBots :many
select sqlc.embed(bots), sqlc.embed(users) from bots join users on bots.owner_id = users.id order by bots.created_at desc;

-- name: PublishedBots :many
select * from bots where published = 1;

-- name: UserVisibleBots :many
select * from bots where owner_id = @owner_id or published = 1;

-- name: Bot :one
select * from bots where id = @id;

-- name: UpdateBotDescription :one
update bots set description = @description, name = @name where id = @id and owner_id = @owner_id returning *;

