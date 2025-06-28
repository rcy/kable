.DEFAULT_GOAL:=all

DATABASE_URL?=postgres://appuser:appuser@localhost:2017/kable_development

oj_production.db:
	fly -a octopusjr ssh sftp get /data/$@
	fly -a octopusjr ssh sftp get /data/$@-shm
	fly -a octopusjr ssh sftp get /data/$@-wal
	sqlite3 $@ "PRAGMA wal_checkpoint(FULL);"

TABLES := users \
quizzes \
attempts \
bots \
codes \
rooms \
messages \
deliveries \
friends \
gradients \
images \
kids_codes \
kids_parents \
notes \
postcards \
questions \
responses \
rooms \
sessions \
threads

all: $(patsubst %,%.copy,$(TABLES))

%.csv: oj_production.db
	sqlite3 $< -header -csv "SELECT * FROM $*;" > $@

.PHONY: %.copy
%.copy: %.csv
	echo "\copy $* FROM '$<' WITH (FORMAT csv, HEADER true);" | psql ${DATABASE_URL} -v ON_ERROR_STOP=1
	echo "SELECT setval(pg_get_serial_sequence('$*', 'id'), (SELECT MAX(id) FROM $*));" | psql ${DATABASE_URL} -v ON_ERROR_STOP=1

clean:
	rm -rf *.csv

