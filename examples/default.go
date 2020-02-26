package main

import (
	"fmt"

	"github.com/rl404/go-malscraper/pkg/malscraper"
)

func main() {
	// Initiate the malscraper with default config.
	mal := malscraper.Default()

	// Get anime id 1 (Cowboy Bebop)
	parser, err := mal.GetAnime(1)

	// Print the data
	fmt.Println(parser.Data, err)
}
