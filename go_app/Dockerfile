FROM golang:1.21.6-alpine

ENV TZ=$TIMEZONE

WORKDIR /go_app

COPY /go_app .

RUN go mod download
RUN go build -o app
