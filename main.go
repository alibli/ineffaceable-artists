package main

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

// just for test
func main() {
	ctx := context.Background()
	// Ensure that you have Redis running on your system
	rdb := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	// Ensure that the connection is properly closed gracefully
	defer rdb.Close()

	result, error := rdb.Get(ctx, "company").Result()
	if error != nil {
		fmt.Println("Error")
	}
	fmt.Println(result)

}
