package adapter

import (
	"context"
	_ "github.com/jackc/pgx"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"time"
)

type PostgresAdapter struct {
	timeoutDbContext time.Duration
	connection       *sqlx.DB
}

func NewPostgresAdapter(timeoutDbContextInSecond int, connectionString string) *PostgresAdapter {
	return &PostgresAdapter{
		connection:       connect(connectionString),
		timeoutDbContext: time.Duration(timeoutDbContextInSecond) * time.Second,
	}
}

func connect(connectionString string) *sqlx.DB {
	conn := sqlx.MustOpen("postgres", connectionString)
	return conn
}

func (p *PostgresAdapter) Close() error {
	return p.connection.Close()
}

func (p *PostgresAdapter) Execute(requestCtx context.Context, sql string, args ...interface{}) error {
	if p.connection == nil {
		return errors.New("[ PostgresAdapter ] Execute: connection is nil")
	}

	ctx, cancel := context.WithTimeout(requestCtx, p.timeoutDbContext)
	defer cancel()

	_, err := p.connection.ExecContext(ctx, sql, args...)
	return err
}

func (p *PostgresAdapter) ExecuteAndGet(requestCtx context.Context, destination interface{}, sql string, args ...interface{}) error {
	if p.connection == nil {
		return errors.New("[ PostgresAdapter ] ExecuteAndGet: connection is nil")
	}

	ctx, cancel := context.WithTimeout(requestCtx, p.timeoutDbContext)
	defer cancel()

	return p.connection.GetContext(ctx, destination, sql, args...)
}

func (p *PostgresAdapter) Query(requestCtx context.Context, destination interface{}, query string, args ...interface{}) error {
	if p.connection == nil {
		return errors.New("[ PostgresAdapter ] Query: connection is nil")
	}

	ctx, cancel := context.WithTimeout(requestCtx, p.timeoutDbContext)
	defer cancel()

	return p.connection.SelectContext(ctx, destination, query, args...)
}
