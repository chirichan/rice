package database

import "time"

type MariaDBOption func(*MariaDB)

// 一个连接空闲的最大时长
func ConnMaxIdleTime_MYSQL(d time.Duration) MariaDBOption {
	return func(c *MariaDB) {
		c.DB.SetConnMaxIdleTime(d)
	}
}

// 一个连接使用的最大时长
func ConnMaxLifetime_MYSQL(d time.Duration) MariaDBOption {
	return func(c *MariaDB) {
		c.DB.SetConnMaxLifetime(d)
	}
}

// 最大闲置的连接数
func MaxIdleConns_MYSQL(size int) MariaDBOption {
	return func(c *MariaDB) {
		c.DB.SetMaxIdleConns(size)
	}
}

// 最大打开的连接数
func MaxOpenConns_MYSQL(size int) MariaDBOption {
	return func(c *MariaDB) {
		c.DB.SetMaxOpenConns(size)
	}
}

// 尝试连接数据库次数
func ConnAttempts_MYSQL(attempts int) MariaDBOption {
	return func(c *MariaDB) {
		c.connAttempts = attempts
	}
}

// 连接数据库超时时间
func ConnTimeout_MYSQL(timeout time.Duration) MariaDBOption {
	return func(c *MariaDB) {
		c.connTimeout = timeout
	}
}
