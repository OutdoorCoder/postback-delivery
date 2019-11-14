# dockerfile for kochava project
FROM ubuntu:14.04

FROM php:7.2-cli
COPY ./test.php /
RUN php test.php
# CMD ["php", "test.php"]

FROM golang:onbuild
COPY ./test.go /
RUN go build -o main .
CMD ./main
