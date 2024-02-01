FROM golang:1.21.7-slim-bookworm

RUN mkdir -p /go/

COPY go/ go/

WORKDIR /go

RUN go mod download
RUN go build -o app
