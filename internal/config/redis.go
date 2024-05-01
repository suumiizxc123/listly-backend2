package config

import "github.com/go-redis/redis"

var RS *redis.Client

func RedisConfig() {

	opt, err := redis.ParseURL("redis://103.48.116.100:6379")
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(opt)

	RS = client
}
