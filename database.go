package rice

import (
	"context"
	"database/sql"
	"time"
)

var (
	_ Pretty   = &sql.DB{}
	_ Pretty   = &sql.Tx{}
	_ PrettyDB = &sql.DB{}
	_ PrettyTx = &sql.Tx{}
)

type Pretty interface {
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Prepare(query string) (*sql.Stmt, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	QueryRow(query string, args ...interface{}) *sql.Row
}

type PrettyDB interface {
	Pretty
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	Begin() (*sql.Tx, error)
}

type PrettyTx interface {
	Pretty
	Commit() error
	Rollback() error
	StmtContext(ctx context.Context, stmt *sql.Stmt) *sql.Stmt
	Stmt(stmt *sql.Stmt) *sql.Stmt
}

type Option func(*sql.DB)

// ConnMaxIdleTime 一个连接空闲的最大时长
func ConnMaxIdleTime(d time.Duration) Option {
	return func(c *sql.DB) {
		c.SetConnMaxIdleTime(d)
	}
}

// ConnMaxLifetime 一个连接使用的最大时长
func ConnMaxLifetime(d time.Duration) Option {
	return func(c *sql.DB) {
		c.SetConnMaxLifetime(d)
	}
}

// MaxIdleConns 最大闲置的连接数
func MaxIdleConns(size int) Option {
	return func(c *sql.DB) {
		c.SetMaxIdleConns(size)
	}
}

// MaxOpenConns 最大打开的连接数
func MaxOpenConns(size int) Option {
	return func(c *sql.DB) {
		c.SetMaxOpenConns(size)
	}
}
