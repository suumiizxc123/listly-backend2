package config

import "github.com/go-redis/redis"

var RS *redis.Client

func RedisConfig() {

	opt, err := redis.ParseURL("redis://206.189.82.163:6379")
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(opt)

	RS = client
}
