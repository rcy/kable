-include .env

start: clean-temp
	air

clean-temp:
	rm -rf ${TEMP}/go-build*

build:
	go build .

sql:
	sqlite3 ${SQLITE_DB}

psql:
	psql postgres://appuser:appuser@localhost:2017/kable_development

version.%:
	echo "update migration_version set version = $*" | sqlite3 ${SQLITE_DB}

deploy:
	flyctl deploy

drop:
	-rm ${SQLITE_DB}{,-shm,-wal}

seed: db/seed.sql
	sqlite3 ${SQLITE_DB} < $<

db/schema-fixed.sql: db/schema.sql
	sed -e 's/\"//g' $< > $@

db/pgschema.sql:
	docker exec kable-postgres-1 pg_dump --dbname=kable_development --schema public --user=postgres --schema-only > /tmp/schema
	mv /tmp/schema $@

generate: db/schema-fixed.sql db/pgschema.sql
	go tool github.com/sqlc-dev/sqlc/cmd/sqlc generate

getproddb:
	fly ssh sftp get /data/oj_production.db
	fly ssh sftp get /data/oj_production.db-shm
	fly ssh sftp get /data/oj_production.db-wal

test:
	. ./.env.test && go test ./...

migrate:
	cd migrations && tern migrate

migrate.%:
	cd migrations && tern migrate -d $*

compose-up:
	docker compose up -d
compose-stop:
	docker compose stop
compose-down:
	docker compose down
