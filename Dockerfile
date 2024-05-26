FROM golang:1.22.2-alpine AS builder


WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o app ./cmd

CMD ["./app"]