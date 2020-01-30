package config

import (
	"testing"
	"time"

	"github.com/go-redis/redis/v7"
)

// TestInitConfig to test init config.
func TestInitConfig(t *testing.T) {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	configs := []*Config{
		// empty config
		{},
		// no redis
		{
			Verbose: true,
		},
		// redis without cache time
		{
			RedisConfig: &redis.Options{
				Addr: "localhost:6379",
			},
			Verbose: true,
		},
		// redis invalid port
		{
			RedisConfig: &redis.Options{
				Addr: "localhost:62379",
			},
			Verbose: true,
		},
		// redis with cache time
		{
			RedisConfig: &redis.Options{
				Addr: "localhost:6379",
			},
			CacheTime: 5 * time.Second,
			Verbose:   true,
		},
		// assigned redis client
		{
			RedisClient: client,
		},
	}

	for _, c := range configs {
		c.Init()
	}
}
