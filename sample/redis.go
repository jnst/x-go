package sample

import (
	"fmt"

	redis "gopkg.in/redis.v5"
)

// Connect is sample code
func Connect() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
}
