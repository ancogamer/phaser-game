FROM alpine3.13 AS base 
WORKDIR /app

RUN go build -o app main.go
