package database

import (
	"database/sql"
	"log"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

// const (
// 	_defaultConnMaxIdleTime = 10 * time.Second // 一个连接空闲的最大时长
// 	_defaultConnMaxLifetime = 10 * time.Second // 一个连接使用的最大时长
// 	_defaultMaxIdleConns    = 2                // 最大闲置的连接数
// 	_defaultMaxOpenConns    = 4                // 最大打开的连接数
// )

var (
	once    sync.Once
	mariadb *MariaDB
)

type MariaDB struct {
	connAttempts int
	connTimeout  time.Duration
	*sql.DB
}

func GetMariaDB() *MariaDB {
	return mariadb
}

func NewMariaDB(url string, opts ...MariaDBOption) *MariaDB {

	once.Do(func() {

		db, err := sql.Open("mysql", url)
		if err != nil {
			log.Fatalf("mariadb连接失败: %v\n", err)
		}

		mariadb = &MariaDB{DB: db}

		for _, opt := range opts {
			opt(mariadb)
		}
	})

	return mariadb
}

func (db *MariaDB) Close() {
	if db.DB != nil {
		db.DB.Close()
	}
}
