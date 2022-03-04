package rice

import (
	"context"
	"sync"

	"github.com/go-redis/redis/v8"
)

const (
	_defaultPoolSize = 10
	_defaultPassword = ""
	_defaultDB       = 0
)

var (
	redisOnce   sync.Once
	redisClient *redis.Client
)

func NewRedis(addr string, opts ...RedisOption) (*redis.Client, error) {

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

		redisClient = redis.NewClient(options)

		_, err = redisClient.Ping(context.Background()).Result()
	})

	return redisClient, err
}

type RedisOption func(*redis.Options)

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

func NewRedisDB() *redis.Client { return redisClient }
