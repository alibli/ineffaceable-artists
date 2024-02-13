package main

import (
	"context"
	"fmt"
	"log"
	"github.com/redis/go-redis/v9"
)

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

	// Perform basic diagnostic to check if the connection is working
	// Expected result > ping: PONG
	// If Redis is not running, error case is taken instead
	status, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalln("Redis connection was refused")
	}
	fmt.Println(status)
	result, error := rdb.Get(ctx, "company").Result()
	if error != nil {
		fmt.Println("not found company from redis")
		result = "Ramzinex"
		// ttl is in nanosecond
		rdb.Set(ctx, "company", result, 8000000000)
	} else {
		fmt.Println(result)
	}
}
