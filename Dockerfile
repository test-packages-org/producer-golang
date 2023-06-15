FROM golang:1.18-alpine as build_producer

ENV CGO_ENABLED 0
COPY . /service

WORKDIR /service
RUN go mod download

WORKDIR /service/app/services/producer
RUN go build

