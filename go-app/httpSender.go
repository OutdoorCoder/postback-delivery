package main

import (
		"fmt"
		"log"
		"encoding/json"
		"strings"
		//"net/http"
		"github.com/gomodule/redigo/redis"
)

type Postback struct {
	RequestMethod string
	Url string
	Data []map[string]string
}

func sendHttpRequest(pback Postback){

	switch pback.RequestMethod {
	case "GET":
			//resp, err := http.Get("http://example.com/")
			for _, dataMap := range pback.Data {

				httpUrl := pback.Url
				for k, v := range dataMap {
					fmt.Println("k:", k, "v:", v)

					httpUrl = strings.Replace(httpUrl, "{" + k + "}", v, 1)
					fmt.Println("Request String: " + httpUrl)
					//TODO: for each pair send an http request
				}
			}



			fmt.Println("GET")
		case "POST":
			fmt.Println("POST")
		default:
			fmt.Println("ERROR")
	}

	//resp, err := http.PostForm("http://example.com/form",
	//url.Values{"key": {"Value"}, "id": {"123"}})
}

func main() {

	conn, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	defer conn.Close()

	//pull first value from Redis
	val, err := redis.String(conn.Do("RPOP", "postback-list"))
	if err != nil {
	    fmt.Println(err)
	}
	//fmt.Println(val)

	var valByte []byte = []byte(val)

	var pback Postback
	err = json.Unmarshal(valByte, &pback)

	//Send http request
	sendHttpRequest(pback)

	//fmt.Println(pback)
}
