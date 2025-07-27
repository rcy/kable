-- name: CreateChessMatch :one
insert into chess_matches(white_user_id, black_user_id, pgn)
values (@white_user_id, @black_user_id, @pgn)
returning *;

-- name: ChessMatchesBetweenUsers :many
select * from chess_matches
where
  (white_user_id = @user1_id and black_user_id = @user2_id)
or
  (white_user_id = @user2_id and black_user_id = @user1_id)
order by updated_at desc;

-- name: ChessMatchByID :one
select * from chess_matches where id = @id;

-- name: UpdateChessMatchPGN :one
update chess_matches set pgn = @pgn where id = @id returning *;
