package main

import (
		"fmt"
		"log"
		"encoding/json"
		"time"
		"strings"
		"errors"
		"net/http"
		"io/ioutil"
		"github.com/tcnksm/go-httpstat"
		"github.com/gomodule/redigo/redis"
)

type Postback struct {
	RequestMethod string
	Url string
	Data []map[string]string
}

func buildHttpUrls(pback Postback) []string {

	requestUrls := make([]string,len(pback.Data))

	for index, dataMap := range pback.Data {
		var httpUrl string = pback.Url
		for k, v := range dataMap {
			r := strings.NewReplacer("{" + k + "}", v)
			httpUrl = r.Replace(httpUrl)
		}
		requestUrls[index] = httpUrl
	}

	return requestUrls
}

func sendHttpRequest(httpUrl string, requestMethod string) {

	if requestMethod != "POST" && requestMethod != "GET" {
		log.Fatal(errors.New("Bad request method: " + requestMethod))
	}

	req, err := http.NewRequest(requestMethod, httpUrl, nil)
	if err != nil {
		log.Fatal(err)
	}
	//log delivery time, respose code, response time, and response body
	var result httpstat.Result
	ctx := httpstat.WithHTTPStat(req.Context(), &result)
	req = req.WithContext(ctx)

	client := http.DefaultClient
	res, err := client.Do(req)
	if err != nil {
			log.Fatal(err)
	}

	if res.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
				log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		log.Printf(bodyString)
	}

	res.Body.Close()
	result.End(time.Now())

	log.Printf(res.Status)
	log.Printf("%+v", result)
}

func main() {

	conn, err := redis.Dial("tcp", "redis:6379")
	if err != nil {
		fmt.Println(err)
		log.Fatal(err)
	}

	defer conn.Close()

	for true {

		redisVal, err := redis.String(conn.Do("RPOP", "postback-list"))
		if err != nil {
		    fmt.Println(err)
		}

		if len(redisVal) > 0 {
			fmt.Println("Grabbed request")
			var valByte []byte = []byte(redisVal)

			var pback Postback
			err = json.Unmarshal(valByte, &pback)

			//Send http request
			var requestUrls []string = buildHttpUrls(pback)
			//fmt.Println(requestUrls)

			for _, url := range requestUrls {
				sendHttpRequest(url, pback.RequestMethod)
			}
		}else {
			fmt.Println("No requests left")
			time.Sleep(1 * time.Second)
		}

	}
}
