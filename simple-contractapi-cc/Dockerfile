# Copyright the Hyperledger Fabric contributors. All rights reserved.
# SPDX-License-Identifier: Apache-2.0

ARG GO_VER=1.19
ARG ALPINE_VER=3.16

FROM golang:${GO_VER}-alpine${ALPINE_VER}

WORKDIR /home/shashank/Code/reference/open-source/simple-contractapi-cc
COPY . .

RUN go get -d -v ./...
RUN go install -v ./...

EXPOSE 9999
CMD ["/go/bin/simple"]
