package rice

import (
	"context"
	"database/sql"
)

var (
	_ Customer   = &sql.DB{}
	_ Customer   = &sql.Tx{}
	_ CustomerDB = &sql.DB{}
	_ CustomerTx = &sql.Tx{}
)

type Customer interface {
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Prepare(query string) (*sql.Stmt, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
	QueryRow(query string, args ...interface{}) *sql.Row
}

type CustomerDB interface {
	Customer
	BeginTx(ctx context.Context, opts *sql.TxOptions) (*sql.Tx, error)
	Begin() (*sql.Tx, error)
}

type CustomerTx interface {
	Customer
	Commit() error
	Rollback() error
	StmtContext(ctx context.Context, stmt *sql.Stmt) *sql.Stmt
	Stmt(stmt *sql.Stmt) *sql.Stmt
}

func NewCustomerDB() CustomerDB {
	return mariadb
}

func NewCustomerTx() CustomerTx {
	tx, _ := mariadb.Begin()
	return tx
}
