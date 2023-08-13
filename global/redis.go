package global

import (
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
}

func Redis() *redis.Client {
	return redisClient

}
