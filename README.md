# postback-delivery
Service to function as a small scale simulation of distributing data to third parties in real time.

This app is built in Docker using docker-compose. This requires the file docker-compose.yml to organize the app
There are three containers in the app:

  1. PHP server to take in http requests in the following format:

   {
    "endpoint":{
      "method":"GET", //can also be POST
      "url":"https://postman-echo.com/get?foo1={key}&foo2={key2}&foo-n={key-n}"
    },
    "data":[
        {
          "key":"value",
          "key2":"value2"
          ....
          "key-n":"value-n"
        },
        {
          "key":"value",
          "key2":"value2"
          ....
          "key-n":"value-n"
        }
      ]
  }

  Each key listed in the url will be replaced by that keys corresponding value in the data map.
  Rhis format specifies http requests to be sent by the golang app (the third piece).
  These requests are turned into Postback objects were are pushed to Redis.

  2. Redis stack to take in Postback objects

  3. Golang app to send http requests pulled from Redis
  
  This app will continually check Redis for Postback objects. For each list of key value pairs in     the Postback objects data map the app will send an http request, after inserting the values into   their corresponding places in the url.
  
  This app logs the requests delivery time, response time, and response body
  
  
File Structure:

  Folder: go-app    - Holds files for the Golang application
    Dockerfile      - build instructions for the Golang container
    httpSender.go   - golang code for pulling Postback objects from Redis and then sending http                           requests built from those postback objects
    
  Folder: www      - Holds php files. Folder must be named www to work with apache server
    index.php      - PHP server. Takes in http requests, turns them into Postback objects, then                          pushes them to Redis




