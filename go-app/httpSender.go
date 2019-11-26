package main

import (
		"fmt"
		"log"
		"encoding/json"
		"github.com/gomodule/redigo/redis"
)

type Postback struct {
	RequestMethod string
	Url string
	Mascot string
	Location string
}

func main() {

	conn, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	defer conn.Close()

	val, err := redis.String(conn.Do("RPOP", "postback-list"))
	if err != nil {
	    fmt.Println(err)
	}

	var valByte []byte = []byte(val)

	var pback Postback
	err = json.Unmarshal(valByte, &pback)

	fmt.Println(pback)
}
