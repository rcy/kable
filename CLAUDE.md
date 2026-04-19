# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Commands

```bash
make start        # Hot-reload dev server (air)
make build        # Build binary
make test         # Run tests with .env.test
make generate     # Run sqlc to regenerate database models/queries
make compose-up   # Start PostgreSQL (port 2017) and MinIO (ports 9010-9011)
make migrate      # Run pending database migrations
make reset        # Drop, recreate, and migrate database
make deploy       # Deploy to Fly.io
```

Single test: `go test ./handlers/chat/... -run TestSomething`

## Architecture

**Kable** is a Go web app (internal package name: `oj`) for social connections, messaging, quizzes, chess, and postcards. Chi v5 router, PostgreSQL with sqlc-generated queries, server-side rendering via Gomponents (Go component library — not traditional HTML templates).

### Request Flow

```
main.go → handlers/router.go → handlers/<feature>/ → services/ → api/ (sqlc)
```

- `handlers/` — HTTP route handlers, organized by feature (chat, connect, me, u, postoffice, fun/, admin/, etc.)
- `services/` — Business logic (email, family relationships, reachability, room, storage)
- `api/` — sqlc-generated type-safe database layer; **do not edit `*.sql.go` files directly**
- `queries/` — SQL query definitions that feed sqlc; edit here, then run `make generate`
- `worker/` — Background jobs via neoq (delivery/friend notifications)
- `internal/` — Shared packages: config, middleware, AI integration, text processing

### Database

- Migrations live in `migrations/` (tern v2 format, embedded in binary)
- Schema defined in `schema.sql`
- Add queries in `queries/*.sql`, then run `make generate` to update `api/`
- PostgreSQL connection via `pg-service.conf` service definitions

### Templating

Uses [Gomponents](https://www.gomponents.com/) — Go functions that return HTML nodes, not `.gohtml` files. Look for `g.El`, `html.Div`, `c.Classes` patterns. Shared layout in `handlers/layout/`.

### Environment

- `.env` — local dev vars
- `.env.test` — test vars (loaded automatically by `make test`)
- Nix shell (`nix-shell shell.nix`) provides Go, air, golangci-lint, gopls, PostgreSQL, flyctl
