# postback-delivery
Service to function as a small scale simulation of distributing data to third parties in real time.

There are three pieces to the app:

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
  
  This app will continually check Redis for Postback objects. For each list of key value pairs in the Postback objects data map the app will send an http request, after inserting the values into their corresponding places in the url.




