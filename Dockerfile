# Golang base image
FROM golang:latest as go_compiler

WORKDIR /go/src/github.com/rl404/go-malscraper

COPY . .

RUN go get ./...

WORKDIR /go/src/github.com/rl404/go-malscraper/cmd/malscraper

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o go-malscraper


# New stage from scratch
FROM alpine:3.10 as go_mal_api_image

RUN apk add --no-cache bash
RUN apk add --no-cache ca-certificates

WORKDIR /docker/bin

COPY --from=go_compiler /go/src/github.com/rl404/go-malscraper/cmd/malscraper/go-malscraper go-malscraper

CMD ["/docker/bin/go-malscraper"]

EXPOSE 8005
