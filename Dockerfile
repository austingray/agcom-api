FROM golang:1.15.7-alpine3.13
RUN apk --update add build-base bash
RUN mkdir /app
COPY . /app
COPY ./docker/db/docker-entrypoint-initdb.d/init.sql /docker-entrypoint-initdb.d/init.sql
WORKDIR /app
EXPOSE 8080
RUN go build -o server . 