# syntax=docker/dockerfile:1
FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o /iot-server

EXPOSE 8080

CMD ["/iot-server"]
