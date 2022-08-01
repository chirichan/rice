# rice

```go
import "github.com/latext/rice"
```

```bash
go mod tidy
```

## Usage

**Database**

`Postgresql` 初始化

```go
pg, err := rice.NewPostgres("postgres://postgres:123456@localhost:5432/test?sslmode=disable")
if err != nil {
	return err
}
defer pg.Close()
```

`MariaDB` 初始化

```go
mariadb, err := rice.NewMaria("root:root@tcp(localhost:3306)/test?parseTime=True&loc=Local&charset=utf8mb4")
if err != nil {
	return err
}
defer mariadb.Close()
```

`Redis` 初始化

```go
rdb, err := rice.NewRedis("localhost:6379")
if err != nil {
	return err
}
defer rdb.Close()
```

`Pretty` 封装了 `database/sql` 标准库里 `sql.DB` 和 `sql.Tx` 共有的一些方法，使用 `Pretty` 时可以在 repo 的上一层初始化，使 repo 层既可以执行 `sql.DB` 的方法，也可以执行 `sql.Tx` 的方法，下面是示例。

```go
package repo

type Repo struct {
	Pretty
}

func NewRepo(db Pretty) Repo {
	return Repo{Pretty: db}
}

func (r Repo) WithTx(tx Pretty) Repo {
	return Repo{Pretty: tx}
}

func (r Repo) Create() error { return nil }

func (r Repo) Update() error { return nil }

func (r Repo) Find() error { return nil }
```

```go
package usecase

func UseCaseDB() {

	db := NewMariaDB()

	dbRepo := NewRepo(db)

	dbRepo.Find()
}

func UseCaseTx() error {

	tx, _ := NewMariaTx()
	defer tx.Rollback()

	txRepo := NewRepo(tx)

	err := txRepo.Create()
	if err != nil {
		return err
	}

	err = txRepo.Update()
	if err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	return nil
}
```

