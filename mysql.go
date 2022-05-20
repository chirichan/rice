package rice

import (
	"context"
	"database/sql"
	"sync"

	_ "github.com/go-sql-driver/mysql"
)

// const (
// 	_defaultConnMaxIdleTime = 10 * time.Second // 一个连接空闲的最大时长
// 	_defaultConnMaxLifetime = 10 * time.Second // 一个连接使用的最大时长
// 	_defaultMaxIdleConns    = 2                // 最大闲置的连接数
// 	_defaultMaxOpenConns    = 4                // 最大打开的连接数
// )

var (
	mariaOnce sync.Once
	mariadb   = &MariaDB{}
)

type MariaDB struct {
	// connAttempts int
	// connTimeout  time.Duration
	*sql.DB
}

func (db *MariaDB) Close() {
	if db.DB != nil {
		db.DB.Close()
	}
}

func NewMaria(dsn string, opts ...Option) (*MariaDB, error) {

	var err error

	mariaOnce.Do(func() {

		mariadb.DB, err = sql.Open("mysql", dsn)
		if err != nil {
			return
		}

		for _, opt := range opts {
			opt(mariadb.DB)
		}

		err = mariadb.Ping()
	})

	return mariadb, err
}

func NewMariaDB() PrettyDB { return mariadb }

func NewMariaTx() (PrettyTx, error) { return mariadb.Begin() }

func NewMariaTxContext(ctx context.Context) (PrettyTx, error) { return mariadb.BeginTx(ctx, nil) }
