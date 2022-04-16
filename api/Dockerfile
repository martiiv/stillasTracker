FROM golang:1.18 AS build-env

ADD . /go1.18beta1.linux-amd64.tar.gz/src/app
WORKDIR /go1.18beta1.linux-amd64.tar.gz/src/app



RUN go mod tidy
RUN go mod download
RUN go mod vendor
RUN go mod verify
CMD ["go", "run", "cmd/main.go"]