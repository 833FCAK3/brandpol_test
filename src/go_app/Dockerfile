FROM golang:1.21.6-alpine

RUN apk add --no-cache tzdata

WORKDIR /src/go_app/

COPY /src/go_app .

RUN go mod download
RUN go build -o app

ADD https://github.com/ufoscout/docker-compose-wait/releases/download/2.11.0/wait /wait
RUN chmod +x /wait
