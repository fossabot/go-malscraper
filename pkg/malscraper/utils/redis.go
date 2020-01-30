package utils

import (
	"encoding/json"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/rl404/go-malscraper/pkg/malscraper/model/common"
)

// GetValueFromRedis to get value from redis. Assume redis is already connected.
func GetValueFromRedis(client *redis.Client, key string) (value string, err error) {
	if client != nil {
		return client.Get(key).Result()
	}
	return value, common.ErrMissingRedis
}

// UnmarshalFromRedis to unmarshal redis value to model.
func UnmarshalFromRedis(client *redis.Client, key string, model interface{}) (isFound bool, err error) {
	value, err := GetValueFromRedis(client, key)
	if err != nil && err != redis.Nil {
		return false, err
	}

	if err == redis.Nil {
		return false, nil
	}

	return true, json.Unmarshal([]byte(value), &model)
}

// SaveToRedis to save value to redis. Assume redis is already connected.
func SaveToRedis(client *redis.Client, key string, value interface{}, expiredTime time.Duration) error {
	if client != nil {
		jsonData, err := json.Marshal(value)
		if err != nil {
			return err
		}

		// To keep the time when data saved to redis.
		go saveRedisTime(client, key, expiredTime)

		return client.Set(key, jsonData, expiredTime).Err()
	}
	return common.ErrMissingRedis
}

// saveRedisTime to save saved time and expired time of the key to redis.
func saveRedisTime(client *redis.Client, key string, expiredTime time.Duration) {
	now := time.Now()
	client.Set(key+":time", redisTimeStr(now), expiredTime)
	client.Set(key+":expired", redisTimeStr(now.Add(expiredTime)), expiredTime)
}

// redisTimeStr to convert time format to string.
func redisTimeStr(t time.Time) string {
	return t.Format("2006-01-02 15:04:05")
}
