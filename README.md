# rice

封装了一些常用但又记不住名字的方法。

```go
import "github.com/woxingliu/rice"
```

```bash
go mod tidy
```

## Usage

**Database**

postgresql

```go
pg, err := rice.NewPostgresDB("postgres://postgres:123456@localhost:5432/user?sslmode=disable")
if err != nil {
	return err
}
defer pg.Close()
```

mariadb

```go
mariaDB := rice.NewMariaDB("root:root@tcp(localhost:3306)/user?parseTime=True&loc=Local&charset=utf8mb4")
defer mariaDB.Close()
```
