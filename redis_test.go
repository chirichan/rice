package rice

import (
	"context"
	"encoding/json"
	"log"
	"testing"
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

func TestNewRedis(t *testing.T) {

	c, err := NewRedis("127.0.0.1:6379")
	if err != nil {
		t.Error(err)
	}
	t.Error(c.Ping(context.Background()).Result())

	user := User{
		Id:   2,
		Name: "test",
	}

	err = c.SetStruct("u1", &user)
	if err != nil {
		log.Println(err)
		return
	}

	var user2 User
	err = c.GetStruct("u3", &user2)
	if err != nil {
		log.Println(err)
		return
	}

	var user3 User
	err = c.HSetStruct("user", "u6", &User{Id: 4, Name: "4444"})
	if err != nil {
		log.Println(err)
		return
	}
	err = c.HGetStruct("user", "u7", &user3)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(user2.Name)

	log.Println(user3.Name)
}
