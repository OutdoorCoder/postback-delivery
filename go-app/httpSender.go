package main

import (
		"fmt"
		"log"
		"encoding/json"
		"time"
		"strings"
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

func sendHttpRequest(pback Postback){

	switch pback.RequestMethod {
	case "GET":
			var httpUrl string = pback.Url
			for _, dataMap := range pback.Data {
				for k, v := range dataMap {
					//fmt.Println("k:", k, "v:", v)

					httpUrl = strings.Replace(httpUrl, "{" + k + "}", v, 1)
					//fmt.Println("Request String: " + httpUrl)
				}

				//send http request
				req, err := http.NewRequest("GET", httpUrl, nil)
				if err != nil {
					fmt.Println(err)
					log.Fatal(err)
				}
				//log delivery time, respose code, response time
				//and response body

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
				log.Printf(result.Total(time.Now()).String())
				log.Printf(res.Status)
				//fmt.Println(result.Total(time.Now()))

				// Show results
				//log.Printf("%+v", result)

				fmt.Println(req)
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
