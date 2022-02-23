# rice

## Usage

**Database**

```go
import "github.com/woxingliu/rice"
```

```bash
go mod tidy
```

postgresql

```go
pg, err := rice.NewPostgresDB(
	"postgres://postgres:123456@localhost:5432/user?sslmode=disable",
	rice.ConnAttempts_Postgres(10),
	rice.MaxPoolSize_Postgres(4),
)
if err != nil {
	return err
}
defer pg.Close()
```

mariadb

```go
mariaDB := rice.NewMariaDB(
	"root:root@tcp(localhost:3306)/user?parseTime=True&loc=Local&charset=utf8mb4&collation=utf8mb4_unicode_ci",
	rice.MaxIdleConns_MYSQL(4),
	rice.MaxOpenConns_MYSQL(4),
)
defer mariaDB.Close()
```
