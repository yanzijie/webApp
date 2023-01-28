package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/yanzijie/webApp/settings"
)

var rdb *redis.Client

func Init(cfg *settings.RedisConfig) (err error) {
	addr := fmt.Sprintf("%s:%d",
		cfg.Host,
		cfg.Port,
	)

	rdb = redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.Password,
		DB:       cfg.Db,
		PoolSize: cfg.PoolSize,
	})
	_, err = rdb.Ping().Result()
	return
}

func Close() {
	_ = rdb.Close()
}
