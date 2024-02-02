FROM golang:1.21.6-alpine

ENV TZ=$TIMEZONE

WORKDIR /src/go_app/

COPY /src/go_app .

RUN go mod download
RUN go build -o app
