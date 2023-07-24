package postgres

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/willian2s/clean-go/config"

	_ "github.com/golang-migrate/migrate/v4/database/pgx" //driver pgx used to run migrations
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

// PoolInterface is an wraping to PgxPool to create test mocks
type PoolInterface interface {
	Close()
	Exec(ctx context.Context, sql string, arguments ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	QueryFunc(
		ctx context.Context,
		sql string,
		args []interface{},
		scans []interface{},
		f func(pgx.QueryFuncRow) error,
	) (pgconn.CommandTag, error)
	SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults
	Begin(ctx context.Context) (pgx.Tx, error)
	BeginFunc(ctx context.Context, f func(pgx.Tx) error) error
	BeginTxFunc(ctx context.Context, txOptions pgx.TxOptions, f func(pgx.Tx) error) error
}

// GetConnection returns a *pgxpool.Pool object that represents a connection to the database.
//
// It takes a context.Context object as a parameter.
// It returns a *pgxpool.Pool object.
func GetConnection(context context.Context) *pgxpool.Pool {
	databaseURL := config.EnvConfigs.DatabaseUrl

	conn, err := pgxpool.Connect(context, "postgres"+databaseURL)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	return conn
}

// RunMigrations runs the database migrations.
//
// It does not take any parameters.
// It does not return anything.
func RunMigrations() {
	databaseURL := config.EnvConfigs.DatabaseUrl
	m, err := migrate.New("file://database/migrations", "pgx"+databaseURL)
	if err != nil {
		log.Println(err)
	}

	if err := m.Up(); err != nil {
		log.Println(err)
	}
}
