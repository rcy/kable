SHELL=/bin/bash -o pipefail

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

drop:
	cat etc/drop-database.sql | psql service=admin -v ON_ERROR_STOP=1

create:
	cat etc/docker-entrypoint-initdb.d/*.sql | psql service=admin -v ON_ERROR_STOP=1

reset: drop create migrate

deploy:
	flyctl deploy

generate:
	go tool github.com/sqlc-dev/sqlc/cmd/sqlc generate

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
