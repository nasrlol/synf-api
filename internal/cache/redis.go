// Package cache contains the redis setup for caching device resources
package cache

import (
	"github.com/redis/go-redis/v9"
)

func Init() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:8050",
		Password: "", // No password set
		DB:       0,  // Use default DB
		Protocol: 2,  // Connection protocol
	})

	return client
}
