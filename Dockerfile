# dockerfile for kochava project
FROM ubuntu:14.04
FROM php:7.2-cli

COPY ./test.php /

RUN php test.php
