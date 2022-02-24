package tests

import (
	"testing"

	"github.com/woxingliu/rice"
)

func TestNewPostgresDB(t *testing.T) {

	pg, err := rice.NewPostgresDB(
		"postgres://postgres:123456@localhost:5432/user?sslmode=disable",
		rice.ConnAttempts_Postgres(10),
		rice.MaxPoolSize_Postgres(4),
	)
	if err != nil {
		t.Error(err)
	}
	defer pg.Close()
}
