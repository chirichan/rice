package tests

import (
	"testing"

	"github.com/woxingliu/ha/database"
)

func TestNewPostgresDB(t *testing.T) {

	pg, err := database.NewPostgresDB(
		"postgres://postgres:123456@localhost:5432/user?sslmode=disable",
		database.ConnAttempts_Postgres(10),
		database.MaxPoolSize_Postgres(4),
	)
	if err != nil {
		t.Error(err)
	}
	defer pg.Close()
}
