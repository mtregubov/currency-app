# Build front
FROM node:16.13.0-alpine3.14 AS front-builder
WORKDIR /build
COPY ./curr-ui .
RUN npm install
RUN npm run build

# Build go binary
FROM golang:1.19 AS builder
WORKDIR /build
COPY . .
RUN go mod download
ENV CGO_ENABLED=1
ENV GOOS=linux

RUN go build -o curcli -ldflags '-linkmode external -extldflags "-static"' ./cmd/main.go
COPY --from=front-builder /build/build /build/http/build
RUN go build -o cursrv -ldflags '-linkmode external -extldflags "-static"' ./http/server.go

# Build final image
FROM alpine:latest
RUN apk update --no-cache && apk add --no-cache ca-certificates
WORKDIR /app
RUN mkdir -p /app/data

COPY --from=builder /build/curcli /app/curcli
COPY --from=builder /build/cursrv /app/cursrv