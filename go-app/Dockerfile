FROM golang:alpine

RUN mkdir /app
ADD . /app/
WORKDIR /app

RUN \
       apk add --no-cache bash git openssh && \
       go get -u github.com/gomodule/redigo/redis && \
       go build -o main .

CMD ["go","run","httpSender.go"]
