package tools

import "github.com/go-redis/redis"

var _rdb *redis.Client

func init() {
	_rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})
}

func GetRedisDB() *redis.Client {
	return _rdb
}
