# This Dockerfile builds grsd from source and creates a small (55 MB) docker container based on alpine linux.
#
# Clone this repository and run the following command to build and tag a fresh grsd amd64 container:
#
# docker build . -t yourregistry/grsd
#
# You can use the following command to buid an arm64v8 container:
#
# docker build . -t yourregistry/grsd --build-arg ARCH=arm64v8
#
# For more information how to use this docker image visit:
# https://github.com/Groestlcoin/grsd/tree/master/docs
#
# 1331  Mainnet Groestlcoin peer-to-peer port
# 1444  Mainet RPC port

ARG ARCH=amd64

FROM golang:1.16-alpine3.12 AS build-container
RUN apk add --no-cache git
RUN apk add --no-cache make

ARG ARCH
ARG GITHUB_TOKEN
ENV GO111MODULE=on

ADD . /app
WORKDIR /app
RUN set -ex \
  && git config --global url."https://gh:${GITHUB_TOKEN}@github.com/".insteadOf "https://github.com/" \
  && if [ "${ARCH}" = "amd64" ]; then export GOARCH=amd64; fi \
  && if [ "${ARCH}" = "arm32v7" ]; then export GOARCH=arm; fi \
  && if [ "${ARCH}" = "arm64v8" ]; then export GOARCH=arm64; fi \
  && echo "Compiling for $GOARCH" \
  && make install

FROM $ARCH/alpine:3.12

COPY --from=build-container /go/bin /bin

VOLUME ["/root/.grsd"]

EXPOSE 1331 1444

ENTRYPOINT ["grsd"]
