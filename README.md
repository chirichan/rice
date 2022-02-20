# rice

## Usage

**Database**

```go
import "github.com/woxingliu/rice/database"
```

```bash
go mod tidy
```

postgresql

```go
pg, err := database.NewPostgresDB(
	"postgres://postgres:123456@localhost:5432/user?sslmode=disable",
	database.ConnAttempts_Postgres(10),
	database.MaxPoolSize_Postgres(4),
)
if err != nil {
	return err
}
defer pg.Close()
```

mariadb

```go
mariaDB := database.NewMariaDB(
	"root:root@tcp(localhost:3306)/user?parseTime=True&loc=Local&charset=utf8mb4&collation=utf8mb4_unicode_ci",
	database.MaxIdleConns_MYSQL(4),
	database.MaxOpenConns_MYSQL(4),
)
defer mariaDB.Close()
```
