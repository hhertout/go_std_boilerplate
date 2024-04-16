FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

RUN go install github.com/cosmtrek/air@latest

COPY . .

RUN go build -o main cmd/api/main.go
