package rice

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/jackc/pgx/v4/pgxpool"
)

const (
	_defaultMaxPoolSize  = 1
	_defaultConnAttempts = 2
	_defaultConnTimeout  = time.Second
)

// Postgres -.
type Postgres struct {
	maxPoolSize  int
	connAttempts int
	connTimeout  time.Duration

	Builder squirrel.StatementBuilderType
	Pool    *pgxpool.Pool
}

var (
	pg *Postgres
)

// NewPostgresDB -.
func NewPostgresDB(url string, opts ...PostgresOption) (*Postgres, error) {

	if pg == nil {

		pg = &Postgres{
			maxPoolSize:  _defaultMaxPoolSize,
			connAttempts: _defaultConnAttempts,
			connTimeout:  _defaultConnTimeout,
		}

		// Custom options
		for _, opt := range opts {
			opt(pg)
		}

		pg.Builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)

		poolConfig, err := pgxpool.ParseConfig(url)
		if err != nil {
			return nil, fmt.Errorf("postgres - NewPostgres - pgxpool.ParseConfig: %w", err)
		}

		poolConfig.MaxConns = int32(pg.maxPoolSize)

		for pg.connAttempts > 0 {
			pg.Pool, err = pgxpool.ConnectConfig(context.Background(), poolConfig)
			if err == nil {
				break
			}

			log.Printf("Postgres is trying to connect, attempts left: %d", pg.connAttempts)

			time.Sleep(pg.connTimeout)

			pg.connAttempts--
		}

		if err != nil {
			return nil, fmt.Errorf("postgres - NewPostgres - connAttempts == 0: %w", err)
		}

		return pg, nil
	} else {
		return pg, nil
	}
}

func GetPostgresDB() *Postgres {
	return pg
}

// Close -.
func (p *Postgres) Close() {
	if p.Pool != nil {
		p.Pool.Close()
	}
}

// Option -.
type PostgresOption func(*Postgres)

// MaxPoolSize -.
func MaxPoolSize(size int) PostgresOption {
	return func(c *Postgres) {
		c.maxPoolSize = size
	}
}

// ConnAttempts -.
func ConnAttempts(attempts int) PostgresOption {
	return func(c *Postgres) {
		c.connAttempts = attempts
	}
}

// ConnTimeout -.
func ConnTimeout(timeout time.Duration) PostgresOption {
	return func(c *Postgres) {
		c.connTimeout = timeout
	}
}
