# Golang base image
FROM golang:alpine

RUN apk --no-cache add ca-certificates

WORKDIR /go/src/github.com/rl404/go-malscraper

COPY . .

RUN cp config/config_for_docker.json config/config.json

WORKDIR /go/src/github.com/rl404/go-malscraper/cmd/malscraper

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o go-malscraper

CMD ["./go-malscraper"]