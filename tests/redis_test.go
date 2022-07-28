package tests

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"github.com/woxingliu/rice"
)

type User struct {
	Id   int64
	Name string
}

func (user *User) Marshal() ([]byte, error) {
	return json.Marshal(user)
}

func (user *User) Unmarshal(data []byte) error {
	return json.Unmarshal(data, user)
}

func TestNewRedis2(t *testing.T) {

	c, err := rice.NewRedis("127.0.0.1:6379")
	if err != nil {
		t.Error(err)
	}
	t.Error(c.Ping(context.Background()).Result())

}

func TestNewRedis(t *testing.T) {

	c, err := rice.NewRedis("127.0.0.1:6379")
	if err != nil {
		t.Error(err)
	}
	t.Error(c.Ping(context.Background()).Result())

	user := User{
		Id:   2,
		Name: "test123",
	}

	err = rice.HSetStruct("user", "u8", user)
	if err != nil {
		t.Error(err)
	}

	u2, _ := rice.HGetStruct[User]("user", "u8")
	fmt.Printf("u2: %v\n", u2)

	// err = c.SetStruct("u1", &user)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// var user2 User
	// err = c.GetStruct("u3", &user2)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }

	// var user3 User
	// err = c.HSetStruct("user", "u6", &User{Id: 4, Name: "4444"})
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// err = c.HGetStruct("user", "u7", &user3)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// log.Println(user2.Name)

	// log.Println(user3.Name)

	u, err := rice.GetStruct[User]("u1")
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("u: %v\n", u)

}
