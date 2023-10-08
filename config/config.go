package config

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var (
	RedisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6380",
	})
	Ctx = context.Background()
)
