FROM golang:alpine

ENV SRC_DIR=/go/src/github.com/so-heee/graphql-example/plan2
ENV  GO111MODULE=on

WORKDIR $SRC_DIR

RUN apk update && apk add git

RUN go mod init graphql-example/plan2
RUN GO111MODULE=off go get github.com/oxequa/realize