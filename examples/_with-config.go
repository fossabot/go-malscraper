package main

import (
	"fmt"
	"time"

	"github.com/go-redis/redis/v7"
	"github.com/rl404/go-malscraper/pkg/malscraper"
	"github.com/rl404/go-malscraper/pkg/malscraper/config"
)

func main() {
	// Declare the config
	cfg := config.Config{
		RedisConfig: &redis.Options{
			Addr:     "localhost:6379",
			Password: "",
			DB:       0,
		},
		CacheTime:     15 * time.Second,
		CleanImageURL: true,
		CleanVideoURL: true,
		Verbose:       true,
	}

	// Initiate the malscraper with config
	newMal := malscraper.New(cfg)

	// Get anime id 21 (One Piece)
	parser, err := newMal.GetAnime(21)

	// Print the data
	fmt.Println(parser.Data, err)

	// Add sleep so goroutine for redis is
	// finished before the app exit because this
	// sample code is very short.
	time.Sleep(3 * time.Second)

	// Run this sample code again to get the anime
	// data from redis. This time, this sample code
	// will exit faster.
}
