package rice

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

func (db *MariaDB) Close() {
	if db.DB != nil {
		db.DB.Close()
	}
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

func GetMariaDB() *MariaDB {
	return mariadb
}

type MariaDBOption func(*MariaDB)

// ConnMaxIdleTime_MYSQL 一个连接空闲的最大时长
func ConnMaxIdleTime_MYSQL(d time.Duration) MariaDBOption {
	return func(c *MariaDB) {
		c.DB.SetConnMaxIdleTime(d)
	}
}

// ConnMaxLifetime_MYSQL 一个连接使用的最大时长
func ConnMaxLifetime_MYSQL(d time.Duration) MariaDBOption {
	return func(c *MariaDB) {
		c.DB.SetConnMaxLifetime(d)
	}
}

// MaxIdleConns_MYSQL 最大闲置的连接数
func MaxIdleConns_MYSQL(size int) MariaDBOption {
	return func(c *MariaDB) {
		c.DB.SetMaxIdleConns(size)
	}
}

// MaxOpenConns_MYSQL 最大打开的连接数
func MaxOpenConns_MYSQL(size int) MariaDBOption {
	return func(c *MariaDB) {
		c.DB.SetMaxOpenConns(size)
	}
}

// ConnAttempts_MYSQL 尝试连接数据库次数
func ConnAttempts_MYSQL(attempts int) MariaDBOption {
	return func(c *MariaDB) {
		c.connAttempts = attempts
	}
}

// ConnTimeout_MYSQL 连接数据库超时时间
func ConnTimeout_MYSQL(timeout time.Duration) MariaDBOption {
	return func(c *MariaDB) {
		c.connTimeout = timeout
	}
}
