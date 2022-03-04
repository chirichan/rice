package rice

import (
	"context"
	"database/sql"
	"sync"
	"time"

	_ "github.com/lib/pq"
)

const (
	_defaultMaxPoolSize  = 1
	_defaultConnAttempts = 2
	_defaultConnTimeout  = time.Second
)

var (
	pgOnce   sync.Once
	postgres = &Postgres{}
)

type Postgres struct {
	*sql.DB
}

func (db *Postgres) Close() {
	if db.DB != nil {
		db.DB.Close()
	}
}

func NewPostgres(url string, opts ...Option) (*Postgres, error) {

	var err error

	pgOnce.Do(func() {
		postgres.DB, err = sql.Open("postgres", url)
		if err != nil {
			return
		}

		for _, opt := range opts {
			opt(postgres.DB)
		}

		err = postgres.Ping()
	})

	return postgres, err
}

func NewPostgresDB() PrettyDB { return postgres }

func NewPostgresTx() (PrettyTx, error) { return postgres.Begin() }

func NewPostgresTxContext(ctx context.Context) (PrettyTx, error) { return postgres.BeginTx(ctx, nil) }
