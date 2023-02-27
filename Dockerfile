FROM golang:1.20.0-alpine

RUN apk update && apk add git && apk add alpine-sdk

ENV APP_ROOT /go/src/

RUN mkdir $APP_ROOT

WORKDIR $APP_ROOT

ENV CGO_ENABLED=1 \
  GOOS=linux \
  GOARCH=arm64 \
  GO111MODULE=on

COPY . $APP_ROOT

EXPOSE 8080