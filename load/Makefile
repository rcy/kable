load-schema: oj-fixed.db
	pgloader ./schema.load

load-data: oj-fixed.db
	pgloader ./data.load

oj-fixed.db: dump.sql
	sqlite3 $@ < $<

dump.raw.sql: oj_production.db
	sqlite3 $< .dump > $@

dump.timestamps.sql: dump.raw.sql
	sed -e "s/default current_timestamp//g" < $< > $@

dump.references.sql: dump.timestamps.sql
	sed -E "s/\breferences ([a-zA-Z_][a-zA-Z0-9_]+)([, ])/references \1(id)\2/g" < $< > $@ 

dump.idtypes.sql: dump.references.sql
	sed -E "s/id references/id integer references/g" < $< > $@ 

dump.idtypes1.sql: dump.idtypes.sql
	sed -E "s/room_id text/room_id integer/g" < $< > $@ 

dump.sql: dump.idtypes1.sql
	cp $< $@

oj_production.db:
	fly -a octopusjr ssh sftp get /data/$@
	fly -a octopusjr ssh sftp get /data/$@-shm
	fly -a octopusjr ssh sftp get /data/$@-wal
	sqlite3 $@ "PRAGMA wal_checkpoint(FULL);"

clean:
	rm -f dump.sql dump.*.sql oj-fixed.db oj_production.db
