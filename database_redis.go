package rice

import (
	"context"
	"log"
	"sync"

	"github.com/go-redis/redis/v8"
)

var (
	redis_once  sync.Once
	redisClient *Redis
)

type Redis struct {
	*redis.Client
	ctx context.Context
}

func NewRedisClient(addr, pwd string, db, poolSize int) *Redis {

	redis_once.Do(func() {
		redisClient = &Redis{
			Client: redis.NewClient(&redis.Options{
				Addr:     addr,
				Password: pwd,
				DB:       db,
				PoolSize: poolSize,
			}),
			ctx: context.Background(),
		}
		s, err := redisClient.Ping(redisClient.ctx).Result()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("redis连接成功: %v\n", s)
	})

	return redisClient
}

func (r *Redis) Close() {
	if r.Client != nil {
		err := r.Client.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
}
