# syntax=docker/dockerfile:1

# Build

FROM golang:1.21.3 AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o ico .

# Deploy 

FROM ubuntu:latest AS deploy

WORKDIR /
COPY json/api_province_with_amphure_tambon.json json/

COPY --from=build /app/ico /usr/local/bin

EXPOSE 1323

ENTRYPOINT [ "ico" ]