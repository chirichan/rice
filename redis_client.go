package rice

import (
	"github.com/redis/go-redis/v9"
)

type RedisOption func(*redis.Options)

func NewRedisClient(addr string, opts ...RedisOption) *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	for _, opt := range opts {
		opt(rdb.Options())
	}
	return rdb
}
