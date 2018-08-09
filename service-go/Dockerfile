FROM golang:1.10.2-alpine as builder
LABEL maintainer="will835559313@163.com"
COPY . /app
WORKDIR /app
RUN apk update && apk add git \
    && go get github.com/olivere/elastic \
    && go build \
    && rm -rf /go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/app .
CMD ["./app"]