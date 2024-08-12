package config

import "github.com/go-redis/redis"

var RS *redis.Client

func RedisConfig() {

	opt, err := redis.ParseURL("rediss://default:AVNS_PLeE9ZTbH05bv1NRWz3@redis-do-user-16975301-0.h.db.ondigitalocean.com:25061")
	if err != nil {
		panic(err)
	}

	client := redis.NewClient(opt)

	RS = client
}
