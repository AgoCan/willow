package redis

import "github.com/go-redis/redis"

type Redis struct {
	Client *redis.Client
}

func New(host, password string, db int) *Redis {
	client := redis.NewClient(&redis.Options{
		Addr:     host,
		Password: password,
		DB:       db,
	})

	return &Redis{
		Client: client,
	}
}
