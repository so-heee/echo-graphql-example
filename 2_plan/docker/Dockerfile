FROM golang:alpine

ENV SRC_DIR=/go/src/github.com/so-heee/graphql-example/2_plan
ENV  GO111MODULE=on

WORKDIR $SRC_DIR

RUN apk update && apk add git

RUN go mod init graphql-example/2_plan
RUN GO111MODULE=off go get github.com/oxequa/realize