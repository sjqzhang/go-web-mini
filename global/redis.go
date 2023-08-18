package global

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go-web-mini/config"
)

var redisClient *redis.Client

//从配置中初始化redis

func InitRedis() {
	var opts redis.Options
	opts.Addr = config.Conf.Redis.Addr
	opts.DB = config.Conf.Redis.DB
	opts.Password = config.Conf.Redis.Password
	redisClient = redis.NewClient(&opts)
	if config.Conf.Redis.Enable {
		if _, err := redisClient.Ping(context.Background()).Result(); err != nil {
			panic(err)
		}
	}
}

func Redis() *redis.Client {
	return redisClient

}
