package database

import "time"

// Option -.
type PostgresOption func(*Postgres)

// MaxPoolSize -.
func MaxPoolSize_Postgres(size int) PostgresOption {
	return func(c *Postgres) {
		c.maxPoolSize = size
	}
}

// ConnAttempts -.
func ConnAttempts_Postgres(attempts int) PostgresOption {
	return func(c *Postgres) {
		c.connAttempts = attempts
	}
}

// ConnTimeout -.
func ConnTimeout_Postgres(timeout time.Duration) PostgresOption {
	return func(c *Postgres) {
		c.connTimeout = timeout
	}
}
