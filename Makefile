-include .env

export PGSERVICE?=local

start: clean-temp
	air

clean-temp:
	rm -rf ${TEMP}/go-build*

build:
	go build .

sql:
	psql service=local

sql.prod:
	psql service=prod

version.%:
	echo "update migration_version set version = $*" | sqlite3 ${SQLITE_DB}

deploy:
	flyctl deploy

drop:
	-rm ${SQLITE_DB}{,-shm,-wal}

generate:
	go tool github.com/sqlc-dev/sqlc/cmd/sqlc generate

getproddb:
	fly ssh sftp get /data/oj_production.db
	fly ssh sftp get /data/oj_production.db-shm
	fly ssh sftp get /data/oj_production.db-wal

test:
	. ./.env.test && go test ./...

migrate:
	cd migrations && tern migrate && tern status

migrate.%:
	cd migrations && tern migrate -d $* && tern status

compose-up:
	docker compose up -d
compose-stop:
	docker compose stop
compose-down:
	docker compose down
