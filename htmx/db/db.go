package db

import (
	"log"
	"os"

	"github.com/BurntSushi/migration"
	"github.com/jmoiron/sqlx"
	_ "modernc.org/sqlite"
)

var DB *sqlx.DB

func init() {
	var err error

	dbfile, ok := os.LookupEnv("SQLITE_DB")
	if !ok {
		panic("SQLITE_DB not defined")
	}

	// migrate the database
	mDB, err := migration.Open("sqlite", dbfile, migrations())
	if err != nil {
		log.Fatalf("MIGRATION: %v", err)
	}

	// create sqlx connection from migrated database
	DB = sqlx.NewDb(mDB, "sqlite")

	// dump schema to schema.sql -- TODO: don't do this if database didn't change (compare migration_versions before and after)
	MustDump()

	log.Printf("Connected to database: %s", dbfile)
}

func migrations() (ms []migration.Migrator) {
	for i, sql := range Migrations {
		i, sql := i, sql
		ms = append(ms, func(tx migration.LimitedTx) error {
			log.Printf("MIGRATION: %d %s", i, sql)
			_, err := tx.Exec(sql)
			return err
		})
	}

	return
}
