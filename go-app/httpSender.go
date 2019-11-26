package main

import (
		"fmt"
		"github.com/go-redis/redis"
)

func main() {

	client := redis.NewClient(&redis.Options{
		Addr: "redis:6379",
		Password: "",
		DB: 0,
	})

	pong, err := client.Ping().Result()

	fmt.Println(pong, err)

	val, err := client.Do("RPOP", "postback-list").Result()
	if err != nil {
	    fmt.Println(err)
	}

	fmt.Println(val)
}
