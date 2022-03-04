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

redis

```go
rdb := rice.NewRedisClient("localhost:6379")
defer rdb.Close()
```

customer

```go

type Repo struct {
	Customer
}

func NewRepo(customer Customer) Repo {
	return Repo{Customer: customer}
}

func (r Repo) Create() error { return nil }

func (r Repo) Update() error { return nil }

func UsecaseDB() {

	db := NewCustomerDB()

	dbRepo := NewRepo(db)

	dbRepo.Create()
}

func UsecaseTx() error {

	tx := NewCustomerTx()
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
```
