package mysql

import (
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func TestNew(t *testing.T) {

	for i := 0; i < 1000; i++ {
		go func() {
			mdb := New("root:root@tcp(localhost:3306)/user?parseTime=True&loc=Local&charset=utf8mb4&collation=utf8mb4_unicode_ci", MaxIdleConns(4), MaxOpenConns(4))
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
