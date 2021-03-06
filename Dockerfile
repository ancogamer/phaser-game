FROM golang:alpine3.13 AS base 
WORKDIR /app
COPY . .
RUN go mod download

RUN go build -o app main.go

