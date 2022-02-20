package tests

import (
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/woxingliu/rice/database"
)

func TestNewMariaDB(t *testing.T) {

	for i := 0; i < 1000; i++ {
		go func() {
			mdb := database.NewMariaDB(
				"root:root@tcp(localhost:3306)/user?parseTime=True&loc=Local&charset=utf8mb4&collation=utf8mb4_unicode_ci",
				database.MaxIdleConns_MYSQL(4),
				database.MaxOpenConns_MYSQL(4),
			)
			defer mdb.Close()
			err := mdb.Ping()
			if err != nil {
				t.Error(err)
			} else {
				t.Error("success")
			}
		}()
	}

	time.Sleep(3 * time.Second)

}
