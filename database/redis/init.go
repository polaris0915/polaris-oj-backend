package redis

import (
	"log"

	"github.com/go-redis/redis"
)

var RedisDB = InitRedis()

// 处理redis数据库连接
func InitRedis() *redis.Client {
	url := "redis://root:@localhost:6379/0?protocol=3"
	opts, err := redis.ParseURL(url)
	if err != nil {
		log.Println("redis init error: ", err)
	}

	return redis.NewClient(opts)
}
