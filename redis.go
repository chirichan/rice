package rice

import (
	"context"
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/go-redis/redis/v9"
)

const (
	// _defaultAddr     = "127.0.0.1:6379"
	// _defaultPoolSize = 10
	_defaultPassword = ""
	_defaultDB       = 0
)

type Redis struct{ *redis.Client }

var (
	ctx       = context.Background()
	redisOnce sync.Once
	rdb       = &Redis{}
)

func NewRedis(addr string, opts ...RedisOption) (*Redis, error) {

	var err error

	redisOnce.Do(func() {
		options := &redis.Options{
			Addr:     addr,
			Password: _defaultPassword,
			DB:       _defaultDB,
		}

		for _, opt := range opts {
			opt(options)
		}

		rdb.Client = redis.NewClient(options)

		_, err = rdb.Ping(context.Background()).Result()
	})

	return rdb, err
}

type RedisOption func(*redis.Options)

func RedisAddr(addr string) RedisOption {
	return func(options *redis.Options) {
		options.Addr = addr
	}
}

func RedisPoolSize(poolSize int) RedisOption {
	return func(options *redis.Options) {
		options.PoolSize = poolSize
	}
}

func RedisPassword(pwd string) RedisOption {
	return func(options *redis.Options) {
		options.Password = pwd
	}
}

func RedisDB(db int) RedisOption {
	return func(options *redis.Options) {
		options.DB = db
	}
}

func NewRedisDB() *Redis { return rdb }

func SetStruct(key string, data any, expiration ...time.Duration) error {

	var (
		expire time.Duration
		b      []byte
		err    error
	)

	if len(expiration) == 1 {
		expire = expiration[0]
	} else {
		expire = 0
	}

	if data == nil {
		b = nil
	} else {
		b, err = json.Marshal(data)
		if err != nil {
			return err
		}
	}

	_, err = rdb.Set(ctx, key, b, expire).Result()

	return err
}

// GetStruct 如果 key 不存在，返回 error; 如果 key 为空，返回零值
func GetStruct[T any](key string) (T, error) {

	var data T

	b, err := rdb.Get(ctx, key).Bytes()
	if err != nil {
		return data, err
	}

	if len(b) == 0 {
		return data, err
	}

	err = json.Unmarshal(b, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

func DeleteStructByPrefix(prefix string) error {

	iter := rdb.Scan(ctx, 0, prefix+"*", 0).Iterator()

	for iter.Next(ctx) {
		err := rdb.Del(ctx, iter.Val()).Err()
		if err != nil {
			return err
		}
	}

	return nil
}

func DeleteStruct(key string) error {

	_, err := rdb.Del(ctx, key).Result()
	return err
}

func SetString(key, value string, expiration ...time.Duration) error {

	var expire time.Duration

	if len(expiration) == 1 {
		expire = expiration[0]
	} else {
		expire = 0
	}

	_, err := rdb.Set(ctx, key, value, expire).Result()

	return err
}

func GetString(key string) string { return rdb.Get(ctx, key).String() }

func HSetStruct(key, field string, data any) error {

	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = rdb.HSet(ctx, key, field, b).Result()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func HGetStruct[T any](key, field string) (T, error) {

	var data T

	b, err := rdb.HGet(ctx, key, field).Bytes()
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(b, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}
