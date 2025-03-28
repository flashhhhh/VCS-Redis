package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func NewRedisClient(addr string) *redis.Client {
	fmt.Println("Connecting to Redis at", addr)
	client := redis.NewClient(&redis.Options{
		Addr: addr,
	})

	// Test the connection
	if err := client.Ping(context.Background()).Err(); err != nil {
		panic(err)
	}

	return client
}