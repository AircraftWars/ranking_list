package tools

import "github.com/go-redis/redis"

var _rdb *redis.Client

func init() {
	_rdb = redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
		DB:   0,
	})
}

func GetRedisDB() *redis.Client {
	return _rdb
}
