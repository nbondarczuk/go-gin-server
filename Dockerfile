FROM golang:1.22.0-alpine as builder

ARG VERSION

ENV GOPATH /go

ENV GOCACHE /go/caches/go-build

RUN apk add git alpine-sdk

ADD . /workpath

WORKDIR /workpath

# build the source
RUN make go/tidy
RUN make

# use a minimal alpine image
FROM alpine:latest

# add ca-certificates in case you need them
RUN apk update && apk add ca-certificates

# set working directory
WORKDIR /work

COPY --from=builder /workpath/bin/go-gin-server /work/go-gin-server
COPY config/config.yaml /work/config/config.yaml
USER 1001

# run the binary
CMD ["./go-gin-server"]
