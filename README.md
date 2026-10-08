# go-databases

Database helpers for Go projects: a `database/sql` connection service with transactions, pgx/pgxpool helpers, GORM constraint and join-table helpers, and MongoDB collection, index and transaction helpers. Requires Go 1.25 (per `go.mod`).

## Installation

```bash
go get github.com/ralvarezdev/go-databases
```

Main dependencies: `pgx/v5`, `mongo-driver`, `gorm` and `golang.org/x/net`.

## Packages

- **`godatabases`** (root) — shared errors such as `ErrNilConfig`, `ErrNilConnection`, `ErrConnectionFailed`, `ErrPingFailed`, `ErrNotConnected`.
- **`sql`** — `Config` / `NewConfig`, `Handler` and `Service` interfaces, `NewDefaultHandler`, `NewDefaultService`, `CreateTransaction`, `RunQueriesConcurrently`, `RunQueriesConcurrentlyWithCancel`.
- **`sql/pgx`** — `IsUniqueViolationError(err)`, which also returns the constraint name.
- **`sql/pgxpool`** — `CreateTransaction(ctx, pool, fn)`.
- **`sql/gorm`** — `NewModelConstraints`, `HasConstraint`, `CreateModelConstraints`, `CreateModelsConstraints`, `NewJoinField`, `SetupJoinTable`, `SetupJoinTables`.
- **`mongodb`** — `Collection` / `NewCollection` with index creation, `NewFieldIndex`, `NewUniqueIndex`, `NewTTLIndex`, `NewCompoundFieldIndex`, `GetObjectIDFromString`, `Prepare*Options` helpers, `CreateSession`, `CreateTransaction`.

## Usage

```go
import (
    "time"

    godbsql "github.com/ralvarezdev/go-databases/sql"
)

cfg, err := godbsql.NewConfig(
    "pgx", dsn,      // driver name, data source name
    10, 5,           // max open / max idle connections
    5*time.Minute,   // connection max idle time
    30*time.Minute,  // connection max lifetime
)
svc, err := godbsql.NewDefaultService(cfg)
_, err = svc.Connect()
defer svc.Disconnect()

q := "SELECT id FROM users WHERE email = $1"
row, err := svc.QueryRow(&q, "user@example.com")
var id int
err = svc.ScanRow(row, &id)
```

The caller imports the `database/sql` driver. `Service` also exposes `IsConnected`, `DB`, `CreateTransaction`, `Exec`/`ExecWithCtx` and `QueryRowWithCtx`.

## Development

```bash
go build ./...
go vet ./...
```

There are no tests.

## License

GNU General Public License v3.0 (see [LICENSE](LICENSE)).
