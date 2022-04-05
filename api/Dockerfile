FROM golang:1.18 AS build-env

ADD . /go/src/app

WORKDIR /go/src/app

RUN go mod tidy
RUN go mod download
RUN go mod vendor
RUN go mod verify
CMD["go", "run", "main.go"]