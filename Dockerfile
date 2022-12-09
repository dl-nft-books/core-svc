FROM golang:1.18-alpine as buildbase

RUN apk add git build-base

WORKDIR /go/src/gitlab.com/tokend/nft-books/generator-svc
COPY vendor .
COPY . .

RUN GOOS=linux go build  -o /usr/local/bin/generator-svc /go/src/gitlab.com/tokend/nft-books/generator-svc


FROM alpine:3.9

COPY --from=buildbase /usr/local/bin/generator-svc /usr/local/bin/generator-svc
RUN apk add --no-cache ca-certificates

ENTRYPOINT ["generator-svc"]
