package config

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
)

// Config is config model for go-malscraper.
type Config struct {
	// Redis client and configuration for caching parsing data.
	// RedisClient will be assigned automatically if RedisConfig is set.
	// RedisClient can also be assigned directly without RedisConfig.
	// For more information, read `github.com/go-redis/redis/v7`.
	RedisClient *redis.Client
	RedisConfig *redis.Options

	// Expired time limit of the cached data (in seconds).
	CacheTime time.Duration

	// Does malscraper need to automatically clean any image url using
	// pkg/malscraper/utils.ImageURLCleaner() function.
	CleanImageURL bool

	// Does malscraper need to automatically clean any video url using
	// pkg/malscraper/utils.VideoURLCLeaner() function.
	CleanVideoURL bool

	// Not used...yet.
	Logging     bool
	LoggingPath string

	// Using or expressed in more detailed information to console.
	Verbose bool
}

var (
	// DefaultCacheTime is redis caching time.
	// Default value is 1 day (24 hours) if redis client is set.
	DefaultCacheTime = 24 * time.Hour

	// DefaultCleanImageURL for cleaning image url.
	DefaultCleanImageURL = true

	// DefaultCleanVideoURL for cleaning video url.
	DefaultCleanVideoURL = true

	// Default verbose boolean.
	DefaultVerbose = true

	// DefaultConfig for default malscraper config with its default field value.
	DefaultConfig = Config{
		RedisClient:   nil,
		RedisConfig:   nil,
		CacheTime:     0,
		CleanImageURL: DefaultCleanImageURL,
		CleanVideoURL: DefaultCleanVideoURL,
		Verbose:       DefaultVerbose,
	}
)

// Init to initiate config value.
func (c *Config) Init() {
	if c.RedisClient == nil && c.RedisConfig == nil {
		return
	}

	if c.RedisClient == nil && c.RedisConfig != nil {
		c.RedisClient = redis.NewClient(c.RedisConfig)
	}

	if c.RedisClient != nil && c.CacheTime == 0 {
		c.CacheTime = DefaultCacheTime
	}

	if c.Verbose {
		_, err := c.RedisClient.Ping().Result()
		if err != nil {
			fmt.Printf("Redis connection: failed (%v)\n", err.Error())
		} else {
			fmt.Println("Redis connection: success")
		}
	}
}
