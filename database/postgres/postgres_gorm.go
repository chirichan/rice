package postgres

import (
	"database/sql"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Close(db *gorm.DB) {
	d, _ := db.DB()
	if d != nil {
		d.Close()
	}
}

// "host=localhost user=postgres password=111111 dbname=user port=5432 sslmode=disable TimeZone=Asia/Shanghai"
func NewGormDB(dsn string, opts ...GormOption) *gorm.DB {

	once.Do(func() {
		var err error
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Fatal(err)
		}

		sqlDB, err := db.DB()
		if err != nil {
			log.Fatal(err)
		}

		// SetMaxIdleConns 设置空闲连接池中连接的最大数量
		sqlDB.SetMaxIdleConns(10)

		// SetMaxOpenConns 设置打开数据库连接的最大数量。
		sqlDB.SetMaxOpenConns(100)

		// SetConnMaxLifetime 设置了连接可复用的最大时间。
		sqlDB.SetConnMaxLifetime(time.Hour)

		for _, opt := range opts {
			opt(sqlDB)
		}

		err = sqlDB.Ping()
		if err != nil {
			log.Print(err)
		}
		log.Println("🎉🎉🎉")
	})

	return db
}

type GormOption func(*sql.DB)

func MaxIdleConns(n int) GormOption {
	return func(db *sql.DB) {
		db.SetMaxIdleConns(n)
	}
}
func MaxOpenConns(n int) GormOption {
	return func(db *sql.DB) {
		db.SetMaxOpenConns(n)
	}
}
func ConnMaxLifetime(d time.Duration) GormOption {
	return func(db *sql.DB) {
		db.SetConnMaxLifetime(d)
	}
}
