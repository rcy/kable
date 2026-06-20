-- name: UserMusicTracks :many
select * from music_tracks where user_id = @user_id order by created_at desc;

-- name: MusicTrackByID :one
select * from music_tracks where id = @id;

-- name: InsertMusicTrack :one
insert into music_tracks(user_id, url, title, uploader, filename, status, error)
values(@user_id, @url, @title, @uploader, @filename, @status, @error)
returning *;

-- name: UpdateMusicTrack :exec
update music_tracks set title = @title, uploader = @uploader, filename = @new_filename, status = @status, error = @error
where user_id = @user_id and filename = @old_filename;

-- name: UpdateMusicTrackByID :exec
update music_tracks set title = @title, uploader = @uploader, status = @status, error = @error
where id = @id;
