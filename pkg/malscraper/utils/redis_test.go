package utils

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/rl404/go-malscraper/pkg/malscraper/model/common"
)

// redis key and value for testing.
const RedisKeyTest = "gotest"
const RedisValueTest = "gotest value"

// TestSaveToRedis to test saving value to redis.
func TestSaveToRedis(t *testing.T) {
	type scenario struct {
		Client        *redis.Client
		Key           string
		Value         interface{}
		ExpiredTime   time.Duration
		ExpectedError bool
	}

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	scenarios := []scenario{
		{
			Key:           RedisKeyTest,
			Value:         RedisValueTest,
			ExpiredTime:   30 * time.Second,
			ExpectedError: true,
		},
		{
			Client:        client,
			Key:           RedisKeyTest,
			Value:         make(chan int),
			ExpiredTime:   30 * time.Second,
			ExpectedError: true,
		},
		{
			Client:        client,
			Key:           RedisKeyTest,
			Value:         RedisValueTest,
			ExpiredTime:   30 * time.Second,
			ExpectedError: false,
		},
	}

	for _, s := range scenarios {
		err := SaveToRedis(s.Client, s.Key, s.Value, s.ExpiredTime)
		if (err != nil) != s.ExpectedError {
			t.Errorf("Expected %v got %v", s.ExpectedError, err)
		}
	}
}

// TestGetValue to test getting value from redis.
func TestGetValue(t *testing.T) {
	type scenario struct {
		Client        *redis.Client
		Key           string
		ExpectedValue string
		ExpectedError error
	}

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	scenarios := []scenario{
		{
			Key:           RedisKeyTest,
			ExpectedValue: "",
			ExpectedError: common.ErrMissingRedis,
		},
		{
			Client:        client,
			Key:           RedisKeyTest,
			ExpectedValue: RedisValueTest,
			ExpectedError: nil,
		},
	}

	var strValue string
	for _, s := range scenarios {
		value, err := GetValueFromRedis(s.Client, s.Key)
		json.Unmarshal([]byte(value), &strValue)
		if err != s.ExpectedError {
			t.Errorf("Expected error %v got %v", s.ExpectedError, err)
		}
		if strValue != s.ExpectedValue {
			t.Errorf("Expected value %v got %v", s.ExpectedValue, value)
		}
	}
}

// TestUnmarshalFromRedis to test unmarshaling redis value to model.
func TestUnmarshalFromRedis(t *testing.T) {
	var unmarshaledValue string

	type scenario struct {
		Client        *redis.Client
		Key           string
		Model         interface{}
		ExpectedFound bool
		ExpectedError error
	}

	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	scenarios := []scenario{
		{
			Key:           RedisKeyTest,
			Model:         unmarshaledValue,
			ExpectedFound: false,
			ExpectedError: common.ErrMissingRedis,
		},
		{
			Client:        client,
			Key:           "randomKey",
			Model:         unmarshaledValue,
			ExpectedFound: false,
			ExpectedError: nil,
		},
		{
			Client:        client,
			Key:           RedisKeyTest,
			Model:         unmarshaledValue,
			ExpectedFound: true,
			ExpectedError: nil,
		},
	}

	for _, s := range scenarios {
		isFound, err := UnmarshalFromRedis(s.Client, s.Key, &s.Model)
		if err != s.ExpectedError {
			t.Errorf("Expected error %v got %v", s.ExpectedError, err)
		}

		if isFound != s.ExpectedFound {
			t.Errorf("Expected found %v got %v", s.ExpectedFound, isFound)
		}
	}
}
