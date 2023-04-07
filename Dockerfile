FROM golang:1.18-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/github.com/dl-nft-books/core-svc
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/core-svc /go/src/github.com/dl-nft-books/core-svc


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/core-svc /usr/local/bin/core-svc
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["generator-svc"]
