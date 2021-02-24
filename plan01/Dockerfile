FROM golang:alpine

ENV SRC_DIR=/go/src/github.com/so-heee/echo-graphql-example
ENV  GO111MODULE=on

WORKDIR $SRC_DIR

RUN apk update && apk add git

RUN go mod init echo-graphql-example
RUN go get github.com/labstack/echo/v4
RUN GO111MODULE=off go get github.com/oxequa/realize