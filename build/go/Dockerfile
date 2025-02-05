FROM golang:1.23.3-alpine

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download
RUN apk add build-base

COPY ./ /app/