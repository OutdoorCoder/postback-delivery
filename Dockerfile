# dockerfile for kochava project
FROM ubuntu:14.04 as base


FROM golang:onbuild as gobase
WORKDIR /build
# This could be optimized by just copying the necessary files to run go
COPY . .
RUN CGO_ENABLED=0 go build -o main .



FROM php:7.2-cli
COPY ./test.php /
COPY ./wrapper_script.sh /wrapper_script.sh
COPY --from=gobase /build /

CMD ["./wrapper_script.sh"]
