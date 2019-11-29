package main

import (
		"fmt"
		"log"
		"encoding/json"
		"time"
		"strings"
		"net/http"
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
			var httpUrl string = pback.Url
			for _, dataMap := range pback.Data {
				for k, v := range dataMap {
					fmt.Println("k:", k, "v:", v)

					httpUrl = strings.Replace(httpUrl, "{" + k + "}", v, 1)
					fmt.Println("Request String: " + httpUrl)
				}

				//send http request
				resp, err := http.Get(httpUrl)
				if err != nil {
					fmt.Println(err)
					log.Fatal(err)
				}
				fmt.Println(resp)
			}

		case "POST":
			//fmt.Println("POST")
		default:
			//fmt.Println("ERROR")
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

	for true {

		val, err := redis.String(conn.Do("RPOP", "postback-list"))
		if err != nil {
		    fmt.Println(err)
		}

		if len(val) > 0 {
			fmt.Println("Grabbed request")
			var valByte []byte = []byte(val)

			var pback Postback
			err = json.Unmarshal(valByte, &pback)

			//Send http request
			sendHttpRequest(pback)
		} else{
			fmt.Println("No requests left")
			time.Sleep(1 * time.Second)
		}

	}
}
