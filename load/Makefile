load: oj-fixed.sqlite
	pgloader ./db.load

oj-fixed.sqlite: dump.sql
	sqlite3 $@ < $<

dump.raw.sql: ./oj.sqlite
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

clean:
	rm -f dump.sql dump.*.sql oj-fixed.sqlite
