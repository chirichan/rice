package rice

import (
	"context"
	"github.com/go-redis/redis/v8"
	"log"
)

const (
	_defaultPoolSize = 10
	_defaultPassword = ""
	_defaultDB       = 0
)

var redisClient *redis.Client

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

func NewRedisClient(addr string, opts ...RedisOption) *redis.Client {

	once.Do(func() {
		options := &redis.Options{
			Addr:     addr,
			Password: _defaultPassword,
			DB:       _defaultDB,
		}

		for _, opt := range opts {
			opt(options)
		}

		redisClient = redis.NewClient(options)

		s, err := redisClient.Ping(context.Background()).Result()
		if err != nil {
			log.Printf("redis连接失败: %v\n", err)
		}
		log.Printf("redis连接成功: %v\n", s)
	})

	return redisClient
}

func GetRedisClient() *redis.Client { return redisClient }

func CloseRedisClient() error {
	if redisClient != nil {
		err := redisClient.Close()
		return err
	} else {
		return nil
	}
}
