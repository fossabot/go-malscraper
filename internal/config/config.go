package config

import (
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/rl404/go-malscraper/pkg/malscraper/config"
	"github.com/spf13/viper"
)

// GetConfig to parse config from config.json.
// GetConfig will looking for file config.json in /go-malscraper/config/ folder.
func GetConfig() (cfg config.Config, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("./../../config/")
	if err = viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			// If config.json file is not found, it will use the default config from
			// config_default.json. config_default.json should not be modified or deleted.
			viper.SetConfigName("config_default")
			viper.MergeInConfig()
		} else {
			return cfg, err
		}
	}

	if viper.IsSet("redis") {
		cfg.RedisConfig = &redis.Options{
			Addr:     viper.GetString("redis.address"),
			Password: viper.GetString("redis.password"),
			DB:       viper.GetInt("redis.db"),
		}
	}

	cfg.CacheTime = time.Duration(viper.GetInt("cacheTime")) * time.Second
	cfg.CleanImageURL = viper.GetBool("cleanImageURL")
	cfg.CleanVideoURL = viper.GetBool("cleanVideoURL")
	cfg.Verbose = viper.GetBool("verbose")

	return cfg, nil
}
