package rice

import (
	"context"
	"encoding/json"
	"log"
	"sync"
	"time"

	"github.com/go-redis/redis/v8"
)

const (
	_defaultAddr     = "127.0.0.1:6379"
	_defaultPoolSize = 10
	_defaultPassword = ""
	_defaultDB       = 0
)

type Redis struct{ *redis.Client }

var (
	ctx         = context.Background()
	redisOnce   sync.Once
	redisClient = &Redis{}
)

// func NewRedisCluster(addrs []string, opts ...RedisOption) {
//
// 	redis.NewClusterClient(&redis.ClusterOptions{
// 		Addrs: addrs,
// 	})
// }

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

		redisClient.Client = redis.NewClient(options)

		_, err = redisClient.Ping(context.Background()).Result()
	})

	return redisClient, err
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

func NewRedisDB() *Redis { return redisClient }

func SetStruct(key string, data any, expiration ...time.Duration) error {

	var expire time.Duration

	if len(expiration) == 1 {
		expire = expiration[0]
	} else {
		expire = 0
	}

	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = redisClient.Set(ctx, key, b, expire).Result()
	if err != nil {
		return err
	}

	return nil
}

func DeleteStructByPrefix(prefix string) error {

	iter := redisClient.Scan(ctx, 0, prefix+"*", 0).Iterator()

	for iter.Next(ctx) {
		err := redisClient.Del(ctx, iter.Val()).Err()
		if err != nil {
			return err
		}
	}

	return nil
}

func DeleteStruct(key string) error {

	_, err := redisClient.Del(ctx, key).Result()
	return err
}

func GetStruct[T any](key string) (T, error) {

	var data T

	b, err := redisClient.Get(ctx, key).Bytes()
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(b, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}

func HSetStruct(key, field string, data any) error {

	b, err := json.Marshal(data)
	if err != nil {
		return err
	}

	_, err = redisClient.HSet(ctx, key, field, b).Result()
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func HGetStruct[T any](key, field string) (T, error) {

	var data T

	b, err := redisClient.HGet(ctx, key, field).Bytes()
	if err != nil {
		return data, err
	}

	err = json.Unmarshal(b, &data)
	if err != nil {
		return data, err
	}
	return data, nil
}
