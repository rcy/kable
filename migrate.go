package main

import (
	"context"
	"embed"
	"fmt"
	"io/fs"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/jackc/tern/v2/migrate"
)

//go:embed migrations/*.sql
var migrationFiles embed.FS

func migrateSchema(ctx context.Context, pool *pgxpool.Pool) error {
	conn, err := pool.Acquire(ctx)
	if err != nil {
		return fmt.Errorf("Acquire: %w", err)
	}
	defer conn.Release()

	migrator, err := migrate.NewMigratorEx(
		context.Background(), conn.Conn(), "public.schema_version",
		&migrate.MigratorOptions{
			DisableTx: false,
		})
	if err != nil {
		return fmt.Errorf("NewMigratorEx: %w", err)
	}

	migrationRoot, _ := fs.Sub(migrationFiles, "migrations")
	err = migrator.LoadMigrations(migrationRoot)
	if err != nil {
		return fmt.Errorf("LoadMigrations: %w", err)
	}

	current, err := migrator.GetCurrentVersion(ctx)
	if err != nil {
		return fmt.Errorf("GetCurrentVersion: %w", err)
	}

	log.Printf("migrateSchema: %d/%d", current, len(migrator.Migrations))

	err = migrator.Migrate(ctx)
	if err != nil {
		return fmt.Errorf("Migrate: %w", err)
	}

	return nil
}
