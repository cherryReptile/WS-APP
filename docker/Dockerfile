FROM golang:1.19-alpine

WORKDIR /usr/src/app

RUN go install github.com/cosmtrek/air@latest
RUN go install github.com/go-delve/delve/cmd/dlv@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN apk update
RUN apk add net-tools lsof

RUN rm -rf /var/cache/apk/*