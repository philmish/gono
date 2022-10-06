# syntax=docker/dockerfile:1

FROM golang:1.18-alpine

RUN apk add build-base

WORKDIR /app
ENV GO111MODULE=on

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . .

RUN cd cmd && go build -o app

ENTRYPOINT ["/app/cmd/app"]
