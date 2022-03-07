package rice

import (
	"gorm.io/driver/mysql"
	pg "gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewMariaGorm(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	return db, err
}

func NewPostgresGorm(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(pg.Open(dsn), &gorm.Config{})
	return db, err
}
