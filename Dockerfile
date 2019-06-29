FROM golang:1.12-alpine3.10

RUN apk update && \
    apk upgrade && \
    apk add --no-cache \
    gcc \
    libc-dev \
    git \
    czmq-dev \
    libzmq \
    libsodium

ENV GO111MODULE=on
EXPOSE 5563

RUN mkdir /go/src/sudare_contents
WORKDIR /go/src/sudare_contents

ENTRYPOINT ["./docker-entrypoint.sh"]
